package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"net/http"
	// "net/http/cookiejar"
	// "net/url"
	// "os"
	// "errors"
)

func main() {

	// jar, _ := cookiejar.New(nil)

	// client := http.Client{
	// 	Jar: jar,
	// }

	// resp, err := client.PostForm("https://www.xcontest.org/2024/world/en/", loginDetails)
	// if err != nil {
	// 	fmt.Println("Error logging in:", err)
	// 	return
	// }
	// defer resp.Body.Close()

	// if resp.StatusCode != http.StatusOK {
	// 	fmt.Println("Login failed with status:", resp.Status)
	// 	return
	// }
	// fmt.Println("Logged in successfully!")

	resp2, err := http.Get("https://www.xcontest.org/world/en/flights-search/")

	resp2.Header.Set("cookie", "8d86ad15e2da572369602591985c19e2")

	if err != nil {
		fmt.Println("Oops an error occured! :(")
	}
	defer resp2.Body.Close()
	body := bufio.NewScanner(resp2.Body)

	oupfile, _ := os.Create("temp.txt")

	oup := bufio.NewWriter(oupfile)
	start := 0

	for body.Scan() {

		line := body.Text()

		if strings.Contains(line, "<table") {
			start = 1
		}

		if start == 1 {
			oup.Write([]byte(line))
			oup.Write([]byte("\n"))

		}
		if strings.Contains(line, "</table") {
			start = 0
		}

	}

}

// https://www.xcontest.org/2024/world/en/flights/detail:fialka/29.09.2024/1

// func takeHref(line string) string{

// 	startIndex := strings.Index(line, "href=")
// 	if startIndex == -1{
// 		return ""
// 	}

// 	endIndex := strings.Index(line[startIndex:], "\"")

// }
