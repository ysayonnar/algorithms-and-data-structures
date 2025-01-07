package leetcode

func count(senate string) int {
	counter := 0

	for i := 0; i < len(senate); i++ {
		if senate[i] == 'D' {
			counter--
		} else {
			counter++
		}
	}

	return counter
}

func predictPartyVictory(senate string) string {
	for {
		counter := count(senate)
		if counter > 0 && len(senate) == counter {
			return "Radiant"
		} else if counter < 0 && len(senate) == -1*counter {
			return "Dire"
		}

		current := senate[0]
		for i := 1; i < len(senate); i++ {
			if current != senate[i] {
				senate = senate[1:i] + senate[i+1:] + string(current)
				break
			}
		}
	}
}
