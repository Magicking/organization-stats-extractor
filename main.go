// Command screenshot is a chromedp example demonstrating how to take a
// screenshot of a specific element.
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"

	"github.com/chromedp/chromedp"
)

func screenshotRepository(username, repo string) {
	// create context
	ctxt, cancel := context.WithCancel(context.Background())
	defer cancel()

	// create chrome instance
	c, err := chromedp.New(ctxt, chromedp.WithLog(log.Printf))
	if err != nil {
		log.Fatal(err)
	}

	// run task list
	var activityMaster, activityDetail []byte
	err = c.Run(ctxt, screenshot(`https://github.com/`+username+`/`+repo+`/graphs/commit-activity`, `#commit-activity-master`, `#commit-activity-detail`, &activityMaster, &activityDetail))
	if err != nil {
		log.Fatal(err)
	}

	// shutdown chrome
	err = c.Shutdown(ctxt)
	if err != nil {
		log.Fatal(err)
	}

	// wait for chrome to finish
	err = c.Wait()
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile(repo+"activityMaster.png", activityMaster, 0644)
	if err != nil {
		log.Fatal(err)
	}
	err = ioutil.WriteFile(repo+"activityDetail.png", activityDetail, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {

	// Open our jsonFile
	jsonFile, err := os.Open("repo_url.json")

	// if we os.Open returns an error then handle it
	if err != nil {
		log.Fatal(err)
	}
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var result []string
	json.Unmarshal([]byte(byteValue), &result)

	for _, fullName := range result {
		fullNameSplit := strings.Split(fullName, "/")
		if len(fullNameSplit) != 2 {
			log.Fatalf("Expected Org/repo, got %v", fullName)
		}
		org, repo := fullNameSplit[0], fullNameSplit[1]
		fmt.Println(org, repo)
		screenshotRepository(org, repo)
	}
}

func screenshot(urlstr, sel, sel2 string, res, res2 *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(urlstr),
		chromedp.Sleep(2 * time.Second),
		chromedp.WaitVisible(sel, chromedp.ByID),
		chromedp.WaitVisible(sel2, chromedp.ByID),
		//		chromedp.WaitNotVisible(`div.v-middle > div.la-ball-clip-rotate`, chromedp.ByQuery),
		chromedp.Screenshot(sel, res, chromedp.NodeVisible, chromedp.ByID),
		chromedp.Screenshot(sel2, res2, chromedp.NodeVisible, chromedp.ByID),
		chromedp.Stop(),
	}
}
