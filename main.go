package main

import (
	"flag"
	"fmt"
	"github.com/jszwec/csvutil"
	"math/rand"
	"os"
	"time"
)

type Participant struct {
	Nome      string `csv:"Nome"`
	Sobrenome string `csv:"Sobrenome"`
	Email     string `csv:"Email"`
}

func main() {
	num := flag.Int("num", 3, "Quantos numeros sortear?")
	file := flag.String("file", "participantes.csv", "Qual o arquivo com os nomes?")
	flag.Parse()
	dat, err := os.ReadFile(*file)
	if err != nil {
		panic(err)
	}
	var p []Participant
	if err := csvutil.Unmarshal(dat, &p); err != nil {
		fmt.Println("error:", err)
	}

	n := getRandomNumbers(*num, 0, len(p))
	for _, i := range n {
		fmt.Printf("Número sorteado: %d Nome:%s %s\n", i, p[i].Nome, p[i].Sobrenome)
	}

}

func getRandomNumbers(qtd, min, max int) []int {
	var r []int
	for i := 0; i < qtd; i++ {
		rand.Seed(time.Now().UnixNano())
		r = append(r, rand.Intn(max-min+1)+min)
	}
	return r
}
