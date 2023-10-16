package problem88

func merge(nums1 []int, m int, nums2 []int, n int) {
	n1, n2, w, leftIsOver := m-1, n-1, len(nums1)-1, false

	if n1 < 0 {
		n1 = 0
		leftIsOver = true
	}

	for n2 >= 0 {
		if nums2[n2] >= nums1[n1] || leftIsOver {
			nums1[w] = nums2[n2]
			n2--
		} else {
			nums1[w], nums1[n1] = nums1[n1], nums1[w]
			n1--
		}

		w--
		if leftIsOver || n1 < 0 {
			n1, leftIsOver = w, true
		}
	}
}
