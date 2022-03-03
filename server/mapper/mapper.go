package mapper

import (
	"github.com/zhenyiya/server/task"
)

type Mapper interface {
	Map(t task.Task) (map[int64]task.Task, error)
}
