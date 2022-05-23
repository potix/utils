package server

import (
	"fmt"
	"time"
	"github.com/gin-gonic/gin"
        "testing"
)

type httpHandler struct {
}

func (h *httpHandler) Start() (error) {
	return nil
}

func (h *httpHandler) Stop() {
}

func (h *httpHandler) SetRouting(route *gin.Engine) {
}

func TestHttpRun(t *testing.T) {
	fmt.Printf("--- http test ---\n")
	to := 30 * time.Second
	opt1 := HttpServerVerbose(true)
	opt2 := HttpServerMode("debug")
	opt3 := HttpServerTls("", "")
	opt4 := HttpServerReadHeaderTimeout(to)
	opt5 := HttpServerReadTimeout(to)
	opt6 := HttpServerWriteTimeout(to)
	opt7 := HttpServerIdleTimeout(to)
	opt8 := HttpServerShutdownTimeout(to)
	server, err :=  NewHttpServer(
		"127.0.0.1:12345",
		&httpHandler{},
		opt1,
		opt2,
		opt3,
		opt4,
		opt5,
		opt6,
		opt7,
		opt8,
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
