package models

type VirtualMachineAzure struct {
 ResourceGroupName string `json:"resourceGroupName"`
 VmName string `json:"vmName"`
 VmSize string `json:"vmSize"`
 AdminUsername string `json:"adminUsername"`
 AdminPassword string `json:"adminPassword""`
 Location string `json:"location"`
}