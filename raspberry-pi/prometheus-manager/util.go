package main

func FindIndexByValue[T comparable](list []T, target T) int {
	for index, arg := range list {
		if arg == target {
			return index
		}
	}

	return -1
}

func RemoveItemInSlice[T comparable](list []T, target T) []T {
	for index, item := range list {
		if item == target {
			return append(list[:index], list[index+1:]...)
		}
	}

	return list
}
