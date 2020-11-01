# uselessd

A [uselessd](https://github.com/golang/go/wiki/GitHubCodeLayout) application

```
mkdir -p uselessd
cd uselessd

go mod init github.com/victorskl/go-tute/uselessd
go get golang.org/x/net/websocket
go get github.com/victorskl/go-tute/useless
tree -L 2 $GOPATH/pkg/mod/github.com/victorskl
tree -L 2 $GOPATH/src/github.com/victorskl

vi uselessd.go

go install

tree $GOPATH/bin
which uselessd
uselessd

wscat -o ws://127.0.0.1 -c ws://127.0.0.1:1234/useless
websocat --origin ws://127.0.0.1 ws://127.0.0.1:1234/useless
```
