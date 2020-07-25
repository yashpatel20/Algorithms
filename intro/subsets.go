package intro

//Given a set of integers without repeated elements, return all possible subsets (power set) of the array

//Subsets find the subsets for a set
func Subsets(nums []int) [][]int {
	//final result
	result := make([][]int, 0)
	//intermediate results
	list := make([]int, 0)
	//fill up the result array
	backtrack(nums, 0, list, &result)
	//return the filled result
	return result
}

//nums given set
//pos the index of the element position to be added to the collection next time
//list temporary result (copy and save each time)
//result final result
func backtrack(nums []int, pos int, list []int, result *[][]int) {
	//copy the current list values to ans array
	ans := make([]int, len(list))
	copy(ans, list)
	//push the ans array to result
	*result = append(*result, ans)

	//Select , process the result and then cancel the selection
	for i := pos; i < len(nums); i++ {
		list = append(list, nums[i])
		backtrack(nums, i+1, list, result)
		list = list[0 : len(list)-1]
	}

}
