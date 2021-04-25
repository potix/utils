package signal

import (
	"fmt"
	"os"
	"syscall"
	"testing"
	"time"
)

func hupFunc() {
	fmt.Println("hup")
}

func TestWaitINT(t *testing.T) {
	go func() {
		pid := os.Getpid()
		syscall.Kill(pid, syscall.SIGHUP)
		syscall.Kill(pid, syscall.SIGPIPE)
		time.Sleep(2 * time.Second)
		syscall.Kill(pid, syscall.SIGINT)
	}()
	SignalWait(hupFunc)
}

func TestWaitTERM(t *testing.T) {
	go func() {
		pid := os.Getpid()
		syscall.Kill(pid, syscall.SIGUSR1)
		syscall.Kill(pid, syscall.SIGCHLD)
		time.Sleep(2 * time.Second)
		syscall.Kill(pid, syscall.SIGTERM)
	}()
	SignalWait(nil)
}

func TestWaitQUIT(t *testing.T) {
	go func() {
		pid := os.Getpid()
		syscall.Kill(pid, syscall.SIGUSR2)
		syscall.Kill(pid, syscall.SIGALRM)
		time.Sleep(2 * time.Second)
		syscall.Kill(pid, syscall.SIGQUIT)
	}()
	SignalWait(nil)
}
