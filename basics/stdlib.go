package basics

import (
	"fmt"
	"math"
	"sort"
	"strconv"
)

//Sort in go
func Sort() {
	intSlice := []int{3, 2, 5, 6, 1}
	//int sort
	sort.Ints(intSlice)
	fmt.Println(intSlice)

	stringSlice := []string{"c", "b", "d", "a"}
	//sort string
	sort.Strings(stringSlice)
	fmt.Println(stringSlice)

	intSlice1 := []int{3, 2, 5, 6, 1}
	//sort custom
	sort.Slice(intSlice1, func(i int, j int) bool {
		return intSlice1[i] < intSlice1[j]
	})
	fmt.Println(intSlice1)
}

//Math basics in go
func Math() {
	fmt.Println("int32 min", math.MinInt32)
	fmt.Println("int32 max", math.MaxInt32)
	fmt.Println("int64 min", math.MinInt64)
	fmt.Println("int64 max", math.MaxInt64)
}

//Typeconversion basics in go
func Typeconversion() {
	s := "12345" //s[0] is of type byte
	fmt.Println(s[0])

	//byte to number
	num := int(s[0] - '0')
	fmt.Println(num)

	//byte to string
	str := string(s[0])
	fmt.Println(str)

	//number to byte
	b := byte(num + '0')
	fmt.Println(b)

	//format string print
	fmt.Printf("num : %d , str : %s , b : %c", num, str, b)
	fmt.Println()

	//string to number
	numb, _ := strconv.ParseInt(s, 10, 64)
	fmt.Printf("string to number %d", numb)
	fmt.Println()
	numToStr := strconv.FormatInt(12345, 10)
	fmt.Printf("number to string %s", numToStr)
	fmt.Println()

}
