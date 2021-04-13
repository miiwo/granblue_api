var dotenv = require('dotenv');
var express = require('express');
var cors = require('cors');


// Initializing
var app = express();
dotenv.config();
app.use(cors());
app.set('query parser', 'simple');

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
        res.status(401).json({error: 'Request denied. Please put proper authorization.'});
    }
};

// Make a generateAuthTokenFunction sometime

// Routes
app.use('/weapons', auth_request, weapons);

// 404 Page
app.get('*', function(req, res) {
    res.status(404).json('Sorry this is an invalid URL.');
});

// PORT to listen to
app.listen(process.env.PORT, function() {
    console.log('Cors-enabled webserver listening on: ' + process.env.PORT);
});
