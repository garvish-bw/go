package main
import (
"fmt"
"math/big"
)
func factorialIterative(n int64) *big.Int {
result := big.NewInt(1)
for i := int64(2); i <= n; i++ {
result.Mul(result, big.NewInt(i))
}
return result
}
func factorialRecursive(n int64) *big.Int {
if n == 0 || n == 1 {
return big.NewInt(1)
}
result := big.NewInt(n)
return result.Mul(result, factorialRecursive(n-1))
}
func main() {
var num int64
fmt.Print("Enter a number: ")
fmt.Scan(&num)
// Iterative factorial
fmt.Printf("Factorial of %d (Iterative): %v\n", num, factorialIterative(num))
// Recursive factorial
fmt.Printf("Factorial of %d (Recursive): %v\n", num, factorialRecursive(num))
}
