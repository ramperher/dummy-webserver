# dummy-webserver

Dummy webserver written in Golang, to practice a little bit with the language.

## Spec

The webserver exposes one endpoint, `/hello`, using [Chi routing framework](https://go-chi.io/#/). This is defined in the [routes](pkg/routes) package. Then, the main executable, called [webserver.go](cmd/webserver.go), makes use of it.

The `/hello` endpoint expects an HTTP request with the following JSON structure:

```
{
    "name": "...",
    "msg": "..."
}
```

Where `name` and `msg` fields must be String values.

The endpoint will behave in the following way depending on the received input:

- If no JSON is received, an HTTP 400 message will be returned.
- If JSON is received but it's incomplete, an HTTP 404 message will be returned.
- If no JSON is received and it's complete, an HTTP 200 message will be returned.

## Usage

Launch it with the following command:

```
$ go run cmd/webserver.go
2024/05/21 17:55:17 Set endpoints on webserver
2024/05/21 17:55:17 Start webserver
```

And test it with the following command in a separate terminal:

```
$ curl localhost:8080/hello
{"status":"error found: EOF"}
$ curl -X GET -d '{"nam":"bob"}' localhost:8080/hello
{"status":"not found"}
$ curl -X GET -d '{"msg":"hi all"}' localhost:8080/hello
{"status":"not found"}
$ curl -X GET -d '{"name":"bob","msg":"hi all"}' localhost:8080/hello
{"status":"correct"}
```

Following logs would be printed in the terminal where the application is running:

```
Query received in /hello
Error in request
2024/05/21 17:55:32 "GET http://localhost:8080/hello HTTP/1.1" from [::1]:38430 - 400 30B in 71.11µs
Query received in /hello
Received: {"name": "", "msg": ""}
Not found case
2024/05/21 17:55:39 "GET http://localhost:8080/hello HTTP/1.1" from [::1]:41164 - 404 23B in 66.688µs
Query received in /hello
Received: {"name": "", "msg": "hi all"}
Not found case
2024/05/21 17:55:46 "GET http://localhost:8080/hello HTTP/1.1" from [::1]:59690 - 404 23B in 43.775µs
Query received in /hello
Received: {"name": "bob", "msg": "hi all"}
Correct case
2024/05/21 17:55:50 "GET http://localhost:8080/hello HTTP/1.1" from [::1]:59702 - 200 21B in 49.764µs
```
