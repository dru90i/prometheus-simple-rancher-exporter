package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

var (
	metricsPath     = getEnv("METRICS_PATH", "/metrics") // Путь для получения метрик
	listenAddress   = getEnv("LISTEN_ADDRESS", ":9191")  // Порт для получения метрик
	rancherURL      = os.Getenv("CATTLE_URL")            // URL rancher сервера. Пример: https://rancher.example.com
	accessKey       = os.Getenv("CATTLE_ACCESS_KEY")     // Access Key для Rancher API
	secretKey       = os.Getenv("CATTLE_SECRET_KEY")     // Secret Key для Rancher API
	resourceLimit   = getEnv("API_LIMIT", "100")         // Лимит ресурсов Rancher API (по-умолчанию: 100)
	namespaceFilter = getEnv("NAMESPACES", "all")        // Фильтр по namespace (по-умолчанию все. Пример: "kube-system,default,test")
)

var (
	clusterStates = []string{"active", "cordoned", "degraded", "disconnected", "drained", "draining", "healthy", "initializing", "locked", "purged", "purging", "reconnecting", "reinitializing", "removed", "running", "unavailable", "unhealthy", "upgraded", "upgrading"}
	nodeStates    = []string{"active", "cordoned", "drained", "draining", "provisioning", "registering", "unavailable"}
	podStates     = []string{"Pending", "Running", "Succeeded", "Failed", "Unknown", "Unavailable"}
	contStates    = []string{"waiting", "running", "terminated"}
)

func main() {
	if rancherURL == "" {
		log.Fatal("CATTLE_URL must be set and non-empty")
	}
	if accessKey == "" {
		log.Fatal("CATTLE_ACCESS_KEY must be set and non-empty")
	}
	if secretKey == "" {
		log.Fatal("CATTLE_SECRET_KEY must be set and non-empty")
	}
	http.HandleFunc(metricsPath, getMetrics)
	err := http.ListenAndServe(listenAddress, nil)
	if err != nil {
		log.Fatal("ListenAndServer: ", err)
	}
}

// getEnv - получаем переменную оркужения, если её нет - возвращаем значение по-умолчанию
func getEnv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

