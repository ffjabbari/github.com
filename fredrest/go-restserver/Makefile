default: docker

docker: echo-linux
	docker build -t ffjabbari/go-rest:1.0 -t ffjabbari/go-rest .

echo-linux:
	env GOOS=linux GOARCH=amd64 go build -v \
		-o bin/rest \
		github.com/fredrest/go-restserver/cmd/rest

run: docker
	docker run -t --name go-rest -p 3333:3333 ffjabbari/go-rest:1.0

curl-host:
	curl -XGET $(shell docker-machine ip):3333

curl-docker:
	docker run -t --rm --link go-rest appropriate/curl go-rest:3333
