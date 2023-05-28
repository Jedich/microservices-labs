const fluentLogger = require('fluent-logger');
require('dotenv').config();

const logger = fluentLogger.createFluentSender('movie-mserv', {
    host: process.env.FLUENT_HOST,
    port: 24224,
    timeout: 3000,
});

var log = (req, res, next) => {
    const oldWrite = res.write;
    const oldEnd = res.end;
    const chunks = [];

    res.write = function (chunk) {
        chunks.push(Buffer.from(chunk));
        oldWrite.apply(res, arguments);
    };

    res.end = function (chunk) {
        if (chunk) {
            chunks.push(Buffer.from(chunk));
        }

        const body = Buffer.concat(chunks).toString('utf8');

        // Log the response body to Fluent
        logger.emit('response', {
            url: req.url,
            method: req.method,
            body,
        });

        oldEnd.apply(res, arguments);
    };

    next();
}

module.exports = {
    log: log
}