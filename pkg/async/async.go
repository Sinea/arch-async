package async

import (
	"errors"
)

type Message struct {
	Kind    string `json:"kind"`
	Payload []byte `json:"payload"`
}

type Writer interface {
	Write(kind string, payload interface{}) error
}

type Reader interface {
	Read() <-chan Message
}

type Pipe interface {
	Writer
	Reader
}

func New(config interface{}) (Pipe, error) {
	switch config.(type) {
	case RabbitConfig:
		return newRabbit(config.(RabbitConfig))
	default:
		return nil, errors.New("unknown config type")
	}
}
