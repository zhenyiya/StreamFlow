package service

import (
	"github.com/zhenyiya/StreamFlow/artifacts/card"
)

type Registry struct {
	Cards []card.Card `json:"cards"`
}
