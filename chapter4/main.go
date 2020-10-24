package main

import "gitlab.wcxst.com/jormin/go-tools/log"

func main() {
	arr1 := [5]int{1, 2, 3, 4, 5}
	log.Info("arr1: %+v", arr1)
	var arr2 [5]int
	arr2 = arr1
	log.Info("arr2: %+v", arr2)
	//var arr3 [4]int
	//arr3 = arr1
	//log.Info("arr3: %+v", arr3)

	arr3 := [5]*int{new(int), new(int), new(int), new(int), new(int)}
	*arr3[0] = 1
	*arr3[1] = 2
	log.Info("arr3: %+v", arr3)
	var arr4 [5]*int
	arr4 = arr3
	log.Info("arr4: %+v", arr4)
	log.Info("*arr3[0]=%d, *arr4[0]=%d", *arr3[0], *arr4[0])
	*arr3[0] = 10
	log.Info("*arr3[0]=%d, *arr4[0]=%d", *arr3[0], *arr4[0])

	slice1 := make([]int, 5)
	log.Info("slice1: %+v, len: %d, cap: %d", slice1, len(slice1), cap(slice1))
	slice2 := make([]int, 5, 10)
	log.Info("slice2: %+v, len: %d, cap: %d", slice2, len(slice2), cap(slice2))

	slice3 := [3]int{1, 2, 3}
	for i, v := range slice3 {
		log.Info("v: %d, i-addr:%X, v-addr: %X, element-addr: %X", v, &i, &v, &slice3[i])
	}

	slice4 := [][]int{{10}, {100, 200}}
	log.Info("slice4: %+v", slice4)
	slice4 = append(slice4, []int{1000})
	log.Info("slice4: %+v", slice4)
	slice4[0] = append(slice4[0], 20)
	log.Info("slice4: %+v", slice4)

	m1 := make(map[string]string)
	for k := range m1 {
		log.Info("k: %v", k)
	}
	for k, exists := range m1 {
		log.Info("k: %v, exists: %v", k, exists)
	}

	m2 := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
		"e": 5,
	}
	for k := range m2 {
		log.Info("k: %v, v: %v", k, m2[k])
	}
	log.Info("after delete:")
	delete(m2, "d")
	for k := range m2 {
		log.Info("k: %v, v: %v", k, m2[k])
	}
	log.Info("m3:")
	m3 := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
		"d": 4,
		"e": 5,
	}
	for k := range m3 {
		log.Info("k: %v, v: %v", k, m3[k])
	}
	log.Info("len(m3): %d", len(m3))
	changeMap(m3)
	log.Info("after change:")
	for k := range m3 {
		log.Info("k: %v, v: %v", k, m3[k])
	}
	log.Info("len(m3): %d", len(m3))

}

func changeMap(m map[string]int) {
	m["f"] = 6
}
