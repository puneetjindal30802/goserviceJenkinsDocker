FROM jenkins/jenkins:lts-alpine

USER root

RUN apk add --no-cache \
		ca-certificates

# set up nsswitch.conf for Go's "netgo" implementation
# - https://github.com/golang/go/blob/go1.9.1/src/net/conf.go#L194-L275
# - docker run --rm debian:stretch grep '^hosts:' /etc/nsswitch.conf
RUN [ ! -e /etc/nsswitch.conf ] && echo 'hosts: files dns' > /etc/nsswitch.conf

ENV GOLANG_VERSION 1.13.5

RUN set -eux; \
	apk add --no-cache --virtual .build-deps \
		bash \
		gcc \
		musl-dev \
		openssl \
		go \
	; \
	export \
# set GOROOT_BOOTSTRAP such that we can actually build Go
		GOROOT_BOOTSTRAP="$(go env GOROOT)" \
# ... and set "cross-building" related vars to the installed system's values so that we create a build targeting the proper arch
# (for example, if our build host is GOARCH=amd64, but our build env/image is GOARCH=386, our build needs GOARCH=386)
		GOOS="$(go env GOOS)" \
		GOARCH="$(go env GOARCH)" \
		GOHOSTOS="$(go env GOHOSTOS)" \
		GOHOSTARCH="$(go env GOHOSTARCH)" \
	; \
# also explicitly set GO386 and GOARM if appropriate
# https://github.com/docker-library/golang/issues/184
	apkArch="$(apk --print-arch)"; \
	case "$apkArch" in \
		armhf) export GOARM='6' ;; \
		x86) export GO386='387' ;; \
	esac; \
	\
	wget -O go.tgz "https://golang.org/dl/go$GOLANG_VERSION.src.tar.gz"; \
	echo '27d356e2a0b30d9983b60a788cf225da5f914066b37a6b4f69d457ba55a626ff *go.tgz' | sha256sum -c -; \
	tar -C /usr/local -xzf go.tgz; \
	rm go.tgz; \
	\
	cd /usr/local/go/src; \
	./make.bash; \
	\
	rm -rf \
# https://github.com/golang/go/blob/0b30cf534a03618162d3015c8705dd2231e34703/src/cmd/dist/buildtool.go#L121-L125
		/usr/local/go/pkg/bootstrap \
# https://golang.org/cl/82095
# https://github.com/golang/build/blob/e3fe1605c30f6a3fd136b561569933312ede8782/cmd/release/releaselet.go#L56
		/usr/local/go/pkg/obj \
	; \
	apk del .build-deps; \
	\
	export PATH="/usr/local/go/bin:$PATH"; \
	go version

ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH
ENV CGO_ENABLED=0

RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"
WORKDIR $GOPATH

ADD repo-key /
RUN cd src && \
  chmod 600 /repo-key && \
  echo "IdentityFile /repo-key" >> /etc/ssh/ssh_config && \
  echo -e "StrictHostKeyChecking no" >> /etc/ssh/ssh_config && \
  git clone git@github.com:puneetjindal30802/goserviceJenkinsDocker.git && \
  cd goserviceJenkinsDocker && \
  go get && \
  go build

# RUN echo 'http://dl-cdn.alpinelinux.org/alpine/v3.6/main' >> /etc/apk/repositories
# RUN echo 'http://dl-cdn.alpinelinux.org/alpine/v3.6/community' >> /etc/apk/repositories
# RUN apk update
# RUN apk add mongodb=3.4.4-r0

# RUN mongo --version

# USER root

# RUN apk update \
# 		g++ \
# 		gcc \
# 		libc6-dev \
# 		make \
# 	&& rm -rf /var/lib/apt/lists/*

# RUN apk update && \
#     apk upgrade && \
#     apk add git

# RUN apk update && \
#     apk add --no-cache openssh

# ENV GOLANG_VERSION 1.13.5
# ENV GOLANG_DOWNLOAD_URL https://golang.org/dl/go$GOLANG_VERSION.linux-amd64.tar.gz
# ENV GOLANG_DOWNLOAD_SHA1 97f9ec90d54f3a580789f1f855b17282e7dbccb69a44b20a20c2167e907db800

# RUN curl -fsSL "$GOLANG_DOWNLOAD_URL" -o golang.tar.gz \
# 	# && echo "$GOLANG_DOWNLOAD_SHA1  golang.tar.gz" | sha1sum -c - \
# 	&& tar -C /usr/local -xzf golang.tar.gz \
# 	&& rm golang.tar.gz && \
# 	go

# ENV GOPATH /go
# ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH

# ADD repo-key /

# RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"
# RUN go
