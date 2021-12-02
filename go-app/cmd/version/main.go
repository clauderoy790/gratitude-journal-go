package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"runtime"

	"github.com/coreos/go-semver/semver"
)

func main() {
	verPtr := flag.String("vi", "", "version increment")
	flag.Parse()
	_, filename, _, _ := runtime.Caller(0)
	filename = path.Join(filename, "../../../VERSION")

	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	v, err := semver.NewVersion(string(bytes))
	if err != nil {
		panic("invalid version: " + string(bytes))
	}

	switch *verPtr {
	case "major":
		v.BumpMajor()
	case "minor":
		v.BumpMinor()
	case "patch":
		v.BumpPatch()
	default:
		panic("invalid vi flag: " + *verPtr)
	}
	err = os.WriteFile(filename, []byte(v.String()), 0644)
	if err != nil {
		panic("failed to update version: " + err.Error())
	} else {
		fmt.Println("successfully updated version to: ", v)
	}
}
