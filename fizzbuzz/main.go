import "strconv"

func fizzBuzz(n int) []string {
    var results []string
    for i:=1; i<=n;i++ {
        results = append(results, "")
        if i % 3 == 0 {
            results[i-1] = results[i-1] + "Fizz"
        }
        if i % 5 == 0 {
            results[i-1] = results[i-1] + "Buzz"
        }
        if i % 3 != 0 && i % 5 != 0 {
            results[i-1] = results[i-1] + strconv.Itoa(i)
        }
    }
    return results
}