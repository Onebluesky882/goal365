package lib

func FormScore(match int, form string) int {

	// set
	score := 0

	// running
	count := 0

	for i := len(form) - 1; i >= 0 && count < match; i-- {
		switch form[i] {
		case 'W':
			score += 3
		case 'D':
			score += 1
		case 'L':
		}
		count++
	}
	return score
}
