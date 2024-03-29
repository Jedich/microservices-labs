const express = require("express");
const http = require('http');
// const { Worker } = require('worker_threads');
require('dotenv').config();
const app = express();

middleware = require("./middleware");
app.use(middleware.log);

app.get("/api/max-service", function (request, response) {
    response.send("Hello from Express!").end();
});

app.get("/api/send-requests", function (request, response) {
    var ok = 0;
    var result = "";
    var r1 = http.request(`http://${process.env.USER_SERVICE_NAME}/api/golang-service/`, (res) => {
        let data = '';
        res.on('data', (chunk) => {
            data += chunk;
        });
        res.on('end', () => {
            result += data;
            if (res.statusCode === 200) {
                ok++;
            }
            if (ok === 2) {
                response.status(200).send(result);
            }
        });
    }).on('error', (err) => {
        response.status(400).send('Error: ' + err.message);
    });

    var r2 = http.request(`http://${process.env.ADMIN_SERVICE_NAME}/api/titov-service/`, (res) => {
        let data = '';
        res.on('data', (chunk) => {
            data += chunk;
        });
        res.on('end', () => {
            result += data;
            if (res.statusCode === 200) {
                ok++;
            }
            if (ok === 2) {
                response.status(200).send(result);
            }
        });
    }).on('error', (err) => {
        response.status(400).send('Error: ' + err.message);
    });

    // Send the requests
    r1.end();
    r2.end();

});

// const kafkaConsumerThread = new Worker('./consumer.js');

// // Listen for messages from the Kafka consumer thread
// kafkaConsumerThread.on('message', (message) => {
//     console.log('Received message from Kafka:', message);
// });

// // Handle errors from the Kafka consumer thread
// kafkaConsumerThread.on('error', (err) => {
//     console.error('Error in Kafka consumer thread:', err);
// });

app.listen(8081);
