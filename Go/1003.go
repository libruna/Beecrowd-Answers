package main
 
import (
    "fmt"
)
 
func main() {
    var a, b int
	
	fmt.Scanf("%d", &a)
	fmt.Scanf("%d", &b)
	
	soma := a + b
	fmt.Printf("SOMA = %d\n", soma)
}
