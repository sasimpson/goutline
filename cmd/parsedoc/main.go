package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/sasimpson/goutline"
)

func main() {
	file, err := os.Open("example.gd")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	var nodes []*goutline.Node
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		node, err := goutline.ParseLine(scanner.Text())
		if node != nil {
			fmt.Printf("%#v\n", node)
			nodes = append(nodes, node)
		} else {
			fmt.Printf("%s", err.Error())
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
