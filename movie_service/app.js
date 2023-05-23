
const express = require("express");
const http = require('http');
require('dotenv').config();
const app = express();

app.get("/api/max-service", function (request, response) {
    response.end("Hello from Express!");
});

app.get("/api/send-requests", function (request, response) {
    http.request(`http://${process.env.USER_SERVICE_NAME}/api/golang-service`, (res) => {
        res.on('error', (err) => {
            response.status(400).send('error: ', err.message);
        });
    });
    http.request(`http://${process.env.ADMIN_SERVICE_NAME}/api/titov-service`, (res) => {
        res.on('error', (err) => {
            response.status(400).send('error: ', err.message);
        });
    });
    response.status(200).send();

});

app.listen(8081);
