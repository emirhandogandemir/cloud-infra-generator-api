package models

type AzureAccessModel struct {
	ID         uint32    `gorm:"primary_key;auto_increment" json:"id"`
	SubscriptionId string `json:"subscriptionId"`
	TenantId string `json:"tenantId"`
	ClientID string `json:"clientID"`
	ClientSecret string `json:"clientSecret"`
}
