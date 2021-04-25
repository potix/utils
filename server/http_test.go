package server

import (
	"time"
	"github.com/gin-gonic/gin"
        "testing"
)

type handler struct {
}

func (h *handler) Start() (error) {
	return nil
}

func (h *handler) Stop() {
}

func (h *handler) SetRouting(route *gin.Engine) {
}

func TestLen(t *testing.T) {
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
		&handler{},
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
