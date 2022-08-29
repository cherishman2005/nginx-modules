package main 
  
// Importing fmt and sync/atomic 
import ( 
    "fmt"
    "sync/atomic"
) 
  
// Main function 
func main() { 
  
    // Assigning value to uint32 
    var x uint32 = 18384411 
  
    // Using SwapUint32 method  
    // with its parameters 
    var old_val = atomic.SwapUint32(&x, 324233535) 
  
    // Prints new and old value 
    fmt.Println("Stored new value:", 
         x, ", Old value:", old_val) 
}