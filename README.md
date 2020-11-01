# go-tute

Assorted Go tutes!

- First, install:
```
brew search go
brew info go
brew install go
brew upgrade go

which go
go version
```

- Optionally, set:
```
export GOPATH="${HOME}/go"
export GOROOT="$(brew --prefix golang)/libexec"
export PATH="$PATH:${GOPATH}/bin:${GOROOT}/bin"

go get golang.org/x/lint/golint
which golint
golint -h
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
[source.go, code.go] < package < module < repository
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

## Notes

- [Go Module](https://blog.golang.org/using-go-modules) `go.mod` take care of dependencies.
- Go codes are organised in **packages** which then arranged in a **module**. 
- A Go module is published in a **repository**.
- A repository can contain [multi-modules](https://github.com/golang/go/wiki/Modules#faqs--multi-module-repositories). 🙋‍♂️ _This `go-tute` is a multi-modules Go repo, for example!_
- Go executable entrypoint bears a package name "**main**" and `func main()`.
- Go build use parent directory name as an executable name, if `-o` flag is absent.
- Go standard community proposes some [project structure](https://github.com/golang-standards/project-layout) and naming convention. This structure is known to be [a traditional monolith structure](https://tutorialedge.net/golang/go-project-structure-best-practices/). 
- You can, however, arrange your Go project whichever way you like! 
- Though, `go.mod` favours _fan-out_ multiple repositories; with each repository contain only one Go module; and `go.mod` take care of inter-dependencies between these Go modules.
