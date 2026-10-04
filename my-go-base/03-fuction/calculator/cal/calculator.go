package cal
//import "fmt"
import "errors"
func Add(a int, b int) float64 {
	return float64(a + b)
}
func Sub(a int, b int) float64 {
	return float64(a - b)
}
func Mul(a int, b int) float64 {
	return float64(a * b)
}
func Div(a int, b int) (float64,error) {
	if b == 0 {
		return 0,errors.New("division by zero")
	}
	return float64(a / b),nil
}