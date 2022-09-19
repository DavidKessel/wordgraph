package graph

import (
	"strings"
)

/*
 * WordGraph is an undirected unweighted graph
 * that will hold our results.
 */
type WordGraph struct {
	numOfNodes  int // unique words
	wordcount   int //total words
	words       map[string]int
	connections map[string]map[string]int
	result      map[string]Result
}

/*
 * Result stores the analysis for each word.
 * @TODO add wordcount
 */
type Result struct {
	word         string
	dgCentrality float32
}

/*
 * Constructor for the WordGraph.
 * This is necesary to initialize the map.
 */
func NewWordGraph() *WordGraph {
	var wg WordGraph
	wg.numOfNodes = 0
	wg.wordcount = 0
	wg.words = make(map[string]int)
	wg.connections = make(map[string]map[string]int)
	wg.result = make(map[string]Result)
	return &wg
}

/*
 * Getter method for the WordGraph connections.
 */
func (wg *WordGraph) GetConnections() map[string]map[string]int {
	return wg.connections
}

/*
 * Getter method to get the number of nodes.
 */
func (wg *WordGraph) GetNumOfNodes() int {
	return wg.numOfNodes
}

/*
 * Getter method to get the results.
 */
func (wg *WordGraph) GetResults() map[string]Result {
	return wg.result
}

/*
 * Getter method for the degree centrality.
 */
func (r *Result) GetDegreeCentrality() float32 {
	return r.dgCentrality
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
				wg.makeConnections(strings.ToLower(word), strings.ToLower(sent[i]))
			}
		}
		wg.wordcount++
	}
}

/*
 * Method implements a set to store all related keywords.
 */
func (wg *WordGraph) makeConnections(key string, word string) {
	if _, ok := wg.connections[key]; ok {
		set := wg.connections[key]
		set[word]++
		wg.connections[key] = set
	} else {
		m := make(map[string]int)
		m[word] = 1
		wg.connections[key] = m
		wg.numOfNodes++
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
				stack = append(stack, neighbour)
				numOfPaths++
			}
		}
		dgCentrality := float32(numOfPaths) / float32(wg.numOfNodes) //calculate degree centrality.
		makeResult(wg, node, dgCentrality)
	}
}

/*
 * Method creates a result struct and appends it to our
 * WordGraph.
 */
func makeResult(wg *WordGraph, word string, numOfPaths float32) Result {
	var r Result
	r.word = word
	r.dgCentrality = numOfPaths
	wg.result[word] = r
	return r
}
