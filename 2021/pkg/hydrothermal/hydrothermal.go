package hydrothermal

import (
	"fmt"
	"strconv"
	"strings"
)

type Puzzle struct {
	gridSize int
	grid     [][]int
	lines    []Line
	overlaps map[string]int
}

func (p *Puzzle) PartOne() string {
	p.gridSize = 0
	p.overlaps = make(map[string]int)
	p.inputToLineArray(inputs)
	// make dat grid
	p.grid = make([][]int, p.gridSize)
	for i := range p.grid {
		p.grid[i] = make([]int, p.gridSize)
	}

	for _, line := range p.lines {
		if line.startX == line.endX {
			p.drawVerticalLine(line)
		}

		if line.startY == line.endY {
			p.drawHorizontalLine(line)
		}
	}

	result := len(p.overlaps)

	return fmt.Sprintf("%v", result)
}

func (p *Puzzle) PartTwo() string {
	p.gridSize = 0
	p.overlaps = make(map[string]int)
	p.inputToLineArray(inputs)
	// make dat grid
	p.grid = make([][]int, p.gridSize)
	for i := range p.grid {
		p.grid[i] = make([]int, p.gridSize)
	}

	for _, line := range p.lines {
		if line.startX == line.endX {
			p.drawVerticalLine(line)
			continue
		}

		if line.startY == line.endY {
			p.drawHorizontalLine(line)
			continue
		}

		if line.startX+line.endX == line.startY+line.endY {
			p.drawDiagonalLine(line)
		}
	}

	result := len(p.overlaps)

	return fmt.Sprintf("%v", result)
}

type Line struct {
	startX int
	startY int
	endX   int
	endY   int
}

func (p *Puzzle) inputToLineArray(sArr []string) {
	p.gridSize = 0
	p.lines = nil

	for _, s := range sArr {
		coordPair := strings.Split(s, " -> ")
		startCoord, endCoord := coordPair[0], coordPair[1]
		startPair := strings.Split(startCoord, ",")
		endPair := strings.Split(endCoord, ",")
		startXStr, startYStr := startPair[0], startPair[1]
		endXStr, endYStr := endPair[0], endPair[1]
		startX, _ := strconv.Atoi(startXStr)
		if startX > p.gridSize {
			p.gridSize = startX
		}
		startY, _ := strconv.Atoi(startYStr)
		if startY > p.gridSize {
			p.gridSize = startY
		}
		endX, _ := strconv.Atoi(endXStr)
		if endX > p.gridSize {
			p.gridSize = endX
		}
		endY, _ := strconv.Atoi(endYStr)
		if endY > p.gridSize {
			p.gridSize = endY
		}
		line := Line{
			startX: startX,
			startY: startY,
			endX:   endX,
			endY:   endY,
		}

		p.lines = append(p.lines, line)
	}
}

func (p *Puzzle) drawHorizontalLine(line Line) {
	startX, endX := 0, 0
	if line.startX < line.endX {
		startX = line.startX
		endX = line.endX
	} else {
		startX = line.endX
		endX = line.startX
	}

	for i := startX; i <= endX; i++ {
		p.grid[i][line.startY]++
		if p.grid[i][line.startY] > 1 {
			p.overlaps[fmt.Sprintf("%v,%v", i, line.startY)] = p.grid[i][line.startY]
		}
	}
}

func (p *Puzzle) drawVerticalLine(line Line) {
	startY, endY := 0, 0
	if line.startY < line.endY {
		startY = line.startY
		endY = line.endY
	} else {
		startY = line.endY
		endY = line.startY
	}

	for i := startY; i <= endY; i++ {
		p.grid[line.startX][i]++
		if p.grid[line.startX][i] > 1 {
			p.overlaps[fmt.Sprintf("%v,%v", line.startX, i)] = p.grid[line.startX][i]
		}
	}
}

func (p *Puzzle) drawDiagonalLine(line Line) {
	startX := line.startX
	endX := line.endX
	startY := line.startY
	endY := line.endY
	if line.startX > line.endX {
		startX = line.endX
		endX = line.startX
		startY = line.endY
		endY = line.startY
	}

	y := startY
	yDirection := 0
	if startY < endY {
		yDirection = 1
	} else {
		yDirection = -1
	}

	for x := startX; x <= endX; x++ {
		p.grid[x][y]++
		if p.grid[x][y] > 1 {
			p.overlaps[fmt.Sprintf("%v,%v", x, y)] = p.grid[x][y]
		}
		y = y + yDirection
	}
}
