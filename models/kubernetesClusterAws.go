package models

type KubernetesParamsAws struct {
	ClusterName     string `json:"clusterName"`
	ClusterVersion string `json:"clusterVersion"`
}