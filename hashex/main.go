package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"os"
	"strings"
)

const codePipelineGitHubTokenHashPrefix = "hash-"

func hashCodePipelineGitHubToken(token string) string {
	// Without this check, the value was getting encoded twice
	if strings.HasPrefix(token, codePipelineGitHubTokenHashPrefix) {
		return token
	}
	sum := sha256.Sum256([]byte(token))
	return codePipelineGitHubTokenHashPrefix + hex.EncodeToString(sum[:])
}

func main() {

	// https://gobyexample.com/command-line-arguments
	// fmt.Printf("%T\n", os.Args)

	// https://gobyexample.com/arrays
	// fmt.Println(len(os.Args))

	// https://gobyexample.com/if-else

	if len(os.Args) > 1 {
		arg := os.Args[1]
		// fmt.Println(arg)
		fmt.Println(hashCodePipelineGitHubToken(arg))
	} else {
		fmt.Println(hashCodePipelineGitHubToken(""))
	}
}
