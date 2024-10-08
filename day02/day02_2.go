package day02

func main_2(lines []string, verbose bool) (n int, err error) {
	sum_power_sets := 0
	for _, line := range lines {
		line, err = simplifyLine(line)
		if err != nil {
			return -1, err
		}
		games, err := lineToGames(line)
		if err != nil {
			return -1, err
		}
		fewest := gameToFewest(games)
		power_set := fewest.red * fewest.green * fewest.blue
		sum_power_sets += power_set
	}
	return sum_power_sets, nil
}

func gameToFewest(games []game) (g game) {
	for _, game := range games {
		g.red = max(g.red, game.red)
		g.green = max(g.green, game.green)
		g.blue = max(g.blue, game.blue)
	}
	return g
}
