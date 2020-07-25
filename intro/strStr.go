package intro

/*
Given a haystack string and a needle string, find the first position(starting from 0) where the needle
string appears in the haystack string. If it does not exist, -1 is returned.
*/

//StrStr returns the first index of the needle in haystring
func StrStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}

	var i, j int
	//outer loop is for iterating on haystack
	for i = 0; i < len(haystack)-len(needle); i++ {
		//inner loop is for checking every haystack char with needle char
		for j = 0; j < len(needle); j++ {
			if haystack[i+j] != needle[j] {
				break
			}
		}
		if j == len(needle) {
			return i
		}
	}

	return -1

}
