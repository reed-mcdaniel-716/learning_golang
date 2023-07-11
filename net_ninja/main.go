package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getInput(r *bufio.Reader, prompt string) (string, error) {
	fmt.Print(prompt)

	val, err := r.ReadString('\n')

	val = strings.TrimSpace(val)

	return val, err
}

func promptOptions(b bill) {
	reader := bufio.NewReader(os.Stdin)
	opt, _ := getInput(reader, "Choose option (a - add item, s - save bill, t - add tip): ")

	switch opt {
	case "a":
		fmt.Println("You choose to add an item.")
		itemName, _ := getInput(reader, "What is the name of the item? ")
		itemPrice, _ := getInput(reader, "What is the price of the item? ")

		castItemPrice, err := strconv.ParseFloat(itemPrice, 64)

		if err != nil {
			fmt.Println("Error parsing price as float64")
			promptOptions(b)
		}

		b.addItem(itemName, castItemPrice)
		fmt.Printf("You have added %q for $%v \n", itemName, castItemPrice)
		promptOptions(b)

	case "t":
		fmt.Println("You choose to add a tip.")
		tip, _ := getInput(reader, "What tip amount? ")

		castTip, err := strconv.ParseFloat(tip, 64)

		if err != nil {
			fmt.Println("Error parsing tip as float64")
			promptOptions(b)
		}

		b.updateTip(castTip)
		fmt.Printf("You have added a tip of $%v \n", tip)
		promptOptions(b)

	case "s":
		b.save()
		fmt.Printf("You choose to save the bill %q \n", b.name)
	default:
		fmt.Println("Your selection is not a valid option. Please try again.")
		promptOptions(b)
	}
}

func createBill() bill {
	// new reader on terminal input
	reader := bufio.NewReader(os.Stdin)

	name, _ := getInput(reader, "Create a new bill name: ")

	b := newBill(name)
	fmt.Println("Created the bill - ", b.name)

	return b
}

func main() {
	myBill := createBill()
	promptOptions(myBill)

	//fmt.Println("New bill:", myBill)
}
