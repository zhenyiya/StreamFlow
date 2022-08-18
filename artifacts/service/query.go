package service

import (
	"github.com/zhenyiya/artifacts/card"
)

type Query struct {
	Agent card.Card `json:"card"`
}
