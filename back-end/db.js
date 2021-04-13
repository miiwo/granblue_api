const postgres = require('postgres');
var dotenv = require('dotenv');

dotenv.config();
const sql = postgres();

module.exports = sql;
