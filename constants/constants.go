package constants

import (
	"errors"
	"time"
)

// system vars
const (
	CollaboratorModeAbbr   = "clbt"
	CollaboratorMode       = "collaborator"
	CoordinatorModeAbbr    = "cdnt"
	CoordinatorMode        = "coordinator"
	DebugInactivated       = false
	DebugActivated         = true
	DefaultListenPort      = 8080
	DefaultContactBookPath = "./contact.json"
	DefaultLogPath         = "./history.log"
	DefaultDataStorePath   = "./streamflow.dat"
	DefaultLogPrefix       = "zhenyiya:"
	CleanHistory           = true
	DefaultNotCleanHistory = false
	Mapper                 = "Mapper"
	Reducer                = "Reducer"
	Function               = "Function"
	Shared                 = "Shared"
	Local                  = "Local"
	ProjectPath            = "ProjectPath"
)

// master/worker setting
const (
	DefaultWorkerPerMaster = 10
	DefaultHost            = "localhost"
)

var (
	DefaultReadTimeout         = 15 * time.Second
	DefaultPeriodShort         = 500 * time.Millisecond
	DefaultPeriodLong          = 2000 * time.Millisecond
	DefaultPeriodRoutineDay    = 24 * time.Hour
	DefaultPeriodRoutineWeek   = 7 * 24 * time.Hour
	DefaultPeriodRoutine30Days = 30 * DefaultPeriodRoutineDay
	DefaultPeriodPermanent     = 0 * time.Second
	DefaultTaskExpireTime      = 30 * time.Second
)

// communication types
const (
	ArgTypeInteger              = "integer"
	ArgTypeNumber               = "number"
	ArgTypeString               = "string"
	ArgTypeObject               = "object"
	ArgTypeBoolean              = "boolean"
	ArgTypeNull                 = "null"
	ArgTypeArray                = "array"
	ConstraintTypeMax           = "maximum"
	ConstraintTypeMin           = "minimum"
	ConstraintTypeXMin          = "exclusiveMinimum"
	ConstraintTypeXMax          = "exclusiveMaximum"
	ConstraintTypeUniqueItems   = "uniqueItems"
	ConstraintTypeMaxProperties = "maxProperties"
	ConstraintTypeMinProperties = "minProperties"
	ConstraintTypeMaxLength     = "maxLength"
	ConstraintTypeMinLength     = "minLength"
	ConstraintTypePattern       = "pattern"
	ConstraintTypeMaxItems      = "maxItems"
	ConstraintTypeMinItems      = "minItems"
	ConstraintTypeEnum          = "enum" // value of interface{} should accept a slice
	ConstraintTypeAllOf         = "allOf"
	ConstraintTypeAnyOf         = "anyOf"
	ConstraintTypeOneOf         = "oneOf"
)

// errors
var (
	ErrUnknownCmdArg      = errors.New("zhenyiya: unknown commandline argument, please enter -h to check out")
	ErrConnectionClosed   = errors.New("zhenyiya: connection closed")
	ErrUnknown            = errors.New("zhenyiya: unknown error")
	ErrAPIError           = errors.New("zhenyiya: api error")
	ErrNoCollaborator     = errors.New("zhenyiya: collaborator does not exist")
	ErrCollaboratorExists = errors.New("zhenyiya: collaborator already exists")
	ErrNoService          = errors.New("zhenyiya: service of id does not exist")
	ErrConflictService    = errors.New("zhenyiya: found conflict, service of id already exists")
	ErrNoRegister         = errors.New("zhenyiya: register does not exist")
	ErrConflictRegister   = errors.New("zhenyiya: found conflict, provider of the service already exists")
	ErrNoSubscriber       = errors.New("zhenyiya: subscriber does not exist")
	ErrConflictSubscriber = errors.New("zhenyiya: found conflict, subscriber of the service already exists")
	ErrTimeout            = errors.New("zhenyiya: task timeout error")
	ErrNoPeers            = errors.New("zhenyiya: no peer appears in the contact book")
	ErrFunctNotExist      = errors.New("zhenyiya: no such function found in store")
)

type Header struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// header
var (
	Header200OK        = Header{"200", "OK"}
	Header201Created   = Header{"201", "Created"}
	Header202Accepted  = Header{"202", "Accepted"}
	Header204NoContent = Header{"204", "NoContent"}
	Header403Forbidden = Header{"403", "Forbidden"}
	Header404NotFound  = Header{"404", "NotFound"}
	Header409Conflict  = Header{"409", "Conflict"}
)

// restful
const (
	JSONAPIVersion = `{"version":"1.0"}`
)

var (
	ProjectDir = "github.com/zhenyiya/"
)
