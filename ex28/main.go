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

	strReader := strings.NewReader(article.Text)
	sc = bufio.NewScanner(strReader)

	italicBoldRegexp := regexp.MustCompile(`'''''(.*)'''''`)
	article.Text = italicBoldRegexp.ReplaceAllString(article.Text, "$1")
	boldRegexp := regexp.MustCompile(`'''(.*)'''`)
	article.Text = boldRegexp.ReplaceAllString(article.Text, "$1")
	italicsRegexp := regexp.MustCompile(`''(.*)''`)
	article.Text = italicsRegexp.ReplaceAllString(article.Text, "$1")

	outerLinkRegexp := regexp.MustCompile(`\[http.*\]`)
	article.Text = outerLinkRegexp.ReplaceAllString(article.Text, "")
	fileRegexp := regexp.MustCompile(`\[\[ファイル:.*\]\]`)
	article.Text = fileRegexp.ReplaceAllString(article.Text, "")
	categoryRegexp := regexp.MustCompile(`\[\[Category:(.*?)(\|.*)?\]\]`)
	article.Text = categoryRegexp.ReplaceAllString(article.Text, "$1")
	innerLinkRegexp := regexp.MustCompile(`\[\[([^#|]*?)(#[^|]*?)?(\|.*)?\]\]`)
	article.Text = innerLinkRegexp.ReplaceAllString(article.Text, "$1")

	commentOutRegexp := regexp.MustCompile(`<!--.*-->`)
	article.Text = commentOutRegexp.ReplaceAllString(article.Text, "")

	headingRegexp := regexp.MustCompile(`={2,5}(.*?)={2,5}`)
	article.Text = headingRegexp.ReplaceAllString(article.Text, "$1")
	bulletRegexp := regexp.MustCompile(`\*{1,2} (.*?)`)
	article.Text = bulletRegexp.ReplaceAllString(article.Text, "$1")
	numberedBulletRegexp := regexp.MustCompile(`#{1,2} (.*?)`)
	article.Text = numberedBulletRegexp.ReplaceAllString(article.Text, "$1")
	definedBulletRegexp := regexp.MustCompile(`;.*?\n:(.*?)`)
	article.Text = definedBulletRegexp.ReplaceAllString(article.Text, "$1")
	horizontalRegexp := regexp.MustCompile(`----`)
	article.Text = horizontalRegexp.ReplaceAllString(article.Text, "")
	templateRegexp := regexp.MustCompile(`{{.*}}`)
	article.Text = templateRegexp.ReplaceAllString(article.Text, "")

	fmt.Printf("%v\n", article)
}
