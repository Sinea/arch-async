package async

import (
	"encoding/json"
	"github.com/streadway/amqp"
)

type RabbitConfig struct {
	URL   string
	Queue string
}

type rabbit struct {
	config     RabbitConfig
	connection *amqp.Connection
	channel    *amqp.Channel
	queue      *amqp.Queue
	out        chan Message
	err        chan error
}

// Dirty internals
func (r *rabbit) Write(tag string, payload interface{}) error {
	data, err := bundleMessage(tag, payload)

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
			r.config.Queue,
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

	return r.channel.Publish(
		"",
		r.queue.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        data,
		})
}

func (r *rabbit) Read() (<-chan Message, <-chan error) {
	if r.out == nil {
		r.out = make(chan Message)
		r.err = make(chan error, 1)

		if delivery, err := r.getDelivery(); err == nil {
			go r.deliver(delivery)
		}
	}

	return r.out, r.err
}

func (r *rabbit) endWithError(err error) {
	r.err <- err
	close(r.err)
	close(r.out)
}

func (r *rabbit) getDelivery() (<-chan amqp.Delivery, error) {
	ch, err := r.connection.Channel()
	if err != nil {
		r.endWithError(err)
		return nil, err
	}

	q, err := ch.QueueDeclare(
		r.config.Queue,
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		r.endWithError(err)
		return nil, err
	}

	delivery, err := ch.Consume(
		q.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		r.endWithError(err)
		return nil, err
	}

	return delivery, nil
}

func (r *rabbit) deliver(delivery <-chan amqp.Delivery) {
	for d := range delivery {
		m := Message{}
		if err := json.Unmarshal(d.Body, &m); err != nil {
			r.endWithError(err)
			break
		}

		r.out <- m

		if err := d.Ack(false); err != nil {
			r.endWithError(err)
			break
		}
	}

	if err := r.channel.Close(); err != nil {
		r.endWithError(err)
	}
}

func bundleMessage(tag string, payload interface{}) ([]byte, error) {
	payloadBytes, err := json.Marshal(payload)

	if err != nil {
		return nil, err
	}

	m := Message{Tag: tag, Payload: payloadBytes}
	return json.Marshal(m)
}

func newRabbit(config RabbitConfig) (r *rabbit, err error) {
	conn, err := amqp.Dial(config.URL)

	if err != nil {
		return nil, err
	}

	return &rabbit{config: config, connection: conn}, nil
}
