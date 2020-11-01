curl -XPOST -d '{"query":"query test {\n person(firstName:\"Pedro\") {\n lastName\n }\n}\n","variables":null,"operationName":"test"}' localhost:8080/query


curl -XPOST -d '{"query":"query test {\n person(firstName:\"John\") {\n firstName\n lastName\n }\n}\n","variables":null,"operationName":"test"}' localhost:8080/query