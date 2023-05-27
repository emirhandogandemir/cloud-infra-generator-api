package models

import "gorm.io/gorm"

type AzureAccessModel struct {
	gorm.Model
	SubscriptionId string `json:"subscriptionId"`
	TenantId string `json:"tenantId"`
	ClientID string `json:"clientID"`
	ClientSecret string `json:"clientSecret"`

}
