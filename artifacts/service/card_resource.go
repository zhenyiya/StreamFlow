package service

// import (
// 	. "github.com/zhenyiya/StreamFlow/artifacts/card"
// 	. "github.com/zhenyiya/StreamFlow/artifacts/resources"
// )

// type CardResource struct {
// 	*Resource
// 	Attributes *Card `json:"attributes"`
// }

// func NewHeartbeatResource(heartbeat *Card) *CardResource {
// 	return &CardResource{
// 		&Resource{
// 			Id:            heartbeat.GetFullIP(),
// 			Type:          "heartbeat",
// 			Relationships: map[string]*Relationship{},
// 		},
// 		heartbeat,
// 	}
// }

// func NewQueryResource(query *Card) *CardResource {
// 	return &CardResource{
// 		&Resource{
// 			Id:            query.GetFullIP(),
// 			Type:          "query",
// 			Relationships: map[string]*Relationship{},
// 		},
// 		query,
// 	}
// }

// func NewRegistryResource(registry *Card) *CardResource {
// 	return &CardResource{
// 		&Resource{
// 			Id:            registry.GetFullEndPoint(),
// 			Type:          "registry",
// 			Relationships: map[string]*Relationship{},
// 		},
// 		registry,
// 	}
// }
