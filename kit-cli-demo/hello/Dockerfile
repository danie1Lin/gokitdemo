FROM golang

RUN mkdir -p /go/src//Users/daniellin/gokitdemo/kit-cli-demo

ADD . /go/src//Users/daniellin/gokitdemo/kit-cli-demo

RUN go get  -t -v ./...
RUN go get  github.com/canthefason/go-watcher
RUN go install github.com/canthefason/go-watcher/cmd/watcher

ENTRYPOINT  watcher -run /Users/daniellin/gokitdemo/kit-cli-demo/hello/cmd  -watch /Users/daniellin/gokitdemo/kit-cli-demo/hello
