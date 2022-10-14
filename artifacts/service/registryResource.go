package service

import (
	"github.com/zhenyiya/StreamFlow/artifacts/restful"
)

type RegistryResource struct {
	*restful.Resource
	Attributes Registry `json:"attributes"`
}

type RegistryPayload struct {
	Data     []RegistryResource `json:"data"`
	Included []RegistryResource `json:"included,omitempty"`
	Links    restful.Links      `json:"links,omitempty"`
}
