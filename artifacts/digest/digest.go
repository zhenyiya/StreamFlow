package digest

import (
	"github.com/zhenyiya/artifacts/card"
)

type Digest struct {
	Cards_ map[string]card.Card `json:"cards"`
	Ts_    int64                `json:"timestamp"`
}

func (dgst *Digest) Cards() map[string]card.Card {
	return dgst.Cards_
}

func (dgst *Digest) TimeStamp() int64 {
	return dgst.Ts_
}

func (dgst *Digest) SetCards(cards map[string]card.Card) {
	dgst.Cards_ = cards
}

func (dgst *Digest) SetTimeStamp(ts int64) {
	dgst.Ts_ = ts
}
