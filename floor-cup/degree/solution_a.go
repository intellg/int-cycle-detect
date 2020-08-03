package degree

func InnerCalculateA(floor, cup int) (degree int) {
	// 1.1 Prepare the init list (all 1)
	list := make([]int, cup)
	for c := 0; c < cup; c++ {
		list[c] = 1
	}

	// 1.2 Calculate the list
	sum := 1
	for degree = 1; sum < floor; degree++ {
		calList := make([]int, cup)
		calList[0] = 1
		for c := 1; c < cup; c++ {
			calList[c] = list[c] + list[c-1]
			counter++
		}
		list = calList
		sum += calList[cup-1]
	}
	return
}
