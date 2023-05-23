
const express = require("express");
const http = require('http');
require('dotenv').config();
const app = express();

app.get("/api/max-service", function (request, response) {
    response.end("Hello from Express!");
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
            if (res.statusCode === 200)
                ok++;
            if (ok === 2) {
                response.status(200).send(result);
            }
        });
    }).on('error', (err) => {
        response.status(400).send('error: ', err.message);
    });

    var r2 = http.request(`http://${process.env.ADMIN_SERVICE_NAME}/api/titov-service/`, (res) => {
        let data = '';
        res.on('data', (chunk) => {
            data += chunk;
        });
        res.on('end', () => {
            result += data;
            if (res.statusCode === 200)
                ok++;
            if (ok === 2) {
                response.status(200).send(result);
            }
        });
    }).on('error', (err) => {
        response.status(400).send('error: ', err.message);
    });

    // Send the requests
    r1.end();
    r2.end();

});

app.listen(8081);
