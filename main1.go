type ListNode struct{
        Val *int
        Next *ListNode
    }


func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    
    var listHead *ListNode;
    var listTail *ListNode;
    
    r:=999;
    firstNode := ListNode{
        Val: &r,
        Next: nil,
    }
    listHead =&firstNode;
    listTail =&firstNode;
    
    current1 := 0; // array1's pointer
    current2 := 0; // array2's pointer
    
    total_count :=0;
    //median_pos :=0
    
    //n1 := 0; // current number in array 1
    //n2 := 0; // current number in array 2
    
    l1 := len(nums1);
    l2 := len(nums2);
    
    // newN := 0;  // the current number taken
    medianL := listHead;
    medianR := listHead;
    
    for (current1<l1 || current2<l2){
        
        
            // move the median number to right
        /*
        if(total_count == 1){
            medianR = listHead.Next;
            medianL = listHead.Next;
        }else if(total_count%2==0){ //even 
            medianR = medianR.Next;
        }else{
            medianL = medianL.Next; // odd
        }*/
        
        
        
        
        
        // number already exists, skip
        //temp := listTail.Val;
        /*
        if(current1 < l1 && *temp == nums1[current1]){
            current1++;
            continue;
        }
        if(current2 < l2 && *temp == nums2[current2]){
            current2++;
            continue;
        }*/
        // if current 1 finished
        if(current1 == l1){ 
            //newN = nums2[current2];
            
            // move pointer
            newNode := ListNode{
                Val: &nums2[current2],
                Next: nil,
            }
            listTail.Next = &newNode;
            listTail = listTail.Next;
            
            current2++;
            
        } else if(current2 == l2){ // current 2 finished
            newNode := ListNode{
                Val: &nums1[current1],
                Next: nil,
            }
            listTail.Next = &newNode;
            listTail = listTail.Next;
            
            current1++;
            
        } else if(nums1[current1] <= nums2[current2]){ // one is smaller, has result
            newNode := ListNode{
                Val: &nums1[current1],
                Next: nil,
            }
            listTail.Next = &newNode;
            listTail = listTail.Next;
            
           // newN = nums1[current1];
            current1++; // array advance
           // fmt.Println(newNode);
            
        } else if( nums1[current1] > nums2[current2]){ // take n2
             newNode := ListNode{
                Val: &nums2[current2],
                Next: nil,
            }
            listTail.Next = &newNode;
            listTail = listTail.Next;
            
            //newN =  nums2[current2];
            current2++;
            
        } else if( nums1[current1] == nums2[current2]){
             newNode := ListNode{
                Val: &nums1[current1],
                Next: nil,
            }
            listTail.Next = &newNode;
            listTail = listTail.Next;
           // newN =  nums1[current1];
            current1++;
            current2++;
            
        }
        total_count++;
        
        //fmt.Printf("%d-",newN);
        
        // print the list
        
        
        
        // move the median number to right
        
        if(total_count == 1){
            medianR = listHead.Next;
            medianL = listHead.Next;
        }else if(total_count%2==0){ //even 
            medianR = medianR.Next;
            //a:=medianR.Val;
            //fmt.Printf("Total %d, move R, now %d\n", total_count, *a);
        }else{
            medianL = medianL.Next; // odd
            //a:=medianR.Val;
            //fmt.Printf("Total %d, move L, now %d\n", total_count, *a);
        }

    }
    mR := medianR.Val;
    mL := medianL.Val;
    //fmt.Printf("[%d:%d]\n", *mR,*mL);
    return float64(*mR+*mL)/2;
    //fmt.Printf("  ||  ");
    //printListNode(listHead);
    //fmt.Println();
    //fmt.Printf("[%d:%d]\n", medianL,medianR);
    //return float64(medianL+medianR)/2;
}

func printListNode(h *ListNode){
    for h !=nil{
        //fmt.Printf("%p->", h.Val);
        
        //p := h.Val;
        //fmt.Printf("%v->", *p);
       // fmt.Println(h);
        h=h.Next;
    }
}
