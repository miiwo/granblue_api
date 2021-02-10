var dotenv = require('dotenv');
var express = require('express');
var cors = require('cors');


// Initializing
var app = express();
dotenv.config();
app.use(cors());

// Routers
var weapons = require('./routes/weapons.js');


// GET MAIN PAGE
app.use('/', express.static('./dist'));
app.use('/weapons', weapons);

app.get('*', function(req, res) {
    res.send('Sorry this is an invalid URL.');
});

// PORT to listen to
app.listen(process.env.PORT, function() {
    console.log('Cors-enabled webserver listening on: ' + process.env.PORT);
});
