FROM golang:latest

RUN mkdir /usr/local/go/src/deli-reco
WORKDIR /usr/local/go/src/deli-reco

# go.mod側に記載するとdownloadしても/go/bin/配下に落ちなくコマンドが実行出来なかった
RUN go get -u github.com/beego/bee/v2

CMD bee run
