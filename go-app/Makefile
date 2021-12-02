test:
	go test ./... -v

update-quotes:
	go run cmd/quotes-updater/main.go

run:
	go run cmd/gratitude-journal/main.go -version $(shell cat VERSION) -commit-hash $(shell git describe --always)

release-major:
	go run cmd/version/main.go -vi major
release-minor:
	go run cmd/version/main.go -vi minor
release-patch:
	go run cmd/version/main.go -vi patch

docker-run:
	docker start (docker ps -a -q --filter="ancestor=mysql/mysql-server")
	# docker run -p 3307:3306  -e MYSQL_ROOT_PASSWORD=admin -d mysql/mysql-server