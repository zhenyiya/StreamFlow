package iremote

import (
	"github.com/zhenyiya/artifacts/task"
)

type ICollaboratable interface {
	SyncDistribute(sources []*task.Task) ([]*task.Task, error)
}
