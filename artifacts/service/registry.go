package service

import (
	"github.com/zhenyiya/artifacts/card"
)

type Registry struct {
	Cards []card.Card `json:"cards"`
}
