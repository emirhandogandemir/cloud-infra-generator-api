package models


type KubernetesParamsAzure struct {
	ResourceGroupName string `json:"resourceGroupName"`
	ClusterName string `json:"clusterName"`
	Location string `json:"location"`
	VmSize string `json:"vmSize"`
	KubernetesVersion string `json:"kubernetesVersion"`
	DnsPrefixName string `json:"dnsPrefixName"`
}