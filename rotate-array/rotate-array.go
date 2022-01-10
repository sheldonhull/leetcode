func rotate(nums []int, k int)  {
    
    length:=len(nums)
    r := make([]int,len(nums))
    
    // fmt.Println("----- set all rotated values -----")
    start:=0
    if k == length { 
        return
    }
    if k > length { 
        k =  k % length
    }
    
    for i := k ; i > 0; i-- {    
        r[start]  = nums[length-i]
        // fmt.Printf("\t\ti %d r : %+v\n",i,r)
        start++
    }
    // fmt.Println("----- now fix all nonrotated values -----")
    for i := start; i < len(nums); i++ { 
        r[i] = nums[i-start]
        // fmt.Printf("\t\ti %d r : %+v\n",i,r)
    }

    // fmt.Printf("\trotated k:%+v\n",r)
    for i:=0; i < len(nums); i ++ { 
        nums[i] = r[i]
    }
    
}