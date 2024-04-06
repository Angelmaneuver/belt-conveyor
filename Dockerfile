FROM    golang:bookworm

WORKDIR /usr/local/src

RUN     apt update -y && apt install -y \
           sudo                         \
           unzip                        \
           build-essential              \
        && apt clean                    \
        && rm -rf /var/lib/apt/lists/*

RUN     go mod init temporary                                                             && \
        go get -u -d gocv.io/x/gocv                                                       && \
        find /go/pkg/mod/cache/download/gocv.io/x/gocv/@v -name "*.zip" -exec cp {} . ';' && \
        unzip ./*.zip                                                                     && \
        cd gocv.io/x/                                                                     && \
        find . -name "gocv@v*" -type d -exec ln -s {} thedir ';'                          && \
        cd thedir                                                                         && \
        make install                                                                      && \
        rm -rf /usr/local/src/*
