package options

import "time"

type ServerRunOptions struct {
	Name      string
	AliveTime time.Duration
	BufChannel   chan interface{}
}

func NewServerRunOptions() *ServerRunOptions {
	s := ServerRunOptions{
		Name:" Mock  Server ",
		AliveTime:  1 * time.Second,
	}
	return &s
}

