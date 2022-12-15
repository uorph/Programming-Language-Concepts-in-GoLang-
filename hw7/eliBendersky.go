// Regina Richardson ik6089
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
)

// This struct object contains a individual Mutex, owner string, balance int for each unique key in our accounts store object
type Account struct {
	sync.Mutex
	owner   string
	balance int
}

// This struct object holds a map with key-value pair of string and account object
type AccountStore struct {
	accountMap map[string]*Account
}

func (cs *AccountStore) get(w http.ResponseWriter, req *http.Request) {
	log.Printf("get %v", req)
	//Getting and setting the proper parameters
	name := req.URL.Query().Get("name")
	//balance := req.URL.Query().Get("balance")
	if _, ok := cs.accountMap[name]; ok {
		//Locking to ensure proprietary access to our account object
		cs.accountMap[name].Lock()
		//Printing in format requested
		fmt.Fprint(w, name, ": owner: ", cs.accountMap[name].owner, " balance: $", cs.accountMap[name].balance)
		cs.accountMap[name].Unlock()
	} else {
		fmt.Fprintf(w, "%s not found\n", name)
	}
}

func (cs *AccountStore) set(w http.ResponseWriter, req *http.Request) {
	log.Printf("set %v", req)
	//Getting and setting the proper parameters
	name := req.URL.Query().Get("name")
	bal := req.URL.Query().Get("balance")
	// Converting balance from a string to an integer and checking for error
	intval, err := strconv.Atoi(bal)
	if err != nil {
		fmt.Fprintf(w, "%s\n", err)
	} else {
		if entry, ok := cs.accountMap[name]; ok {
			//Locking to ensure proprietary access to our account object
			cs.accountMap[name].Lock()
			entry.balance = intval
			cs.accountMap[name] = entry
			fmt.Fprintf(w, "ok\n")
			//Unlocking proprietary access to our account object
			cs.accountMap[name].Unlock()
		}
	}
}

func (cs *AccountStore) inc(w http.ResponseWriter, req *http.Request) {
	log.Printf("inc %v", req)
	//Getting and setting the proper parameters
	name := req.URL.Query().Get("name")
	amount := req.URL.Query().Get("amt")
	// Converting amount from string to integer and checking for error
	intval, err := strconv.Atoi(amount)
	if err != nil {
		fmt.Fprintf(w, "%s\n", err)
	}
	// Checking for valid account and balance and incrementing the balance by amount
	if entry, ok := cs.accountMap[name]; ok {
		//Locking to ensure proprietary access to our account object
		cs.accountMap[name].Lock()
		entry.balance += intval
		cs.accountMap[name] = entry
		fmt.Fprintf(w, "ok\n")
		//Unlocking proprietary access to our account object
		cs.accountMap[name].Unlock()
	} else {
		fmt.Fprintf(w, "%s not found\n", name)
	}
}

func (cs *AccountStore) pay(w http.ResponseWriter, req *http.Request) {
	log.Printf("set %v", req)
	//Getting and setting the proper parameters
	from := req.URL.Query().Get("from")
	to := req.URL.Query().Get("to")
	amt := req.URL.Query().Get("amt")
	// Converting amount from string to integer and checking for error
	intval, err := strconv.Atoi(amt)
	if err != nil {
		fmt.Fprintf(w, "%s\n", err)
	}
	// Conditional to first check if both accounts are valid
	if fromEntry, ok := cs.accountMap[from]; ok {
		//Locking to ensure proprietary access to our account object from account object which is sending money
		cs.accountMap[from].Lock()
		if toEntry, ok := cs.accountMap[to]; ok {
			//Locking to ensure proprietary access to our account object from account object which is receiving money
			cs.accountMap[to].Lock()
			// If the account has enough to take out and add to the other account
			if cs.accountMap[from].balance >= intval {
				fromEntry.balance -= intval
				toEntry.balance += intval
				cs.accountMap[from] = fromEntry
				cs.accountMap[to] = toEntry
				fmt.Fprint(w, "ok")
			} else { // The balance in the account is less than the amount to be transferred
				fmt.Fprint(w, from, " only has a balance of $", cs.accountMap[from].balance, ", cannot pay $", amt)
			}
			//Unlocking both once transfer has completed
			cs.accountMap[to].Unlock()
			cs.accountMap[from].Unlock()
		}
	} else {
		fmt.Fprint(w, "error")
	}
}

func main() {
	//Initializing our store variable to contain an account store object which contains starter account objects
	store := AccountStore{accountMap: map[string]*Account{"A": {sync.Mutex{}, "Adu", 0}, "B": {sync.Mutex{}, "Lago", 0}}}
	http.HandleFunc("/get", store.get)
	http.HandleFunc("/set", store.set)
	http.HandleFunc("/inc", store.inc)
	http.HandleFunc("/pay", store.pay)

	portnum := 8080
	if len(os.Args) > 1 {
		portnum, _ = strconv.Atoi(os.Args[1])
	}
	log.Printf("Going to listen on port %d\n", portnum)
	log.Fatal(http.ListenAndServe("localhost:"+strconv.Itoa(portnum), nil))
}
