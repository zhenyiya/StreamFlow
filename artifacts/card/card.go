package card

import (
	"github.com/zhenyiya/utils"
	"strconv"
)

// Card is the network config of server
type Card struct {
	IP    string `json:"ip"`
	Port  int    `json:"port"`
	Alive bool   `json:"alive"`
	API   string `json:"api,omitempty"`
}

func (c *Card) GetFullIP() string {
	return c.IP + ":" + strconv.Itoa(c.Port)
}

func (c *Card) GetFullExposureAddress() string {
	return utils.MapToExposureAddress(c.IP) + ":" + strconv.Itoa(c.Port)
}

func (c *Card) GetFullExposureCard() Card {
	return Card{utils.MapToExposureAddress(c.IP), c.Port, c.Alive, c.API}
}

func (c *Card) GetFullEndPoint() string {
	return c.IP + ":" + strconv.Itoa(c.Port) + "/" + c.API
}

func (c *Card) IsEqualTo(another *Card) bool {
	return c.GetFullIP() == another.GetFullIP() || c.GetFullExposureAddress() == another.GetFullExposureAddress()
}

// current RPC port
func Default() *Card {
	return &Card{utils.GetLocalIP(), utils.GetPort(), true, ""}
}
