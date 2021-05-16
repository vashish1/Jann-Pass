# Jann-Pass
    
    The situation of India is alarming, due to massive increase in Covid19 cases, Indian government has announced lockdown And if we have to fight this pandemic successfylly this must be extended, but along with the Extension of Lockdown, the problems of common people will only increase.

The problem is people are not understanding the deapth       of the issue and roaming around casually. Also people do need their daily rashans. 

There need to be a system that allows limited number of people during lockdown to move out of their home.

#### Jann-Pass is an E-Pass Portal, Which will help in limiting publc movement.

## Features of Jann-Pass

    •	E-pass will be issued in three time slots for each day during lockdown i.e.
       - 11:00 am - 01:00 pm
       - 02:00 pm - 04:00 pm
       - 05:00 pm - 07:00 pm

    •	For every week  limited passes, eg:50, will be available.

    •	Once pass is issued, same person cannot issue another for that day.

    •	One person allowed per pass and pass is issued with Aadhaar number.

    •	 Police will have alternate login path.

    •	They will have scanner built in their page which will verify the pass.

### Front End:

    The Project has its templates and all the static source ready.
    Due to modifications the API's has not been Integrated in the Frontend yet. 

### Back End
  Golang is used as a primary language
  
  MongoDB is used for Database

#### Version of API :  V1.0
    
# API Documentation

 All responses come in standard JSON. All requests must include a `content-type` of `application/json` and the body must be valid JSON.


### Response Codes
```
200: Success
201: Created
400: Bad request
401: Unauthorized
404: Cannot be found
405: Method not allowed
50x: Server Error
```
### Error and Success Message Example

```json
  {
    "error":"message" 
  }
  
  {
      "success":"message", //subjective
      "any-data-type": "data-sent-in response"  //w.r.t the API
  }
```

## SignUp

**You send:**  You send the details required to signup.

**You get:** An `Error-Message` or a `Success-Message` depending on the status of the account created.

**Endpoint:** 
     /signup

**Authorization Token:** Not required

**Request:**
`POST HTTP/1.1`

```json
Accept: application/json
Content-Type: application/json
Content-Length: xy

{   
    "name": "abc",
    "email": "foo",
    "aadhar": 9876567898,
    "password": "1234567",
}
```

**Successful Response:**
```json
HTTP/1.1 200 OK
Content-Type: application/json
Content-Length: xy

{
   "success": true,
}
```

## Login

**You send:**  Your  login credentials.

**You get:** An `API-Token` and a `Success-Message` with which you can make further actions.

**Endpoint:** 
     /login

**Authorization Token:** Not required

**Request:**
`POST HTTP/1.1`
```json
Accept: application/json
Content-Type: application/json
Content-Length: xy

{
    "email": "foo",
    "password": "1234567" 
}
```

**Successful Response:**
```json
HTTP/1.1 200 OK
Content-Type: application/json
Content-Length: xy

{
   "success":"Log In successful",
   "token": "e3b0...................",
}
```
## Police Login

**You send:**  unique user ID.

**You get:** An `API-Token` and a `Success-Message` with which you can make further actions.

**Endpoint:** 
     /login/police

**Authorization Token:** Not required

**Request:**
`POST HTTP/1.1`
```json
Accept: application/json
Content-Type: application/json
Content-Length: xy

{
    "id":6568 
}
```


**Successful Response:**
```json
HTTP/1.1 200 OK
Content-Type: application/json
Content-Length: xy

{
   "success":"Log In successful",
   "token": "e3b0...................",
}
```

## Get Epass

**You send:**  You send the details required related to Epass such as slot, date, time e.t.c.

**You get:** An `Error-Message` or a png image of QR code in Response

**Endpoint:** 
     /epass

**Authorization Token:** User's token required

**Request:**
`GET HTTP/1.1`

```json
Accept: application/json
Content-Type: application/json
Content-Length: xy

{   
    "slot":"5:00-7:00",
    "date":"16-05-2021",
    "area":"aliganj",
    "area_code":467009
}
```

**Successful Response:**
```json
HTTP/1.1 200 OK
Content-Type: image/png
Content-Length: xy

```
Here's a sample image:

<img title="QR Code" src="sample/qr_code_Sample.png" >

## Check Epass

**You send:**  You send the string encoded into the QR code.

**You get:** An `Error-Message` or a `Success-Message` with which you can make further actions.

**Endpoint:** 
     /epass

**Authorization Token:** Police's auth token required

**Request:**
`GET HTTP/1.1`

```json
Accept: application/json
Content-Type: application/json
Content-Length: xy

{   
    "base_64_string":"zsxdfcgvfcxswefgvcsdfg="
}
```

**Successful Response:**
```json
HTTP/1.1 200 OK
Content-Type: application/json
Content-Length: xy

{
    "success":true,
    "error": "nil",
}
```

## Environment variables

  
  `DbUrl` : to store the url of mongodb cluster or localhost.

  `Secret Key` : to store encryption secret.

  `PORT`     : port on localhost to run application

