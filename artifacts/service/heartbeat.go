package service

import (
	"github.com/zhenyiya/StreamFlow/artifacts/card"
)

type Heartbeat struct {
	Agent card.Card `json:"card"`
}
