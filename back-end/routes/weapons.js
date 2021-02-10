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

router.get('/:id', async (req, res) => {
    console.log('recieved: ' + req.params.id);

    const query = req.params.id.toLowerCase();
    const weapons = await pg`select * from weapons where lower(name) like ${ '%' + query + '%' }`

    res.send(weapons);
});

module.exports = router;
