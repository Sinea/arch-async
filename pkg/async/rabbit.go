package async

import (
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"time"
)

const queue = "events"

type RabbitConfig struct {
	Url string
}

type rabbit struct {
	connection *amqp.Connection
	channel    *amqp.Channel
	queue      *amqp.Queue
	out        chan Message
}

// Dirty internals
func (r *rabbit) Write(kind string, payload interface{}) error {
	data, err := bundleMessage(kind, payload)

	if err != nil {
		return err
	}

	if r.channel == nil {
		ch, err := r.connection.Channel()

		if err != nil {
			return err
		}

		r.channel = ch
	}

	if r.queue == nil {
		q, err := r.channel.QueueDeclare(
			queue,
			false,
			false,
			false,
			false,
			nil,
		)

		if err != nil {
			return err
		}

		r.queue = &q
	}

	err = r.channel.Publish(
		"",
		r.queue.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        data,
		})

	if err != nil {
		return err
	}

	return nil
}

func (r *rabbit) Read() <-chan Message {
	if r.out == nil {
		r.out = make(chan Message)
		go func() {
			ch, _ := r.connection.Channel()

			q, _ := ch.QueueDeclare(
				queue,
				false,
				false,
				false,
				false,
				nil,
			)

			delivery, _ := ch.Consume(
				q.Name,
				"",
				false,
				false,
				false,
				false,
				nil,
			)

			for d := range delivery {
				time.Sleep(time.Millisecond * 30)
				if err := d.Ack(false); err != nil {
					fmt.Println(err)
					continue
				}
				m := Message{}
				if err := json.Unmarshal(d.Body, &m); err != nil {
					log.Fatalf("Bad")
				}
				r.out <- m
			}

			if err := ch.Close(); err != nil {
				log.Fatalf("Also bad")
			}
		}()
	}

	return r.out
}

func bundleMessage(tag string, payload interface{}) (data []byte, err error) {
	payloadBytes, err := json.Marshal(payload)

	if err != nil {
		return
	}

	m := Message{Tag: tag, Payload: payloadBytes}
	data, err = json.Marshal(m)
	if err != nil {
		return
	}

	return
}

func newRabbit(config RabbitConfig) (r *rabbit, err error) {
	conn, err := amqp.Dial(config.Url)

	if err != nil {
		return nil, err
	}

	return &rabbit{connection: conn}, nil
}
