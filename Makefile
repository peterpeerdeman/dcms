all:
	go install storage
	go install resty
	go install mysite
	go install site
	go build src/dcms.go

fetch:
	go get github.com/gorilla/mux
	go get github.com/howeyc/fsnotify
	go get code.google.com/p/go.text/unicode/norm
	go get github.com/extemporalgenome/slug

clean:
	rm data/*
	touch data/empty.txt

test:
	./dcms &
	sleep 5
	curl -f -XGET http://localhost:8080/rest/document
	curl -f -XGET http://localhost:8080/rest/document-type
	killall dcms