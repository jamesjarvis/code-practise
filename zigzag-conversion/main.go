func convert(s string, numRows int) string {
	if numRows == 1 {
			return s
	}
	var newS []byte
	cycleLength := 2 * numRows - 2
	for i := 0; i < numRows; i++ {
			for j := 0; j + i < len(s); j+=cycleLength {
					newS = append(newS, s[j+i])
					if i!= 0 && i != numRows - 1 && j + cycleLength - i < len(s) {
							newS = append(newS, s[j + cycleLength - i])
					}
			}
	}
	return string(newS)
}