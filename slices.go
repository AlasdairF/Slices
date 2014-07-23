/*

INDEX

Array2Map_uint
Array2Map_string
-- Converts a slice to a map

ArrayUniqueRetain_uint
ArrayUniqueRetain_string
-- Very fast method for removing duplicates from a slice, using maps, preserves the order of the original slice

ArrayUnique_uint
ArrayUnique_string
-- Very fast method for removing duplicates from a slice, using maps, even faster but does not preserve the order

ArrayIntersect_uint
ArrayIntersect_string
-- Returns a new slice that consists only of values that existed in all slices

ArrayRemove_uint
ArrayRemove_string
-- Return a new slice that consists of the values in the first slice that don't exist in any subsequent slices

ArraySquish_uint
ArraySquish_string
-- Returns a new slice having combined the values of all slices into one, with duplicates removed and moved higher up the list

NOTES

I am aware that I use the term `array` a lot instead of `slice`. This is because I don't care much for the difference and `array` is easier to think of.
I tend to use uint for most things, I'm not a fan of int, therefore these functions are made to work only with uint and string.

*/


package aslices

import (
	"sort"
	)

// ---------------- STRUCTURES ----------------
//  Structure for converting a map into a slice (usually for the purpose of sorting it)
type keyVal_uint struct {
	Key uint
	Val int
}
type keyVal_string struct {
	Key string
	Val int
}

// ---------------- SORTING FUNCTIONS ----------------
type sorter_uint []keyVal_uint
type sorter_string []keyVal_string
func (a sorter_uint) Len() int           { return len(a) }
func (a sorter_uint) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a sorter_uint) Less(i, j int) bool { return a[i].Val < a[j].Val }
func (a sorter_string) Len() int           { return len(a) }
func (a sorter_string) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a sorter_string) Less(i, j int) bool { return a[i].Val < a[j].Val }

// ---------------- ARRAY to MAP FUNCTIONS ----------------
//  Array2Map_uint makes a map out of a slice of integers: uint=>false
func Array2Map_uint(a []uint) map[uint]bool {
	mc := make(map[uint]bool)
    for _,tok := range a {
         mc[tok]=false
        }
    return mc
}

//  Array2Map_string makes a map out of an slice of strings: word=>false
func Array2Map_string(a []string) map[string]bool {
	mc := make(map[string]bool)
    for _,tok := range a {
         mc[tok]=false
        }
    return mc
}

// ---------------- ARRAY FUNCTIONS ----------------
// ArrayUniqueRetain_uint returns the same slice with all duplicates removed, the order is retained
func ArrayUniqueRetain_uint(a []uint) []uint {
	mc := make(map[uint]uint)
	var i uint
    for _,tok := range a {
		_, ok := mc[tok]
        if !ok {
         mc[tok]=i;
		 i++
        }
    }
	b := make([]uint,len(mc))
	for k,i := range mc {
		b[i] = k
	}
	return b
}

// ArrayUniqueRetain_string returns the same slice with all duplicates removed, the order is retained
func ArrayUniqueRetain_string(a []string) []string {
	mc := make(map[string]uint)
	var i uint
    for _,tok := range a {
		_, ok := mc[tok]
        if !ok {
         mc[tok]=i;
		 i++
        }
    }
	b := make([]string,len(mc))
	for k,i := range mc {
		b[i] = k
	}
	return b
}

// ArrayUnique_uint returns the same slice with all duplicates removed, the order is not kept (faster than retaining order)
func ArrayUnique_uint(a []uint) []uint {
	mc := Array2Map_uint(a)
	b := make([]uint,len(mc))
	var i uint
	for k := range mc {
		b[i] = k
		i++
	}
	return b
}

// ArrayUnique_string returns the same slice with all duplicates removed, the order is not kept (faster than retaining order)
func ArrayUnique_string(a []string) []string {
	mc := Array2Map_string(a)
	b := make([]string,len(mc))
	var i uint
	for k := range mc {
		b[i] = k
		i++
	}
	return b
}

//  ArrayIntersect_uint returns a slice only with values that are in all of the slices passed, the order is as slice1
//  Note that if there are duplicate values in slice1 then these will still occur in the result
func ArrayIntersect_uint(a ...[]uint) []uint {
	num := uint8(len(a)-1)
	mc := make(map[uint]uint8)
	for i := uint8(0); i<num; i++ {
		for _,tok := range a[i+1] {
		if mc[tok]==i {
			mc[tok]++
			}
		}
	}
	var current uint = 0
	new1 := make([]uint,len(a[0]))
	for _,tok := range a[0] {
		if (mc[tok]==num) {
		new1[current]=tok
		current++
		}
	}
	new2 := make([]uint,current)
	copy(new2, new1)
	return new2
}

