package main

import (
	"bufio"
	"flag"
	"os"

	hel "github.com/thejini3/go-helper"
	"golang.org/x/crypto/bcrypt"
)

func main() {

	wordlist := flag.String("wordlist", "", "wordlist file path (Required)")
	hash := flag.String("hash", "", "hash string that need to be found (Required)")
	flag.Parse()

	if *wordlist == "" || *hash == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	file, err := os.Open(*wordlist)

	if err != nil {
		hel.Pl("Error opening file", err)
		os.Exit(1)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		if compareHash(*hash, scanner.Text()) {
			hel.Pl("Found pass", scanner.Text())
			os.Exit(1)
		}
	}
}

// HashAndSalt get a hash for given string
func hashAndSalt(pwd string) string {

	// Use GenerateFromPassword to hash & salt pwd.
	// MinCost is just an integer constant provided by the bcrypt
	// package along with DefaultCost & MaxCost.
	// The cost can be any value you want provided it isn't lower
	// than the MinCost (4)
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)
	if err != nil {
		hel.Pl("error hassing", err)
	}
	// GenerateFromPassword returns a byte slice so we need to
	// convert the bytes to a string and return it
	return string(hash)
}

// ComparePasswords compare a hash for given string
func compareHash(hashedPwd string, pwd string) bool {
	// Since we'll be getting the hashed password from the DB it
	// will be a string so we'll need to convert it to a byte slice
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, []byte(pwd))

	if err != nil {
		// hel.Pl("error CompareHashAndPassword", err)
		return false
	}

	return true
}
