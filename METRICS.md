# Example Metrics


```
# HELP clusters
# TYPE rancher_cluster gauge
rancher_cluster{cluster="local",id="local"} 0
# HELP namespaces cluster
# TYPE rancher_cluster_namespace gauge
rancher_cluster_namespace{cluster="local",id="local",namespace="cattle-fleet-clusters-system"} 1
rancher_cluster_namespace{cluster="local",id="local",namespace="cattle-fleet-local-system"} 1
rancher_cluster_namespace{cluster="local",id="local",namespace="cattle-fleet-system"} 1
rancher_cluster_namespace{cluster="local",id="local",namespace="cattle-global-data"} 1
rancher_cluster_namespace{cluster="local",id="local",namespace="cattle-global-nt"} 1
rancher_cluster_namespace{cluster="local",id="local",namespace="cattle-impersonation-system"} 1
rancher_cluster_namespace{cluster="local",id="local",namespace="cattle-system"} 1
rancher_cluster_namespace{cluster="local",id="local",namespace="cluster-fleet-local-local-1a3d67d0a899"} 1
rancher_cluster_namespace{cluster="local",id="local",namespace="default"} 1
rancher_cluster_namespace{cluster="local",id="local",namespace="fleet-default"} 1
rancher_cluster_namespace{cluster="local",id="local",namespace="fleet-local"} 1
rancher_cluster_namespace{cluster="local",id="local",namespace="kube-node-lease"} 1
rancher_cluster_namespace{cluster="local",id="local",namespace="kube-public"} 1
rancher_cluster_namespace{cluster="local",id="local",namespace="kube-system"} 1
rancher_cluster_namespace{cluster="local",id="local",namespace="local"} 1
rancher_cluster_namespace{cluster="local",id="local",namespace="user-bxncv"} 1
# HELP nodes cluster
# TYPE rancher_cluster_node gauge
rancher_cluster_node{cluster="local",id="local",node="local-node",node_role="Master"} 0
rancher_cluster_node_pods_count{cluster="local",id="local",node="local-node"} 5
# HELP pods cluster
# TYPE rancher_cluster_pod gauge
rancher_cluster_pod{cluster="local",id="local",namespace_pod="cattle-fleet-local-system",name="fleet-agent-7bcc7d69cf-6dr9k"} 1
rancher_cluster_pod_container{cluster="local",id="local",namespace_pod="cattle-fleet-local-system",name_pod="fleet-agent-7bcc7d69cf-6dr9k",name_container="fleet-agent"} 1
rancher_cluster_pod_container_count_restart{cluster="local",id="local",namespace_pod="cattle-fleet-local-system",name_pod="fleet-agent-7bcc7d69cf-6dr9k",name_container="fleet-agent"} 0
rancher_cluster_pod{cluster="local",id="local",namespace_pod="cattle-fleet-system",name="fleet-controller-5cbbd7c4c9-rk5ds"} 1
rancher_cluster_pod_container{cluster="local",id="local",namespace_pod="cattle-fleet-system",name_pod="fleet-controller-5cbbd7c4c9-rk5ds",name_container="fleet-controller"} 1
rancher_cluster_pod_container_count_restart{cluster="local",id="local",namespace_pod="cattle-fleet-system",name_pod="fleet-controller-5cbbd7c4c9-rk5ds",name_container="fleet-controller"} 0
rancher_cluster_pod{cluster="local",id="local",namespace_pod="cattle-fleet-system",name="gitjob-5c5979d844-djqrc"} 1
rancher_cluster_pod_container{cluster="local",id="local",namespace_pod="cattle-fleet-system",name_pod="gitjob-5c5979d844-djqrc",name_container="gitjob"} 1
rancher_cluster_pod_container_count_restart{cluster="local",id="local",namespace_pod="cattle-fleet-system",name_pod="gitjob-5c5979d844-djqrc",name_container="gitjob"} 0
rancher_cluster_pod{cluster="local",id="local",namespace_pod="cattle-system",name="rancher-webhook-5898d78956-btbgn"} 1
rancher_cluster_pod_container{cluster="local",id="local",namespace_pod="cattle-system",name_pod="rancher-webhook-5898d78956-btbgn",name_container="rancher-webhook"} 1
rancher_cluster_pod_container_count_restart{cluster="local",id="local",namespace_pod="cattle-system",name_pod="rancher-webhook-5898d78956-btbgn",name_container="rancher-webhook"} 0
rancher_cluster_pod{cluster="local",id="local",namespace_pod="kube-system",name="coredns-b96499967-zclgc"} 1
rancher_cluster_pod_container{cluster="local",id="local",namespace_pod="kube-system",name_pod="coredns-b96499967-zclgc",name_container="coredns"} 1
rancher_cluster_pod_container_count_restart{cluster="local",id="local",namespace_pod="kube-system",name_pod="coredns-b96499967-zclgc",name_container="coredns"} 0
# HELP events cluster
# TYPE rancher_cluster_events gauge
rancher_cluster_events{cluster="local",id="local"} 0
```
