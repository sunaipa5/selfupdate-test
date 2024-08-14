package main

import "selfupdate-test/updater"

func main() {
	updaterOptions := updater.Options{
		Author:         "sunaipa5",
		Repo:           "soundark",
		CurrentVersion: "0.9",
		TagEnd:         "executable.zip",
	}

	updaterOptions.CheckUpdate()
}
