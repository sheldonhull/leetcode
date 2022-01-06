
func removeDuplicates(nums []int) int {
    // fmt.Printf("original: %+v\n",nums)
    
    // placeholder is set to 999 per constraints of test <= 100 to ensure goes to end
    const placeholder = 999    
    var uniqueCounter int
    
    
    // iterate through array
    for i := range nums { 
        // fmt.Printf("\ti = %d\n",i)
        
        if i+1 > len(nums) {
            // fmt.Printf("\tbreak: reached end of array for compare")
            break
        }
        // while continue to find duplicates, keep setting placeholder and resorting till this
        // condition is no longer met
        lookAhead := 1
        
        // start evaluating from current item forward
        // and look for match until doesn't exist then break to next item in array

        for {
               
            // However, to avoid looking beyond last item in array, ie panic
            // make sure to break out of assessment once is last item
            // avoid panic and ensure exit early if we've reached end of array for compare
            if i+lookAhead >= len(nums) {
                break
            }
            
            if nums[i] == placeholder { 
                // fmt.Printf("\tplaceholder reached, no need to compare further\n")
                break
            }
            if nums[i] != nums[i+lookAhead] { 
                // fmt.Printf("\tno dup found\n")
                break
            }
            
            // if current item equals the next assessed value then set to the placeholder and 
            // mark dupCounter + 1 to continue looking at next one
            // fmt.Printf("\tDUP%d == %d (nums[%d] == nums[%d+%d]\n",
            // nums[i],
            // nums[i+lookAhead],
            // i,
            // i,
            // lookAhead)
            nums[i+lookAhead] = placeholder
            // fmt.Printf("\t\tplaceholder set: nums: %+v\n",nums)
            lookAhead++
         
        }
        
        // fmt.Printf("\tsort and start eval on next")
        // could jump ahead in array 
        // and not sort each time but not optimizing based on constraints right now

    }
    // to further optimize could jump ahead in array based on the lookahead
    // but this works for now based on constraints
    sort.Ints(nums)
    // fmt.Printf("count unique in loop now\n")
    // now that sorting is done, count all the entries until we get to placeholders as unique
    for i := range nums {
        if nums[i] != placeholder { 
            uniqueCounter ++
            // fmt.Printf("\tnums[%d]: %d -> uniqueCounter: %d\n",i,nums[i],uniqueCounter)
            continue
        }
        // fmt.Printf("\treached placeholders, no longer evaluating\n")
        break
    }
    // fmt.Printf("uniqueCounter: %d\tmodified: %+v\n",uniqueCounter, nums)
    return uniqueCounter
}