package remoteshared

import (
	"encoding/json"
	"github.com/zhenyiya/constants"
	"github.com/zhenyiya/logger"
	"github.com/zhenyiya/restful"
	"github.com/zhenyiya/utils"
	"strconv"
)

// agent is the network config of server
type Agent struct {
	IP    string `json:"ip"`
	Port  int    `json:"port"`
	Alive bool   `json:"alive"`
	API   string `json:"api,omitempty"`
}

type Parameter struct {
	Type        string       `json:"type"`
	Description string       `json:"description"`
	Constraints []Constraint `json:"constraints"`
	Required    bool         `json:"required"`
}

type Constraint struct {
	Key   string      `json:"key"`
	Value interface{} `json:"value"`
}

func (par *Parameter) SerializeToJSON() string {
	mal, err := json.Marshal(*par)
	if err != nil {
		logger.LogWarning("Internal error during serialization: " + err.Error())
	}
	return string(mal)
}

func (par *Parameter) Validate(dat *restful.Resource) bool {
	if par.Required && dat == nil {
		return false
	}
	if par.Type != dat.Type {
		return false
	}
	switch par.Type {
	case constants.ArgTypeInteger, constants.ArgTypeNumber:
		for _, cons := range par.Constraints {
			switch cons.Key {
			case constants.ConstraintTypeMax:
			case constants.ConstraintTypeMin:
			case constants.ConstraintTypeEnum:
			default:
				return false
			}
		}
	case constants.ArgTypeString:
		for _, cons := range par.Constraints {
			switch cons.Key {
			case constants.ConstraintTypeMaxLength:
			case constants.ConstraintTypeMinLength:
			case constants.ConstraintTypePattern:
			case constants.ConstraintTypeEnum:
			default:
				return false
			}
		}
	case constants.ArgTypeArray:
		for _, cons := range par.Constraints {
			switch cons.Key {
			case constants.ConstraintTypeMaxItems:
			case constants.ConstraintTypeMinItems:
			case constants.ConstraintTypeUniqueItems:
			case constants.ConstraintTypeAllOf:
			case constants.ConstraintTypeOneOf:
			case constants.ConstraintTypeAnyOf:
			default:
				return false
			}
		}
	case constants.ArgTypeObject:
		for _, cons := range par.Constraints {
			switch cons.Key {
			case constants.ConstraintTypeMaxProperties:
			case constants.ConstraintTypeMinProperties:
			default:
				return false
			}
		}
	default:
		return false
	}
	return true
}

func UnmarshalAgents(original []interface{}) []Agent {
	var agents []Agent
	for _, o := range original {
		oo := o.(map[string]interface{})
		if oo["api"] != nil {
			agents = append(agents, Agent{oo["ip"].(string), int(oo["port"].(float64)), oo["alive"].(bool), oo["api"].(string)})
			continue
		}
		agents = append(agents, Agent{oo["ip"].(string), int(oo["port"].(float64)), oo["alive"].(bool), ""})
	}
	return agents
}

func (c *Agent) GetFullIP() string {
	return c.IP + ":" + strconv.Itoa(c.Port)
}

func (c *Agent) GetFullExposureAddress() string {
	return utils.MapToExposureAddress(c.IP) + ":" + strconv.Itoa(c.Port)
}

func (c *Agent) GetFullEndPoint() string {
	return c.IP + ":" + strconv.Itoa(c.Port) + "/" + c.API
}

func (c *Agent) IsEqualTo(another *Agent) bool {
	return c.GetFullIP() == another.GetFullIP() || c.GetFullExposureAddress() == another.GetFullExposureAddress()
}
