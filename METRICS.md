# Example Metrics


```
# HELP clusters
# TYPE rancher_cluster gauge
rancher_cluster{cluster="local",id="local"} 0
# HELP namespaces cluster
# TYPE rancher_cluster_namespace gauge
rancher_cluster_namespace{cluster="local",id="local",namespace="cattle-dashboards"} 1
rancher_cluster_namespace{cluster="local",id="local",namespace="cattle-fleet-clusters-system"} 1
rancher_cluster_namespace{cluster="local",id="local",namespace="cattle-fleet-local-system"} 1
rancher_cluster_namespace{cluster="local",id="local",namespace="cattle-fleet-system"} 1
rancher_cluster_namespace{cluster="local",id="local",namespace="cattle-global-data"} 1
rancher_cluster_namespace{cluster="local",id="local",namespace="cattle-global-nt"} 1
rancher_cluster_namespace{cluster="local",id="local",namespace="cattle-impersonation-system"} 1
rancher_cluster_namespace{cluster="local",id="local",namespace="cattle-monitoring-system"} 1
rancher_cluster_namespace{cluster="local",id="local",namespace="cattle-system"} 1
rancher_cluster_namespace{cluster="local",id="local",namespace="cluster-fleet-local-local-1a3d67d0a899"} 1
rancher_cluster_namespace{cluster="local",id="local",namespace="default"} 1
rancher_cluster_namespace{cluster="local",id="local",namespace="fleet-default"} 1
rancher_cluster_namespace{cluster="local",id="local",namespace="fleet-local"} 1
rancher_cluster_namespace{cluster="local",id="local",namespace="kube-node-lease"} 1
rancher_cluster_namespace{cluster="local",id="local",namespace="kube-public"} 1
rancher_cluster_namespace{cluster="local",id="local",namespace="kube-system"} 1
rancher_cluster_namespace{cluster="local",id="local",namespace="local"} 1
rancher_cluster_namespace{cluster="local",id="local",namespace="monitoring"} 1
rancher_cluster_namespace{cluster="local",id="local",namespace="user-bxncv"} 1
# HELP nodes cluster
# TYPE rancher_cluster_node gauge
rancher_cluster_node{cluster="local",id="local",node="local-node",node_role="Master"} 0
rancher_cluster_node_pods_count{cluster="local",id="local",node="local-node"} 18
rancher_cluster_node_stat{cluster="local",id="local",status_node="active"} 1
# HELP pods cluster
# TYPE rancher_cluster_pod gauge
rancher_cluster_pod{cluster="local",id="local",namespace_pod="kube-system",name="coredns-b96499967-zclgc"} 1
rancher_cluster_pod_container{cluster="local",id="local",namespace_pod="kube-system",name_pod="coredns-b96499967-zclgc",name_container="coredns"} 1
rancher_cluster_pod_container_count_restart{cluster="local",id="local",namespace_pod="kube-system",name_pod="coredns-b96499967-zclgc",name_container="coredns"} 4
rancher_cluster_pod_stat{cluster="local",id="local",status_pod="Running"} 1
rancher_cluster_container_stat{cluster="local",id="local",status_container="running"} 1
# HELP events cluster
# TYPE rancher_cluster_events gauge
rancher_cluster_events{cluster="local",id="local"} 7
```
