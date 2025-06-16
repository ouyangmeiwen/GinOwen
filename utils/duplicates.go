package utils

import (
	"fmt"
)

// 判断是否包含重复元素
func HasDuplicate[T any](arr []T, equals func(a, b T) bool) bool {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if equals(arr[i], arr[j]) {
				return true
			}
		}
	}
	return false
}

// 返回第一个重复元素（第二次出现的位置），第二返回值标记是否找到
func FirstDuplicate[T any](arr []T, equals func(a, b T) bool) (T, bool) {
	var zero T
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if equals(arr[i], arr[j]) {
				return arr[j], true
			}
		}
	}
	return zero, false
}

// 返回所有重复元素（每种只返回一次）
func AllDuplicates[T any](arr []T, equals func(a, b T) bool) []T {
	result := []T{}
	seen := []T{}

	for i := 0; i < len(arr); i++ {
		isDup := false
		// 如果之前已经收集过这个重复元素，就跳过
		for _, s := range seen {
			if equals(arr[i], s) {
				isDup = true
				break
			}
		}
		if isDup {
			continue
		}

		for j := i + 1; j < len(arr); j++ {
			if equals(arr[i], arr[j]) {
				result = append(result, arr[i])
				seen = append(seen, arr[i])
				break
			}
		}
	}
	return result
}

type Person struct {
	ID   int
	Name string
	Age  int
}

func Testmain() {
	people := []Person{
		{ID: 1, Name: "Alice", Age: 30},
		{ID: 2, Name: "Bob", Age: 25},
		{ID: 3, Name: "Alice", Age: 30},
		{ID: 4, Name: "Eve", Age: 40},
		{ID: 5, Name: "Bob", Age: 25},
	}

	equals := func(a, b Person) bool {
		return a.Name == b.Name && a.Age == b.Age
	}

	fmt.Println("是否有重复：", HasDuplicate(people, equals))
	fmt.Println("是否有重复：", HasDuplicate(people, func(a, b Person) bool {
		return a.Name == b.Name && a.Age == b.Age
	}))

	if dup, found := FirstDuplicate(people, equals); found {
		fmt.Println("第一个重复元素：", dup)
	}
	if dup, found := FirstDuplicate(people, func(a, b Person) bool {
		return a.Name == b.Name && a.Age == b.Age
	}); found {
		fmt.Println("第一个重复元素：", dup)
	}

	fmt.Println("所有重复元素：", AllDuplicates(people, equals))
	fmt.Println("所有重复元素：", AllDuplicates(people, func(a, b Person) bool {
		return a.Name == b.Name && a.Age == b.Age
	}))

}
