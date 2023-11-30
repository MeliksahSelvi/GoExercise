package brokers

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

type RabbitMQ struct {
	conn *amqp.Connection
}

func NewRabbitMQ() *RabbitMQ {
	log.Printf("dialing %q", "amqp://guest:guest@localhost:5672/")
	connection, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		fmt.Println(fmt.Errorf("Dial: %s", err))
		return nil
	}
	return &RabbitMQ{
		conn: connection,
	}
}

func (rabbit RabbitMQ) Publish(body []byte) {

	log.Printf("got Connection, getting Channel")
	channel, err := rabbit.conn.Channel()
	if err != nil {
		fmt.Println(err)
		return
	}

	if err := channel.ExchangeDeclare(
		"gotr-city-exchange",
		"fanout",
		true,
		false,
		false,
		false,
		nil,
	); err != nil {
		fmt.Println(err)
	}

	log.Printf("declared Exchange, publishing %dB body (%q)", len(body), body)
	if err = channel.Publish(
		"gotr-city-exchange",
		"",
		false,
		false,
		amqp.Publishing{
			Headers:         amqp.Table{},
			ContentType:     "text/plain",
			ContentEncoding: "",
			Body:            body,
			DeliveryMode:    amqp.Transient,
			Priority:        0,
		},
	); err != nil {
		fmt.Println(err)
	}
}
