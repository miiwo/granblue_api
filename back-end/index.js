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

// Authorization
function auth_request(req, res, next) {
    console.log("Checking Authorization...");

    if (req.get('Authorization') === process.env.API_KEY) {
        console.log("Authorization granted");
        next();
    } else {
        console.log('Rejecting request...');
        res.status(403).json({error: 'Request denied. Please put proper Authorization.'});
    }
};

// Make a generateAuthTokenFunction sometime

// Routes
app.use('/weapons', auth_request, weapons);

app.get('*', function(req, res) {
    res.send('Sorry this is an invalid URL.');
});

// PORT to listen to
app.listen(process.env.PORT, function() {
    console.log('Cors-enabled webserver listening on: ' + process.env.PORT);
});
