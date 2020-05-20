
const solution = (arr) => {
  let breadth = 2
  let floor = 1
  let left = 0
  let right = 0
  
  while (floor < arr.length) {
      let level = arr.slice(floor,(floor+breadth))
      let pivot = breadth / 2
      
      left += level.slice(0,pivot).reduce((a, b) => a + b, 0)
      right += level.slice(pivot).reduce((a, b) => a + b, 0)
      
      floor += breadth
      breadth *= 2
  }
  
  if (left < right) {
      return "Right"
  }
  if (left > right) {
      return "Left"
  }
  return ""
};
