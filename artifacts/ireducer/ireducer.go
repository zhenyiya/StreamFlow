package ireducer

import (
	"github.com/zhenyiya/StreamFlow/artifacts/task"
)

type IReducer interface {
	Reduce(sources map[int]*task.Task) (map[int]*task.Task, error)
}

type DefaultReducer struct {
}

func Default() *DefaultReducer {
	return new(DefaultReducer)
}

func (rd *DefaultReducer) Reduce(sources map[int]*task.Task) (map[int]*task.Task, error) {
	return map[int]*task.Task{}, nil
}
