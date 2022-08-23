package service

import (
	"github.com/zhenyiya/artifacts/card"
)

type Heartbeat struct {
	Agent card.Card `json:"card"`
}
