package main

func isReapeted(set map[int]int) bool {
	for _, count := range set {
		if count > 1 {
			return true
		}
	}
	return false
}

func lenMaxNonRepeatingElements(list []int) int {
	if len(list) < 2 {
		return len(list)
	}

	maxWindow := 0
	set := make(map[int]int)

	for i := 0; i < len(list); i++ {
		// Добавляем в set i-тый элемент
		set[list[i]]++

		// Если в set есть повтор - вычитаем из set i-maxWindow-тый элемент и делаем шаг
		if isReapeted(set) {
			set[list[i-maxWindow]]--
		} else {
			// Если в set нет повторов - расширяем окно и двигаемся дальше
			maxWindow++
		}

	}

	return maxWindow
}

func GetMaxLen(list []int) int {
	set := make(map[int]int)
	prev, best := 0, 0

	for i := 0; i < len(list); i++ {
		prev = max(prev, set[list[i]])
		best = max(best, i-prev+1)
		set[list[i]] = i + 1
	}

	return best
}
