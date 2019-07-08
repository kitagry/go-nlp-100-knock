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
	for i := 0; i < len(lists); i++ {
		ads := lists[i]
		if ads.Pos != "名詞" {
			continue
		}

		j := i + 1
		current := lists[j]
		for current.Pos == "名詞" {
			j++
			current = lists[j]
		}

		if j > i+1 {
			for k := i; k < j; k++ {
				out.WriteString(lists[k].Surface)
			}
			out.WriteString("\n")
		}

		i = j
	}
}
