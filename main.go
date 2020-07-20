package main

import (
	"bufio"
	"flag"
	"os"
	"runtime"
	"sync"
	"time"

	hel "github.com/thejini3/go-helper"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	started := time.Now()
	wordlist := flag.String("wordlist", "", "wordlist file path (Required)")
	hash := flag.String("hash", "", "hash string that need to be found (Required)")
	core := flag.Int("core", -1, "number of cpu core, Default -1 (all core) (Optional)")

	flag.Parse()

	if *wordlist == "" || *hash == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	if *core == 0 {
		*core = 1
	} else if *core > runtime.NumCPU() || *core == -1 {
		*core = runtime.NumCPU()
	}

	hel.Pl("Using cpu core(s):", runtime.GOMAXPROCS(*core))

	file, err := os.Open(*wordlist)

	if err != nil {
		hel.Pl("Error opening file", err)
		os.Exit(1)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	hashByte := []byte(*hash)

	var wg sync.WaitGroup

	for scanner.Scan() {
		wg.Add(1)
		go func(hashByte []byte, pass []byte) {
			if bcrypt.CompareHashAndPassword(hashByte, pass) == nil {
				hel.Pl("Found pass `" + string(pass) + "`")
				os.Exit(1)
			}
			wg.Done()
		}(hashByte, scanner.Bytes())
	}
	wg.Wait()
	hel.Pl("Done in:", time.Since(started))
}

// // HashAndSalt get a hash for given string
// func hashAndSalt(pwd string) string {

// 	// Use GenerateFromPassword to hash & salt pwd.
// 	// MinCost is just an integer constant provided by the bcrypt
// 	// package along with DefaultCost & MaxCost.
// 	// The cost can be any value you want provided it isn't lower
// 	// than the MinCost (4)
// 	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)
// 	if err != nil {
// 		hel.Pl("error hassing", err)
// 	}
// 	// GenerateFromPassword returns a byte slice so we need to
// 	// convert the bytes to a string and return it
// 	return string(hash)
// }

// // ComparePasswords compare a hash for given string
// func compareHash(hashedPwd string, pwd string) bool {
// 	// Since we'll be getting the hashed password from the DB it
// 	// will be a string so we'll need to convert it to a byte slice
// 	return bcrypt.CompareHashAndPassword([]byte(hashedPwd), []byte(pwd)) == nil
// }
