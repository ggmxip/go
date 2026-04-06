package main

import (
	"fmt"
)

func main() {
	INLOGGED := true
	ISAMDIN := true
	HASSUBSCRIPTION := false

	// && operator returns true if both the operands are true
	CANACCESSDASHBOARD := INLOGGED && HASSUBSCRIPTION
	CANDELETEPOST := (INLOGGED && HASSUBSCRIPTION) || ISAMDIN
	fmt.Println("can access dashboard:", CANACCESSDASHBOARD)
	fmt.Println("can delete post:", CANDELETEPOST)

	age := 99
	drivingLicense := 60 > age && age > 18
	fmt.Println("can i drive: ", drivingLicense)
}
