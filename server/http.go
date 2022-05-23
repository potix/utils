package server

import (
        "fmt"
        "log"
        "context"
        "time"
        "net/http"
	"crypto/tls"
        "github.com/gin-gonic/gin"
)

type httpServerOptions struct {
        verbose           bool
        mode              string
        tlsCertPath       string
        tlsKeyPath        string
	readHeaderTimeout time.Duration
	readTimeout       time.Duration
	writeTimeout      time.Duration
	idleTimeout       time.Duration
	shutdownTimeout   time.Duration
}

func httpServerDefaultOptions() *httpServerOptions {
        return &httpServerOptions{
                mode:              gin.ReleaseMode,
                verbose:           false,
                tlsCertPath:       "",
                tlsKeyPath:        "",
		readHeaderTimeout: 5 * time.Second,
		readTimeout:       10 * time.Second,
		writeTimeout:      10 * time.Second,
		idleTimeout:       20 * time.Second,
		shutdownTimeout:   60 * time.Second,
        }
}

type HttpServerOption func(*httpServerOptions)

func HttpServerVerbose(verbose bool) HttpServerOption {
        return func(opts *httpServerOptions) {
                opts.verbose = verbose
        }
}

func HttpServerMode(mode  string) HttpServerOption {
        return func(opts *httpServerOptions) {
                opts.mode = mode
        }
}

func HttpServerTls(tlsCertPath string, tlsKeyPath string) HttpServerOption {
        return func(opts *httpServerOptions) {
                opts.tlsCertPath = tlsCertPath
                opts.tlsKeyPath = tlsKeyPath
        }
}

func HttpServerReadHeaderTimeout(readHeaderTimeout time.Duration) HttpServerOption {
        return func(opts *httpServerOptions) {
                opts.readHeaderTimeout = readHeaderTimeout
        }
}

func HttpServerReadTimeout(readTimeout time.Duration) HttpServerOption {
        return func(opts *httpServerOptions) {
                opts.readTimeout = readTimeout
        }
}

func HttpServerWriteTimeout(writeTimeout time.Duration) HttpServerOption {
        return func(opts *httpServerOptions) {
                opts.writeTimeout = writeTimeout
        }
}

func HttpServerIdleTimeout(idleTimeout time.Duration) HttpServerOption {
        return func(opts *httpServerOptions) {
                opts.idleTimeout = idleTimeout
        }
}

func HttpServerShutdownTimeout(shutdownTimeout time.Duration) HttpServerOption {
        return func(opts *httpServerOptions) {
                opts.shutdownTimeout = shutdownTimeout
        }
}

type HttpHandler interface {
        Start() (error)
        Stop()
        SetRouting(*gin.Engine)
}

type HttpServer struct {
	addrPort string
	opts     *httpServerOptions
        server   *http.Server
        router   *gin.Engine
	handler  HttpHandler
}

func (s *HttpServer) Start() (error){
	if err := s.handler.Start(); err != nil {
		return fmt.Errorf("can not start handelr: %w", err)
	}
        go func() {
		if s.opts.tlsCertPath != "" && s.opts.tlsKeyPath != "" {
			err := s.server.ListenAndServeTLS(s.opts.tlsCertPath, s.opts.tlsKeyPath);
			if err != nil && err != http.ErrServerClosed {
				log.Fatalf("listen: %v", err)
			}
		} else {
			err := s.server.ListenAndServe();
			if err != nil && err != http.ErrServerClosed {
				log.Fatalf("listen: %v", err)
			}
		}
        }()
	return nil
}

func (s *HttpServer) Stop() {
        ctx, cancel := context.WithTimeout(context.Background(), s.opts.shutdownTimeout)
        defer cancel()
        err := s.server.Shutdown(ctx)
        if err != nil {
            log.Printf("Server Shutdown: %v", err)
        }
	s.handler.Stop()
}

func NewHttpServer(addrPort string, handler HttpHandler, opts ...HttpServerOption) (*HttpServer, error) {
	baseOpts := httpServerDefaultOptions()
        for _, opt := range opts {
		if opt == nil {
			continue
		}
                opt(baseOpts)
        }
	gin.SetMode(baseOpts.mode)
        router := gin.Default()
	handler.SetRouting(router)
        s := &http.Server{
		Addr: addrPort,
		Handler: router,
		ReadHeaderTimeout: baseOpts.readHeaderTimeout,
		ReadTimeout: baseOpts.readTimeout,
		WriteTimeout: baseOpts.writeTimeout,
		IdleTimeout : baseOpts.idleTimeout,
        }
	if baseOpts.tlsCertPath != "" && baseOpts.tlsKeyPath != "" {
		s.TLSConfig = &tls.Config{
			MinVersion: tls.VersionTLS12,
		}
	}
        return &HttpServer {
		addrPort: addrPort,
		opts: baseOpts,
		server: s,
		router: router,
		handler: handler,
        }, nil
}
