
![Version](https://img.shields.io/badge/version-0.1.2-orange.svg)  
![Maintained YES](https://img.shields.io/badge/Maintained%3F-YES-green.svg)  
![Ask Me Anything !](https://img.shields.io/badge/Ask%20me-anything-1abc9c.svg)  

# ![logo](https://github.com/jolav/geoip-xyz/blob/master/www/_public/icons/ip48.png?raw=true) **[GEOIP.XYZ](https://geoip.xyz)** 

version 0.1.2

- Free service that provides a public secure API (CORS enabled) to retrieve geolocation from any IP or hostname.  
- 10 request per second. Once reached subsequent requests will result in error 429 until your quota is cleared.  
- This API requires no key or signup.  
- JSON and XML supported
- IPv4 and IPv6 supported  
- CORS support out of the box makes this perfect to your front end apps or webs  


Examples

https://geoip.xyz/v1/json  
https://geoip.xyz/v1/json?q=codetabs.com  
https://geoip.xyz/v1/xml?q=8.8.8.8  
https://geoip.xyz/v1/xml?q=2a00:1450:4006:803::200e  

Response JSON :

```json
{   
  "ip": "172.168.90.240",
  "country_code": "FR",
  "country_name": "France",
  "region_code": "IDF",
  "region_name": "Ile-de-France",
  "city": "Paris",
  "zip_code": "75001",
  "time_zone": "Europe/Paris",
  "latitude": 48.8628,
  "longitude": 2.3292   
}
```

<hr>



## **Acknowledgment**


* This site includes GeoLite2 data created by MaxMind, available from  [maxmind.com](http://maxmind.com)