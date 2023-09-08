buildrun:
	go build -o ./bin/blog-aggregator && ./bin/blog-aggregator
dbupdate:
	cd sql/schema && goose postgres postgres://vahap:@localhost:5432/blogagg up
