package main

import (
	"flag"
	"fmt"
	"log"
)

func seedAccount(store Storage, fname, lname, pw string) *Account {
	acc, err := NewAccount(fname, lname, pw)
	if err != nil {
		log.Fatal(err)
	}
	if err := store.CreateAccount(acc); err != nil {
		log.Fatal(err)
	}
	fmt.Println("account number => ", acc.Number)
	return acc
}

func seedAccounts(store Storage) {
	seedAccount(store, "Anonymous", "Byte", "Tfinity")
}

func main() {

	seed := flag.Bool("seed", false, "seed the db")
	flag.Parse()

	store, err := NewPostgressStore()
	if err != nil {
		log.Fatal(err)
	}

	if err := store.Init(); err != nil {
		log.Fatal(err)
	}

	if *seed {
		fmt.Println("Seeding the db")
		seedAccounts(store)
	}

	server := NewApiServer(":3000", store)
	server.Run()

}
