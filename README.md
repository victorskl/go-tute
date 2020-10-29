# go-tute

Let Go!

- First, install:
```
brew search go
brew info go
brew install go
brew upgrade go

which go
go version
```

- Second, read:
```
go help
go help gopath
go help packages
go help modules
```

- Third, get-started:
```
go env
go env GOROOT
go env GOPATH
```

- Simple World
```
mkdir simple
cd simple
vi main.go
go build .
./simple
go build -o simple.exe
./simple.exe
```

- [Code Organization](https://golang.org/doc/code.html#Organization)
```
package < module < repository
```

- [Module Example](https://golang.org/doc/code.html#Command)

```
mkdir hello
cd hello
go mod init github.com/victorskl/go-tute/hello
vi hello.go

go install github.com/victorskl/go-tute/hello
go install .
go install

tree $GOPATH/bin
which hello
hello
```

- [Local Package](https://golang.org/doc/code.html#ImportingLocal)

```
mkdir morestrings
cd morestrings
vi morestrings
go build
cd ..
go install
hello
```

- [Remote Package](https://golang.org/doc/code.html#ImportingRemote)

```
go get github.com/google/go-cmp/cmp
vi hello.go
go install
hello

go help go.mod
```

- [Testing](https://golang.org/doc/code.html#Testing)
```
cd morestrings
vi reverse_test.go
go test

go help test
```

## Real World

- Developing a `hashex` cli
```
mkdir hashex
cd hashex
go mod init github.com/victorskl/go-tute/hashex
vi main.go
golint
go install
hashex
hashex
hashex me
hashex me
```

- Useless World
```
mkdir useless
go mod init github.com/victorskl/go-tute/useless
vi useless.go
go build github.com/victorskl/go-tute/useless
```
