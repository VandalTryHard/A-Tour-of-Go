package main

import (
	"fmt"

	"../dates"
)

func main() {
	days := 3
	fmt.Println("\nYour appointment is in", days, "days with a follow-up in", days+dates.DaysInWeek, "days")
}
