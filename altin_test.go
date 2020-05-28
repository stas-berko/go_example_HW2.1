package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"runtime"
	"strings"
	"testing"
)


func assert(o bool) {
	if !o {
		fmt.Println(__getRecentLine())
		os.Exit(1)
	}
}
func __getRecentLine() string {
	_, file, line, _ := runtime.Caller(2)
	buf, _ := ioutil.ReadFile(file)
	code := strings.TrimSpace(strings.Split(string(buf), "\n")[line-1])
	return fmt.Sprintf("%v:%d\n%s", path.Base(file), line, code)
}

func TestPigLatin(t *testing.T) {

	assert(pigLatin("ball") == "allbay")
	assert(pigLatin("button") == "uttonbay")
	assert(pigLatin("hello") == "ellohay")
	assert(pigLatin("smile") == "ilesmay")


}
