package task

import (
	"github.com/zhenyiya/utils"
)

type TaskContext struct {
	Context map[string]interface{}
}

func NewTaskContext(ctx interface{}) *TaskContext {
	maps := utils.Map(ctx)
	return &TaskContext{maps}
}

func (this *TaskContext) Entries() map[string]interface{} {
	return this.Context
}
