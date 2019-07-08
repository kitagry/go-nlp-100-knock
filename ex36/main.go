package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
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

	maps := make(map[string]int)
	for _, ad := range lists {
		if ad.Base == "*" {
			continue
		}

		if _, ok := maps[ad.Base]; !ok {
			maps[ad.Base] = 0
		}
		maps[ad.Base] += 1
	}

	type res struct {
		Text  string
		Count int
	}

	result := make([]res, len(maps))
	i := 0
	for key, value := range maps {
		result[i] = res{Text: key, Count: value}
		i++
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].Count > result[j].Count
	})
	fmt.Println(result)

	out := bufio.NewWriter(os.Stdout)
	for _, item := range result {
		out.WriteString(item.Text + "\n")
	}
	out.Flush()
}
