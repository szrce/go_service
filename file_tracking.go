/*
 * created @szrce simple tracking file
 * 
*/

package main

import(
	"log"
    	"log/syslog"
    	"fmt"
    	"os"
    	"time"
)

func main(){
	//fmt.Println("service started...\n")
	syslogger, err := syslog.New(syslog.LOG_DEBUG, " @szrce service is started")

	if err != nil {
		 fmt.Println(err); 
		 return
	}

	for{

		log.SetOutput(syslogger)

		info, err := os.Stat("/root/test.txt")

		if err != nil {
			 fmt.Println(err); 
		}
		x := info.Size()

		if x > 50  {
			log.Println("size limit reach:",x)
			//fmt.Printf("%d %s\n",x,"size limit reach")
			break
		}

		log.Println("file size:",x)
		//fmt.Printf("%s %d\n","file size:",x)
		time.Sleep(5 * time.Second)
	}
}
