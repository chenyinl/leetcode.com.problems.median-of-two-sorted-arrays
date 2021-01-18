func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    l1 := len(nums1);
    l2 := len(nums2);
    l12 := l1+l2;
    even := ((l12)%2==0);
    max := (l12)/2+1;
    // fmt.Printf("Max: %d ",max);
    current1 := 0; // array1's pointer
    current2 := 0; // array2's pointer
    
    total_count :=0;
    
    newN    :=0;
    medianL :=0;
   // medianR :=0;
    
    for ((current1<l1 || current2<l2) && total_count<max){

        // if current 1 finished
        if(current1 == l1  ){ 
            newN = nums2[current2];
            current2++;
            
        } else if(current2 == l2 || nums1[current1] <= nums2[current2] ){ // current 2 finished
            newN = nums1[current1];
            current1++;
            
        } else if( nums1[current1] > nums2[current2]){ 
            newN = nums2[current2];
            current2++;
            
        }/* else if(nums1[current1] <= nums2[current2] ){ // current 2 finished
            newN = nums1[current1];
            current1++;
        }*/
        total_count++;
        
        if(even && total_count == max-1){ // 
            medianL = newN;
        }/* else if(total_count == max && even){
            medianR = newN;
           // return float64(medianL+newN)/2;
            
        }else if(total_count == max && !even){ // odd and this is the last number
            //return float64(newN);
            medianR = newN;
        }*/
    }
    //fmt.Printf("medianR: %d; newN: %d\n", medianR, newN);
    if(even){
        return float64(medianL+newN)/2;
        //return float64(medianL+medianR)/2;
    }else{
        return float64(newN);
        //return float64(medianR);
    }
   // fmt.Printf("[%d:%d]\n", medianL,medianR);
    //return float64(0);
}
