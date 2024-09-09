package day10

import (
	"fmt"
)

func Main(part int, lines []string, verbose bool) (n int, err error) {
	if part == 1 {
		return main_1(lines, verbose)
	} else if part == 2 {
		return main_2(lines, verbose)
	} else {
		return -1, fmt.Errorf("invalid part")
	}
}

func main_1(lines []string, verbose bool) (n int, err error) {
	pipe_map := parseLinesToPipeMap(lines)
	if pipe_map.IsEmpty() {
		return 0, fmt.Errorf("empty pipe map")
	}

	// fmt.Println(pipe_map)
	// fmt.Println("start at", pipe_map.x, pipe_map.y)

	directions := walkFullCircle(&pipe_map)

	// fmt.Println("directions:", directions)

	furthest := len(directions) / 2

	return furthest, nil
}

type Pipe rune

const (
	VERTICAL   Pipe = '|'
	HORIZONTAL Pipe = '-'
	NORTH_EAST Pipe = 'L'
	NORTH_WEST Pipe = 'J'
	SOUTH_EAST Pipe = 'F'
	SOUTH_WEST Pipe = '7'
	GROUND     Pipe = '.'
	START      Pipe = 'S'
)

var pipes = []Pipe{VERTICAL, HORIZONTAL, NORTH_EAST, NORTH_WEST, SOUTH_EAST, SOUTH_WEST, GROUND, START}

func isPipe(r rune) bool {
	for _, p := range pipes {
		if r == rune(p) {
			return true
		}
	}
	return false
}

func (p Pipe) String() string {
	return string(p)
}

// Given a pipe and the direction, return the other direction that is valid for
// that pipe.
func (p Pipe) otherDirection(one_direction direction) direction {
	switch p {
	case VERTICAL:
		if one_direction == NORTH || one_direction == SOUTH {
			return one_direction.opposite()
		} else {
			panic(fmt.Errorf("invalid direction %s for pipe %s", one_direction, p))
		}
	case HORIZONTAL:
		if one_direction == EAST || one_direction == WEST {
			return one_direction.opposite()
		} else {
			panic(fmt.Errorf("invalid direction %s for pipe %s", one_direction, p))
		}
	case NORTH_EAST:
		if one_direction == NORTH {
			return EAST
		} else if one_direction == EAST {
			return NORTH
		} else {
			panic(fmt.Errorf("invalid direction %s for pipe %s", one_direction, p))
		}
	case NORTH_WEST:
		if one_direction == NORTH {
			return WEST
		} else if one_direction == WEST {
			return NORTH
		} else {
			panic(fmt.Errorf("invalid direction %s for pipe %s", one_direction, p))
		}
	case SOUTH_EAST:
		if one_direction == SOUTH {
			return EAST
		} else if one_direction == EAST {
			return SOUTH
		} else {
			panic(fmt.Errorf("invalid direction %s for pipe %s", one_direction, p))
		}
	case SOUTH_WEST:
		if one_direction == SOUTH {
			return WEST
		} else if one_direction == WEST {
			return SOUTH
		} else {
			panic(fmt.Errorf("invalid direction %s for pipe %s", one_direction, p))
		}
	default:
		panic(fmt.Errorf("invalid pipe %d", p))
	}
}

type PipeMap struct {
	m      [][]Pipe
	width  int
	height int
	x      int
	y      int
}

func (m PipeMap) IsEmpty() bool {
	return m.width == 0 || m.height == 0
}

func (m PipeMap) Here() Pipe {
	return m.m[m.y][m.x]
}

func (m PipeMap) String() string {
	s := ""
	for y := 0; y < m.height; y++ {
		for x := 0; x < m.width; x++ {
			s += string(m.m[y][x])
		}
		if y < m.height-1 {
			s += "\n"
		}
	}
	return s
}

type direction int

const (
	NORTH direction = iota
	EAST
	SOUTH
	WEST
)

func (d direction) String() string {
	switch d {
	case NORTH:
		return "NORTH"
	case EAST:
		return "EAST"
	case SOUTH:
		return "SOUTH"
	case WEST:
		return "WEST"
	default:
		panic(fmt.Errorf("invalid direction %d", d))
	}
}

