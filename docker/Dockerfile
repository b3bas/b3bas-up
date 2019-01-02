ARG GOLANG_VERSION=1.11
FROM golang:${GOLANG_VERSION}

# ================================================================================================
#  Inspiration: Docker Framework (https://github.com/zeroc0d3/docker-framework)
#               Dwi Fahni Denni <zeroc0d3.0912@gmail.com>
# ================================================================================================

MAINTAINER "ZeroC0D3Lab <zeroc0d3.team@gmail.com>"

ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH

RUN apt-get install -y bash git

RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"

# Install dep
RUN go get -u github.com/golang/dep/cmd/dep -v

WORKDIR $GOPATH

CMD ["go", "version"]

EXPOSE 22
