package util

import "fmt"

type Set map[interface{}]bool

func isValidForSet(t interface{}) {
	switch v := t.(type) {
	case int:
		fmt.Println("Received an integer:", v)
	case string:
		fmt.Println("Received a string:", v)
	default:
		fmt.Println("Received an unsupported type")
		panic(1)
	}
}

func (set Set) Add(element interface{}) {
	isValidForSet(element)
	set[element] = true
}

func (set Set) Remove(element interface{}) {
	delete(set, element)
}

func (set Set) Contains(element interface{}) bool {
	return set[element]
}
