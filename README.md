# go_simple_rest

based on https://blog.codeship.com/building-minimal-docker-containers-for-go-applications/

```

docker build -t golang_rest -f Dockerfile .

docker run -it  -p 8080:8080 golang_rest

curl http://localhost:8080/info


docker build -t example-scratch -f Dockerfile.scratch .

docker build -t example-multistage -f Dockerfile.multistage .

docker run -it example-multistage

```


regards Rosso0815