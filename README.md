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
2024/08/28 13:32:50 INFO Set endpoints on webserver
2024/08/28 13:32:50 INFO Start webserver
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
2024/08/28 13:33:20 INFO Query received in /hello
2024/08/28 13:33:20 ERROR Error in request
2024/08/28 13:33:20 "GET http://localhost:8080/hello HTTP/1.1" from [::1]:60606 - 400 30B in 192.433µs
2024/08/28 13:33:31 INFO Query received in /hello
2024/08/28 13:33:31 INFO Received: {"name": "", "msg": ""}
2024/08/28 13:33:31 WARN Not found case
2024/08/28 13:33:31 "GET http://localhost:8080/hello HTTP/1.1" from [::1]:47524 - 404 23B in 76.623µs
2024/08/28 13:33:41 INFO Query received in /hello
2024/08/28 13:33:41 INFO Received: {"name": "", "msg": "hi all"}
2024/08/28 13:33:41 WARN Not found case
2024/08/28 13:33:41 "GET http://localhost:8080/hello HTTP/1.1" from [::1]:43470 - 404 23B in 57.894µs
2024/08/28 13:33:51 INFO Query received in /hello
2024/08/28 13:33:51 INFO Received: {"name": "bob", "msg": "hi all"}
2024/08/28 13:33:51 INFO Correct case
2024/08/28 13:33:51 "GET http://localhost:8080/hello HTTP/1.1" from [::1]:33564 - 200 21B in 69.777µs
```

## Testing

Previous tests are also launched using table-driven tests with [webserver_test.go file](cmd/webserver_test.go). Result should be the following:

```
$ go test cmd/webserver_test.go 
ok  	command-line-arguments	0.004s
```
