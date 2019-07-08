package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("../neko.txt.mecab")
	if err != nil {
		log.Println(err)
		return
	}
	defer file.Close()

	sc := bufio.NewScanner(file)

	type ads struct {
		Surface string
		Base    string
		Pos     string
		Pos1    string
	}

	lists := make([]*ads, 0)

	for sc.Scan() {
		t := sc.Text()
		if t == "EOS" {
			continue
		}

		ts := strings.Split(t, "	")
		surface := ts[0]
		adjs := strings.Split(ts[1], ",")

		lists = append(lists, &ads{
			Surface: surface,
			Base:    adjs[6],
			Pos:     adjs[0],
			Pos1:    adjs[1],
		})
	}

	out := bufio.NewWriter(os.Stdout)
	for i, ad := range lists {
		if ad.Surface == "の" && ad.Pos == "助詞" && lists[i-1].Pos == "名詞" && lists[i+1].Pos == "名詞" {
			out.WriteString(lists[i-1].Base + ad.Surface + lists[i+1].Base + "\n")
		}
	}
	out.Flush()
}
