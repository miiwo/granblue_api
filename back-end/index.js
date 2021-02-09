var dotenv = require('dotenv');
var express = require('express');
var cors = require('cors');
var app = express();



// Routers
var weapons = require('./routes/weapons.js');


dotenv.config();

app.use(cors());

// GET MAIN PAGE
//app.get('/', function(req, res) {
//    res.send("Hello World");
//});

app.use('/', express.static('./dist'));
app.use('/weapons', weapons);



app.get('*', function(req, res) {
    res.send('Sorry this is an invalid URL.');
});

// PORT to listen to
app.listen(process.env.PORT, function() {
    console.log('Cors-enabled webserver listening on: ' + process.env.PORT);
});
