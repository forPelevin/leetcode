package array

import "errors"

//https://leetcode.com/problems/queries-on-a-permutation-with-key/
//Given the array queries of positive integers between 1 and m, you have to process all queries[i] (from i=0 to i=queries.length-1) according to the following rules:
//
//In the beginning, you have the permutation P=[1,2,3,...,m].
//For the current i, find the position of queries[i] in the permutation P (indexing from 0) and then move this at the beginning of the permutation P. Notice that the position of queries[i] in P is the result for queries[i].
//Return an array containing the result for the given queries.
//Example 1:
//
//Input: queries = [3,1,2,1], m = 5
//Output: [2,1,2,1]
//Explanation: The queries are processed as follow:
//For i=0: queries[i]=3, P=[1,2,3,4,5], position of 3 in P is 2, then we move 3 to the beginning of P resulting in P=[3,1,2,4,5].
//For i=1: queries[i]=1, P=[3,1,2,4,5], position of 1 in P is 1, then we move 1 to the beginning of P resulting in P=[1,3,2,4,5].
//For i=2: queries[i]=2, P=[1,3,2,4,5], position of 2 in P is 2, then we move 2 to the beginning of P resulting in P=[2,1,3,4,5].
//For i=3: queries[i]=1, P=[2,1,3,4,5], position of 1 in P is 1, then we move 1 to the beginning of P resulting in P=[1,2,3,4,5].
//Therefore, the array containing the result is [2,1,2,1].
//Constraints:
//
//1 <= m <= 10^3
//1 <= queries.length <= m
//1 <= queries[i] <= m

const notFoundIndex = -1
var errInvalidRange = errors.New("invalid interval boundaries")

func processQueries(queries []int, m int) []int {
	p, err := Range(1,m)
	if err != nil {
		return nil
	}

	for i, q := range queries {
		foundI, pVal := find(p, q)
		queries[i] = foundI

		// Move the found val at the beginning
		if len(p) > 1 {
			p = append([]int{pVal}, append(p[:foundI], p[foundI+1:]...)...)
		}
	}

	return queries
}

func find(items []int, search int) (int, int) {
	for i, v := range items {
		if v == search {
			return i, v
		}
	}
	return notFoundIndex, 0
}

func Range(from,to int) ([]int, error) {
	if from > to {
		return nil,nil
	}

	r := make([]int, 0, to-from)
	for i := from; i <= to; i++ {
		r = append(r, i)
	}
	return r, nil
}
