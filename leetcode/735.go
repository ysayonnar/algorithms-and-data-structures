package leetcode

func asteroidCollision(asteroids []int) []int {
	stack := []int{}
	stack = append(stack, asteroids[0])

	for i := 1; i < len(asteroids); i++ {
		if len(stack) != 0 {
			if stack[len(stack)-1] > 0 && asteroids[i] < 0 {
				if stack[len(stack)-1] == asteroids[i]*-1 {
					stack = stack[:len(stack)-1]
				} else if stack[len(stack)-1] < asteroids[i]*-1 {
					stack[len(stack)-1] = asteroids[i]
					stack = asteroidCollision(stack) // хавает много времени
				}
			} else {
				stack = append(stack, asteroids[i])
			}
		} else {
			stack = append(stack, asteroids[i])
		}
	}
	return stack
}
