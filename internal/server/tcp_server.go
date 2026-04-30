package server

import (
	"context"
	"fmt"
	"log/slog"
	"net"
	"sync"
)

type TCPServer struct {
	addr     string
	listener net.Listener
	wg       sync.WaitGroup
	logger   *slog.Logger
}

func New(addr string, logger *slog.Logger) *TCPServer {
	return &TCPServer{
		addr:   addr,
		logger: logger,
	}
}

func (s *TCPServer) Start(ctx context.Context) error {
	ln, err := net.Listen("tcp", s.addr)
	if err != nil {
		return fmt.Errorf("failed to listen on %s: %w", s.addr, err)
	}
	s.listener = ln
	s.logger.Info("server listening", "addr", s.addr)

	// Shutdown when context is cancelled
	go func() {
		<-ctx.Done()
		s.logger.Info("shutting down server")
		err := s.listener.Close()
		if err != nil {
			return
		}
	}()

	for {
		conn, err := ln.Accept()
		if err != nil {
			// Check if we shut down intentionally
			select {
			case <-ctx.Done():
				s.wg.Wait()
				return nil
			default:
				s.logger.Error("accept error", "err", err)
				continue
			}
		}

		s.wg.Add(1)
		go func() {
			defer s.wg.Done()
			handleConnection(conn, s.logger)
		}()
	}
}
