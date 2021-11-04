test:
	go test ./... -v

update-quotes:
	go run cmd/quotes-updater/main.go

run: docker-run
	go run cmd/gratitude-journal/main.go

docker-run:
	docker start (docker ps -a -q --filter="ancestor=mysql/mysql-server")
	# docker run -p 3307:3306  -e MYSQL_ROOT_PASSWORD=admin -d mysql/mysql-server