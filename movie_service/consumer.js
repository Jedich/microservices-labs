const { parentPort } = require('worker_threads');
const Kafka = require('node-rdkafka');
require('dotenv').config();

// Kafka consumer configuration
const consumer = new Kafka.KafkaConsumer({
  'metadata.broker.list': process.env.KAFKA_URL,
  'auto.offset.reset': 'earliest',
});

const topicName = process.env.KAFKA_TOPIC;
consumer.subscribe([topicName]);

consumer.on('message', (message) => {
  parentPort.postMessage(message.value.toString());
});

consumer.connect();

consumer.on('event.error', (err) => {
  console.error('Error in consumer:', err);
});