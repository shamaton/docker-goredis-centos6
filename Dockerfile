From shamaton/golang-centos6:latest

MAINTAINER shamaton

# golang library
RUN go get github.com/garyburd/redigo/redis

# copy test code
COPY test_redis.go .

