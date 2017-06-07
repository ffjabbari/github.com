# go-echoserver

A simple echo server to test a golang server in a docker container.

Once you have the docker-machine running, and golang installed,
you can simply build and run the server:

```bash
$ make run
```

With that server up and running, you can curl to it:

```bash
$ curl -XGET $(docker-machine ip):8080/hello
```

Or you can curl to it from a docker container:

```bash
$ docker run -t --rm --link go-echo:go-echo appropriate/curl go-echo:8080/hello
```
