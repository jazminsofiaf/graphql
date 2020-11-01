curl -XPOST -d '{"query":"query test {\n person(firstName:\"Pedro\") {\n lastName\n }\n}\n","variables":null,"operationName":"test"}' localhost:8080/query


curl -XPOST -d '{"query":"query test {\n movie(title:\"The Big New Movie\") {\n title\n year\n plot\n rating\n }\n}\n","variables":null,"operationName":"test"}' localhost:8080/query