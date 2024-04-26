compose script: `docker compose -f "docker-compose.yml" up -d --build`

exex: `docker exec -it kafka bin/bash`

create topic: `kafka-topics.sh --create --zookeeper zookeeper:2181 --replication-factor 1 --partitions 1 --topic <topic-name>`

list: `kafka-topics.sh --list --zookeeper zookeeper:2181`

producer: `kafka-console-producer.sh --broker-list kafka:9092 --topic mesages`

consumer: `kafka-console-consumer.sh --topic messages --from-beginning --bootstrap-server kafka:9092`