package main

import (
	"math/rand"
	"testing"
	"time"
)


func init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
const sliceSize = 50

func randString() string {
	n := 10
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func randSlice(size int) []string {
	var s []string
	for i:=0; i<size; i++ {
		s = append(s, randString())
	}
	return s
}

func areTheSameSlice(x, y []string) bool {
	if len(x) != len(y) {
		return false
	}

	for i:=0; i<len(x); i++ {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}

func TestDictOrderIsPredictable(t *testing.T) {
	times := 10
	type ssMap = map[string]string
	var allMaps []ssMap

	// create slice of size $times that contains maps of size $sliceSize; add to allMaps
	var m = ssMap{}
	for i:=0; i<times; i++ {
		for _, v := range randSlice(sliceSize) {
			m[v] = "val"
		}
		allMaps = append(allMaps, m)
	}

	if len(allMaps) != times {
		t.Fatal("allMaps is not correct len")
	}

	// from a map for each loop, write all of the keys into a slice
	type sSlice []string
	var allSlices []sSlice
	for _, m := range allMaps {
		var tmpSlice sSlice
		for key := range m {
			tmpSlice = append(tmpSlice, key)
		}
		allSlices = append(allSlices, tmpSlice)
	}

	// this code works, but if we leave it here the slices are sorted and the later test is tainted
	//// verify all of the maps are the same if sorted
	//for _, s := range allSlices {
	//	firstSlice := allSlices[0]
	//	sort.Slice(firstSlice, func(i, j int) bool {return firstSlice[i] < firstSlice[j]})
	//	sort.Slice(s, func(i, j int) bool {return s[i] < s[j]})
	//	if !areTheSameSlice(firstSlice, s) {
	//		t.Fatal("the sorted slices are NOT the same")
	//	}
	//}

	// verify all of the maps are the same if sorted
	for _, s := range allSlices {
		firstSlice := allSlices[0]
		if !areTheSameSlice(firstSlice, s) {
			t.Fatal("the unsorted slices are NOT the same")
			// this fails, non-deterministic
		}
	}
}



