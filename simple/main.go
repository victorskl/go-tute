package main // <-- package main follow by func main makes program entrypoint

import "fmt" // <-- import from std lib

func main() { // <-- func main
	fmt.Println("Simple World")
}

// build statically link all imports to produce a single executable
// go build use directory name of package main as executable output i.e. 'simple' in this case
// go build .
// ./simple
// otherwise define build output
// go build -o simple.exe
// ./simple.exe
