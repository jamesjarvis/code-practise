/** 
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad 
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

 func firstBadVersion(n int) int {
	min := 1
	for min!=n {
			mid := ((n-min) / 2) + min
			bad := isBadVersion(mid)
			if bad {
					n = mid
			} else {
					min = mid + 1
			}
	}
	return n
}
