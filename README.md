API GRAPH-QL
======================

Este repo tiene una api GraphQL.<br>
Seguramente conozcas lo que es una api-Rest, bueno GraphQL es otro diseño de arquitectura de apis. 
GraphQL define una sintaxis para obtener datos. 
Al igual que otros paradigmas se debe definir la interfaz con la que se accedera a los datos.
Luego de pasar esa capa inicial, debera tener una capa de seguridad.
Por debajo la logica de negocio y por ultimo la capa de persistencia. 
El usuario (que seguramente sea otro desarrollador) no necesita saber nada sobre las capaz inferiores, 
el solo va a interactuar con la interfaz. 
Esto permite que podamos modificar las capaz inferiores sin molestar al usuario. 
Por ejemplo podiramos migrar todos los datos desde postgres a dynamo y el usuario no deberia enterarse 
ni tener que hacer ninguna modificacion en su logica para obtener los datos.<br>
La idea de graphQL es representar las relaciones entre las entidades como un grafo. 
Por ejemplo podemos tener la entiedad pelicula y la entidad director y pensar a cada una como un nodo del grafo. 
Un director puede estar conectado con multiples peliculas. 
Se modela el dominio de negocio utilizando un esquema. 
Mas info en la [documentacion oficial](https://graphql.org/learn/thinking-in-graphs/) 

Pero lo que me parecio super util de GraphQL es que es super flexible. 
Si tenemos una entidad con millones de fields, pero solo estamos interesados en uno, 
podemos crear un request pidiendo solo eso, y no nos van a venir datos extras, entonces la respuesta queda super prolija.
Por ejemplo si solo quiero saber el mail de los usuarios la respuesta sera una lista de mails y nada mas.  

<br>Se puede deployar en un lambda de aws y conectarlo con una Apigateway en modo proxy o 
<br>Se puede probar local, conectandose a aws usando las credenciales default ubicadas en el archivo .aws/credentials

Supongamos que en Dynamo tenemos la siguiente tabla Movies

| Title | Year | Plot | Rating |
| ----- | ---- | ---- | ------ |
| The Big New Movie  | 2015  | Se trata de una chica que no hace nada  | 4.5 |
| The Big New Movie  | 2016  | Es una pelicula vieja renovada          | 7.0 |
| La sirenita        | 1994  | Se trata de una sirena                  | 8.1 |


#### Pasos probarlo local
1. clonar el repo 
2. buildearlo
```go build movies-graphql-aws-lambda/src/main.go```
3. levantar el server local en el puerto 8080
```go run movies-graphql-aws-lambda/src/main.go```

##### Ejemplo 1 de request
curl -XPOST -d '{"query":"query test {\n movie(title:\"The Big New Movie\"\n) {\n title\n year\n plot\n }\n}\n","variables":null,"operationName":"test"}' localhost:8080/query

request:
```
query test {
    movie(title: "The Big New Movie") {
        title
        year 
        plot 
    }
}
```
response
```
{
    data: {
        movie: {
            "title": "The Big New Movie",
            "year": 2016,
            "plot": "Es una pelicula vieja renovada"
        }
    }
}
```

##### Ejemplo 2 de request
curl -XPOST -d '{"query":"query test {\n movie(title:\"The Big New Movie\" rating:4.5\n) {\n title\n year\n plot\n rating\n }\n}\n","variables":null,"operationName":"test"}' localhost:8080/query

request:
```
query test {
    movie(title: "The Big New Movie"   rating=4.5) {
        title
        rating
        year 
        plot 
    }
}
```

response:
```
{
    data: {
        movie: {
            "title": "The Big New Movie",
            "year": 2015,
            "plot": "Se trata de una chica que no hace nada",
            "rating": 4.5
        }
    }
}
```
