package iremote

import (
	"github.com/zhenyiya/StreamFlow/artifacts/task"
)

type ICollaboratable interface {
	SyncDistribute(sources []*task.Task) ([]*task.Task, error)
}
