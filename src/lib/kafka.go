package lib

// import (
// 	"errors"
// 	"fmt"
// 	config "heartBeat/src/config"
// 	"os"
// 	"strings"
// 	"time"

// 	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
// 	// "github.com/google/uuid"
// )

// var KafkaConsumer *kafka.Consumer;
// var KafkaProducer *kafka.Producer;
// var Message *kafka.Message;

// func KafkaConnect(eventListner func([]byte)) error {

// 	fmt.Fprintln(os.Stderr,"::[kafka.go]::starting kafka connection...");

// 	if ( config.SERVICE_NAME == "" || config.KAFKA_BROKERS == "" || config.KAFKA_TOPIC_NAME == "" || config.KAFKA_GROUP_ID == "" ) {

// 		fmt.Fprintf(os.Stderr, "SERVICE_CONFIG_NOT_FOUND_IN_ENV :: %v", map[string]interface{}{
// 			"requiredConfigs" : "SERVICE_NAME,KAFKA_BROKERS,KAFKA_TOPIC_NAME,KAFKA_GROUP_ID",
// 		})

// 		return errors.New("SERVICE_CONFIG_NOT_FOUND_IN_ENV");
// 	}
// 	var kafkaTopicName []string = []string{config.KAFKA_TOPIC_NAME};

// 	if (config.KAFKA_OTHER_TOPICS_NAME != "") {
// 		kafkaTopicName = strings.Split(config.KAFKA_TOPIC_NAME + "," + config.KAFKA_OTHER_TOPICS_NAME, ",")
// 	}

// 	KafkaConsumer, err := kafka.NewConsumer(&kafka.ConfigMap{
// 		"bootstrap.servers": config.KAFKA_BROKERS,
// 		"group.id":          config.KAFKA_GROUP_ID,
// 		"auto.offset.reset": "earliest",
// 		"go.application.rebalance.enable": true,
// 	})
// 	if err != nil {
// 		return err;
// 	}

// 	KafkaProducer, err = kafka.NewProducer(&kafka.ConfigMap{
// 		"bootstrap.servers": config.KAFKA_BROKERS,
// 	})
// 	if err != nil {
// 		return err;
// 	}

// 	fmt.Fprintln(os.Stderr,"::[kafka.go]:: kafka consumer connected...");

// 	KafkaConsumer.SubscribeTopics(kafkaTopicName, nil);

// 	fmt.Fprintln(os.Stderr,"::[kafka.go]:: kafka topic subscribe success :: starting msg consumer...");

// 	// start(eventListner);

// 	// A signal handler or similar could be used to set this to false to break the loop.
// 	run := true

// 	go func () {
// 		for run {
// 			msg, err := KafkaConsumer.ReadMessage(time.Second)
// 			if err == nil {
// 				fmt.Fprintln(os.Stderr,"Message recieved on == TopicPartition :", msg.TopicPartition);
// 				go func() {
// 					KafkaConsumer.Commit();
// 			}()
// 			eventListner(msg.Value);
// 			} else if !err.(kafka.Error).IsTimeout() {
// 				fmt.Fprintf(os.Stderr,"Consumer error: %v (%v)\n", err, msg)
// 			}
// 		}
// 	}();

// 	// Delivery report handler for produced messages
// 	go func() {
// 		for e := range KafkaProducer.Events() {
// 			switch ev := e.(type) {
// 			case *kafka.Message:
// 				if ev.TopicPartition.Error != nil {
// 					fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
// 				} else {
// 					fmt.Printf("Delivered message to %v\n", ev.TopicPartition)
// 				}
// 			}
// 		}
// 	}()

// 	return nil;
// }

// func SendKafka(topicName string, message string) {

// 	KafkaProducer.Produce(&kafka.Message{
// 		TopicPartition: kafka.TopicPartition{Topic: &topicName, Partition: kafka.PartitionAny},
// 		Value:          []byte(message),
// 	}, nil)

// }