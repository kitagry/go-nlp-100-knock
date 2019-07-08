package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"sort"
	"strings"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
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

	p, err := plot.New()
	if err != nil {
		panic(err)
	}

	font, err := vg.MakeFont("Times-Bold", 12)
	if err != nil {
		panic(err)
	}
	p.Title.Font = font

	p.Title.Text = "Chart"
	p.X.Label.Text = "Rank"
	p.Y.Label.Text = "Count"

	group := make(plotter.XYs, len(result))
	for i := 0; i < len(result); i++ {
		group[i] = plotter.XY{X: math.Log(float64(i) + 1), Y: math.Log(float64(result[i].Count))}
	}

	scatter, err := plotter.NewScatter(group)
	if err != nil {
		panic(err)
	}
	p.Add(scatter)

	if err = p.Save(5*vg.Inch, 3*vg.Inch, "scatter.png"); err != nil {
		panic(err)
	}
}
