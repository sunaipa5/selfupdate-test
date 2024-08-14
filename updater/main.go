package updater

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/minio/selfupdate"
)

func (options Options) CheckUpdate() {
	resp, err := http.Get("https://api.github.com/repos/" + options.Author + "/" + options.Repo + "/releases/latest")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	jsonDoc, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	var release Release
	if err := json.Unmarshal(jsonDoc, &release); err != nil {
		fmt.Println(err)
		return
	}

	if release.Version != options.CurrentVersion {
		fmt.Println("New version available!")
	} else {
		fmt.Println("App up to date")
		return
	}

	var source Source
	count := 0
	for _, asset := range release.Assets {
		if count > 0 {
			source.Name += " "
			source.Download_Url += " "
		}
		if strings.HasSuffix(asset.Name, options.TagEnd) {
			count++
			source.Name += asset.Name
			source.Download_Url += asset.Download_Url
		}
	}

	if count > 1 {
		fmt.Println("Multiple source found! please change 'TagEnd' :")
		fmt.Println(source)
		return
	} else if count < 1 {
		fmt.Println("Not found any source! plese change 'TagEnd'")
		return
	}

	if err := installUpdate(source.Download_Url); err != nil {
		fmt.Println(err)
		return
	}

	releaseJson, err := json.Marshal(release)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(releaseJson))

}

func installUpdate(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	err = selfupdate.Apply(resp.Body, selfupdate.Options{})
	if err != nil {
		return err
	}
	return err
}
