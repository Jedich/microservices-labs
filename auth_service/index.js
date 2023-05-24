const express = require('express');
const app = express();
const http = require('http');
require('dotenv').config();
const Kafka = require('node-rdkafka');
const hostname = "0.0.0.0";
const port = 8085;

app.post('/api/auth/login', (req, res) => {
	res.statusCode = 200;
	res.send('logged in');
})

const producer = new Kafka.Producer({
	'metadata.broker.list': process.env.KAFKA_URL,
});

producer.connect();

producer.on('ready', () => {
	console.log('Producer ready');
});

producer.on('event.error', (err) => {
	console.error('Error in producer:', err);
});

app.post('/api/auth/register', (request, response) => {
	var r1 = http.request({
		hostname: `http://${process.env.USER_SERVICE_NAME}`,
		path: '/api/user/',
		method: 'POST',
		headers: {
			'Content-Type': 'application/json',
			'Content-Length': Buffer.byteLength(request.body)
		}
	}, (res) => {
		let data = '';
		res.on('data', (chunk) => {
			data += chunk;
		});
		res.on('end', () => {
			if (res.statusCode !== 200) {
				response.status(400).send('error');
			}

			var jwt = createJWT(data.id);

			const topicName = process.env.KAFKA_TOPIC;
			const message = jwt;

			producer.produce(
				topicName,
				null,
				Buffer.from(message),
				null,
				Date.now()
			);
		});
	}).on('error', (err) => {
		response.status(400).send('error: ', err.message);
	});

	response.statusCode = 200;
	response.send('Registered successfully');
})

function createJWT(payload) {
	const secretKey = process.env.JWT_SECRET;
	const options = {
		expiresIn: '1h'
	};
	return jwt.sign(payload, secretKey, options);
}

app.listen(port);
