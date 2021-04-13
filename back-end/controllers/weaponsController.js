const pg = require('../db.js');

async function getWeaponByName(req, res) {
    const query = req.params.id.toLowerCase();
    const weapons = await pg`select * from weapons where lower(name) like ${ '%' + query + '%' }`
    res.send(weapons);
}

async function getAllWeapons(req, res) {
    const weapons = await pg`select * from weapons`
    res.send(wewapons);
}

async function getWeaponByQuery(query) {
    const queryweps = await pg`select * from weapons where `
    res.send(queryweps);
}

// Check whether the query requested by the user is safe to pass from sql injection
function sanitizeId(req, res, next) {
    next();
}

// Check whether the query requested by the user is safe to pass from sql injection
function sanitizeQuery(req, res, next) {
    next();
}

module.exports = {
    sanitizeId,
    sanitizeQuery,
    getWeaponByName,
    getAllWeapons,
    getWeaponByQuery
};