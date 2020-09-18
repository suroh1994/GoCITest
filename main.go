package main

import (
	"fmt"
	"log"

	"github.com/blang/semver"
	"github.com/rhysd/go-github-selfupdate/selfupdate"
)

var version = "0.0.0"

func main() {
	Greet()
	doSelfUpdate()
}

func funnyFunction() string {
	return "Smile! :)"
}

func doSelfUpdate() {
	v := semver.MustParse(version)
	latest, err := selfupdate.UpdateSelf(v, "suroh1994/GoCITest")
	if err != nil {
		log.Println("Binary update failed:", err)
		return
	}
	if latest.Version.Equals(v) {
		// latest version is the same as current version. It means current binary is up to date.
		log.Println("Current binary is the latest version", version)
	} else {
		log.Println("Successfully updated to version", latest.Version)
		log.Println("Release note:\n", latest.ReleaseNotes)
	}
}

func Greet() {
	fmt.Printf("hello user, you are running version %s\n", version)

	if true {
		return
	}
	fmt.Println("never printed")
}
