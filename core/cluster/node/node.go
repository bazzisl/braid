package node

import "braid/core/actor"

/*
	init - 初始化进程
	update - 将一堆执行线程丢到node的运行时驱动
	close - 监听退出信号，通知到各执行线程停止接受新处理，等待当前处理结束退出
*/

type INode interface {
	Init(...Option) error
	Update(...actor.IActor)
	WaitClose()

	ID() string
	Name() string
}