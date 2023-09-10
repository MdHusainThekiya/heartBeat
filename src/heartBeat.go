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
	var epoch string = strconv.FormatInt(time.Now().Unix(), 10);
	epochData, err := lib.RedisHGetAll(epoch);

	if (err != nil) {
		fmt.Fprintf(os.Stderr, "failed to hgetall at current epoch : %v,\n epoch: %v\n", epochData, epoch);
	} else if (len(epochData) != 0) {
		epochData["heartBeatEpoch"] = epoch;
		heartBeatAction(epochData);
	}
}

func heartBeatAction(epochData map[string]string) {

	if (epochData["callBackTopicName"] == "" || epochData["tenantDetails"] == "") {
		fmt.Fprintf(os.Stderr, "callBackTopicName or tenantDetails not found in heartBeat epoch epochData : %v", epochData);
	} else {
		
		epochData["kafkaEventName"] = "heartBeat";
		epochData["requesterServiceName"] = config.SERVICE_NAME;

		jsonBytes, err := json.Marshal(epochData);

		if (err != nil) {
			fmt.Fprintf(os.Stderr, "failed to json.marshal epochData : %v\n", epochData);
		}

		strEpochData := string(jsonBytes);

		lib.SendKafka(epochData["callBackTopicName"], strEpochData)

	}

}