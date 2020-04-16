func stringShift(s string, shift [][]int) string {
	var getshifty int
	for _, v := range shift {
			if v[0] == 0 {
					getshifty -= v[1]
			} else if v[0] == 1 {
					getshifty += v[1]
			}
	}
	getshifty = getshifty % len(s)
	tempstring := make([]byte, len(s))
	tempindex := (len(s) + getshifty) % len(s)
	for i:= 0; i < len(s); i++ {
			tempstring[tempindex] = s[i]
			tempindex = (tempindex + 1) % len(s)
	}
	return string(tempstring)
}