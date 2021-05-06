# proglog

### How to run?
1. Fork a docker container - `docker run -it -p 8080:8080 -v ~/:/my_host golang:1.13 /bin/bash`
1. Navigate to your git repo inside the container.
1. Then navigate to main file location - `cd cmd/server/`
1. Run `go run main.server`
1. From your host machine or docker, you can access the service at port `8080`