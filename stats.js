/* */

const path = require('path');
const http = require("http");
const lib = require(path.join(__dirname, '_lib/lib.js'));

function updateStats(req, res, c, next) {
  const test = lib.getIP(req);
  if (test) {
    if (c.app.mode === 'production') {
      console.log(test + " " + req.originalUrl);
      sendHit("geoip");
    } else {
      console.log('SAVE TEST ...');
    }
  } else {
    console.log(test, 'DONT SAVE => ');
  }
  next();
}

function sendHit(service) {
  let path = "http://localhost:3970/addhit/" + service;
  makeHttpRequest(path, function (err, res, data) {
    if (err) {
      console.error('Error with the request:', err.message);
    }
  });
}

function makeHttpRequest(path, callback) {
  http.get(path, function (resp) {
    let data = '';
    resp.on('data', function (chunk) {
      data += chunk;
    });
    resp.on('end', function () {
      try {
        var parsed = JSON.parse(data);
      } catch (err) {
        //console.error('Unable to parse response as JSON', err);
        return callback(err, null, null);
      }
      callback(null, resp, parsed);
    });
  }).on('error', (err) => {
    //console.error('Error with the request:', err.message);
    callback(err, null, null);
  });
}

module.exports = {
  updateStats: updateStats,
};

