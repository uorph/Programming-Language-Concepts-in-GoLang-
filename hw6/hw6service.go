// Copyright 2010 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// An edited version of code from the Go tutorials
// that supports a service that manages email contacts
//This code is edited by Regina Richardson (ik6089)

package main

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"
)

// Describes valid subpaths following
var validPath = regexp.MustCompile(`^/(add|lookup|bid)$`)

// Port for clients to connect to
const PORT_STR = ":8080"

// Takes a function that handles HTTP requests
// Wraps call to this function with code to ensure
// URL is valid
func makeHandler(fn func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r)
	}
}

// Where items, usernames, and bids are stored
var bidTable map[string]map[string]string

// Supports requests to add bids, username, and details
func addHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	desc := r.URL.Query().Get("desc")
	minBid := r.URL.Query().Get("minBid")
	bidTable[name] = make(map[string]string)
	bidTable[name]["bidder"] = ""
	bidTable[name]["desc"] = desc
	bidTable[name]["bid"] = "0"
	bidTable[name]["minBid"] = minBid
}

// Supports looking up an item and prints the map in the format requested in the example
func lookupHandler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	//fmt.Println(name)
	bidder := bidTable[name]["bidder"]
	desc := bidTable[name]["desc"]
	bid := bidTable[name]["bid"]
	minBid := bidTable[name]["minBid"]

	//fmt.Println(email)
	str := desc
	fmt.Fprint(w, "{", str, minBid, "{", bid, bidder, "}", "}")
	// fmt.Fprint(w, name, username, details, bid, minimumBid, bestBid)
}

//This function handles updating the bid, prior to doing so the minimum bid must be checked so that
//the attempted bid is larger. The must also be a check that the bid is larger than the previous bid

func bidHandler(w http.ResponseWriter, r *http.Request) {

	name := r.URL.Query().Get("name")
	bid := r.URL.Query().Get("bid")
	bidder := r.URL.Query().Get("bidder")
	currentBid := bidTable[name]["bid"]
	minBid := bidTable[name]["minBid"]
	intBid, err := strconv.Atoi(bid)
	intCurrentBid, err := strconv.Atoi(currentBid)
	intMinimumBid, err := strconv.Atoi(minBid)

	if err == nil {
		if intBid < intMinimumBid {
			fmt.Fprint(w, bid, " is below the minimum bid of ", minBid)

		} else if intBid <= intCurrentBid {
			fmt.Fprint(w, bid, " is not greater than the current high bid of ", currentBid, "\n")
		} else {
			bidTable[name]["bidder"] = bidder
			bidTable[name]["bid"] = bid
		}
	} else {
		fmt.Fprint(w, "looser")
	}

}

func main() {
	bidTable = make(map[string]map[string]string)
	http.HandleFunc("/lookup", makeHandler(lookupHandler))
	http.HandleFunc("/add", makeHandler(addHandler))
	http.HandleFunc("/bid", makeHandler(bidHandler))

	log.Fatal(http.ListenAndServe(PORT_STR, nil))
}
