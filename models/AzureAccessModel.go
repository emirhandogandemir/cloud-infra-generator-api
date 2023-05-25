package models

type AzureAccessModel struct {
	SubscriptionId string `json:"subscriptionId"`
	TenantId string `json:"tenantId"`
	ClientID string `json:"clientID"`
	ClientSecret string `json:"clientSecret"`
}
