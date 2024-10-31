package algo

/** 
两两比较，每一轮把未排序的前面那部分中的最大元素移到尾部有序部分
*/

func BubbelSort(nums []int) {
	for i:=len(nums)-1; i>0; i-- {
		for j:=1; j<=i; j++ {
			if nums[j] < nums[j-1] {
				nums[j], nums[j-1] = nums[j-1], nums[j]
			}
		}
	}
	
}