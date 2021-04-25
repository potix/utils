package server

import (
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
	server, err :=  NewHttpServer("127.0.0.1:12345", &handler{})
	if err != nil {
		t.Fatalf("can not create server")
	}
	if err := server.Start(); err != nil {
		t.Fatalf("can not start server")
	}
	server.Stop()

}
