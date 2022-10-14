package messageHelper

import (
	"github.com/zhenyiya/StreamFlow/artifacts/iremote"
	"github.com/zhenyiya/StreamFlow/artifacts/message"
	"github.com/zhenyiya/StreamFlow/store"
)

func Exchange(in *message.CardMessage, out *message.CardMessage) error {
	future := message.NewCardMessageFuture(in)
	// push future into message chan
	store.GetMsgChan() <- future
	// wait until future is consumed
	*out = *(future.Done())
	return future.Error()
}

func Compare(a iremote.IDigest, b iremote.IDigest) bool {
	if a.TimeStamp() < b.TimeStamp() {
		return true
	}
	return false
}

func Merge(a iremote.IDigest, b iremote.IDigest) (c iremote.IDigest) {
	if a.TimeStamp() < b.TimeStamp() {
		a.SetCards(b.Cards())
		a.SetTimeStamp(b.TimeStamp())
		return a
	}
	b.SetCards(a.Cards())
	b.SetTimeStamp(a.TimeStamp())
	return b
}
