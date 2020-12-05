//write on 2020/12/5
//package main
//
//import (
//	"os"
//	"fmt"
//)
//
//func main(){
//	var s, seq  string
//	for i := 1; i< len(os.Args); i++{
//		s += seq + os.Args[i]
//		seq = " "
//	}
//	fmt.Println(s)
//}

package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
