# HeartBeat

serviceName : "HeartBeat"

serviceNumber : 3

serviceType : support


## Appendix
to set timer using redis db and data along with it, data conatins callback queue name. once timerHits, this service will send stored data to callback queue name


## Authors

- Md Husain Thekiya
    - [hussainthekiya@gmail.com](mailto:hussainthekiya@gmail.com)
    - [https://github.com/MdHusainThekiya/](https://github.com/MdHusainThekiya/)
    - [https://www.linkedin.com/in/md-husain-thekiya/](https://github.com/MdHusainThekiya/)


## Production Setup
use below docker-compose.yml
```yml
version: '3'
services:
  redis:
    image: redis:7.2.2-alpine
    container_name: redis
    ports:
      - 6379:6379
    environment:
      - TZ=Asia/Kolkata

  rabbitmq:
    image: rabbitmq:3.12-management
    container_name: rabbitmq
    ports:
      - 5672:5672
      - 15672:15672

  heartbeat:
    image: mdhusainthekiya/heartbeat:v2.0.3
    container_name: heartbeat
    ports:
      - 4001:4001
    depends_on:
      - redis
      - rabbitmq
    env_file:
      - ./heartBeat/.env
    environment:
      - TZ=Asia/Kolkata
```
./heartBeat/.env contails following
```bash
# -- SERVICE CONFIG -- #
SERVICE_NAME=heartBeat

# -- KAFKA CONFIG (removed from v2.0.0 onwards) -- #
KAFKA_BROKERS="kafka1:9092,kafka2:9092,kafka3:9092"
KAFKA_TOPIC_NAME=test_topic
KAFKA_GROUP_ID=test_group
KAFKA_OTHER_TOPICS_NAME=<this optional env>"anotherTopic1,anotherTopic2,anotherTopic3"

# -- RABBITMQ CONFIG (available from v2.0.0 onwards) -- #
RABBIT_MQ_URL = "amqp://rabbitmq:5672"
RABBIT_MQ_QUEUE_NAME = "heartBeat"

# -- REDIS CONFIG -- #
REDIS_HOST=redis
REDIS_PORT=6379
REDIS_PASSWORD=
REDIS_DATABASE=0

# -- HEARTBEAT CONFIG -- #
HEARTBEAT_TIME_IN_SEC=5

# -- DAILY TASK QUEUE NAME -- #
RABBIT_MQ_CRON_QUEUE_NAME=apiServer
```


## Local Environment Setup

    0. required go to be installed (recommened 1.19 and available at current path)
    1. required rabbitMQ installed (recommened 3.12, and available at "amqp://rabbitmq:5672")
    2. required redis connection (example : localhost:6379)
    3. run following commands

### 1. Clone project
```bash
git clone https://github.com/MdHusainThekiya/heartBeat.git

cd heartBeat
```
### 2. create environment variables
To run this project, you will need to add the following environment variables to your .env file

refer `sampleEnv.txt`
```bash
cp sampleEnv.txt .env
vim .env
```
modify .env as per your values and execute below commands

```bash
go mod tidy
bash ./build.sh
go run heartBeat_binary
```
