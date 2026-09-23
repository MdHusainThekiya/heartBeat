## [v2.0.5](https://github.com/MdHusainThekiya/heartBeat/tree/v2.0.5) (2026-09-23)

### Security

* Upgraded github.com/rabbitmq/amqp091-go from v1.9.0 to v1.15.0
* Fixes CVE-2026-77405 (Critical) - Missing minimum TLS version in Dial()
* Fixes CVE-2026-77408 (Critical) - Silent data truncation via integer overflow in writeShortstr
* Fixes CVE-2026-77411 (Critical) - Protocol desynchronization via oversized longstr
* Fixes CVE-2026-79921 (High) - Broker-driven heartbeat deadlocks
* Fixes CVE-2026-77404 (High) - Plaintext credential retention in memory post-handshake

## [v2.0.3](https://github.com/MdHusainThekiya/heartBeat/tree/v2.0.3) (2025-04-05)

### Bug

* Added panic on redis or rabbitmq runtime error to make this service restart

## [v2.0.2](https://github.com/MdHusainThekiya/heartBeat/tree/v2.0.2) (2024-12-30)

### Feature

* Added custom event for crons such as daily events

## [v2.0.0](https://github.com/MdHusainThekiya/heartBeat/tree/v2.0.0) (2023-11-12)

### Optimization

* Replaced Kafka.js with RabbitMQ to remove headach of managing kafka

## [v1.0.0](https://github.com/MdHusainThekiya/heartBeat/tree/v1.0.0) (2023-09-25)

### Stable Release

* Stable Release with Kafka.js as a IPC