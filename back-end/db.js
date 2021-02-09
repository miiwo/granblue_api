const postgres = require('postgres');
var dotenv = require('dotenv');

dotenv.config();

// Change this so I don't have my password revealed.
//const sql = postgres('postgres://latka:humR1nger@localhost:5432/granblue_api');
const sql = postgres();

module.exports = sql;
