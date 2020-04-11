func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
 }
 
 type stack []rune
 
 func (s stack) Empty() bool { return len(s) == 0 }
 func (s stack) Peek() rune   { return s[len(s)-1] }
 func (s *stack) Put(i rune)  { (*s) = append((*s), i) }
 func (s *stack) Pop() {
		 (*s) = (*s)[:Max(len(*s)-1,0)]
 }
 
 func remove(s string) string {
		 var t stack
		 for _, v := range s {
				 if (v == rune('#')) {
						 t.Pop()
				 } else {
						 t.Put(v)
				 }
		 }
		 return string(t)
 }
 
 func backspaceCompare(S string, T string) bool {
 //     This is a classic stack problem
		 return remove(S) == remove(T)
 }