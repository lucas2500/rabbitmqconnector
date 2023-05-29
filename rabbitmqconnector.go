package rabbitmqconnector

import (
	"log"
	"os"

	amqp "github.com/rabbitmq/amqp091-go"
)

var (
	RabbitConn *amqp.Connection
)

func ConnectoToRabbitMQ() {

	var err error

	log.Println("- Connecting to RabbitMQ...")

	dsn := os.Getenv("RABBITMQ_STRING_CONNECTION")

	RabbitConn, err = amqp.Dial(dsn)

	if err != nil {
		log.Fatalf("Unable to connect to RabbitMQ -> %s\n", err)
	}
}
