package graph

import (
	"fmt"
	"strings"
)

/*
 * WordGraph is an undirected unweighted graph
 * that will hold our results.
 */
type WordGraph struct {
	numOfNodes        int
	words             map[string]int
	connections       map[string]map[string]bool
	degreeCentrallity map[string]int
	result            []Result
}

/*
 * Result stores the analysis for each word.
 * @TODO add wordcount
 */
type Result struct {
	word              string
	degreeCentrallity int
	count             int
}

/*
 * Constructor for the WordGraph.
 * This is necesary to initialize the map.
 */
func NewWordGraph() *WordGraph {
	var wg WordGraph
	wg.numOfNodes = 0
	wg.words = make(map[string]int)
	wg.connections = make(map[string]map[string]bool)
	wg.degreeCentrallity = make(map[string]int)
	return &wg
}

/*
 * Getter method for the WordGraph connections.
 */
func (wg *WordGraph) GetConnections() map[string]map[string]bool {
	return wg.connections
}

/*
 * Getter method for the degree centrality.
 */
func (wg *WordGraph) GetDegCentrality() map[string]int {
	return wg.degreeCentrallity
}

/*
 * Getter method to get the number of nodes.
 */
func (wg *WordGraph) GetNumOfNodes() int {
	return wg.numOfNodes
}

/*
 * Connect each word with all other words
 * in the same sentence. @TODO creat O(n) algorithm.
 */
func (wg *WordGraph) AddNodes(sentence string) {
	sent := strings.Fields(sentence)
	for ind, word := range sent {
		for i := 0; i < len(sent); i++ { //we want to exclude the current word
			if i != ind {
				wg.makeConnections(word, sent[i])
			}
		}
		wg.numOfNodes++
	}
}

/*
 * Method implements a set to store all related keywords
 */
func (wg *WordGraph) makeConnections(key string, word string) {
	if _, ok := wg.connections[key]; ok {
		set := wg.connections[key]
		set[word] = true
		wg.connections[key] = set
	} else {
		m := make(map[string]bool)
		m[word] = true
		wg.connections[key] = m
	}
}

/*
 * Method traverses the graph with DFS
 * and analyzes the results.
 */
func (wg *WordGraph) DFS() {
	for node := range wg.GetConnections() {
		numOfPaths := 0
		var stack []string = []string{node}
		for _, current := range stack { //visit all the neighbours
			for neighbour := range wg.connections[current] {
				numOfPaths += 1
				startStack := append(stack, neighbour)
				fmt.Sprintln(startStack)
			}
		}
		wg.degreeCentrallity[node] = numOfPaths //number of paths
		makeResult(wg, node, numOfPaths, 1)
	}
}

/*
 * Method creates a result struct and appends it to our
 * WordGraph.
 */
func makeResult(wg *WordGraph, word string, numOfPaths int, count int) {
	var r Result
	r.word = word
	r.degreeCentrallity = numOfPaths
	r.count = 1
	wg.result = append(wg.result, r)
}
