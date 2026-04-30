package server

import (
	"bufio"
	"log/slog"
	"net"
)

func handleConnection(conn net.Conn, logger *slog.Logger) {
	defer conn.Close()
	logger.Info("new connection", "remote", conn.RemoteAddr())

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		line := scanner.Text()
		logger.Debug("received", "data", line)
		conn.Write([]byte("+OK\r\n"))
	}

	if err := scanner.Err(); err != nil {
		logger.Warn("connection error", "remote", conn.RemoteAddr(), "err", err)
	}
	logger.Info("connection closed", "remote", conn.RemoteAddr())
}
