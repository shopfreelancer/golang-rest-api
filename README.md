## Rest API with Golang

A RESTful API with Golang. In contrast to setups with Node.js or PHP no Reverse Proxy like Nginx is needed.

The app starts at webserver at localhost:8080. You can test the routes either with Postman or straight in the console with curl

## Use the routes

Index - list articles
`curl localhost:8080/articles`

CreateArticle - Create a new article resource
`curl -X POST -H "Content-Type: application/json" -d '{"title":"asdasd"}' localhost:8080/article`

ShowArticle show one article resource
`curl localhost:8080/article/5c5328c8bb4a78b2b9d55dde`

DeleteArticle  - delete one article by object id
`curl -X DELETE -H "Content-Type: application/json" -d '{"title":"asdasd"}' localhost:8080/article/5c5328c8bb4a78b2b9d55dde`


## Development
- Create your .env file: `mv .env.example .env`
- Enter your mongodb credentials in the .env file
- Run the app `go run main.go`

## Production
### Build your binary
`go build -o restApiGo`
`./restApiGo`