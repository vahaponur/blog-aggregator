buildrun:
	go build -o ./bin/blog-aggregator && ./bin/blog-aggregator



dbupdate:
	cd sql/schema && goose postgres postgres://vahap:@localhost:5432/blogagg up && cd ...
buildlinux:
	GOOS=linux GOARCH=amd64 go build -o ./bin/blog-aggregator && ./bin/blog-aggregator

doall:
	go install github.com/pressly/goose/cmd/goose@latest && \
	cd sql/schema && goose postgres postgres://vahap:@docker.for.mac.localhost:5432/blogagg?sslmode=disable up
	@echo ok
	GOOS=linux GOARCH=amd64 go build -o ./bin/blog-aggregator && ./bin/blog-aggregator
