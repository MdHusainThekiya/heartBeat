package main

import (
	heartBeat "heartBeat/src"
	config "heartBeat/src/config"
	lib "heartBeat/src/lib"
)

func main() {
	err := config.LoadConfig();

	if (err != nil) {
		panic(err);
	}

 	err = lib.RabbitMQConnect(eventListner);
	if (err != nil) {
		panic(err);
	}

	lib.RedisConnect();

	heartBeat.InitHeartBeat();
}

func eventListner(msg []byte) {
	
}