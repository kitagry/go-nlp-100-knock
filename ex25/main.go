package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
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

	scr := strings.NewReader(article.Text)
	sc = bufio.NewScanner(scr)

	text := ""
	result := make(map[string]string)
	start := false
	for sc.Scan() {
		t := sc.Text()
		// 基礎情報が出てきたら次から読み込みスタートする
		if strings.Contains(t, "基礎情報") {
			start = true
			continue
		}

		if !start {
			continue
		}

		// }}で終了の合図
		if t == "}}" {
			break
		}

		if t[0] == '|' {
			if len(text) > 0 {
				strs := strings.Split(text[1:], " = ")
				result[strs[0]] = strs[1]
			}
			text = t
		} else {
			text += t
		}
	}

	for key, content := range result {
		fmt.Printf("%s: %s\n", key, content)
	}
}
