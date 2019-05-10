/* */

const express = require('express');
const app = express();

const path = require('path');

const job = require(path.join(__dirname, 'geoipTask.js'));
const c = require(path.join(__dirname, '_config.js'));
const stats = require(path.join(__dirname, 'stats.js'));

if (c.app.mode !== 'production') {
  c.app.port = 3000;
}

app.disable('x-powered-by');

app.use(function (req, res, next) {
  stats.updateStats(req, res, c, next);
});

app.get('/v1/json/', function (req, res) {
  job.getGeoData(req, res, 'json');
});

app.get('/v1/xml/', function (req, res) {
  job.getGeoData(req, res, 'xml');
});

app.get('*', function (req, res) {
  if (c.app.mode === "production") {
    res.redirect('https://geoip.xyz/notFound');
  } else {
    res.status(404).send('Not Found');
  }
});

app.listen(c.app.port, function () {
  const time = new Date().toUTCString().split(',')[1];
  console.log('Express server on port ' + c.app.port + ' - ' + time);
  initApp();
});

function initApp() {
  stats.testDB();
  setInterval(function () {
    stats.keepConnectionAlive();
  }, 5000);
}

module.exports = app;
