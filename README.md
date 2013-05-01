dcms
====

```bash
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

```

* sitemap (couples pages to URLs)
* template
* pages (couples template to document)
* channel (couples sitemap to tree and document version)
* tree (contains documents)
* document
* document-type (property of document)
* document-version