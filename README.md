# SensorReadAPI
Practice problem , testing  HTTP,JSON,REST...

### You can clone the repo using:
```
git clone https://github.com/aTTiny73/SensorReadAPI.git
```
Functionality can be tested with CURL or POSTMAN.

JSON data format is used for this demonstration.

SensorValues in JSON are respresented :
```
{
    "id":1,
    "temperature":29,
    "humidity": 20,
    "co2": 320

}
```
Main handlers:
```
http://localhost:8090/postReadings
```
```
http://localhost:8090/getReadings
```
```
http://localhost:8090/getReading?ID=3
```
```
http://localhost:8090/updateReading?ID=3
```
```
http://localhost:8090/deleteReading?ID=3
```
