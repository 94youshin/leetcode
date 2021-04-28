/**
 * 不要纠结局部，看整体
 * 左边比右边小,右边的值减掉
 * 左边比右边大,累加
 */
func romanToInt(s string) int {
	var sum int
	preNum := getValue(s[0])
	for i := 1; i < len(s); i++ {
		curNum := getValue(s[i])
		if preNum < curNum {
			sum -= preNum
		} else {
			sum += preNum
		}
		preNum = curNum
	}
	sum += preNum
	return sum
}

func getValue(n uint8) int {
	switch string(n) {
	case "I":
		return 1
	case "V":
		return 5
	case "X":
		return 10
	case "L":
		return 50
	case "C":
		return 100
	case "D":
		return 500
	case "M":
		return 1000
	default:
		return 0
	}
}

