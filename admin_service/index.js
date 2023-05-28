const express = require('express');
const app = express();
const hostname = "0.0.0.0";
const port = 8082;

middleware = require("./middleware");
app.use(middleware.log);

//Create HTTP server and listen on port 3000 for requests
app.get('/api/titov-service/', (req, res) => {
	res.statusCode = 200;
	res.send('hello biden');
})

app.get('/foo', (req, res) => {
	res.send('footest');
})


//listen for request on port 3000, and as a callback function have the port listened on logged
app.listen(port);
