dcms
====

set GOPATH="d:\Git\Projects\dcms"

go get github.com/alleveenstra/godb
go get github.com/gorilla/mux
go get github.com/zuitek/mymysql/autorc
go get github.com/zuitek/mymysql/mysql
go get github.com/zuitek/mymysql/native

go install resty
go run dcms.go

go build dcms.go
./dcms.exe

