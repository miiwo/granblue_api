var express = require('express');
const pg = require('../db.js');
var router = express.Router();

router.get('/', function(req, res) {
    res.send("This is the weapons endpoint!");
});

router.get('/all', async (req, res) => {
    const weapons = await pg`select * from weapons`

    res.send(weapons);
});

router.get('/:id', function(req, res) {
    res.send('The id you specified is: ' + req.params.id);
});

module.exports = router;
