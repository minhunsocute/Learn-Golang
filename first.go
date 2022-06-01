package main

import (
	"fmt"
	"strings"
)

func main() {
	name1 := "MinHung"
	var name = "Go conference" // Neu nhu khong su dung bien nay thi se thong bao loi
	var firstName string = "Hung"
	var lastName string = "Nguyen"
	const conferenceTickets = 50

	GreatUserPara("Minh Hung")
	greatUsers()
	fmt.Println("Hello ", name1, " World")
	fmt.Println("Get Your tickets % here to attend", name)
	fmt.Printf("Get Your tickets %v here to attend\n", name)

	fmt.Print(name)

	var userName string
	var userTickets int

	fmt.Scan(&userName) // Nhap gia tri cho bien username
	userTickets = 2
	fmt.Printf("User %v booked %v %T tickets.\n", userName, userTickets, userTickets) //%v in ra value %T in ra type cua var
	fmt.Println(userName)                                                             // In ra gia tri cua username
	fmt.Print(&userName)                                                              // In ra dia chi cua UserName

	//arrays
	var bookings = [3]string{"Nana", "Nicola", "Petter"} // 50 la size cua mang
	booking := []string{}                                // Khi dung len gan nay thi khong co var o truoc
	// hoac co the viet :
	// var booking  []string
	// except

	for i := 0; i < 5; i++ {
		booking = append(booking, "")
	}
	booking[0] = "Nana" // Set index 0 for array
	booking[1] = firstName + lastName
	fmt.Println(bookings[0])
	fmt.Println(booking[1])
	fmt.Println(len(booking)) // len lay do dai cua mang

	booking = append(booking, firstName+" "+lastName) // Chi ap dung cho mang chua gan cho gia tri length
	// for - loops
	for i := 0; i < len(booking); i++ {
		for j := 0; j < len(bookings); j++ {
			fmt.Println(booking[i] + " - " + bookings[j])
		}
	}
	// except of for loop
	firstNames := []string{}
	firstNames = printFirstName(booking)
	// for _, value := range bookings {
	// 	//fmt.Println(value)
	// 	var names = strings.Fields(value)
	// 	firstNames = append(firstNames, names[0])
	// }
	for i := 0; i < len(firstNames); i++ {
		fmt.Println(firstNames[i])
	}

	//If else

	//noTicketsRemaining := userTickets == 2
	if userTickets == 2 {
		fmt.Println("Our conference is booked out.Come back next year")
	} else {
		fmt.Println("Okay")
	}

	//switch

	city := "London"

	switch city {
	case "New York":
		//Some code here
	case "Singapore":
		// Some code here
	case "London":
		// Some code here
	default:
		fmt.Print("No CiTy")
	}
	a := float32(0)
	b := float32(0)
	c := float32(0)
	a, b, c = returnManyValues(float32(10), float32(2))
	fmt.Printf("%f %f %f", a, b, c)
}

// Functions
func greatUsers() {
	fmt.Println("Welcome to our conference")
}

//Functions Paramters

func GreatUserPara(confName string) {
	fmt.Printf("Welcome to %v booking application", confName)
}

//Functions with returns values
func printFirstName(booking []string) []string {
	firstNames := []string{}
	for _, value := range booking {
		var names = strings.Fields(value)
		firstNames = append(names)
	}
	return firstNames
}

// Functions return many values
