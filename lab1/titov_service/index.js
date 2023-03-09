const express = require('express');
const app = express();
const hostname = "0.0.0.0";
const port = 8082;

//Create HTTP server and listen on port 3000 for requests
app.get('/', (req, res) => {
	res.statusCode = 200;
	res.send('Hello World!');
})

app.get('/foo', (req, res) => {
	res.send('footest');
})


//listen for request on port 3000, and as a callback function have the port listened on logged
app.listen(port, hostname, () => {
	console.log(`Server running at http://${hostname}:${port}/`);
});
