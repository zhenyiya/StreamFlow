package service

import (
	"github.com/zhenyiya/StreamFlow/artifacts/restful"
)

type SubscriptionResource struct {
	*restful.Resource
	Attributes Subscription `json:"attributes"`
}

type SubscriptionPayload struct {
	Data     []SubscriptionResource `json:"data"`
	Included []SubscriptionResource `json:"included,omitempty"`
	Links    restful.Links          `json:"links,omitempty"`
}
