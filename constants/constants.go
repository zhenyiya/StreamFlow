package constants

import (
	"errors"
	"os"
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
	DefaultCasePath        = "./case.json"
	DefaultLogPath         = "./history.log"
	DefaultDataStorePath   = "./streamflow.dat"
	DefaultLogPrefix       = ""
	CleanHistory           = true
	DefaultNotCleanHistory = false
	Mapper                 = "Mapper"
	Reducer                = "Reducer"
	Executor               = "Executor"
	Function               = "Function"
	HashFunction           = "HashFunction"
	Shared                 = "Shared"
	Local                  = "Local"
	Limit                  = "Limit"
	ProjectPath            = "ProjectPath"
	ProjectUnixPath        = "ProjectUnixPath"
	DefaultWorkerPerMaster = 10
	DefaultHost            = "localhost"
	DefaultGossipNum       = 5
	DefaultCaseID          = "zhenyiyaStandardCase"
	DefaultJobRequestBurst = 1000
)

// store setting
const (
	DefaultHashLength = 12
)

// time consts
var (
	DefaultReadTimeout                = 15 * time.Second
	DefaultPeriodShort                = 500 * time.Millisecond
	DefaultPeriodLong                 = 2000 * time.Millisecond
	DefaultPeriodRoutineDay           = 24 * time.Hour
	DefaultPeriodRoutineWeek          = 7 * 24 * time.Hour
	DefaultPeriodRoutine30Days        = 30 * DefaultPeriodRoutineDay
	DefaultPeriodPermanent            = 0 * time.Second
	DefaultTaskExpireTime             = 30 * time.Second
	DefaultGCInterval                 = 30 * time.Second
	DefaultMaxMappingTime             = 600 * time.Second
	DefaultSyncInterval               = 3 * time.Minute
	DefaultHeartbeatInterval          = 5 * time.Second
	DefaultJobRequestRefillInterval   = 1 * time.Millisecond
	DefaultStatFlushInterval          = 20 * time.Millisecond
	DefaultStatAbstractInterval       = 3 * time.Second
	DefaultCollaboratorExpiryInterval = 10 * time.Minute
)

// executor types
const (
	ExecutorTypeMapper   = "mapper"
	ExecutorTypeReducer  = "reducer"
	ExecutorTypeShuffler = "shuffler"
	ExecutorTypeCombiner = "combiner"
	ExecutorTypeDefault  = "default"
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

const (
	StatsPolicySumOfInt = "StatsPolicySumOfInt"
)

// errors
var (
	ErrUnknownCmdArg                   = errors.New("zhenyiya: unknown commandline argument, please enter -h to check out")
	ErrConnectionClosed                = errors.New("zhenyiya: connection closed")
	ErrUnknown                         = errors.New("zhenyiya: unknown error")
	ErrAPIError                        = errors.New("zhenyiya: api error")
	ErrNoCollaborator                  = errors.New("zhenyiya: collaborator does not exist")
	ErrCollaboratorExists              = errors.New("zhenyiya: collaborator already exists")
	ErrNoService                       = errors.New("zhenyiya: service of id does not exist")
	ErrConflictService                 = errors.New("zhenyiya: found conflict, service of id already exists")
	ErrNoRegister                      = errors.New("zhenyiya: register does not exist")
	ErrConflictRegister                = errors.New("zhenyiya: found conflict, provider of the service already exists")
	ErrNoSubscriber                    = errors.New("zhenyiya: subscriber does not exist")
	ErrConflictSubscriber              = errors.New("zhenyiya: found conflict, subscriber of the service already exists")
	ErrTimeout                         = errors.New("zhenyiya: task timeout error")
	ErrNoPeers                         = errors.New("zhenyiya: no peer appears in the contact book")
	ErrFunctNotExist                   = errors.New("zhenyiya: no such function found in store")
	ErrJobNotExist                     = errors.New("zhenyiya: no sucn job found in store")
	ErrExecutorNotFound                = errors.New("zhenyiya: no such executor found in store")
	ErrLimiterNotFound                 = errors.New("zhenyiya: no such limiter found in store")
	ErrValNotFound                     = errors.New("zhenyiya: no value found with such key")
	ErrCaseMismatch                    = errors.New("zhenyiya: case mismatch error")
	ErrCollaboratorMismatch            = errors.New("zhenyiya: collaborator mismatch error")
	ErrUnknownMsgType                  = errors.New("zhenyiya: unknown message type error")
	ErrMapTaskFailing                  = errors.New("zhenyiya: map operation failing error")
	ErrReduceTaskFailing               = errors.New("zhenyiya: reduce operation failing error")
	ErrExecutorStackLengthInconsistent = errors.New("zhenyiya: executor stack length inconsistent error")
	ErrMessageChannelDirty             = errors.New("zhenyiya: message channel has unconsumed message error")
	ErrTaskChannelDirty                = errors.New("zhenyiya: task channel has unconsumed task error")
	ErrStatTypeNotFound                = errors.New("zhenyiya: stat type not found error")
	ErrCoordinatorNotFound             = errors.New("zhenyiya: coordinator not found error")
	ErrInputStreamCorrupted            = errors.New("zhenyiya: input stream corrupted error")
	ErrInputStreamNotSupported         = errors.New("zhenyiya: input stream type not suppoted error")
	ErrIODecodePointerRequired         = errors.New("zhenyiya: Decode error, the reference instance must be a pointer")
	ErrIODecodeSliceRequired           = errors.New("zhenyiya: Decode error, the reference instance must be a slice")
	ErrIODecodeStructRequired          = errors.New("zhenyiya: Decode error, the reference instance must be a struct")
)

type Header struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// HTTP Status
// var (
// Header200OK                   = Header{"200", "OK"}
// Header201Created              = Header{"201", "Created"}
// Header202Accepted             = Header{"202", "Accepted"}
// Header204NoContent            = Header{"204", "NoContent"}
// Header403Forbidden            = Header{"403", "Forbidden"}
// Header404NotFound             = Header{"404", "NotFound"}
// Header409Conflict             = Header{"409", "Conflict"}
// Header415UnsupportedMediaType = Header{"415", "UnsupportedMediaType"}
// Header422ExceedLimit          = Header{"422", "ExceedLimit"}
// )

// HTTP headers
var (
	HeaderContentTypeJSON     = Header{"Content-Type", "application/json"}
	HeaderContentTypeText     = Header{"Content-Type", "text/html"}
	HeaderCORSEnableAllOrigin = Header{"Access-Control-Allow-Origin", "*"}
)

// Gossip Protocol headers
var (
	GossipHeaderOK                   = Header{"200", "OK"}
	GossipHeaderUnknownMsgType       = Header{"400", "UnknownMessageType"}
	GossipHeaderCaseMismatch         = Header{"401", "CaseMismatch"}
	GossipHeaderCollaboratorMismatch = Header{"401", "CollaboratorMismatch"}
	GossipHeaderUnknownError         = Header{"500", "UnknownGossipError"}
)

// restful
const (
	JSONAPIVersion = `{"version":"1.0"}`
)

var (
	ProjectDir     = ""
	ProjectUnixDir = ""
	LibDir         = "github.com/zhenyiya/"
	LibUnixDir     = os.Getenv("GOPATH") + "/src/github.com/zhenyiya/"
)
