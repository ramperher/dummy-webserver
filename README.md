# dummy-webserver

Dummy webserver written in Golang, to practice a little bit with the language

Launch it with the following command:

```
$ go run main.go
# logs should be printed
2024/04/26 17:45:07 Configure webserver
2024/04/26 17:45:07 Start webserver
```



And test it with the following command in a separate terminal:

```
$ curl 127.0.0.1:8080/hello
# following log should be printed in the webserver execution terminal
Query received in /hello
```
