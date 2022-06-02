package main

import (
	"fmt"
	"strconv"
	"strings"
)

// stuct
type DataOfUser struct {
	firstName string
	lastName  string
	email     string
}

//list of maps
var listUsers = make([]map[string]string, 0)

//list of struct
var listUsers1 = make([]DataOfUser, 0)

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
	//a, b, c = helper.returnManyValues(float32(10), float32(2))
	fmt.Printf("%f %f %f", a, b, c)

	//Maps syntax
	// var a = map[keytype]valueType{key1: value1, key2: value2 , key3: value3}
	// b := map[keytype]valueType{key1: value1, key2: value2 , key3: value3}
	// c:= make(map[keyType]valueType)
	// Call map
	// a[keytype]
	var userData = make(map[string]string)

	userData["firstname"] = "Nguyen"
	userData["lastName"] = "Hung"
	userData["email"] = "hungnguyen.201102ak@gmail.com"

	fmt.Println(userData["firstname"])

	//list of maps
	var count int = 0
	//var listUsers = make([]map[string]string, 0)
	for {
		if count > 2 {
			break
		} else if count == 1 {
			go pintGo()
		}
		var first_name, last_name, email = getUserInput(count)
		var user_data = make(map[string]string)
		user_data["firstName"] = first_name
		user_data["lastName"] = last_name
		user_data["Email"] = email

		listUsers = append(listUsers, user_data)
		count++
	}
	printListUserOfMap()
	count = 0
	for {
		if count > 2 {
			break
		}
		var first_name, last_name, email = getUserInput(count)
		var userOfData = DataOfUser{firstName: first_name,
			lastName: last_name,
			email:    email}
		listUsers1 = append(listUsers1, userOfData)
		count++
	}
	printListUserOfStruct()
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

// func returnManyValues(a float32, b float32) (float32, float32, float32) {
// 	return (a + b), (a - b), (a / b)
// }

func getUserInput(count int) (string, string, string) {
	var firstName string
	var lastName string
	var email string

	fmt.Println("Enter your first name of " + strconv.Itoa(count) + ": ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last nameof " + strconv.Itoa(count) + ": ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email addressof " + strconv.Itoa(count) + ": ")
	fmt.Scan(&email)

	for {
		if strings.Contains(email, "@") {
			break
		} else {
			fmt.Println("Enter again email because don't have @ in emain:")
			fmt.Scan(&email)
		}
	}
	return firstName, lastName, email
}

func printListUserOfMap() {
	for _, value := range listUsers {
		fmt.Println(value["firstName"])
		fmt.Println(value["lastName"])
		fmt.Println(value["Email"])
		fmt.Println("-------------------")
	}
}

func printListUserOfStruct() {
	for _, value := range listUsers1 {
		fmt.Println(value.firstName)
		fmt.Println(value.lastName)
		fmt.Println(value.email)
		fmt.Println("-------------------")

	}
}

// sẽ được gọi sau khi gọi hàm go

func pintGo() {
	fmt.Print(0)
}
