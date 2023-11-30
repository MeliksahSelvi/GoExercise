package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		fmt.Println(err)
		return
	}

	go func() { //burası connection kapandığı zaman bize bilgi atıyor
		fmt.Printf("closing: %s", <-conn.NotifyClose(make(chan *amqp.Error)))
	}()

	log.Printf("got Connection, getting Channel")
	channel, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
		return
	}

	log.Printf("declared Exchange, declaring Queue %q", "long-running-task-queue")
	queue, err := channel.QueueDeclare(
		"long-running-task-queue",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		fmt.Println(err)
		return
	}

	log.Printf("declared Queue (%q %d messages, %d consumers), binding to Exchange (key %q)",
		queue.Name, queue.Messages, queue.Consumers, "")

	if err = channel.QueueBind(
		queue.Name,
		"",
		"gotr-city-exchange",
		false,
		nil,
	); err != nil {
		fmt.Println(err)
		return
	}

	log.Printf("Queue bound to Exchange, starting Consume (consumer tag %q)", "city-consumer")
	deliveries, err := channel.Consume(
		queue.Name,
		"city-consumer",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		message := <-deliveries
		fmt.Println("read data")
		fmt.Println(string(message.Body))
		message.Ack(false)
	}
}
