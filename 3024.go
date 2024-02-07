func triangleType(nums []int) string {
	if !((nums[0]+nums[1] > nums[2]) && (nums[1]+nums[2] > nums[0]) && (nums[0]+nums[2] > nums[1])) {
		return "none"
	}
	sch := 0
	if nums[0] == nums[1] {
		sch++
	}
	if nums[1] == nums[2] {
		sch++
	}
	if nums[2] == nums[0] {
		sch++
	}
	if sch == 3 {
		return "equilateral"
	} else if sch == 1 {
		return "isosceles"
	} else {
		return "scalene"
	}
}