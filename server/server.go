package server

import (
	"fmt"
	"github.com/liuhuanjs/MyDB/config"
	"log"
	"net"
)

//mysql protocol server
type Server struct {
	cfg      *config.Config
	listener net.Listener
}

//run the server
func (s *Server) Run() error {
	for {
		//waiting for pkg send from client
		conn, _ := s.listener.Accept()
		clientConn := s.newConn(conn)
		go s.onConn(clientConn)
	}

	return nil
}

//handles queries in its own goroutine
func (s *Server) onConn(conn *clientConn) {
	log.Println("create new connection remoteAddr 127.x.x.1")
}

//creates a new *clientConn from a net.Conn
func (s *Server) newConn(conn net.Conn) *clientConn {
	return &clientConn{}
}

//create a new server
func NewServer(cfg *config.Config) (*Server, error) {
	s := &Server{
		cfg: cfg,
	}

	var err error
	addr := fmt.Sprintf("%s:%d", s.cfg.Host, s.cfg.Port)
	// init server's listener , listen an address:port
	if s.listener, err = net.Listen("tcp", addr); err == nil {
		log.Printf("server is running MySQL protocol on %s\n", addr)
	}

	return s, nil
}
