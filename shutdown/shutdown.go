package shutdown

import (
	"os"
	"os/signal"
	"syscall"
)

/**
 * @Author: BoolDesign
 * @Email: booldesign@163.com
 * @Date: 2021/1/27 15:44
 * @Desc: 优雅关闭
 */

// Hook a graceful shutdown hook, default with signals of SIGINT and SIGTERM
type Hook interface {
	WithSignals(signals ...syscall.Signal) Hook // WithSignals add more signals into hook
	Close(funcs ...func())                      // Close register shutdown handles
}

var _ Hook = (*hook)(nil)

type hook struct {
	singleChan chan os.Signal
}

func NewHook() Hook {
	hook := &hook{
		singleChan: make(chan os.Signal, 1),
	}

	return hook.WithSignals(syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)
}

func (h *hook) WithSignals(signals ...syscall.Signal) Hook {
	for _, s := range signals {
		signal.Notify(h.singleChan, s)
	}

	return h
}

func (h *hook) Close(funcs ...func()) {
	select {
	case <-h.singleChan:
	}
	signal.Stop(h.singleChan)

	for _, f := range funcs {
		f()
	}
}
