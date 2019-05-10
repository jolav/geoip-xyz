/* */

const path = require('path');

const lib = require(path.join(__dirname, '_lib/lib.js'));
const p = require(path.join(__dirname, '_private.js'));

const TABLE = p.mysql.table3;

const mysql = require('mysql');
const con = mysql.createConnection({
  host: p.mysql.host,
  user: p.mysql.user,
  password: p.mysql.password,
  database: p.mysql.db3
});

function updateStats(req, res, c, next) {
  const test = lib.getIP(req);
  if (test) {
    const time = new Date().toISOString().split('T')[0];
    if (c.app.mode === 'production') {
      console.log(test + " " + req.originalUrl);
      insertHit(time);
    } else {
      console.log('SAVE TEST ...');
    }
  } else {
    console.log(test, 'DONT SAVE => ');
  }
  next();
}

function insertHit(time) {
  let sql = 'INSERT INTO ?? (time, geoip)';
  sql += ' VALUES (?, 0)';
  sql += ` ON DUPLICATE KEY UPDATE geoip = geoip + 1;`;
  const inserts = [TABLE, time];
  sql = mysql.format(sql, inserts);
  con.query(sql, function (err, rows) {
    if (err) {
      console.log('Insert HIT error =>', err);
      // throw err
    } else {
      //console.log(rows);
    }
  });
}

function keepConnectionAlive() {
  let sql = 'SELECT 1';
  sql = mysql.format(sql);
  con.query(sql, function (err, rows) {
    if (err) {
      console.log('Error Keeping connection Alive =>', err);
      // throw err
    } else {
      //console.log('Keep connection Alive =>', rows);
    }
  });
}

function testDB() {
  // console.log(con)
  con.connect(function (err) {
    if (err) {
      console.log('Error connecting to DB => ', err);
    } else {
      console.log('Connection to DB ...... OK');
    }
  });
}

module.exports = {
  updateStats: updateStats,
  keepConnectionAlive: keepConnectionAlive,
  testDB: testDB
};

