package src

import (
	"encoding/json"
	"fmt"
	"heartBeat/src/config"
	lib "heartBeat/src/lib"
	"os"
	"strconv"
	"time"

	"github.com/robfig/cron/v3"
)

func InitHeartBeat() {

	c := cron.New(
		cron.WithParser(
			cron.NewParser(
				cron.SecondOptional | cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow)));
	_, err := c.AddFunc(fmt.Sprintf("*/%v * * * * *", config.HEARTBEAT_TIME_IN_SEC), cronListner)

	if err != nil {
		panic(err);
	}

	c.Start();
	fmt.Fprintf(os.Stderr,"::[heartBeat.go]::[%v] heartBeat started for every: %v seconds...\n",time.Now().Format("2006-01-02 15:04:05"), config.HEARTBEAT_TIME_IN_SEC);

	select {}
}

func cronListner() {
	now := time.Now();
	var epoch string = strconv.FormatInt(now.Unix(), 10);
	epochData, err := lib.RedisHGetAll(epoch);

	fmt.Fprintf(os.Stderr, "::[heartBeat.go]::cronListner::[%v]::epochData: %v\n", epoch, epochData);
	
	if (err != nil) {
		fmt.Fprintf(os.Stderr, "failed to hgetall at current epoch : %v,\n epoch: %v\n", epochData, err);
		panic(err);
	} else if (len(epochData) != 0) {

		// delete epoch key
		_, err := lib.RedisDEL(epoch);
		if err != nil {
			fmt.Fprintf(os.Stderr, "failed to RedisDEL current epoch : %v \n error : %v", epoch, err);
		}

		heartBeatAction(epoch, epochData);
	}

	// daily tasks at exactly 12:00 AM
	if now.Hour() == 0 && now.Minute() == 0 {
		sendDailyCronEvent("daily_cron_event", epoch);
	}

	// tasks of every 5 mins
	if now.Minute() % 5 == 0 {
		sendDailyCronEvent("five_min_cron_event", epoch);
	}
}

func heartBeatAction(epoch string, epochData map[string]string) {
		
	for uuid, srtData := range epochData {

		// delete uuid epoch mapping
		_, rdelError := lib.RedisDEL(uuid);
		if rdelError != nil {
			fmt.Fprintf(os.Stderr, "failed to RedisDEL current uuid-epoch mapping : %v \n error : %v", epoch, uuid);
		}

		var data map[string]interface{};

		err := json.Unmarshal([]byte(srtData), &data);
		if (err != nil){
			fmt.Fprintf(os.Stderr, "epoch data parse error,  uuid : %v,\n strData: %v\n, err: %v", uuid, srtData, err);
			continue;
		}

		if (data["callBackQueueName"] == "" || data["tenantDetails"] == "") {
			fmt.Fprintf(os.Stderr, "callBackQueueName or tenantDetails not found for this UUID, uuid : %v,\n data: %v\n", uuid, data);
			continue;
		}

		data["eventName"] = "heartBeat";
		data["requesterServiceName"] = config.SERVICE_NAME;

		sendErr := lib.SendToQueue(fmt.Sprintf("%v", data["callBackQueueName"]), data);
		if sendErr != nil {
			fmt.Fprintf(os.Stderr, "data SendToQueue error, epoch : %v,\n data: %v\n err: %v\n", epoch, data, err);
		}

	}
	
	

}

func sendDailyCronEvent(eventName string, epoch string) {
	var data map[string]interface{} = make(map[string]interface{})
	data["callBackQueueName"] = config.RABBIT_MQ_CRON_QUEUE_NAME
	data["eventName"] = "heartBeat";
	data["subEventName"] = eventName;
	data["requesterServiceName"] = config.SERVICE_NAME;

	sendErr := lib.SendToQueue(fmt.Sprintf("%v", data["callBackQueueName"]), data);
	if sendErr != nil {
		fmt.Fprintf(os.Stderr, "data SendToQueue error, epoch : %v,\n data: %v\n sendErr: %v\n", epoch, data, sendErr);
	}
}