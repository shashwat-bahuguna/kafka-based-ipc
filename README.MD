# Kafka Based Consumer Producer Example

## Description
A kafka based ipc example wherein, a producer, 3 consumers are in action. The consumers 1, 2 are from the same consumer group i.e., "cons-grp-1" while the consumer3 is from "cons-group-3".
All consumers are reading from the same topic, which is the one the producer is writing to i.e., "example-123"

## ScreenShot Description
The first screenshot is from the time when the consumer1 is recieving the messages from the "cons-grp-1". This consumer is then interrupted to ensure that the new incoming messages are now receiced by consumer2 on behlalf of "cons-grp-1". Consumer3 being from a different group receives all the written messages.