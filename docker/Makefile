.PHONY: reader_docker writer_docker up

up: reader_docker writer_docker
	docker-compose up

reader_docker: bin/reader
	docker build -t distributed_course_docker:reader -f reader/Dockerfile .

writer_docker: bin/writer
	docker build -t distributed_course_docker:writer -f writer/Dockerfile .

bin/reader: reader/*.go
	go build -o bin/reader ./reader

bin/writer: writer/*.go
	go build -o bin/writer ./writer
