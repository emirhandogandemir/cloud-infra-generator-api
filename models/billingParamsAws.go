package models

type BillingParamsAws struct {
	Start       string `form:"start"`
	End         string `form:"end"`
	Region string `json:"region"`
}
