package lib

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	config "heartBeat/src/config"
	"log"
	"os"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

var amqpCTX = context.Background();
var rabbitMQConsumerChannel *amqp.Channel;
var rabbitMQProducerChannel *amqp.Channel;

func retriableMQConnect() (*amqp.Connection, error ) {

	for i := 0; i < config.RABBIT_MQ_CONNECTION_RETRIES; i++ {
		conn, err := amqp.Dial(config.RABBIT_MQ_URL)
		if err == nil {
			fmt.Fprintln(os.Stderr,"::[rabbitMQ.go]::got rabbitMQ connection...");
			return conn, nil
		}

		fmt.Fprintln(os.Stderr,"::[rabbitMQ.go]::amqp connection error...", err);
		
		if (i+1 < config.RABBIT_MQ_CONNECTION_RETRIES) {
			time.Sleep(time.Duration(config.RABBIT_MQ_CONNECTION_RETRY_SEC) * time.Second)
			fmt.Fprintln(os.Stderr,"::[rabbitMQ.go]::retiry to connect amqp...");
			fmt.Fprintf(os.Stderr,"::[rabbitMQ.go]::retryCount : %v, retryingInterval : %v sec", i+1, config.RABBIT_MQ_CONNECTION_RETRY_SEC);
		}
	}

	return nil, fmt.Errorf("failed to connect to RabbitMQ after %d retries", config.RABBIT_MQ_CONNECTION_RETRIES);
}

func RabbitMQConnect(eventListner func([]byte)) error {

	fmt.Fprintln(os.Stderr,"::[rabbitMQ.go]::starting rabbitMQ connection...");
	
	if ( config.SERVICE_NAME == "" || config.RABBIT_MQ_URL == "" || config.RABBIT_MQ_QUEUE_NAME == "" ) {

		fmt.Fprintf(os.Stderr, "SERVICE_CONFIG_NOT_FOUND_IN_ENV :: %v", map[string]interface{}{
			"requiredConfigs" : "SERVICE_NAME,RABBIT_MQ_URL,RABBIT_MQ_QUEUE_NAME",
		})
		
		return errors.New("SERVICE_CONFIG_NOT_FOUND_IN_ENV");
	}

	// Assign channel for consumer
	consumerConn, err := retriableMQConnect();
	if (err != nil) {
		return err
	}
	consumerChannel, err := consumerConn.Channel();
	if (err != nil) {
		log.Panicln("::[rabbitMQ.go]::amqp get channel error...", err);
		defer consumerConn.Close()
		return err
	}
	rabbitMQConsumerChannel = consumerChannel;
	fmt.Fprintln(os.Stderr,"::[rabbitMQ.go]::got consumer connection...");

	// Assign channel for producer
	producerConn, err := retriableMQConnect();
	if (err != nil) {
		defer consumerChannel.Close()
		defer consumerConn.Close()
		return err
	}
	producerChannel, err := producerConn.Channel();
	if (err != nil) {
		log.Panicln("::[rabbitMQ.go]::amqp get channel error...", err);
		defer consumerChannel.Close()
		defer consumerConn.Close()
		defer producerConn.Close()
		return err
	}
	rabbitMQProducerChannel = producerChannel;
	fmt.Fprintln(os.Stderr,"::[rabbitMQ.go]::got producer connection...");


	// start consumer
	consumerQ, err := rabbitMQConsumerChannel.QueueDeclare(
		config.RABBIT_MQ_QUEUE_NAME, // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	if (err != nil) {
		log.Panicln("::[rabbitMQ.go]::declare consumerQ error...", err);
		defer consumerChannel.Close()
		defer consumerConn.Close()
		defer producerChannel.Close()
		defer producerConn.Close()
		return err
	}
	fmt.Fprintln(os.Stderr,"::[rabbitMQ.go]::declared consumer queue...");

	consumerMessage, err := rabbitMQConsumerChannel.Consume(
		consumerQ.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if (err != nil) {
		log.Panicln("::[rabbitMQ.go]::declare consumerQ error...", err);
		defer consumerChannel.Close()
		defer consumerConn.Close()
		defer producerChannel.Close()
		defer producerConn.Close()
		return err
	}

	go func () {
		var forever chan struct{}
		
		go func() {
			for d := range consumerMessage {
				log.Printf("Received a message: %s", d.Body)
				eventListner(d.Body);
			}
		}()
	
		<-forever

	}();

	return nil;
}


func SendToQueue(queueName string, sendingData map[string]interface{}) error {

	sendingData["eventStamp"] = time.Now().UnixMilli();
	jsonBytes, err := json.Marshal(sendingData);

	if (err != nil) {
		fmt.Fprintf(os.Stderr, "failed to json.marshal epochData : %v\n", sendingData);
		return nil
	}

	strEpochData := string(jsonBytes);

	err = rabbitMQProducerChannel.PublishWithContext(amqpCTX,
		"",     // exchange
		queueName, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing {
			ContentType: "text/plain",
			Body:        []byte(strEpochData),
		})

	if (err != nil) {
		fmt.Fprintf(os.Stderr, "::[rabbitMQ.go]::message publish error... %v\n", err);
		panic(err);
	}
	
	return nil

}