// Given a direction, return the opposite direction.
func (d direction) opposite() direction {
	switch d {
	case NORTH:
		return SOUTH
	case EAST:
		return WEST
	case SOUTH:
		return NORTH
	case WEST:
		return EAST
	default:
		panic(fmt.Errorf("invalid direction %d", d))
	}
}

// Walk in the given direction. Returns the pipe we end up on or GROUND if we
// can't walk in that direction. The move is only valid if the pipe we are
// walking from a pipe that allows us to walk in that direction and the pipe we
// are walking to allows us to come from the opposite direction. We can't walk
// outside the map boundaries.
func (m *PipeMap) Walk(d direction) Pipe {
	here := m.Here()
	switch d {
	case NORTH:
		if m.y > 0 {
			// If we are on a pipe that allows us to walk north
			if here == VERTICAL || here == NORTH_EAST || here == NORTH_WEST || here == START {
				there := m.m[m.y-1][m.x]
				// And the pipe we are walking to allows us to come from the south
				if there == VERTICAL || there == SOUTH_EAST || there == SOUTH_WEST {
					m.y--
					return there
				}
			}
		}
	case EAST:
		if m.x < m.width-1 {
			there := m.m[m.y][m.x+1]
			if here == HORIZONTAL || here == NORTH_EAST || here == SOUTH_EAST || here == START {
				if there == HORIZONTAL || there == NORTH_WEST || there == SOUTH_WEST {
					m.x++
					return there
				}
			}
		}
	case SOUTH:
		if m.y < m.height-1 {
			there := m.m[m.y+1][m.x]
			if here == VERTICAL || here == SOUTH_EAST || here == SOUTH_WEST || here == START {
				if there == VERTICAL || there == NORTH_EAST || there == NORTH_WEST {
					m.y++
					return there
				}
			}
		}
	case WEST:
		if m.x > 0 {
			there := m.m[m.y][m.x-1]
			if here == HORIZONTAL || here == NORTH_WEST || here == SOUTH_WEST || here == START {
				if there == HORIZONTAL || there == NORTH_EAST || there == SOUTH_EAST {
					m.x--
					return there
				}
			}
		}
	default:
		panic(fmt.Errorf("invalid direction %d", d))
	}
	return GROUND
}

func (m *PipeMap) WalkSteps(d direction, steps int) Pipe {
	var p Pipe = m.Here()
	for i := 0; i < steps; i++ {
		p = m.Walk(d)
		if p == GROUND {
			return GROUND
		}
	}
	return m.Here()
}

func (m *PipeMap) GoTo(x, y int) {
	m.x = x
	m.y = y
}

func parseLinesToPipeMap(lines []string) (pipe_map PipeMap) {
	height := len(lines)
	if height == 0 {
		return PipeMap{}
	}
	width := len(lines[0])
	if width == 0 {
		return PipeMap{}
	}

	// Allocate the map
	pipe_map = PipeMap{make([][]Pipe, height), width, height, 0, 0}
	for y := 0; y < height; y++ {
		pipe_map.m[y] = make([]Pipe, width)
	}

	var sx, sy int

	// Parse the map
	n_starting_points := 0
	for y, line := range lines {
		if len(line) != width {
			panic(fmt.Errorf("line %d has length %d, expected %d", y, len(line), width))
		}
		for x, r := range line {
			if !isPipe(r) {
				panic(fmt.Errorf("invalid character %c at (%d,%d)", r, x, y))
			}
			p := Pipe(r)
			pipe_map.m[y][x] = p
			if p == START {
				n_starting_points++
				sx = x
				sy = y
			}
		}
	}

	if n_starting_points != 1 {
		panic(fmt.Errorf("expected 1 starting point, found %d", n_starting_points))
	}

	pipe_map.x = sx
	pipe_map.y = sy

	// Find the starting pipe
	start_pipe := findStartingPipe(&pipe_map)
	pipe_map.m[sy][sx] = start_pipe

	return pipe_map
}

