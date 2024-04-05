package main

func calculateCost(cost []int) int {
	size := len(cost)
	array := make([]int, size)
	array[0], array[1] = cost[0], cost[1]
	for i := 2; i < size; i++ {
		array[i] = cost[i] + min(array[i-2], array[i-1])
	}
	return minimum(array[size-2], array[size-1])
}

func minimum(x, y int) int {
	if x > y {
		return y
	}
	return x
}
