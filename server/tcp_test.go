package server

import (
	"fmt"
	"time"
	"net"
        "testing"
)

type tcpHandler struct {
}

func (h *tcpHandler) Start() (error) {
	return nil
}

func (h *tcpHandler) Stop() {
}

func (h *tcpHandler) OnAccept(conn net.Conn) {
}

func TestTcpRun(t *testing.T) {
	fmt.Printf("--- tcp test ---\n")
	opt1 := TcpServerVerbose(true)
	opt2 := TcpServerTls("", "")
	opt3 := TcpServerSkipVerify(true)
	server, err :=  NewTcpServer(
		"127.0.0.1:12345",
		&tcpHandler{},
		opt1,
		opt2,
		opt3,
	)
	if err != nil {
		t.Fatalf("can not create server")
	}
	if err := server.Start(); err != nil {
		t.Fatalf("can not start server")
	}
	time.Sleep(2 * time.Second)
	server.Stop()
}
