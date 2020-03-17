package test

func sanitizeIndex(index int, min int, max int) int {
	if index > max {
		return max
	}
	if index < min {
		return min
	}
	return index
}

func binSearch(nums []int, target int) int {
	low := 0
	high := len(nums) - 1

	for low <= high {
		if low == high {
			break
		}

		median := (low + high) / 2
		value := nums[median]

		if value == target {
			return median + 1
		}
		if value < target {
			low = median + 1
		} else {
			high = median - 1
		}
	}

	if target < nums[low] {
		return low
	}
	return low + 1
}

func median(nums []int) float64 {
	size := len(nums)
	m := size / 2

	if size%2 == 1 {
		return float64(nums[m])
	}

	return float64(nums[m-1]+nums[m]) / 2
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	size1 := len(nums1)
	size2 := len(nums2)
	// low and high iterators
	l1 := 0
	l2 := 0
	h1 := size1 - 1
	h2 := size2 - 1

	// empty conditions
	if h1 == -1 {
		return median(nums2)
	}
	if h2 == -1 {
		return median(nums1)
	}

	// extra conditions
	if h1 == h2 && h1 == 0 {
		return float64(nums1[0]+nums2[0]) / 2
	}

	// target indexes
	var t1 int = (size1 + size2) / 2
	t2 := -1

	if (h1+h2)%2 == 0 {
		t2 = t1
		t1--
	}

	var r1, r2 int
	r1Found := false
	r2Found := false
	turn := true

	// fmt.Println("Targets: ", t1, t2)

	// for i := 0; i < 7; i++ {
	for {
		if turn {
			//calc
			m := (l1 + h1) / 2
			s := binSearch(nums2, nums1[m])
			// fmt.Println("1: ", l1, h1, l2, h2, m, s)

			// check
			if m+s == t1 {
				r1Found = true
				r1 = nums1[m]

				if !r2Found && t2 != -1 {
					// search is equal condition
					s0 := sanitizeIndex(s, 1, size2)
					m0 := sanitizeIndex(m, 0, size1-1)
					v1 := nums1[m0]
					v2 := nums2[s0-1]

					if v1 == v2 {
						if m0 < h1 {
							v1 = nums1[m0+1]
						}
						if s0 <= h2 {
							v2 = nums2[s0]
						}
						r2Found = true
						if v1 < v2 {
							r2 = v1
						} else {
							r2 = v2
						}
					}
					// next is equal condition
					if m < h1 && nums1[m] == nums1[m+1] {
						r2Found = true
						r2 = nums1[m]
					}
				}
			} else if m+s == t2 {
				r2Found = true
				r2 = nums1[m]
			}

			// it
			if m+s > t1 {
				h1 = m - 1
			} else {
				l1 = m + 1
			}

		} else {
			// calc
			m := (l2 + h2) / 2
			s := binSearch(nums1, nums2[m])
			// fmt.Println("2: ", l1, h1, l2, h2, m, s)

			// check
			if m+s == t1 {
				r1Found = true
				r1 = nums2[m]

				if !r2Found && t2 != -1 {
					// search is equal condition
					s0 := sanitizeIndex(s, 0, size1-1)
					m0 := sanitizeIndex(m, 1, size2-1)
					v1 := nums1[s0]
					v2 := nums2[m0-1]

					if v1 == v2 {
						if m0 < h2 {
							v2 = nums2[m0]
						}
						if s <= h1 {
							v1 = nums1[s0+1]
						}
						r2Found = true
						if v1 < v2 {
							r2 = v1
						} else {
							r2 = v2
						}
					}
					// next is equal condition
					if m < h1 && nums1[m] == nums1[m+1] {
						r2Found = true
						r2 = nums1[m]
					}
				}
			} else if m+s == t2 {
				r2Found = true
				r2 = nums2[m]
			}

			// it
			if m+s > t1 {
				h2 = m - 1
			} else {
				l2 = m + 1
			}
		}

		// fmt.Println("R:", r1Found, r2Found, r1, r2)
		if r1Found && (r2Found || t2 == -1) {
			break
		}
		if l1 > h1 && l2 > h2 {
			if turn {
				if !r1Found {
					r1 = nums1[l1]
				} else {
					r2 = nums1[l1]
				}
			} else {
				if !r1Found {
					r1 = nums2[l2]
				} else {
					r2 = nums2[l2]
				}
			}
			r2 = nums2[l2]
			break
		}
		turn = !turn
	}

	if t2 == -1 {
		return float64(r1)
	}
	return (float64(r1 + r2)) / 2
}
