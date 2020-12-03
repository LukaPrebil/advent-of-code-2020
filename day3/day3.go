package day3

import (
	"bufio"
	"log"
	"os"
)

// CountTrees counts number of trees in a given path through a map
func CountTrees() int {
	topography := parseMap()
	return topography.countTrees(3, 1)
}

// MinimizeStops counts trees on a few different angles and returns the product of all counts
func MinimizeStops() int {
	topography := parseMap()
	result := 1
	for _, angle := range [][2]int{{1, 1}, {3, 1}, {5, 1}, {7, 1}, {1, 2}} {
		result *= topography.countTrees(angle[0], angle[1])
	}
	return result
}

// mapData stores some metadata, and the 2D slice representing the local topology
type mapData struct {
	width, height int
	topography    [][]byte
}

// TopoAnalyser counts met trees on a particular topography based on the angle of attack
type topoAnalyser interface {
	countTrees(deltaX, deltaY int) int
}

// Implement TopoAnalyser for mapData
func (mapData mapData) countTrees(deltaX, deltaY int) int {
	count, x, y := 0, 0, 0
	for y < mapData.height-1 {
		x = (x + deltaX) % (mapData.width) // move deltaX right and wrap around
		y += deltaY
		//log.Printf("X: %d, Y: %d, Tree: %s", x, y, string(mapData.topography[y][x]))
		if mapData.topography[y][x] == '#' {
			count++
		}
	}
	return count
}

// ------- Helpers ---------

func parseMap() mapData {
	// Open input file
	file, err := os.Open("./day3/input.txt")
	handleError(err)
	defer file.Close()

	// Setup file scanner to read lines
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	topography := make([][]byte, 0)

	// Parse each line as byte slice
	for scanner.Scan() {
		topography = append(topography, []byte(scanner.Text()))
	}

	return mapData{len(topography[0]), len(topography), topography}
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
