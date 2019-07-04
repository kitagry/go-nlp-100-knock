package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type Article struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

func main() {
	file, err := os.Open("../jawiki-country.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	sc := bufio.NewScanner(file)
	buf := make([]byte, 10000)
	sc.Buffer(buf, 1000000)

	var article Article
	for sc.Scan() {
		var art Article
		if err := json.Unmarshal(sc.Bytes(), &art); err != nil {
			fmt.Println(err)
			return
		}
		if art.Title == "イギリス" {
			article = art
		}
	}

	scReader := strings.NewReader(article.Text)
	sc = bufio.NewScanner(scReader)

	re := regexp.MustCompile("^=+.*?=+$")
	for sc.Scan() {
		t := sc.Text()
		if t = re.FindString(t); t != "" {
			level := 0
			for t[level] == '=' {
				level++
			}
			fmt.Printf("level: %d, keyword: %s\n", level, t[level:len(t)-level])
		}
	}
}
