package mathematics

//https://leetcode.com/explore/interview/card/top-interview-questions-easy/102/math/744/
//Count the number of prime numbers less than a non-negative number, n.
//
//Example:
//
//Input: 10
//Output: 4
//Explanation: There are 4 prime numbers less than 10, they are 2, 3, 5, 7.
func countPrimes(n int) int {
	nonPrime := make(map[int]struct{}, n)
	count := 0
	for i := 2; i < n; i++ {
		if _, ok := nonPrime[i]; !ok {
			count++
			for j := 2; i*j < n; j++ {
				nonPrime[i*j] = struct{}{}
			}
		}
	}
	return count
}
