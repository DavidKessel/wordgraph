package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"gokw/graph"
	"io"
	"os"
)

func main() {
	/*
	 * Declaring variables for the cli arguments
	 */
	var (
		inPath, outPath, seperator string
		verbose                    bool
		colIndex                   int
	)

	/*
	 * Defining the cli flags
	 */
	flag.StringVar(&inPath, "in", "./data.csv", "path to the input csv file.")
	flag.StringVar(&outPath, "out", "./result.csv", "path to the output csv file.")
	flag.StringVar(&seperator, "s", ",", "the separator in your csv file.")
	flag.BoolVar(&verbose, "v", false, "verbose: prints output to terminal.")
	flag.IntVar(&colIndex, "col", 0, "index of the column in your csv.")
	flag.Parse()

	/*
	 * Reading the input csv based on cli arguments
	 * @TODO add logic to pass multiple separators.
	 */
	data, err := os.Open(inPath)
	checkError(err)
	r := csv.NewReader(data)
	r.Comma = []rune(seperator)[0]

	graph := graph.NewWordGraph() //declare a new wordgraph to store results

	/*
	 * Read each row and add all words as nodes
	 */
	for {
		cell, err := r.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
		}
		var sentence string = cell[colIndex]
		graph.AddNodes(sentence)
	}

	graph.DFS() //traverse the graph

	/*
	 * now writing output to file
	 */
	csvFile, err := os.Create(outPath)
	checkError(err)
	writer := csv.NewWriter(csvFile)

	/*
	 * looping over results and getting the degree centrality
	 */
	tmp := [][]string{{"Word, Deg. Centrality"}}
	for key, result := range graph.GetResults() {
		dg := fmt.Sprintf("%.2f", result.GetDegreeCentrality())
		row := []string{key, dg}
		tmp = append(tmp, row)
	}

	if verbose {
		fmt.Println(tmp[0])
		for word, result := range graph.GetResults() {
			fmt.Printf("%s %v\n", word, result.GetDegreeCentrality())
		}
	}

	writer.WriteAll(tmp)
	writer.Flush()
	csvFile.Close()
	fmt.Println("🔥 DONE: wrote file to", outPath)
}

func checkError(err error) {
	if err != nil {
		fmt.Println("💣💥 Something went wrong!")
		fmt.Println(err)
		os.Exit(1)
	}
}
