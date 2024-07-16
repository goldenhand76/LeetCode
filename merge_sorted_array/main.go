package main

func main() {
	merge([]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	p1 := m - 1
	p2 := n - 1
	for i := 0; i <= n+m-1; i++ {
		if p2 == -1 {
			break
		}
		if p1 >= 0 {
			if nums1[p1] > nums2[p2] {
				nums1[m+n-1-i] = nums1[p1]
				p1--
			} else if nums1[p1] < nums2[p2] {
				nums1[m+n-1-i] = nums2[p2]
				p2--
			} else {
				nums1[m+n-1-i] = nums2[p2]
				if p1 < p2 {
					p2--
				} else {
					p1--
				}
			}
		} else {
			nums1[m+n-1-i] = nums2[p2]
			p2--
		}

	}
}
