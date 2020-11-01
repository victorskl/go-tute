# websrv

Getting started with REST Web Service in Go!

```
mkdir -p websrv && cd websrv
go mod init github.com/victorskl/go-tute/websrv

mkdir -p cmd/simplesrv
vi cmd/simplesrv/main.go
go run cmd/simplesrv/main.go
curl -s http://localhost:8000/hello

mkdir -p cmd/echosrv
vi cmd/echosrv/main.go
go get -u -v github.com/labstack/echo/v4
go run cmd/echosrv/main.go

curl -s http://localhost:1323 | jq
curl -s http://localhost:1323/version | jq
curl -s http://localhost:1323/users | jq

curl -s -X POST http://localhost:1323/users | jq
curl -s -X PUT -H "Content-Type: application/json" -d '{"name":"apple"}' http://localhost:1323/users/1 | jq
curl -s -X GET http://localhost:1323/users/1 | jq

curl -s -X POST -H "Content-Type: application/json" -d '{"name":"mango"}' http://localhost:1323/users | jq
curl -s -X POST -H "Content-Type: application/json" -d '{"name":"banana"}' http://localhost:1323/users | jq
curl -s -X GET http://localhost:1323/users | jq

curl -s -X DELETE http://localhost:1323/users/1 | jq
curl -s -X GET http://localhost:1323/users | jq
```
