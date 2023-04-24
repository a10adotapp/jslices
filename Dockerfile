FROM ubuntu:22.04

RUN apt update && apt upgrade -y

RUN apt install -y \
    autoconf \
    curl \
    gcc \
    gettext \
    jq \
    libcurl4-gnutls-dev \
    libexpat1-dev \
    libssl-dev \
    libz-dev \
    libpcre3-dev \
    make

RUN curl -L https://mirrors.edge.kernel.org/pub/software/scm/git/git-2.39.2.tar.gz | tar -C /usr/local/src -zx \
 && cd /usr/local/src/git-2.39.2 \
 && make configure \
 && ./configure \
 && make \
 && make install \
 && cd ~ \
 && rm -rf /usr/local/src/git-2.39.2

RUN curl -L https://go.dev/dl/go1.20.3.linux-amd64.tar.gz | tar -C /usr/local -zx

RUN go install github.com/cweill/gotests/gotests@latest
RUN go install github.com/fatih/gomodifytags@latest
RUN go install github.com/go-delve/delve/cmd/dlv@latest
RUN go install github.com/haya14busa/goplay/cmd/goplay@latest
RUN go install github.com/josharian/impl@latest
RUN go install github.com/ramya-rao-a/go-outline@latest
RUN go install golang.org/x/tools/gopls@latest
RUN go install honnef.co/go/tools/cmd/staticcheck@latest

ENV PATH ${PATH}:/usr/local/go/bin
ENV PATH ${PATH}:/root/go/bin

ENTRYPOINT ["tail"]
CMD ["-f", "/dev/null"]
