var express = require('express');
var router = express.Router();
var wepctrl = require('../controllers/weaponsController.js');


router.get('/', function(req, res) { res.send("This is the weapons endpoint!"); });
router.get('/all', wepctrl.getAllWeapons);
router.get('/weapon', wepctrl.sanitizeQuery, wepctrl.getWeaponByQuery);
router.get('/:id', wepctrl.sanitizeId, wepctrl.getWeaponByName);


module.exports = router;
