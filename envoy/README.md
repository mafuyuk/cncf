# Envoy
https://www.envoyproxy.io/

```
$ docker build -t cncf/envoy:v1 .
$ docker run -d -p 8080:10000 -p 9901:9901 cncf/envoy:v1
```