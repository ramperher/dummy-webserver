module example.com/webserver

go 1.21.5

replace example.com/routes => ../routes

require example.com/routes v0.0.0-00010101000000-000000000000

require github.com/go-chi/chi/v5 v5.0.12 // indirect
