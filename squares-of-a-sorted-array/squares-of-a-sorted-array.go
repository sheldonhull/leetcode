func sortedSquares(nums []int) []int {
    // fmt.Printf("array: %+v\n", array)
	
	squaredArr := make([]int,len(nums))
	for idx, val := range nums { 
		squaredArr[idx] = val * val
		// fmt.Printf("\tidx: %d val: %d\t%.2f\n", idx, val, sqval)
		
	}
	sort.Ints(squaredArr)
	return squaredArr
}


