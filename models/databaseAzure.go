package models

type DatabaseAzure struct {
	ResourceGroup string `json:"resourceGroup"`
	Location string `json:"location"`
	ServerName string`json:"serverName"`
	AdminName string `json:"adminName"`
	AdminPassword string `json:"adminPassword"`
	Version string `json:"version"`
}