func findStartingPipe(pipe_map *PipeMap) Pipe {
	// We don't know what direction is the pipe under the starting point, so we
	// try all directions and see which one is valid. Its enough to try to walk
	// one step in each direction.

	// Original starting point
	sx := pipe_map.x
	sy := pipe_map.y

	if pipe_map.Here() != START {
		panic(fmt.Errorf("starting point is not START"))
	}

	is_valid_map := map[Pipe]bool{
		VERTICAL:   true,
		HORIZONTAL: true,
		NORTH_EAST: true,
		NORTH_WEST: true,
		SOUTH_EAST: true,
		SOUTH_WEST: true,
	}

	// For each valid pipe option, try to walk in that direction and see if the path is consistent
	for p := range is_valid_map {
		if p == VERTICAL || p == NORTH_EAST || p == NORTH_WEST {
			// Try to walk north
			pipe_map.GoTo(sx, sy)
			got_to := pipe_map.WalkSteps(NORTH, 1)

			if got_to == GROUND {
				is_valid_map[p] = false
				continue
			}
		}
		if p == HORIZONTAL || p == NORTH_EAST || p == SOUTH_EAST {
			// Try to walk east
			pipe_map.GoTo(sx, sy)
			if pipe_map.WalkSteps(EAST, 1) == GROUND {
				is_valid_map[p] = false
				continue
			}
		}
		if p == VERTICAL || p == SOUTH_EAST || p == SOUTH_WEST {
			// Try to walk south
			pipe_map.GoTo(sx, sy)
			if pipe_map.WalkSteps(SOUTH, 1) == GROUND {
				is_valid_map[p] = false
				continue
			}
		}
		if p == HORIZONTAL || p == NORTH_WEST || p == SOUTH_WEST {
			// Try to walk west
			pipe_map.GoTo(sx, sy)
			if pipe_map.WalkSteps(WEST, 1) == GROUND {
				is_valid_map[p] = false
				continue
			}
		}
	}

	// Go back to the original starting point
	pipe_map.GoTo(sx, sy)

	// Check if we have a single valid pipe option. If so that must be it.
	N_valid_pipes := 0
	valid_pipe := GROUND
	for p, v := range is_valid_map {
		if v {
			N_valid_pipes++
			valid_pipe = p
		}
	}

	if N_valid_pipes != 1 {
		panic(fmt.Errorf("expected 1 valid pipe, found %d", N_valid_pipes))
	}

	return valid_pipe
}

// Starting at the current pipe, start walking in some direction until we reach
// the same position again.
func walkFullCircle(pipe_map *PipeMap) []direction {

	sx := pipe_map.x
	sy := pipe_map.y

	here := pipe_map.Here()
	if here == GROUND || here == START {
		// We cant walk from GROUND because it whoudl not be in the map at all,
		// at least not in the circle. We can't walk from START because we
		// should have found the starting pipe already.
		panic(fmt.Errorf("can't walk in a circle from GROUND or START"))
	}

	// For each pipe type just pick one direction to walk in
	var starting_direction direction
	switch here {
	case VERTICAL:
		starting_direction = NORTH
	case HORIZONTAL:
		starting_direction = EAST
	case NORTH_EAST:
		starting_direction = NORTH
	case NORTH_WEST:
		starting_direction = WEST
	case SOUTH_EAST:
		starting_direction = EAST
	case SOUTH_WEST:
		starting_direction = SOUTH
	}

	walk_direction := starting_direction
	directions := []direction{}
	for {
		// Walk in the current direction
		pipe_map.Walk(walk_direction)
		directions = append(directions, walk_direction)
		here = pipe_map.Here()

		if here == GROUND || here == START {
			panic(fmt.Errorf("found GROUND or START while walking in a circle"))
		}

		// If we are back at the starting point, we are done
		if pipe_map.x == sx && pipe_map.y == sy {
			break
		}

		// Figure out the next direction to walk in
		come_from := walk_direction.opposite()
		walk_direction = here.otherDirection(come_from)
	}

	if pipe_map.x != sx || pipe_map.y != sy {
		panic(fmt.Errorf("did not end up at the starting point"))
	}

	return directions
}
