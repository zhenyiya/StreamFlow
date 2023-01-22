package collaborator

import (
	"encoding/json"
	"github.com/zhenyiya/StreamFlow/artifacts/card"
	"github.com/zhenyiya/StreamFlow/artifacts/digest"
	"github.com/zhenyiya/StreamFlow/artifacts/iremote"
	"github.com/zhenyiya/StreamFlow/artifacts/message"
	"github.com/zhenyiya/StreamFlow/constants"
	"github.com/zhenyiya/StreamFlow/helpers/messageHelper"
	"github.com/zhenyiya/StreamFlow/logger"
	"github.com/zhenyiya/StreamFlow/store"
	"io/ioutil"
	"os"
	"sync"
	"time"
)

var mu = sync.Mutex{}

type Case struct {
	CaseID string `json:"caseid,omitempty"`
	*Exposed
	*Reserved
}

type Exposed struct {
	Cards     map[string]*card.Card `json:"cards,omitempty"`
	TimeStamp int64                 `json:"timestamp,omitempty"`
}

type Reserved struct {
	// local is the Card of localhost
	Local       card.Card `json:"local,omitempty"`
	Coordinator card.Card `json:"coordinator,omitempty"`
}

func (c *Case) readStream() error {
	bytes, err := ioutil.ReadFile(constants.DEFAULT_CASE_PATH)
	if err != nil {
		panic(err)
	}
	// unmarshal, overwrite default if already existed in config file
	if err := json.Unmarshal(bytes, &c); err != nil {
		logger.LogError(err.Error())
		return err
	}
	return nil
}

func (c *Case) writeStream() error {
	mu.Lock()
	defer mu.Unlock()
	mal, err := json.Marshal(&c)
	err = ioutil.WriteFile(constants.DEFAULT_CASE_PATH, mal, os.ModeExclusive)
	return err
}

func (c *Case) Action() {
	go func() {
		for {
			select {
			case future := <-store.GetMsgChan():
				defer future.Close()
				out, err := c.HandleMessage(future.Receive())
				if err != nil {
					logger.LogError(err)
				}
				future.Return(out)
			default:
				continue
			}
		}
	}()
}

func (c *Case) Stamp() *Case {
	c.TimeStamp = time.Now().Unix()
	return c
}

func (c *Case) GetCluster() string {
	return c.CaseID
}

func (c *Case) GetDigest() iremote.IDigest {
	return &digest.Digest{c.Cards, c.TimeStamp}
}

func (c *Case) Update(dgst iremote.IDigest) {
	c.Cards = dgst.GetCards()
	c.TimeStamp = dgst.GetTimeStamp()
}

func (c *Case) Terminate(key string) *Case {
	mu.Lock()
	defer mu.Unlock()
	delete(c.Cards, key)
	return c
}

func (c *Case) ReturnByPos(pos int) *card.Card {
	mu.Lock()
	defer mu.Unlock()
	if l := len(c.Cards); pos > l {
		pos = pos % l
	}
	counter := 0
	for _, a := range c.Cards {
		if counter == pos {
			return a
		}
		counter++
	}
	return &card.Card{}
}

func (c *Case) HandleMessage(in *message.CardMessage) (*message.CardMessage, error) {
	// return if message is wrongly sent
	var (
		out *message.CardMessage = message.NewCardMessage()
		err error                = nil
	)

	if err = c.Validate(in, out); err != nil {
		return out, err
	}
	var (
		// local digest
		ldgst = c.GetDigest()
		// remote digest
		rdgst = in.GetDigest()
		// feedback digest
		fbdgst = ldgst
	)
	switch in.GetType() {
	case message.CardMessage_SYNC:
		// msg has a more recent timestamp
		if messageHelper.Compare(ldgst, rdgst) {
			fbdgst = messageHelper.Merge(ldgst, rdgst)
			// update digest to local
			c.Update(fbdgst)
		}
		// update digest to feedback
		out.Update(fbdgst)

		// return ack message
		out.SetType(message.CardMessage_ACK)
		out.SetStatus(constants.GOSSIP_HEADER_OK)
	case message.CardMessage_ACK:
		// msg has a more recent timestamp
		if messageHelper.Compare(ldgst, rdgst) {
			fbdgst = messageHelper.Merge(ldgst, rdgst)
			// update digest to local
			c.Update(fbdgst)
		}
		// return ack message
		out.SetType(message.CardMessage_ACK2)
		out.SetStatus(constants.GOSSIP_HEADER_OK)
	case message.CardMessage_ACK2:
		// return ack message
		out.SetType(message.CardMessage_ACK3)
		out.SetStatus(constants.GOSSIP_HEADER_OK)
	case message.CardMessage_ACK3:
		// do nothing
	default:
		out.SetStatus(constants.GOSSIP_HEADER_UNKNOWN_MSG_TYPE)
		err = constants.ERR_UNKNOWN_MSG_TYPE
	}
	out.SetTo(in.GetFrom())
	out.SetFrom(in.GetTo())
	out.SetCluster(c.GetCluster())
	return out, nil
}

func (c *Case) Validate(in *message.CardMessage, out *message.CardMessage) error {
	if c.GetCluster() != in.GetCluster() {
		out.SetStatus(constants.GOSSIP_HEADER_CASE_MISMATCH)
		return constants.ERR_CASE_MISMATCH
	}
	if to := in.GetTo(); !c.Local.IsEqualTo(to) {
		logger.LogError(c)
		out.SetStatus(constants.GOSSIP_HEADER_COLLABORATOR_MISMATCH)
		return constants.ERR_COLLABORATOR_MISMATCH
	}
	return nil
}
