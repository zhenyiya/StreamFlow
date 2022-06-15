package remoteshared

import (
	"github.com/zhenyiya/server/task"
)

type Collaboratable interface {
	SyncDistribute(sources []*task.Task) ([]*task.Task, error)
}
