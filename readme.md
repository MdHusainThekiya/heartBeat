
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
    1. required kafka with minimum single broker and topicName created with it
    2. required redis connection
    3. run following commands

```bash
  go mod tidy
  bash ./build.sh
  go run heartBeat
```