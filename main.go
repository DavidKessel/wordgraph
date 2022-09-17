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
	if seperator != "," {
		r.Comma = ';'
	}

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

	tmp := [][]string{{"Word, Deg. Centrality"}}
	for word, centrality := range graph.GetDegCentrality() {
		row := []string{word, fmt.Sprintf("%d", centrality)}
		tmp = append(tmp, row)
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
