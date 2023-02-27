package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	clusterStates = []string{"active", "cordoned", "degraded", "disconnected", "drained", "draining", "healthy", "initializing", "locked", "purged", "purging", "reconnecting", "reinitializing", "removed", "running", "unavailable", "unhealthy", "upgraded", "upgrading"}
	nodeStates    = []string{"active", "cordoned", "drained", "draining", "provisioning", "registering", "unavailable"}
	podStates     = []string{"Pending", "Running", "Succeeded", "Failed", "Unknown", "Unavailable"}
	contStates    = []string{"waiting", "running", "terminated"}
)

func addMetrics() map[string]*prometheus.GaugeVec {
	gaugeVecs := make(map[string]*prometheus.GaugeVec)

	gaugeVecs["cluster"] = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "rancher_cluster",
			Help: "Сlusters",
		}, []string{"cluster", "id"})

	gaugeVecs["cluster_node"] = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "rancher_cluster_node",
			Help: "cluster nodes",
		}, []string{"cluster", "id", "node", "node_role"})

	gaugeVecs["cluster_node_pods_count"] = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "rancher_cluster_node_pods_count",
			Help: "cluster node pods count",
		}, []string{"cluster", "id", "node"})

	gaugeVecs["cluster_node_stat"] = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "rancher_cluster_node_stat",
			Help: "cluster node status",
		}, []string{"cluster", "id", "status_node"})

	gaugeVecs["cluster_pod"] = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "rancher_cluster_pod",
			Help: "cluster pods",
		}, []string{"cluster", "id", "namespace_pod", "name"})

	gaugeVecs["cluster_pod_container"] = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "rancher_cluster_pod_container",
			Help: "cluster pods containers",
		}, []string{"cluster", "id", "namespace_pod", "name_pod", "name_container"})

	gaugeVecs["cluster_pod_container_count_restart"] = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "rancher_cluster_pod_container_count_restart",
			Help: "cluster pods containers count restart",
		}, []string{"cluster", "id", "namespace_pod", "name_pod", "name_container"})

	gaugeVecs["cluster_pod_stat"] = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "rancher_cluster_pod_stat",
			Help: "cluster pods status",
		}, []string{"cluster", "id", "status_pod"})

	gaugeVecs["cluster_container_stat"] = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "rancher_cluster_container_stat",
			Help: "cluster containers status",
		}, []string{"cluster", "id", "status_container"})

	gaugeVecs["cluster_events"] = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "rancher_cluster_events",
			Help: "cluster events",
		}, []string{"cluster", "id"})

	return gaugeVecs
}

func (collector *Exporter) Describe(ch chan<- *prometheus.Desc) {
	for _, m := range collector.gaugeVecs {
		m.Describe(ch)
	}
}

func (collector *Exporter) Collect(ch chan<- prometheus.Metric) {
	// Clusters
	var cluster = new(Cluster)
	getJSON(rancherURL+"/v3/cluster?limit="+resourceLimit, accessKey, secretKey, &cluster)
	for _, x := range cluster.Cluster {
		for i, y := range clusterStates {
			if y == x.State {
				collector.gaugeVecs["cluster"].With(prometheus.Labels{"cluster": x.Name, "id": x.ID}).Set(float64(i))
			}
		}
	}
	// Nodes
	for _, x := range cluster.Cluster {
		if x.State == "active" {
			var node = new(Nodes)
			var stat = make(map[string]int)
			getJSON(rancherURL+"/v3/clusters/"+x.ID+"/nodes?limit="+resourceLimit, accessKey, secretKey, &node)
			for _, y := range node.Nodes {
				var nodeRole string = ""
				if y.Master {
					nodeRole = "Master"
				}
				if y.Worker {
					nodeRole = "Worker"
				}
				for i, z := range nodeStates {
					if z == y.State {
						collector.gaugeVecs["cluster_node"].With(prometheus.Labels{"cluster": x.Name, "id": x.ID, "node": y.Hostname, "node_role": nodeRole}).Set(float64(i))
						stat[y.State]++
					}
				}
			}
			for i, n := range stat {
				collector.gaugeVecs["cluster_node_stat"].With(prometheus.Labels{"cluster": x.Name, "id": x.ID, "status_node": i}).Set(float64(n))
			}
		}
	}
	// Pods
	var nsFilter = strings.Split(namespaceFilter, ",")
	for _, x := range cluster.Cluster {
		if x.State == "active" {
			var pod = new(Pods)
			var statP = make(map[string]int)
			var statC = make(map[string]int)
			var podCount = make(map[string]int)
			getJSON(rancherURL+"/k8s/clusters/"+x.ID+"/v1/pods?limit="+resourceLimit, accessKey, secretKey, &pod)
			for _, z := range pod.Pods {
				podCount[z.Spec.NodeName]++
				for _, ns := range nsFilter {
					if ns == "all" || ns == z.PodsInfo.Namespace {
						for i, m := range podStates {
							if m == z.PodsStatus.State {
								collector.gaugeVecs["cluster_pod"].With(prometheus.Labels{"cluster": x.Name, "id": x.ID, "namespace_pod": z.PodsInfo.Namespace, "name": z.PodsInfo.Name}).Set(float64(i))
								statP[z.PodsStatus.State]++
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
									collector.gaugeVecs["cluster_pod_container"].With(prometheus.Labels{"cluster": x.Name, "id": x.ID, "namespace_pod": z.PodsInfo.Namespace, "name_pod": z.PodsInfo.Name, "name_container": p.Name}).Set(float64(i))
									statC[state]++
								}
							}
							collector.gaugeVecs["cluster_pod_container_count_restart"].With(prometheus.Labels{"cluster": x.Name, "id": x.ID, "namespace_pod": z.PodsInfo.Namespace, "name_pod": z.PodsInfo.Name, "name_container": p.Name}).Set(float64(p.RestartCount))
						}
					}
				}
			}
			for i, n := range statP {
				collector.gaugeVecs["cluster_pod_stat"].With(prometheus.Labels{"cluster": x.Name, "id": x.ID, "status_pod": i}).Set(float64(n))
			}
			for i, n := range statC {
				collector.gaugeVecs["cluster_container_stat"].With(prometheus.Labels{"cluster": x.Name, "id": x.ID, "status_container": i}).Set(float64(n))
			}
			for i, n := range podCount {
				collector.gaugeVecs["cluster_node_pods_count"].With(prometheus.Labels{"cluster": x.Name, "id": x.ID, "node": i}).Set(float64(n))
			}
		}
	}
	// Events
	for _, x := range cluster.Cluster {
		if x.State == "active" {
			var event = new(Events)
			getJSON(rancherURL+"/k8s/clusters/"+x.ID+"/v1/events?limit="+resourceLimit, accessKey, secretKey, &event)
			var numEvent int = 0
			for _, y := range event.Events {
				if y.Type == "Warning" {
					for _, ns := range nsFilter {
						if ns == "all" || ns == y.Namespace.Namespace {
							numEvent++
						}
					}
				}
			}
			collector.gaugeVecs["cluster_events"].With(prometheus.Labels{"cluster": x.Name, "id": x.ID}).Set(float64(numEvent))
		}
	}
	for _, m := range collector.gaugeVecs {
		m.Collect(ch)
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
