// Client for hw6 -- command line program to manage bids for items

package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

// port of service where we send requests
const PORT = 8080

// Prompt user for item to bid on and for their bid
func bid(scanner *bufio.Scanner, username string) {
	fmt.Print("Please enter the name: ")
	scanner.Scan()
	name := scanner.Text()
	fmt.Print("Amount to bid, in $US: ")
	scanner.Scan()
	bid := scanner.Text()
	//Formatting string for request API
	request := fmt.Sprintf("http://localhost:8080/bid?name=%v&bid=%v&bidder=%v", url.PathEscape(name), url.PathEscape(bid), url.PathEscape(username))
	//Requesting a bid offer and chechking to see if any error occur
	resp, err := http.Get(request)
	if err != nil {
		fmt.Printf("%v", err)
		return
	}
	//Reading repsonse made by API
	body, err := io.ReadAll(resp.Body)
	desc := string(body)
	//Printing any response sent back by the API
	fmt.Printf(desc)
}

// Lookup a item and recive the details, current bid, minimum bid, and bidder
func lookup(scanner *bufio.Scanner, username string) {
	fmt.Print("Please enter the name: ")
	scanner.Scan()
	name := scanner.Text()
	//Formatting string for request API
	request := fmt.Sprintf("http://localhost:8080/lookup?name=%v", url.PathEscape(name))
	//Requesting a item information and checking to see if any error occur
	resp, err := http.Get(request)
	if err != nil {
		fmt.Printf("%v", err)
		return
	}
	//Reading repsonse made by API
	body, err := io.ReadAll(resp.Body)
	desc := string(body)
	//Printing any response sent back by the API depending on the conditions
	if "" == desc {
		fmt.Printf("%v has no item details\n", name)
	} else {
		fmt.Printf("The item details are: %v\n", string(body))
	}
}

func main() {
	//client := new(http.Client)

	var cmd string
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Printf("Please enter your username: ")
	scanner.Scan()
	username := scanner.Text()

	quit := "no"
	for quit == "no" {
		fmt.Printf("Bid or Lookup (b/l)? ")
		fmt.Scanf("%s ", &cmd)
		if "b" == cmd {
			bid(scanner, username)
		} else if "l" == cmd {
			lookup(scanner, username)
		}

		fmt.Print("Do you want to quit? (yes/no): ")
		fmt.Scanf("%s ", &quit)
	}

}
