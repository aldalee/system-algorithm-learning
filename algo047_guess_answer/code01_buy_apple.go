package algo047_guess_answer

func MinBagAwesome(apple int) int {
	if apple&1 != 0 { // 奇数返回-1
		return -1
	}
	switch apple {
	case 6, 8:
		return 1
	case 12, 14, 16:
		return 2
	case 2, 4, 10:
		return -1
	default:
		return (apple-18)/8 + 3
	}
}

func MinBag(apple int) int {
	bag8 := apple / 8
	rest := apple - bag8*8
	for bag8 >= 0 {
		if rest%6 == 0 {
			return bag8 + rest/6
		} else {
			bag8--
			rest += 8
		}
	}
	return -1
}
