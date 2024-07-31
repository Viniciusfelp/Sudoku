package main

import (
	"encoding/json"
	"log"
	"sudoku/utils"
	"github.com/streadway/amqp"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"sudoku_queue",
		false,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to declare a queue")

	sudokuRequest := utils.BuildSudokuRequest()
	body, err := json.Marshal(sudokuRequest.Grid)
	failOnError(err, "Failed to marshal JSON")

	err = ch.Publish(
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		})
	failOnError(err, "Failed to publish a message")

	log.Printf(" [x] Sent %s", body)
}
