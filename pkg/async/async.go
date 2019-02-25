package async

import (
	"errors"
)

type Message struct {
	Tag     string `json:"tag"`
	Payload []byte `json:"payload"`
}

type Writer interface {
	Write(tag string, payload interface{}) error
}

type Reader interface {
	Read() <-chan Message
}

type Pipe interface {
	Writer
	Reader
}

func New(config interface{}) (Pipe, error) {
	switch c := config.(type) {
	case RabbitConfig:
		return newRabbit(c)
	default:
		return nil, errors.New("unknown config type")
	}
}
