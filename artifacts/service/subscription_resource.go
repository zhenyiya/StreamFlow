package service

import (
	. "github.com/zhenyiya/StreamFlow/artifacts/resources"
)

type SubscriptionResource struct {
	*Resource
}

func NewSubscriptionResource(token string) *SubscriptionResource {
	return &SubscriptionResource{
		&Resource{
			Id:            token,
			Type:          "subscription",
			Relationships: map[string]*Relationship{},
		},
	}
}
