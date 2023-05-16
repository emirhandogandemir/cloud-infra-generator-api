package models

type VirtualMachineAzure struct {
 ResourcegroupName string `json:"clusterName"`
 VmName string `json:"clusterName"`
 VmSize string `json:"vmSize"`
 AdminUsername string `json:"adminUsername"`
 AdminPassword string `json:"adminPassword""`
 Location string `json:"location"`
}