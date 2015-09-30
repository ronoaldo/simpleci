# Copyright 2014 Ronoaldo JLP <ronoaldo@gmail.com>
# Licensed under the Apache License, Version 2.0

FROM debian:jessie
MAINTAINER Ronoaldo JLP <ronoaldo@gmail.com>

# Install the basic tools for common build scenarios
RUN apt-get update && apt-get -y dist-upgrade &&\
    DEBIAN_FRONTEND=noninteractive apt-get -y --no-install-recommends install \
    subversion mercurial git \
    build-essential devscripts \
    curl wget \
    npm npm2deb nodejs nodejs-legacy \
    openjdk-7-jdk ant \
    python-dev python3-dev \
    ruby ruby-dev ruby-compass

# Install updated Maven 3
RUN MAVEN_BIN="http://ftp.unicamp.br/pub/apache/maven/maven-3/3.3.3/binaries/apache-maven-3.3.3-bin.tar.gz" ;\
    mkdir -p /opt/maven ; \
    curl -s $MAVEN_BIN | tar -xzf - -C /opt/maven ; \
    ln -s /opt/maven/apache-maven-3.3.3/bin/mvnDebug /usr/local/bin/mvnDebug ; \
    ln -s /opt/maven/apache-maven-3.3.3/bin/mvn /usr/local/bin/mvn

# Install front-end development tools
RUN npm -g install grunt-cli@0.1.13 gulp-cli@0.3.0 bower@1.5.3 less@2.5.3

# Install updated Go toolchain
RUN GO_BIN="https://storage.googleapis.com/golang/go1.5.1.linux-amd64.tar.gz" ; \
    curl $GO_BIN | tar -C /usr/local -xzf -

# Environment settings
RUN echo -n 'export PATH=$PATH:/usr/local/go/bin' >> /etc/profile
ENV PATH   $PATH:/usr/local/go/bin
ENV GOPATH /var/tmp/go

RUN go get ronoaldo.gopkg.net/simpleci/cmd/buildbot ; \
    go build -o /usr/local/bin/buildbot ronoaldo.gopkg.net/simpleci/cmd/buildbot

# Main entry point
CMD /usr/bin/env buildbot