package main

import (
   "fmt"
   "os"
   "ubernetes/lib"
)

func main(){
   if(len(os.Args) == 1){
      fmt.Println("unknown args. run ubernetes help for help")
      return
   }
   switch os.Args[1] {
      case "run":
	 fmt.Println("running for real")
	 lib.Make_Container()
      case "help":
	 fmt.Println("helping")
      default:
	 fmt.Println("unknown args. run ubernetes help for help")
   }
}
