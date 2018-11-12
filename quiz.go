package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"os"
)


func quiz() (correct, total int, err error) {
	filename := flag.String("file", "problems.csv", "CSV file for questions, solutions")
	flag.Parse()
	file, err := os.Open(*filename)
	if err != nil {
		return 0,0, fmt.Errorf("problem opening %q", *filename)
	}
	defer file.Close()

	r := csv.NewReader(file)
	questions, err := r.ReadAll()
	if err != nil {
		return 0,0, fmt.Errorf("problem reading %q", *filename)
	}

	for _, l := range questions{
		buf := bufio.NewReader(os.Stdin)
		fmt.Printf("%s ", l[0])
		a, p, err := buf.ReadLine()
		if p {
			continue
		}
		if err != nil {
			fmt.Println(err)
			continue
		}
		if string(a) == l[1]{
			correct++
		}
	}
	return correct,len(questions), nil
}

func main() {
	correct, total, err := quiz()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%d out of %d correct\n", correct, total)
}