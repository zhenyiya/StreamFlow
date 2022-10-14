package service

import (
	"github.com/zhenyiya/StreamFlow/artifacts/card"
)

type Query struct {
	Agent card.Card `json:"card"`
}
