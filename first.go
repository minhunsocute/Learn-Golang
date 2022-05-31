package main

import "fmt"

func main() {
	var name = "Go conference" // Neu nhu khong su dung bien nay thi se thong bao loi
	const conferenceTickets = 50
	fmt.Println("Hello ", name, " World")
	fmt.Println("Get Your tickets % here to attend", name)
	fmt.Printf("Get Your tickets %v here to attend\n", name)

	fmt.Print(name)
}
