package helpers

import "math/rand"

func ArrRemoveRepeatedString(arr []string) (newArr []string) {
	newArr = make([]string, 0)
	for i := 0; i < len(arr); i++ {
		repeat := false
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				repeat = true
				break
			}
		}
		if !repeat {
			newArr = append(newArr, arr[i])
		}
	}
	return
}

// ArrRemoveRepeatedInt Int数组去重
func ArrRemoveRepeatedInt(array []int) []int {
	newArray := make([]int, 0)

	for _, i := range array {
		if len(newArray) == 0 {
			newArray = append(newArray, i)
		} else {
			for k, v := range newArray {
				if i == v {
					break
				}
				if k == len(newArray)-1 {
					newArray = append(newArray, i)
				}
			}
		}
	}
	return newArray
}

// ArrShuffle 打乱数组(字符串型数组)
func ArrShuffle(strings []string) string {
	for i := len(strings) - 1; i > 0; i-- {
		num := rand.Intn(i + 1)
		strings[i], strings[num] = strings[num], strings[i]
	}

	str := ""
	for i := 0; i < len(strings); i++ {
		str += strings[i]
	}
	return str
}

func ArrInString(data []string, item string) bool {
	for _, v := range data {
		if v == item {
			return true
		}
	}
	return false
}
