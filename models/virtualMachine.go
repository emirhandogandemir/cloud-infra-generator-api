package models

type VirtualMachine struct {
	ImageId      string `json:"imageId"`
	InstanceType string `json:"instanceType"`
}