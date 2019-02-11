# Kubernetes Pod Load Balancer Demo

[#56 - Kubernetes General Explanation (screencast)](https://sysadmincasts.com/episodes/56-kubernetes-general-explanation)

### Usage

[Live Demo](https://sysadmindemo.com/)

##### Build go app for linux (from mac)

```sh
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-w' -o web ./main.go
```

##### Build/push container

```sh
docker build -t jweissig/alpine-k8s-pod-lb-demo .
docker push jweissig/alpine-k8s-pod-lb-demo
```

##### Run container

https://hub.docker.com/r/jweissig/alpine-k8s-pod-lb-demo

```sh
docker pull jweissig/alpine-k8s-pod-lb-demo
docker run -it -p 5005:5005 --rm alpine-k8s-pod-lb-demo
```

##### Connect to localhost

http://localhost:5005/
