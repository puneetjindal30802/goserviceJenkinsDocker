FROM jenkins/jenkins

USER root

# FROM golang
# gcc for cgo
RUN apt-get update && apt-get install -y --no-install-recommends \
		g++ \
		gcc \
		libc6-dev \
		make \
	&& rm -rf /var/lib/apt/lists/*

ENV GOLANG_VERSION 1.13.5
ENV GOLANG_DOWNLOAD_URL https://golang.org/dl/go$GOLANG_VERSION.linux-amd64.tar.gz
ENV GOLANG_DOWNLOAD_SHA1 97f9ec90d54f3a580789f1f855b17282e7dbccb69a44b20a20c2167e907db800

RUN curl -fsSL "$GOLANG_DOWNLOAD_URL" -o golang.tar.gz \
	# && echo "$GOLANG_DOWNLOAD_SHA1  golang.tar.gz" | sha1sum -c - \
	&& tar -C /usr/local -xzf golang.tar.gz \
	&& rm golang.tar.gz

ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH

RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"

# install the ec2 cli tools
RUN apt-get update && \
    apt-get install -yq --no-install-recommends  awscli groff-base python-pip && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

RUN pip install --upgrade setuptools \
  && pip install awsebcli

USER jenkins