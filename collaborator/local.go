package collaborator

import (
	"github.com/zhenyiya/artifacts/iworkable"
	"github.com/zhenyiya/artifacts/message"
	"github.com/zhenyiya/artifacts/task"
	"github.com/zhenyiya/logger"
	"github.com/zhenyiya/wrappers/messageHelper"
)

type LocalMethods struct {
	workable iworkable.Workable
}

func NewLocalMethods(wk iworkable.Workable) *LocalMethods {
	return &LocalMethods{wk}
}

func (l *LocalMethods) Exchange(in *message.CardMessage, out *message.CardMessage) error {
	logger.LogNormal("Card message from another Collaborator received")
	err := messageHelper.Exchange(in, out)
	return err
}

func (l *LocalMethods) DistributeSync(source *map[int]*task.Task, result *map[int]*task.Task) error {
	logger.LogNormal("Task from another Collaborator received")
	s := *source
	err := l.workable.DoneMulti(s)
	*result = s
	return err
}
