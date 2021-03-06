package binarysearch

import (
	"fmt"
	"sort"
)

// Define a map to contain parallel lines for keeping the values
var dataRails map[int][]int

// Define a map to contain parallel lines for keeping the keys
// each follows the order of the same key inside dataRails
var keysRails map[int][]int

// RunParallelDynamic is called externally from the cmd package
// generates the data, set up the target and runs the logic for "parallel-dynamic"
// implementation of the binary search
func RunParallelDynamic(t int) int {
	generateData()
	target = t

	return pds()
}

// pds runs the logic for parallelly dynamically updating the data while searching
// the function creates two subslices - "rails" of the data that needs to be searched
// with binary implementation - instead of using one list of values - use two
func pds() int {
	// Do not proceed with logic if data length is zero or one
	if len(sData) == 0 {
		return -1
	}

	if len(sData) == 1 {
		if sData[0] == target {
			return 0
		}

		return -1
	}

	// Sort list of values
	sort.Ints(sData)
	dData = sData

	// Create two lines for searching - not mirroring each other
	// in order to keep the sorting direction of values
	dDataLine1 := dData[:len(dData)/2]
	dDataLine2 := dData[len(dData)/2:]

	dataRails = map[int][]int{
		0: dDataLine1,
		1: dDataLine2,
	}

	// Get keys and set up two slices - rails for keeping the keys
	// corresponding to the data rails
	getKeys(sData)
	keysLine1 := keys[:len(keys)/2]
	keysLine2 := keys[len(keys)/2:]

	keysRails = map[int][]int{
		0: keysLine1,
		1: keysLine2,
	}

	// Search through the data using rails with flip-step mode
	// Set up pointer position
	
	po = len(dataRails[0]) -1
	if len(dataRails[0]) > len(dataRails[1]) {
		po = len(dataRails[1]) -1
	}

	return getKeypds()
}

// getKeypds runs the parallel-dynamic logic implementation for the binary search and 
// returns the target key or -1
func getKeypds() int {
	// We need to lower the p to 1, not 0, because we handle cases for data length
	// separately in the next check
	fmt.Println(po)
	if po > 1 {
		po = po/2
	}
	fmt.Println(po)

	// Check boundary conditions
	// Check if both rails are with zero lenth - return -1 - nothing is found
	if len(dataRails[0]) == 0 && len(dataRails[1]) == 0 {
		return -1
	}

	// If one of the rails length is one - return the value if equal to target
	// otherwise return -1
	// This check references the fact that if one of the rails is with length 0,
	// the other one will be 1
	if (len(dataRails[0]) == 1 && len(dataRails[1]) == 1) {
		if dataRails[0][0] == target {
			return keysRails[0][0]
		}
		if dataRails[1][0] == target {
			return keysRails[1][0]
		}
		return -1
	}

	// If on the the rails is with length zeo, compare the one - if any,
	// with length one and return its key if equals the target
	// This check references the fact that if one of the rails is with length 0,
	// the other one will be 1
	if len(dataRails[0]) == 0 || len(dataRails[1]) == 0 {
		if len(dataRails[1]) == 1 && dataRails[1][0] == target {
			return keysRails[1][0]
		}

		if len(dataRails[0]) == 1 && dataRails[0][0] == target {
			return keysRails[0][0]
		}

		return -1
	}

	// Cover all cases in the switch case or return -1
	switch {
	case target == dataRails[0][po-1] || target == dataRails[1][po-1]:
		if target == dataRails[0][po-1] {
			return keysRails[0][po-1]
		}

		if target == dataRails[1][po-1] {
			return keysRails[1][po-1]
		}

	case target > dataRails[0][po-1] && target > dataRails[1][po-1]:
		updateRails(0, 1)
		updateRails(1, 1)

		return getKeypds()

	case target < dataRails[0][po-1] && target < dataRails[1][po-1]:
		updateRails(0, -1)
		updateRails(1, -1)

		return getKeypds()

	case target < dataRails[0][po-1] && target > dataRails[1][po-1]:
		updateRails(0, -1)
		updateRails(1, 1)

		return getKeypds()

	case target > dataRails[0][po-1] && target < dataRails[1][po-1]:
		updateRails(0, 1)
		updateRails(1, -1)

		return getKeypds()

	default:
		return -1
	}

	return -1
}

// updateRails sets up the new slices that are potential ranges for next step
// * if the target is larger than the values at p position at both rails - get their right parts
// * if the target is smaller than the values at p position at both rails - get their left parts
// * if target is smaller than the value at p position in one slice and greater than the value in the other 
// get corresponding sides and create the new rails for the next step
// cases are handled by step0d and step1d values
func updateRails(key, step int) {
	if step == -1 {
		dataRails[key] = dataRails[key][:po]
		keysRails[key] = keysRails[key][:po]
	} else {
		dataRails[key] = dataRails[key][po:]
		keysRails[key] = keysRails[key][po:]
	}
}

// pdsDescription outputs the description of the "parallel-dynamic" search logic
// matches the doc file
func PdsDescription() string {
	return fmt.Sprint(
`"parallel-dynamic" implementation of the binary search algorithm uses:
* a map of two sub-slices of the data to make the checks at each step
sub-slices are copy of the original generated list (slice) and each of them contains half of the data
* a map of two sub-slices of data value keys from the original list; each slice in the map is half
of the size of the original data at the beginning, and matches the values in the oher map
with their original keys
* a pointer to move across the two subsices - rails of the data copy and compare values at their positions
with the target

At each step continuously updates each "rail" of the map containing the data,
as well as the "rails" inside the map with the keys
If the target is greater than both values at p position inside both rails with data:
* the data "rails" with greater values are kept - rest are moved
* keys "rails" are also updated to contain the keys for the values inside the data map slices
* pointer is moved to the right with half of its previous size and points to the new value
at the laf of the data map slices

If the target is smaller than both values at p position inside both rails:
* the data "rails" with smaller values are kept - rest are removed
* keys "rails" are also updated to contain the keys for the values inside the data map slices
* pointer is moved to the right with half of its previous size and points to the new value
at the laf of the data map slices
		   
If the target is smaller than value at p position in one of the "rails" and greater than the other:
* the corresponding halves of each "rail" are taken and the data map is updated to 
contain the new data
* keys map is updated as well to mirror the data map - to hold the keys of the values from the original slice
* pointer is moved with half of its previous value

After that the check of p with the values at p is done again and data is cut in halves each time
When the target is met - its key is extracted from the keys map is returned, otherwise -1
`)
}