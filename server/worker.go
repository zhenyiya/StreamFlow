package server

import (
	"fmt"
	"github.com/zhenyiya/logger"
)

type Element interface {
	Start()
	Quit()
	GetOwner() Master
}

type Worker struct {
	ID          uint64
	Owner       Master
	Alive       bool
	baseTasks   chan Task
	lowTasks    chan Task
	mediumTasks chan Task
	highTasks   chan Task
	urgentTasks chan Task
	quit        chan bool
}

func (w *Worker) Start() {
	go func() {
		for {
			select {
			case <-w.quit:
				return
			case tk := <-w.urgentTasks:
				w.Owner.Logger.LogNormal(fmt.Sprintf("Worker%v:, Task Level:%v", w.ID, tk.Priority))
				logger.LogNormal(fmt.Sprintf("Worker%v:, Task Level:%v", w.ID, tk.Priority))
				tk.Consumable()
			default:
				select {
				case tk := <-w.highTasks:
					w.Owner.Logger.LogNormal(fmt.Sprintf("Worker%v:, Task Level:%v", w.ID, tk.Priority))
					logger.LogNormal(fmt.Sprintf("Worker%v:, Task Level:%v", w.ID, tk.Priority))
					tk.Consumable()
				default:
					select {
					case tk := <-w.mediumTasks:
						w.Owner.Logger.LogNormal(fmt.Sprintf("Worker%v:, Task Level:%v", w.ID, tk.Priority))
						logger.LogNormal(fmt.Sprintf("Worker%v:, Task Level:%v", w.ID, tk.Priority))
						tk.Consumable()
					default:
						select {
						case tk := <-w.lowTasks:
							w.Owner.Logger.LogNormal(fmt.Sprintf("Worker%v:, Task Level:%v", w.ID, tk.Priority))
							logger.LogNormal(fmt.Sprintf("Worker%v:, Task Level:%v", w.ID, tk.Priority))
							tk.Consumable()
						default:
							select {
							case tk := <-w.baseTasks:
								w.Owner.Logger.LogNormal(fmt.Sprintf("Worker%v:, Task Level:%v", w.ID, tk.Priority))
								logger.LogNormal(fmt.Sprintf("Worker%v:, Task Level:%v", w.ID, tk.Priority))
								tk.Consumable()
							default:
								continue
							}
						}
					}
				}
			}
		}
	}()
}

func (w *Worker) GetID() uint64 {
	return w.ID
}

func (w *Worker) Quit() {
	w.quit <- true
}

func (w *Worker) GetOwner() Master {
	return w.Owner
}
