// every go file must start with a package declaration this is how go organizes and reuse code
//it tells go that it's a dtabd alone executable program not a Libary 
package main

// format package important for IO operations like scanf printf, Println
import "fmt"
 

func main(){
  // this is a single line comment
  /* this is a double line comment*/
  fmt.Println("hello world")
}