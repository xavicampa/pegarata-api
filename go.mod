module github.com/xavicampa/xavic-test

go 1.13

require (
	github.com/gorilla/mux v1.7.3
	github.com/xavicampa/xavic-test/myapi v0.0.0-00010101000000-000000000000
)

replace github.com/xavicampa/xavic-test/myapi => ./myapi
