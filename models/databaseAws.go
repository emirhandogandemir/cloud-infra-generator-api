package models

type DatabaseAws struct {
	DbInstanceIdentifier string `json:"dbInstanceIdentifier"`
	DbInstanceClass string `json:"dbInstanceClass"`
	Engine string `json:"engine"`
	MasterUsername string `json:"masterUsername"`
	MasterPassword string `json:"masterPassword"`
	DbName string `json:"DbName"`
}