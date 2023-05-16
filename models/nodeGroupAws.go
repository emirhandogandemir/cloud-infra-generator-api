package models

type NodeGroup struct {
	ClusterName   string `json:"clusterName"`
	KeyName string `json:"KeyName"`
	NodegroupName      string `json:"nodegroupName"`
	DesiredState int32 `json:"desiredState"`
	MinSize int32 `json:"minSize"`
	MaxSize int32 `json:"maxSize"`
}