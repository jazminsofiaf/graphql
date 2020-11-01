curl -XPOST -d '{"query":"query test {\n person(id:\"1000\") {\n id\n firstName\n }\n}\n","variables":null,"operationName":"test"}' localhost:8080/query
