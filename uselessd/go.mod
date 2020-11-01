module github.com/victorskl/go-tute/uselessd

go 1.15

require (
	github.com/victorskl/go-tute/useless v0.0.0-20201029113428-8d7b932ab236
	golang.org/x/net v0.0.0-20201029055024-942e2f445f3c
)

// https://github.com/golang/go/wiki/Modules#when-should-i-use-the-replace-directive
// Just like gradle dependency substitute!
replace github.com/victorskl/go-tute/useless => ../useless
