package utils

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"time"

	model "github.com/captainkie/websync-api/internal/app/models"
	"github.com/captainkie/websync-api/pkg/helpers"
	amqp "github.com/rabbitmq/amqp091-go"
)

func ConnectionTask(tasks []model.ConnectionQueues) {
	// connect RabbitMQ
	rabbitMQURL := os.Getenv("RABBITMQ_URL")
	conn, err := amqp.Dial(rabbitMQURL)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()

	// Create a RabbitMQ channel
	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}
	defer ch.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Declare a queue for this channel
	for _, task := range tasks {
		body, err := json.Marshal(task)
		helpers.ErrorPanic(err)

		err = ch.PublishWithContext(ctx,
			"",                 // Exchange
			"connection_queue", // Routing key (queue name)
			false,              // Mandatory
			false,              // Immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(body),
			})

		if err != nil {
			log.Printf("Failed to publish connection_queue => [CONNECTION] : %v\n : %v", task.TransactionID, err)
		} else {
			log.Printf("Publish connection_queue => [CONNECTION] : %v\n", task.TransactionID)
		}
	}
}

func AddProductsTask(tasks []model.AddQueues) {
	// connect RabbitMQ
	rabbitMQURL := os.Getenv("RABBITMQ_URL")
	conn, err := amqp.Dial(rabbitMQURL)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()

	// Create a RabbitMQ channel
	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}
	defer ch.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Declare a queue for this channel
	for _, task := range tasks {
		body, err := json.Marshal(task)
		helpers.ErrorPanic(err)

		err = ch.PublishWithContext(ctx,
			"",              // Exchange
			"product_queue", // Routing key (queue name)
			false,           // Mandatory
			false,           // Immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(body),
			})

		if err != nil {
			log.Printf("Failed to publish product_queue => [ADD] : %d\n : %v", task.TransactionID, err)
		} else {
			log.Printf("Publish product_queue => [ADD] : %d\n", task.TransactionID)
		}
	}
}

func UpdateProductsTask(tasks []model.UpdateQueues) {
	// connect RabbitMQ
	rabbitMQURL := os.Getenv("RABBITMQ_URL")
	conn, err := amqp.Dial(rabbitMQURL)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()

	// Create a RabbitMQ channel
	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}
	defer ch.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Declare a queue for this channel
	for _, task := range tasks {
		body, err := json.Marshal(task)
		helpers.ErrorPanic(err)

		err = ch.PublishWithContext(ctx,
			"",              // Exchange
			"product_queue", // Routing key (queue name)
			false,           // Mandatory
			false,           // Immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(body),
			})

		if err != nil {
			log.Printf("Failed to publish product_queue => [UPDATE] : %d\n : %v", task.TransactionID, err)
		} else {
			log.Printf("Publish product_queue => [UPDATE] : %d\n", task.TransactionID)
		}
	}
}

func StockTask(tasks []model.StockQueues) {
	// connect RabbitMQ
	rabbitMQURL := os.Getenv("RABBITMQ_URL")
	conn, err := amqp.Dial(rabbitMQURL)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()

	// Create a RabbitMQ channel
	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}
	defer ch.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Declare a queue for this channel
	for _, task := range tasks {
		body, err := json.Marshal(task)
		helpers.ErrorPanic(err)

		err = ch.PublishWithContext(ctx,
			"",              // Exchange
			"product_queue", // Routing key (queue name)
			false,           // Mandatory
			false,           // Immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(body),
			})

		if err != nil {
			log.Printf("Failed to publish product_queue => [STOCK] : %d\n : %v", task.TransactionID, err)
		} else {
			log.Printf("Publish product_queue => [STOCK] : %d\n", task.TransactionID)
		}
	}
}

func StoreTask(tasks []model.StoreQueues) {
	// connect RabbitMQ
	rabbitMQURL := os.Getenv("RABBITMQ_URL")
	conn, err := amqp.Dial(rabbitMQURL)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()

	// Create a RabbitMQ channel
	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}
	defer ch.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Declare a queue for this channel
	for _, task := range tasks {
		body, err := json.Marshal(task)
		helpers.ErrorPanic(err)

		err = ch.PublishWithContext(ctx,
			"",              // Exchange
			"product_queue", // Routing key (queue name)
			false,           // Mandatory
			false,           // Immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(body),
			})

		if err != nil {
			log.Printf("Failed to publish product_queue => [STORE] : %d\n : %v", task.TransactionID, err)
		} else {
			log.Printf("Publish product_queue => [STORE] : %d\n", task.TransactionID)
		}
	}
}

func PostflagTask(tasks []model.PostflagQueues) {
	// connect RabbitMQ
	rabbitMQURL := os.Getenv("RABBITMQ_URL")
	conn, err := amqp.Dial(rabbitMQURL)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()

	// Create a RabbitMQ channel
	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}
	defer ch.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Declare a queue for this channel
	for _, task := range tasks {
		body, err := json.Marshal(task)
		helpers.ErrorPanic(err)

		err = ch.PublishWithContext(ctx,
			"",               // Exchange
			"postflag_queue", // Routing key (queue name)
			false,            // Mandatory
			false,            // Immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(body),
			})

		if err != nil {
			log.Printf("Failed to publish postflag_queue => [POSTFLAG] : %v\n : %v", task.TransactionID, err)
		} else {
			log.Printf("Publish postflag_queue => [POSTFLAG] : %v\n", task.TransactionID)
		}
	}
}

func ImageTask(tasks []model.ImageQueues) {
	// connect RabbitMQ
	rabbitMQURL := os.Getenv("RABBITMQ_URL")
	conn, err := amqp.Dial(rabbitMQURL)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()

	// Create a RabbitMQ channel
	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}
	defer ch.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Declare a queue for this channel
	for _, task := range tasks {
		body, err := json.Marshal(task)
		helpers.ErrorPanic(err)

		err = ch.PublishWithContext(ctx,
			"",            // Exchange
			"image_queue", // Routing key (queue name)
			false,         // Mandatory
			false,         // Immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(body),
			})

		if err != nil {
			log.Printf("Failed to publish image_queue => [IMAGE] : %v\n : %v", task.TransactionID, err)
		} else {
			log.Printf("Publish image_queue => [IMAGE] : %v\n", task.TransactionID)
		}
	}
}

func DailysaleTask(tasks []model.DailysaleQueues) {
	// connect RabbitMQ
	rabbitMQURL := os.Getenv("RABBITMQ_URL")
	conn, err := amqp.Dial(rabbitMQURL)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()

	// Create a RabbitMQ channel
	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}
	defer ch.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Declare a queue for this channel
	for _, task := range tasks {
		body, err := json.Marshal(task)
		helpers.ErrorPanic(err)

		err = ch.PublishWithContext(ctx,
			"",                // Exchange
			"dailysale_queue", // Routing key (queue name)
			false,             // Mandatory
			false,             // Immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(body),
			})

		if err != nil {
			log.Printf("Failed to publish dailysale_queue => [DAILYSALE] : %v\n : %v", task.TransactionID, err)
		} else {
			log.Printf("Publish dailysale_queue => [DAILYSALE] : %v\n", task.TransactionID)
		}
	}
}
