package main

import (
	"encoding/json"
)

type Cluster struct {
	Cluster []struct {
		ID    string `json:"id"`
		Name  string `json:"name"`
		State string `json:"state"`
	} `json:"data"`
}

type Namespace struct {
	Namespace []struct {
		ID string `json:"id"`
	} `json:"data"`
}

type Nodes struct {
	Nodes []struct {
		Hostname      string     `json:"nodeName"`
		Master        bool       `json:"controlPlane"`
		Worker        bool       `json:"worker"`
		State         string     `json:"state"`
		Unschedulable bool       `json:"unschedulable"`
		PodsInfo      *PodsCount `json:"requested"`
	} `json:"data"`
}

type PodsCount struct {
	Pods string `json:"pods"`
}

type Pods struct {
	Pods []struct {
		PodsInfo   *Metadata      `json:"metadata"`
		PodsStatus *Status        `json:"status"`
		Spec       *Specification `json:"spec"`
	} `json:"data"`
}

type Metadata struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
}

type Status struct {
	State    string        `json:"phase"`
	ContInfo []*Containers `json:"containerStatuses"`
}

type Specification struct {
	NodeName string `json:"nodeName"`
}

type Containers struct {
	Name         string `json:"name"`
	State        *Raw   `json:"state"`
	RestartCount int    `json:"restartCount"`
}
type Raw = json.RawMessage

type Events struct {
	Events []struct {
		Type      string     `json:"_type"`
		Count     int        `json:"count"`
		Namespace *EventInfo `json:"involvedObject"`
	} `json:"data"`
}

type EventInfo struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
}
