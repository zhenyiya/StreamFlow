package service

import (
	"github.com/zhenyiya/StreamFlow/artifacts/restful"
)

type QueryResource struct {
	*restful.Resource
	Attributes Query `json:"attributes"`
}

type QueryPayload struct {
	Data     []QueryResource `json:"data"`
	Included []QueryResource `json:"included,omitempty"`
	Links    restful.Links   `json:"links,omitempty"`
}