func getMetrics(w http.ResponseWriter, r *http.Request) {
	var cluster = new(Cluster)
	getJSON(rancherURL+"/v3/cluster?limit="+resourceLimit, accessKey, secretKey, &cluster)

	// Clusters
	fmt.Fprintln(w, "# HELP clusters")
	fmt.Fprintln(w, "# TYPE rancher_cluster gauge")
	for _, x := range cluster.Cluster {
		for i, y := range clusterStates {
			if y == x.State {
				fmt.Fprintln(w, "rancher_cluster{cluster=\""+x.Name+"\",id=\""+x.ID+"\"}", i)
			}
		}
	}

	// Namespace
	fmt.Fprintln(w, "# HELP namespaces cluster")
	fmt.Fprintln(w, "# TYPE rancher_cluster_namespace gauge")
	for _, x := range cluster.Cluster {
		if x.State == "active" {
			var namespace = new(Namespace)
			getJSON(rancherURL+"/v3/clusters/"+x.ID+"/namespaces?limit="+resourceLimit, accessKey, secretKey, &namespace)
			for _, y := range namespace.Namespace {
				fmt.Fprintln(w, "rancher_cluster_namespace{cluster=\""+x.Name+"\",id=\""+x.ID+"\",namespace=\""+y.ID+"\"} 1")
			}
		}
	}

	// Nodes
	fmt.Fprintln(w, "# HELP nodes cluster")
	fmt.Fprintln(w, "# TYPE rancher_cluster_node gauge")
	for _, x := range cluster.Cluster {
		if x.State == "active" {
			var node = new(Nodes)
			getJSON(rancherURL+"/v3/clusters/"+x.ID+"/nodes?limit="+resourceLimit, accessKey, secretKey, &node)
			for _, y := range node.Nodes {
				var nodeRole string = ""
				if y.Master {
					nodeRole = "Master"
				}
				if y.Worker && !y.Unschedulable {
					nodeRole = "Worker"
				}
				for i, z := range nodeStates {
					if z == y.State {
						fmt.Fprintln(w, "rancher_cluster_node{cluster=\""+x.Name+"\",id=\""+x.ID+"\",node=\""+y.Hostname+"\",node_role=\""+nodeRole+"\"}", i)
					}
				}
				fmt.Fprintln(w, "rancher_cluster_node_pods_count{cluster=\""+x.Name+"\",id=\""+x.ID+"\",node=\""+y.Hostname+"\"}", y.PodsInfo.Pods)
			}
		}
	}

	// Pods
	fmt.Fprintln(w, "# HELP pods cluster")
	fmt.Fprintln(w, "# TYPE rancher_cluster_pod gauge")
	var nsFilter = strings.Split(namespaceFilter, ",")
	for _, x := range cluster.Cluster {
		if x.State == "active" {
			var pod = new(Pods)
			getJSON(rancherURL+"/k8s/clusters/"+x.ID+"/v1/pods?limit="+resourceLimit, accessKey, secretKey, &pod)
			for _, z := range pod.Pods {
				for _, ns := range nsFilter {
					if ns == "all" || ns == z.PodsInfo.Namespace {
						for i, m := range podStates {
							if m == z.PodsStatus.State {
								fmt.Fprintln(w, "rancher_cluster_pod{cluster=\""+x.Name+"\",id=\""+x.ID+"\",namespace_pod=\""+z.PodsInfo.Namespace+"\",name=\""+z.PodsInfo.Name+"\"}", i)
							}
						}
						for _, p := range z.PodsStatus.ContInfo {
							j, err := json.Marshal(p.State)
							if err != nil {
								panic(err)
							}
							var state string = strings.Split(string(j), "\"")[1]
							for i, m := range contStates {
								if m == state && z.PodsStatus.State != "Succeeded" {
									fmt.Fprintln(w, "rancher_cluster_pod_container{cluster=\""+x.Name+"\",id=\""+x.ID+"\",namespace_pod=\""+z.PodsInfo.Namespace+"\",name_pod=\""+z.PodsInfo.Name+"\",name_container=\""+p.Name+"\"}", i)
								}
							}
							fmt.Fprintln(w, "rancher_cluster_pod_container_count_restart{cluster=\""+x.Name+"\",id=\""+x.ID+"\",namespace_pod=\""+z.PodsInfo.Namespace+"\",name_pod=\""+z.PodsInfo.Name+"\",name_container=\""+p.Name+"\"}", p.RestartCount)
						}
					}
				}
			}
		}
	}

	// Events "Warning"
	fmt.Fprintln(w, "# HELP events cluster")
	fmt.Fprintln(w, "# TYPE rancher_cluster_events gauge")
	for _, x := range cluster.Cluster {
		if x.State == "active" {
			var event = new(Events)
			getJSON(rancherURL+"/k8s/clusters/"+x.ID+"/v1/events?limit="+resourceLimit, accessKey, secretKey, &event)
			var numEvent int = 0
			for _, y := range event.Events {
				if y.Type == "Warning" {
					numEvent += y.Count
				}
			}
			fmt.Fprintln(w, "rancher_cluster_events{cluster=\""+x.Name+"\",id=\""+x.ID+"\"}", numEvent)
		}
	}
}

// getJSON - получаем из RANCHER API json
func getJSON(url string, accessKey string, secretKey string, target interface{}) error {
	client := &http.Client{}
	http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error Collecting JSON from API: ", err)
	}
	req.SetBasicAuth(accessKey, secretKey)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error Collecting JSON from API: ", err)
	}
	if resp.StatusCode != 200 {
		fmt.Println("Error Collecting JSON from API: ", resp.Status)
	}
	respFormatted := json.NewDecoder(resp.Body).Decode(target)
	resp.Body.Close()
	return respFormatted
}
