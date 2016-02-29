FROM ubuntu:14.04

MAINTAINER Chris Aubuchon <Chris.Aubuchon@gmail.com>

COPY . /go/src/github.com/CiscoCloud/mesos-consul

RUN apt-get update -y
RUN apt-get install -y golang git mercurial
RUN cd /go/src/github.com/CiscoCloud/mesos-consul \
	&& export GOPATH=/go \
	&& go get \
	&& go build -o /bin/mesos-consul \
	&& rm -rf /go \
	&& apk del --purge go git mercurial


ENTRYPOINT [ "/bin/mesos-consul" ]
