package day10

import (
	"fmt"
)

func main_2(lines []string) (n int, err error) {
	pipe_map := parseLinesToPipeMap(lines)
	if pipe_map.IsEmpty() {
		return 0, fmt.Errorf("empty pipe map")
	}

	// fmt.Println(pipe_map)
	// fmt.Println("start at", pipe_map.x, pipe_map.y)

	augmented_pipe_map := augmentPipeMap(&pipe_map)

	augmented_pipe_map.determineAllTileTypes()

	inside_count := augmented_pipe_map.CountTiles(INSIDE)

	fmt.Println(augmented_pipe_map)

	fmt.Println("inside count:", inside_count)

	return inside_count, nil
}

type TileType int

const (
	UNKNOWN TileType = iota
	CYCLE
	INSIDE
	OUTSIDE
)

func (t TileType) String() string {
	switch t {
	case UNKNOWN:
		return "?"
	case CYCLE:
		return "C"
	case INSIDE:
		return "I"
	case OUTSIDE:
		return "O"
	default:
		panic("unknown tile type")
	}
}

type AugmentedPipeMap struct {
	pipe_map  *PipeMap
	tile_type [][]TileType
}

func (m AugmentedPipeMap) String() string {
	s := ""
	for y := 0; y < m.pipe_map.height; y++ {
		for x := 0; x < m.pipe_map.width; x++ {
			switch m.tile_type[y][x] {
			case UNKNOWN:
				s += "?"
			case CYCLE:
				s += string(m.pipe_map.m[y][x])
				// s += "*"
			case INSIDE:
				s += "I"
			case OUTSIDE:
				s += ","
			}
		}
		if y < m.pipe_map.height-1 {
			s += "\n"
		}
	}
	return s
}

func augmentPipeMap(pipe_map *PipeMap) AugmentedPipeMap {

	// Walk full circle around the pipe map to determine where is it
	directions := walkFullCircle(pipe_map)

	// Make the augmented map
	tile_type := make([][]TileType, pipe_map.height)
	for i := range tile_type {
		tile_type[i] = make([]TileType, pipe_map.width)
	}

	// Fill in the tile_type map
	sx, sy := pipe_map.x, pipe_map.y

	tile_type[sy][sx] = CYCLE
	for _, dir := range directions {
		switch dir {
		case NORTH:
			sy--
		case SOUTH:
			sy++
		case EAST:
			sx++
		case WEST:
			sx--
		}
		tile_type[sy][sx] = CYCLE
	}

	return AugmentedPipeMap{
		pipe_map:  pipe_map,
		tile_type: tile_type,
	}
}

// Line to the edge of the tile_type map in the given direction
func (m *AugmentedPipeMap) lineToEdge(d direction) []Pipe {
	x, y := m.pipe_map.x, m.pipe_map.y
	line := []Pipe{}
	for {
		if x < 0 || x >= m.pipe_map.width {
			break
		}
		if y < 0 || y >= m.pipe_map.height {
			break
		}
		pipe := m.pipe_map.m[y][x]
		tile_type := m.tile_type[y][x]
		if tile_type != CYCLE {
			// If we're on the cycle we care about the type of the pipe
			// Otherwise we're outside the cycle and we only care about the
			// fact we're outside. There is no GROUND on the cycle so we
			// can use it as a marker for 'not cycle'
			pipe = GROUND
		}
		line = append(line, pipe)
		switch d {
		case NORTH:
			y--
		case SOUTH:
			y++
		case EAST:
			x++
		case WEST:
			x--
		}
	}
	return line
}

func (augmented *AugmentedPipeMap) determineCurrentTileType() {
	x, y := augmented.pipe_map.x, augmented.pipe_map.y

	tile_type := augmented.tile_type[y][x]
	if tile_type != UNKNOWN {
		// Already determined
		return
	}

	line_to_top_edge := augmented.lineToEdge(NORTH)
	line_to_top_edge = append(line_to_top_edge, GROUND)
	// fmt.Println(line_to_top_edge)
	cycle_crossings := 0
	var entrance_pipe Pipe = GROUND
	for i, pipe := range line_to_top_edge {

		if i == 0 {
			// Internal sanity check. The first tile_type should be unknown so
			// lineToEdge should have reported it as GROUND
			if pipe != GROUND {
				panic("first tile type should be unknown")
			}
			continue
		}

		switch pipe {
		case GROUND:
			if entrance_pipe == GROUND {
				// Whether we're inside or outside the cycle, we've come from a
				// GROUND tile and we're still on a GROUND tile. We're not
				// crossing the cycle
			} else {
				// We've come from a pipe and we're now on a GROUND tile. We've
				// crossed the cycle (leaving it)
				cycle_crossings++
				entrance_pipe = GROUND
			}
		case HORIZONTAL, NORTH_EAST, NORTH_WEST:
			// Pipes not connected from the south
			if entrance_pipe == GROUND {
				// Entrance pipe is ground, therefore we've not been inside of
				// the cycle. Just set the entrance pipe.
			} else {
				// We've entered a pipe from another pipe. We must have
				// crossed a cycle.
				cycle_crossings++
			}
			entrance_pipe = pipe
		case VERTICAL:
			// We're inside the cycle and we're going to stay inside
			// the cycle. Do nothing.
		case SOUTH_EAST:
			if entrance_pipe == NORTH_WEST {
				// We've entered the cycle, walked through a vertical
				// segment, and we're now crossing it
				cycle_crossings++
				entrance_pipe = GROUND
			} else if entrance_pipe == NORTH_EAST {
				// We've entered the cycle, walked through a vertical
				// segment, and we're now leaving it without crossing
				// the cycle
				entrance_pipe = GROUND
			} else {
				fmt.Println("entrance pipe:", entrance_pipe)
				panic("i dont think this should be possible (SOUTH_EAST)")
			}
		case SOUTH_WEST:
			if entrance_pipe == NORTH_EAST {
				// We've entered the cycle, walked through a vertical
				// segment, and we're now crossing it
				cycle_crossings++
				entrance_pipe = GROUND
			} else if entrance_pipe == NORTH_WEST {
				// We've entered the cycle, walked through a vertical
				// segment, and we're now leaving it without crossing
				// the cycle
				entrance_pipe = GROUND
			} else {
				panic("i dont think this should be possible (SOUTH_WEST)")
			}
		}
	}

	// If we crossed the cycle an odd number of times, we're inside
	if cycle_crossings%2 == 1 {
		tile_type = INSIDE
	} else {
		tile_type = OUTSIDE
	}

	// Fill in the tile_type map
	augmented.tile_type[y][x] = tile_type
}

func (augmented *AugmentedPipeMap) determineAllTileTypes() {
	sx, sy := augmented.pipe_map.x, augmented.pipe_map.y

	for y := 0; y < augmented.pipe_map.height; y++ {
		for x := 0; x < augmented.pipe_map.width; x++ {
			augmented.pipe_map.GoTo(x, y)
			augmented.determineCurrentTileType()
		}
	}

	// Go back to the starting position
	augmented.pipe_map.GoTo(sx, sy)
}

func (augmented *AugmentedPipeMap) CountTiles(tile_type TileType) int {
	count := 0
	for y := 0; y < augmented.pipe_map.height; y++ {
		for x := 0; x < augmented.pipe_map.width; x++ {
			if augmented.tile_type[y][x] == tile_type {
				fmt.Println("found", tile_type, "at", x, y)
				count++
			}
		}
	}
	return count
}
