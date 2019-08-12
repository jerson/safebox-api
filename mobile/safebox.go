package safebox

import (
	"google.golang.org/grpc"
)

// SafeBox ...
type SafeBox struct {
	Server  string
	Session *Session
}

// NewSafeBox ...
func NewSafeBox(server string) *SafeBox {
	return &SafeBox{Server: server, Session: &Session{}}
}

func (s SafeBox) dial() (*grpc.ClientConn, error) {
	return grpc.Dial(s.Server, grpc.WithInsecure())
}
