### CURL запросы для проверки Server

- GET
```
    Request: curl -X GET "http://localhost:9000?OzonLearning" 
    Expected answer: <empty>
    Status code: 200 OK
    Server output: Get, query params: [map[OzonLearning:[]]]

    Request: curl -X GET "http://localhost:9000?Hello=Ozon" 
    Expected answer: <empty>
    Status code: 200 OK
    Server output: Get, query params: [map[Hello:[Ozon]]]
```
- POST
```
    Request: curl -X POST "http://localhost:9000" -H "hw-sum: 10"
    Expected answer: <empty>
    Status code: 200 OK
    Server output: Create, headers:[map[Accept:[*/*] Hw-Sum:[10] User-Agent:[curl/7.81.0]]]
                   Sum: [15]

    Request: curl -X POST "http://localhost:9000" -H "hw-sum: abracadabra" 
    Expected answer: <empty>
    Status code: 400 Bad Request
    Server output: Create, headers:[map[Accept:[*/*] Hw-Sum:[abracadabra] User-Agent:[curl/7.81.0]]]
                   hm-sum parameter should be integer value, [abracadabra] was received
```
- DELETE
```
    Request: curl -X DELETE "http://localhost:9000" 
    Expected answer: <empty>
    Status code: 200 OK
    Server output: Delete
```

- PUT
```
    Request: curl -X PUT "http://localhost:9000" -d "Hello, Ozon" 
    Expected answer: <empty>
    Status code: 200 OK
    Server output: Update, body: [Hello, Ozon]
```

### CURL запросы для проверки ServerWithData

- GET
```
    Request: curl -v -X GET "http://localhost:9001?id=1"
    If id=1 exists in the table with some value:
        Expected answer: value
        Status code: 200 OK
        Server output: <empty>
    Else:
        Expected answer: <empty>
        Status code: 404 Not found
        Server output: There is no entry in the data for this Id [1]
    
    Request: curl -v -X GET "http://localhost:9001?id=invalidValue"
    Expected answer: <empty>
    Status code: 404 Not found
    Server output: Id [invalidValue] does not match type uint32

    Request: curl -v -X GET "http://localhost:9001?someParam=someValue"
    Expected answer: <empty>
    Status code: 400 Bad Request
    Server output: Request parameter "id" is missing
```
- POST
```
    Request: curl -v -X POST "http://localhost:9001" -d '{ "id": 1, "value": "example" }'
    If id=1 exists in the table with some value:
        Expected answer: <empty>
        Status code: 409 Conflict
        Server output: This id=[1] already exists
    Else:
        Expected answer: <empty>
        Status code: 200 OK
        Server output: <empty>
    
    Request: curl -v -X POST "http://localhost:9001" -d '{ "id": -100, "value": -100 }'
    Expected answer: <empty>
    Status code: 500 Internal Server Error
    Server output: Error while unmarshalling request body, err: [json: cannot unmarshal number -100 into Go struct field request.id of type uint32]

    Request: curl -v -X POST "http://localhost:9001" -d '{ "id": 1 }'
    Expected answer: <empty>
    Status code: 400 Bad Request
    Server output: Json file format error
```
- DELETE
```
    Request: curl -v -X DELETE "http://localhost:9001?id=1"
    If id=1 exists in the table with some value:
        Expected answer: <empty>
        Status code: 200 OK
        Server output: <empty>
    Else:
        Expected answer: <empty>
        Status code: 404 Not Found
        Server output: There is no entry in the data for this Id [1]

    Request: curl -v -X DELETE "http://localhost:9001?id=invalidValue"
    Expected answer: <empty>
    Status code: 404 Not Found
    Server output: Id [invalidValue] does not match type uint32

    Request: curl -v -X DELETE "http://localhost:9001?someParam=someValue"
    Expected answer: <empty>
    Status code: 400 Bad Request
    Server output: Request parameter "id" is missing
```

- PUT
```
    Request: curl -v -X PUT "http://localhost:9001" -d '{ "id": 1, "value": "example22222" }'
    If id=1 exists in the table with some value:
        Expected answer: <empty>
        Status code: 200 OK
        Server output: <empty>
    Else:
        Expected answer: <empty>
        Status code: 404 Not Found
        Server output: There is no entry in the data for this Id [1]

    Request: curl -v -X PUT "http://localhost:9001" -d '{ "id": 1 }'
    Expected answer: <empty>
    Status code: 400 Bad Request
    Server output: Json file format error

    Request: curl -v -X PUT "http://localhost:9001" -d '{ "id": -100, "value": "example22222" }'
    Expected answer: <empty>
    Status code: 500 Internal Server Error
    Server output: Error while unmarshalling request body, err: [json: cannot unmarshal number -100 into Go struct field request.id of type uint32]
```