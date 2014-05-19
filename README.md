Web-service-using-Go
====================

Create RESTful web service to maintain simple to-dos using Go 


README
***Created by Priyanka Chavan (priyankachavan88@gmail.com) ****
Compiled and tested successfully.
To-do tasks implemetaion using Go :
Steps followed for implementation:
Install Go
Install curl
Install TortoiseHg
Install Git
Tutorials for golang and Json parsing (unmarshal method in GO)and why RESTful web service.
Write Get and Post methods
Use "store" to store Json objects received grom GET methods.
Used a package from github for a framework.Referred at the end of this file.
Run and test using curl commands as follows.

Output:
Curl commands used to test:
Create method :

C:\>curl -i -d "{\"taskdesc\":\"Call\",\"date\":\"02/12/2014\",\"completed\":\"true\"}" http://127.0.0.1:8080/task/add
HTTP/1.1 200 OK
Content-Type: application/json
X-Powered-By: go-json-rest
Date: Thu, 24 Apr 2014 00:09:14 GMT
Content-Length: 71
{
  "Taskdesc": "Call",
  "Date": "02/12/2014",
  "Completed": "true"
}

Read method : (with task Desc as id)

C:\>curl -i http://127.0.0.1:8080/task/Call
HTTP/1.1 200 OK
Content-Type: application/json
X-Powered-By: go-json-rest
Date: Thu, 24 Apr 2014 00:10:43 GMT
Content-Length: 71

{
  "Taskdesc": "Call",
  "Date": "02/12/2014",
  "Completed": "true"
}


Read method : (Read all)

C:\>curl -i http://127.0.0.1:8080/task/list
HTTP/1.1 200 OK
Content-Type: application/json
X-Powered-By: go-json-rest
Date: Thu, 24 Apr 2014 00:14:22 GMT
Content-Length: 168

[
  {
    "Taskdesc": "pay",
    "Date": "12/12/2014",
    "Completed": "true"
  },
  {
    "Taskdesc": "Call",
    "Date": "02/11/2014",
    "Completed": "false"
  }
]

Update method:(Read all to check the effect of update)

C:\>curl -i -d "{\"taskdesc\":\"Meet\",\"date\":\"03/06/2014\",\"completed\":\"t
rue\"}" http://127.0.0.1:8080/task/pay
HTTP/1.1 200 OK
Content-Type: application/json
X-Powered-By: go-json-rest
Date: Thu, 24 Apr 2014 00:16:04 GMT
Content-Length: 71

{
  "Taskdesc": "Meet",
  "Date": "03/06/2014",
  "Completed": "true"
}
C:\>curl -i http://127.0.0.1:8080/task/list
HTTP/1.1 200 OK
Content-Type: application/json
X-Powered-By: go-json-rest
Date: Thu, 24 Apr 2014 00:16:08 GMT
Content-Length: 169

[
  {
    "Taskdesc": "Meet",
    "Date": "03/06/2014",
    "Completed": "true"
  },
  {
    "Taskdesc": "Call",
    "Date": "02/11/2014",
    "Completed": "false"
  }
]
Delete method:

C:\>curl -i -X DELETE http://127.0.0.1:8080/task/Call
HTTP/1.1 200 OK
Date: Thu, 24 Apr 2014 00:17:13 GMT
Content-Length: 0
Content-Type: text/plain; charset=utf-8


C:\>curl -i http://127.0.0.1:8080/task/list
HTTP/1.1 200 OK
Content-Type: application/json
X-Powered-By: go-json-rest
Date: Thu, 24 Apr 2014 00:17:17 GMT
Content-Length: 85

[
  {
    "Taskdesc": "Meet",
    "Date": "03/06/2014",
    "Completed": "true"
  }
]

Issues faced: 
Passing JSON object to unmarshal method via curl command(resolved using StackOverflow.com)
Data types of Date and boolean in Go
Test cases like empty value in JSON object ,update whole object for "update" method.

References:
https://gowalker.org/github.com/ant0ine/go-json-rest
http://golang.org/ref/spec