//  ArrayIntersect_string returns a slice only with values that are in all of the slices passed, the order is as slice1
//  Note that if there are duplicate values in slice1 then these will still occur in the result
func ArrayIntersect_string(a ...[]string) []string {
	num := uint8(len(a)-1)
	mc := make(map[string]uint8)
	for i := uint8(0); i<num; i++ {
		for _,tok := range a[i+1] {
		if mc[tok]==i {
			mc[tok]++
			}
		}
	}
	var current uint = 0
	new1 := make([]string,len(a[0]))
	for _,tok := range a[0] {
		if (mc[tok]==num) {
		new1[current]=tok
		current++
		}
	}
	new2 := make([]string,current)
	copy(new2, new1)
	return new2
}

//  ArrayRemove_uint returns slice1 with all values from the following slices removed from it, the order is as slice1
func ArrayRemove_uint(a ...[]uint) []uint {
	num := uint8(len(a)-1)
	mc := make(map[uint]bool)
	for i := uint8(0); i<num; i++ {
		for _,tok := range a[i+1] {
			mc[tok]=true
		}
	}
	var current uint = 0
	new1 := make([]uint,len(a[0]))
	for _,tok := range a[0] {
		_, ok := mc[tok]
		if !ok {
			new1[current]=tok
			current++
		}
	}
	new2 := make([]uint,current)
	copy(new2, new1)
	return new2
}

//  ArrayRemove_string returns slice2 with all values from the following slices removed from it, the order is as slice 1
func ArrayRemove_string(a ...[]string) []string {
	num := uint8(len(a)-1)
	mc := make(map[string]bool)
	for i := uint8(0); i<num; i++ {
		for _,tok := range a[i+1] {
			mc[tok]=true
		}
	}
	var current uint = 0
	new1 := make([]string,len(a[0]))
	for _,tok := range a[0] {
		_, ok := mc[tok]
		if !ok {
			new1[current]=tok
			current++
		}
	}
	new2 := make([]string,current)
	copy(new2, new1)
	return new2
}

// ArraySquish_uint merges together multiple slices in the form: slice1[0], slice2[0], slice1[1], slice2[1], etc. Duplicates are removed and given higher priority (they move up the list)
// Note: the purpose of this function is for combining together different search results into one, hence why duplicates are removed and given priority
// This maximum length is set with the first argument `maxlen`, if maxlen is given as 0 then the length of the output is not limited.
func ArraySquish_uint(maxlen int, a ...[]uint) []uint {
	mc := make(map[uint]int)
	// Add the values and weights into a hashtable with value as key
	for i:=0; i<len(a); i++ {
		for i2,tok := range a[i] {
			_, ok := mc[tok]
			if (ok) {
				mc[tok]-=100000-i2
			} else {
				mc[tok]=i2
			}
		}
	}
	// Now convert from map into structured slice
	array := make(sorter_uint,len(mc))
	var current int = 0
	for Key,Val := range mc {
		array[current]=keyVal_uint{Key,Val}
		current++
	}
	sort.Sort(array)
	// Now turn the sorted structure back into a simple slice
	if current<maxlen || maxlen==0 {
		maxlen=current
	}
	new2 := make([]uint,maxlen)
	for i:=0; i<maxlen; i++ {
		new2[i]=array[i].Key
	}
	return new2
}

// ArraySquish_string merges together multiple slices in the form: slice1[0], slice2[0], slice1[1], slice2[1], etc. Duplicates are removed and given higher priority (they move up the list)
// Note: the purpose of this function is for combining together different search results into one, hence why duplicates are removed and given priority
// This maximum length is set with the first argument `maxlen`, if maxlen is given as 0 then the length of the output is not limited.
func ArraySquish_string(maxlen int, a ...[]string) []string {
	mc := make(map[string]int)
	// Add the values and weights into a hashtable with value as key
	for i:=0; i<len(a); i++ {
		for i2,tok := range a[i] {
			_, ok := mc[tok]
			if (ok) {
				mc[tok]-=100000-i2
			} else {
				mc[tok]=i2
			}
		}
	}
	// Now convert from map into structured slice
	array := make(sorter_string,len(mc))
	var current int = 0
	for Key,Val := range mc {
		array[current]=keyVal_string{Key,Val}
		current++
	}
	sort.Sort(array)
	// Now turn the sorted structure back into a simple slice
	if current<maxlen || maxlen==0 {
		maxlen=current
	}
	new2 := make([]string,maxlen)
	for i:=0; i<maxlen; i++ {
		new2[i]=array[i].Key
	}
	return new2
}
