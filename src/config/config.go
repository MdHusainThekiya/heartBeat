package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

var SERVICE_NAME                   string	= "";
var RABBIT_MQ_URL                  string	= "";
var RABBIT_MQ_CLIENT_NAME          string	= "";
var RABBIT_MQ_QUEUE_NAME           string	= "heartBeat";
var RABBIT_MQ_CONNECTION_RETRIES   int		= 60;
var RABBIT_MQ_CONNECTION_RETRY_SEC int		= 5;
var REDIS_HOST        						 string	= "";
var REDIS_PORT        						 string	= "";
var REDIS_PASSWORD        				 string	= "";
var REDIS_DATABASE        				 int		= 0;
var HEARTBEAT_TIME_IN_SEC 				 int		= 1;
var RABBIT_MQ_CRON_QUEUE_NAME			 string	= "heartBeat";

func LoadConfig() error {

	err := godotenv.Load(".env")

	if err != nil {
		fmt.Fprintln(os.Stderr,"Error loading .env file", err);
		return err;
	}

	SERVICE_NAME 				 	           = os.Getenv("SERVICE_NAME");
	RABBIT_MQ_URL 				           = os.Getenv("RABBIT_MQ_URL");
	RABBIT_MQ_CLIENT_NAME	           = SERVICE_NAME + "_" + uuid.New().String();
	RABBIT_MQ_QUEUE_NAME 	           = os.Getenv("RABBIT_MQ_QUEUE_NAME");
	REDIS_HOST				 			 		 		 = os.Getenv("REDIS_HOST");
	REDIS_PORT				 			 		 		 = os.Getenv("REDIS_PORT");
	REDIS_PASSWORD				 	 		 		 = os.Getenv("REDIS_PASSWORD");
	REDIS_DATABASE, _		  	 		 		 = strconv.Atoi(os.Getenv("REDIS_DATABASE"));
	HEARTBEAT_TIME_IN_SEC, _ 		 		 = strconv.Atoi(os.Getenv("HEARTBEAT_TIME_IN_SEC"));
	RABBIT_MQ_CRON_QUEUE_NAME 	     = os.Getenv("RABBIT_MQ_CRON_QUEUE_NAME");

	fmt.Fprintln(os.Stderr,"::[config.go]:: env loading completed.....");

	return nil;
}