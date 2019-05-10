module.exports = {
  apps: [{
    name: 'geoip',
    script: 'geoip.js',
    output: './../logs/hits.log',
    error: './../logs/errors.log',
    instances: 2,
    merge_logs: true,
    log_date_format: 'YYYY-MM-DD HH:mm:ss',
  }]
};