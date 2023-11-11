
# HeartBeat

serviceName : "HeartBeat"

serviceNumber : 3

serviceType : support


## Appendix
to set timer using redis db and data along with it, data conatins callback kafka topic. once timerHits, this service will send stored data to callback kafka topic


## Authors

- Md Husain Thekiya
    - [hussainthekiya@gmail.com](mailto:hussainthekiya@gmail.com)
    - [https://github.com/MdHusainThekiya/](https://github.com/MdHusainThekiya/)
    - [https://www.linkedin.com/in/md-husain-thekiya/](https://github.com/MdHusainThekiya/)


## Environment Variables

To run this project, you will need to add the following environment variables to your .env file

refer `sampleEnv.txt`


## Installation

    0. required go to be installed
    1. required rabbitMQ installed
    2. required redis connection
    3. run following commands

```bash
go mod tidy
bash ./build.sh
go run heartBeat_binary
```

## Optional docker-compose.yml
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
```