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

func solveSudoku(grid *[utils.N][utils.N]int) bool {
	return utils.SolveSudoku(grid)
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

	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			var grid [utils.N][utils.N]int
			err := json.Unmarshal(d.Body, &grid)
			failOnError(err, "Failed to unmarshal JSON")

			if solveSudoku(&grid) {
				log.Println("Sudoku solved successfully")
			} else {
				log.Println("Failed to solve Sudoku")
			}

			utils.PrintGrid(grid)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
