package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

var SERVICE_NAME           	string = "";
var KAFKA_BROKERS          	string = "";
var KAFKA_TOPIC_NAME       	string = "";
var KAFKA_OTHER_TOPICS_NAME string = "";
var KAFKA_GROUP_ID         	string = "";
var KAFKA_CLIENT_ID        	string = "";
var REDIS_HOST        			string = "";
var REDIS_PORT        			string = "";
var REDIS_PASSWORD        	string = "";
var REDIS_DATABASE        	int = 0;
var HEARTBEAT_TIME_IN_SEC   int = 1;

func LoadConfig() error {

	err := godotenv.Load(".env")

	if err != nil {
		fmt.Fprintln(os.Stderr,"Error loading .env file", err);
		return err;
	}

	SERVICE_NAME 				 	 	 = os.Getenv("SERVICE_NAME");
	KAFKA_BROKERS					 	 = os.Getenv("KAFKA_BROKERS");
	KAFKA_TOPIC_NAME			 	 = os.Getenv("KAFKA_TOPIC_NAME");
	KAFKA_OTHER_TOPICS_NAME  = os.Getenv("KAFKA_OTHER_TOPICS_NAME");
	KAFKA_GROUP_ID				 	 = os.Getenv("KAFKA_GROUP_ID");
	KAFKA_CLIENT_ID				 	 = SERVICE_NAME + "_" + uuid.New().String();
	REDIS_HOST				 			 = os.Getenv("REDIS_HOST");
	REDIS_PORT				 			 = os.Getenv("REDIS_PORT");
	REDIS_PASSWORD				 	 = os.Getenv("REDIS_PASSWORD");
	REDIS_DATABASE, _		  	 = strconv.Atoi(os.Getenv("REDIS_DATABASE"));
	HEARTBEAT_TIME_IN_SEC, _ = strconv.Atoi(os.Getenv("HEARTBEAT_TIME_IN_SEC"));

	fmt.Fprintln(os.Stderr,"::[config.go]:: env loading completed.....");

	return nil;
}