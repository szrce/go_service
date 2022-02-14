/*
 * created @szrce simple tracking file like a  service
 * 
*/

package main

import "fmt"
import "os"
import "time"

func main(){

	fmt.Println("service started...\n")
	for{

		info, err := os.Stat("/root/test.txt")
		if err != nil {
			 fmt.Println(err); 
		}
		x := info.Size()

		if x > 50  {
			fmt.Printf("%d %s\n",x,"size limit reach")
			
			break
		}

		fmt.Printf("%s %d\n","file size:",x,)
		time.Sleep(5 * time.Second)
	}
}
