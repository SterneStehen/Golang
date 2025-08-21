package main

import "fmt"


 type ListNode struct {
      Val int
      Next *ListNode
  }

// func ft_push(res *ListNode,  list *ListNode, dig int) *ListNode{
	
// }

func ft_push(resuList **ListNode , Val int){
	
	newList := &ListNode{Val: Val, Next: nil}
	if *resuList == nil{
		fmt.Println(Val)
		*resuList = newList
		return
	}
	curr := *resuList
	for  curr.Next != nil{ 
		curr = curr.Next
	}
	curr.Next = newList
	fmt.Println(Val)

}
 
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    //result := list1
	var resList *ListNode
	//topList2 := list2
	// i := list1
	// n := list2


	for list1 != nil && list2 != nil{
		if list1.Val < list2.Val{
			ft_push(&resList, list1.Val)
			list1 = list1.Next
			continue
		} else if list1.Val > list2.Val{
			ft_push(&resList, list2.Val)
			list2 = list2.Next
			continue
		} else if list1.Val == list2.Val{
			ft_push(&resList, list2.Val)
			ft_push(&resList, list1.Val)
			list1 = list1.Next
			list2 = list2.Next
			continue
		}
	}
	if list1 != nil{
		for i := list1; i != nil; i = i.Next{
			ft_push(&resList, i.Val)
		}
	}
	if list2 != nil{
		for i := list2; i != nil; i = i.Next{
			ft_push(&resList, i.Val)
		}
	}
	return  resList
}
	
	
	

func main(){
	
	list1 := &ListNode{Val:1}
	list1.Next = &ListNode{Val:3}
	list1.Next.Next = &ListNode{Val:6}

	list2 := &ListNode{Val:1}
	list2.Next = &ListNode{Val:7}
	list2.Next.Next = &ListNode{Val:8}

	_ = mergeTwoLists(list1, list2)

	// for n := listResult; n != nil; n = n.Next{
	// 	fmt.Println(n.Val)
	// } 
	

}