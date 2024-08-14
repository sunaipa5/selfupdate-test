package main

import (
	"fmt"
	"selfupdate-test/updater"
	"time"
)

func main() {
	fmt.Println("APP VERSION 0.0.3")

	updaterOptions := updater.Options{
		Author:         "sunaipa5",
		Repo:           "selfupdate-test",
		CurrentVersion: "0.0.3",
		TagEnd:         "linux_amd64.tar.gz",
	}

	updaterOptions.CheckUpdate()

	for i := 0; i < 20; i++ {
		fmt.Println(updaterOptions.CurrentVersion)
		time.Sleep(time.Second * 1)
	}
}
