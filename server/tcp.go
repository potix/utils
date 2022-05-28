package server

import (
        "fmt"
        "log"
        "net"
	"crypto/tls"
	"sync"
)

type tcpServerOptions struct {
        verbose     bool
        tlsCertPath string
        tlsKeyPath  string
	skipVerify  bool
}

func tcpServerDefaultOptions() *tcpServerOptions {
        return &tcpServerOptions{
                verbose:     false,
                tlsCertPath: "",
                tlsKeyPath:  "",
                skipVerify:  false,
        }
}

type TcpServerOption func(*tcpServerOptions)

func TcpServerVerbose(verbose bool) TcpServerOption {
        return func(opts *tcpServerOptions) {
                opts.verbose = verbose
        }
}

func TcpServerTls(tlsCertPath string, tlsKeyPath string) TcpServerOption {
        return func(opts *tcpServerOptions) {
                opts.tlsCertPath = tlsCertPath
                opts.tlsKeyPath = tlsKeyPath
        }
}

func TcpServerSkipVerify(skipVerify bool) TcpServerOption {
        return func(opts *tcpServerOptions) {
                opts.skipVerify = skipVerify
        }
}

type TcpHandler interface {
        Start() (error)
        Stop()
        OnAccept(net.Conn)
}

type TcpServer struct {
	addrPort  string
	tlsConfig *tls.Config
	opts      *tcpServerOptions
	handler   TcpHandler
	listen    net.Listener
	wg        *sync.WaitGroup
	stopped   chan int
}

func (s *TcpServer) acceptLoop() {
log.Printf("start accept loop")
	defer s.wg.Done()
	for {
log.Printf("try accept")
		conn, err := s.listen.Accept()
		if err != nil {
			select {
			case <- s.stopped:
log.Printf("accept stop")
				return
			default:
				log.Printf("can not accept: %v ", err)
				continue
			}
		}
log.Printf("accepted")
		s.wg.Add(1)
		go func() {
log.Printf("start on accept")
			s.handler.OnAccept(conn)
log.Printf("end on accept")
			s.wg.Done()
		}()
	}
log.Printf("end accept loop")
}

func (s *TcpServer) Start() (error){
log.Printf("tcp server start")
	if err := s.handler.Start(); err != nil {
		return fmt.Errorf("can not start handelr: %w", err)
	}
	if s.tlsConfig != nil {
		l, err := tls.Listen("tcp", s.addrPort, s.tlsConfig)
		if err != nil {
			return fmt.Errorf("can not listen: %w", err)
		}
		s.listen = l
	} else {
		l, err := net.Listen("tcp", s.addrPort)
		if err != nil {
			return fmt.Errorf("can not listen: %w", err)
		}
		s.listen = l
	}
	s.wg.Add(1)
        go s.acceptLoop()
	return nil
}

func (s *TcpServer) Stop() {
log.Printf("tcp servert stop")
	close(s.stopped)
	s.listen.Close()
log.Printf("try wg stop")
	s.wg.Wait()
log.Printf("done wg stop")
	s.handler.Stop()
}

func NewTcpServer(addrPort string, handler TcpHandler, opts ...TcpServerOption) (*TcpServer, error) {
	baseOpts := tcpServerDefaultOptions()
        for _, opt := range opts {
		if opt == nil {
			continue
		}
                opt(baseOpts)
        }
	var tlsConfig *tls.Config = nil
	if baseOpts.tlsCertPath != "" && baseOpts.tlsKeyPath != "" {
		cert, err := tls.LoadX509KeyPair(baseOpts.tlsCertPath, baseOpts.tlsKeyPath)
		if err != nil {
			return nil, fmt.Errorf("can not load certs: %w", err)
		}
		tlsConfig = &tls.Config{
			MinVersion: tls.VersionTLS12,
			Certificates: []tls.Certificate{cert},
			InsecureSkipVerify: baseOpts.skipVerify,
		}
	}
        return &TcpServer {
		addrPort: addrPort,
		tlsConfig: tlsConfig,
		opts: baseOpts,
		handler: handler,
		listen: nil,
		wg: &sync.WaitGroup{} ,
		stopped: make(chan int),
        }, nil
}
