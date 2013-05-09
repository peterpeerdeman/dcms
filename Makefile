all:
	go install mysite
	go install site
	go install resty
	go build src/dcms.go