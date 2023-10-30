package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study")

	presentDate := time.Now()

	fmt.Println(presentDate) //2023-10-26 09:30:31.6386761 +0530 IST m=+0.004658801

	//Must use 01-02-2006 for time and 15-04-05 for date
	fmt.Println(presentDate.Format("01-02-2006"))          //10-26-2023
	fmt.Println(presentDate.Format("01-02-2006 15-04-05")) //10-26-2023 09-32-27

	//To create a new date
	craeteDate := time.Date(2021, time.March, 30, 9, 1, 12, 0, time.Local)
	fmt.Println(craeteDate) //2021-03-30 09:01:12 +0530 IST

	fmt.Println(craeteDate.Format("01-02-2006 15-04-05 Monday")) //03-30-2021 09-01-12 Tuesday

}
