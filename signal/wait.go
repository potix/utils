package signal

import (
    "os"
    "os/signal"
    "syscall"
)

func SignalWait(hupFunc func()) {
        sigChan := make(chan os.Signal, 10)
        signal.Notify(sigChan,
                syscall.SIGHUP,
                syscall.SIGINT,
                syscall.SIGQUIT,
                syscall.SIGTERM)
        for {
                sig := <-sigChan
                switch sig {
                case syscall.SIGINT:
                        fallthrough
                case syscall.SIGQUIT:
                        fallthrough
                case syscall.SIGTERM:
                        return
		case syscall.SIGHUP:
			if hupFunc != nil {
				hupFunc()
			}
                default:
                }
        }
}
