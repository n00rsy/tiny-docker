package main

import (
   "fmt"
   "os"
   "ubernetes/container"
)

func main(){
   if(len(os.Args) == 1){
      fmt.Println("unknown args. run ubernetes help for help")
      return
   }
   switch os.Args[1] {
      case "run":
	 fmt.Println("Initializing container...")
	 container.Init()
      case "child":
         fmt.Println("Launching container...")
	 container.Run()
      case "help":
	 fmt.Println("helping")
      default:
	 fmt.Println("unknown args. run ubernetes help for help")
   }
}
