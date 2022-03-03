package reducer

import (
	"github.com/zhenyiya/server/task"
)

type Reducer interface {
	Reduce(sources map[int64]task.Task, result task.Task) error
}
