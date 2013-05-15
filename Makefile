all:
	go install mysite
	go install site
	go install resty
	go build src/dcms.go

test:
	./dcms &
	sleep 5
	curl -f -XGET http://localhost:8080/rest/document
	curl -f -XGET http://localhost:8080/rest/document-type
	killall dcms