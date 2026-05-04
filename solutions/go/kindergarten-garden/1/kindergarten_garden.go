package kindergartengarden

import (
	"errors"
	"sort"
	"strings"
)

type Garden struct {
	row1     []rune
	row2     []rune
	children []string
}

var defaultChildren = []string{
	"Alice", "Bob", "Charlie", "David", "Eve", "Fred",
	"Ginny", "Harriet", "Ileana", "Joseph", "Kincaid", "Larry",
}

var plantMap = map[rune]string{
	'G': "grass",
	'C': "clover",
	'R': "radishes",
	'V': "violets",
}

func NewGarden(diagram string, children []string) (*Garden, error) {
	// Check leading new line
    if !strings.HasPrefix(diagram, "\n") {
		return nil, errors.New("invalid diagram format")
	}
    
    lines := strings.Split(strings.TrimSpace(diagram), "\n")

	// Must have exactly 2 rows
	if len(lines) != 2 {
		return nil, errors.New("invalid diagram format")
	}

	row1 := []rune(lines[0])
	row2 := []rune(lines[1])

	// Rows must be same length
	if len(row1) != len(row2) {
		return nil, errors.New("rows must be equal length")
	}

	// Number of cups must be even
	if len(row1)%2 != 0 {
		return nil, errors.New("odd number of cups")
	}

	// Validate plant codes
	valid := map[rune]bool{'G': true, 'C': true, 'R': true, 'V': true}
	for _, r := range append(row1, row2...) {
		if !valid[r] {
			return nil, errors.New("invalid plant code")
		}
	}

	// Handle children
	if len(children) == 0 {
		children = append([]string{}, defaultChildren...)
	} else {
		children = append([]string{}, children...)
		sort.Strings(children)

		// Check duplicates
		for i := 1; i < len(children); i++ {
			if children[i] == children[i-1] {
				return nil, errors.New("duplicate child")
			}
		}
	}

	// Check if enough cups for children
	if len(row1)/2 < len(children) {
		return nil, errors.New("not enough cups")
	}

	return &Garden{
		row1:     row1,
		row2:     row2,
		children: children,
	}, nil
}

func (g *Garden) Plants(child string) ([]string, bool) {
	// Find child index
	index := -1
	for i, c := range g.children {
		if c == child {
			index = i
			break
		}
	}
	if index == -1 {
		return nil, false
	}

	start := index * 2

	plants := []rune{
		g.row1[start], g.row1[start+1],
		g.row2[start], g.row2[start+1],
	}

	result := make([]string, 0, 4)
	for _, p := range plants {
		result = append(result, plantMap[p])
	}

	return result, true
}