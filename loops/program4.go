/* To Print below pattern: 
   *
   **
   ***
   ****
   *****
   ******
   *******
*/

package main 

import "fmt"

func main() {
	var n int = 7
	
	for i := 0; i < n; i++ {

		for j := 0; j<=i ; j++{
			fmt.Print("*")
		} 
		fmt.Println()
	}
}