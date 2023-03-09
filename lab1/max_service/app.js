
const express = require("express");
const app = express();
 
app.get("/api/max-service", function(request, response){
    response.end("Hello from Express!");
});

app.listen(8081);