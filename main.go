package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"
	"sync"
	"time"

	hel "github.com/thejini3/go-helper"
	"golang.org/x/crypto/bcrypt"
)

var started time.Time

// 12345 = $2y$12$APew2qEmu/1YDnHmdPUT5.idVsU3lN2gE17srB3lC7Jiqsdf2qg9m

func main() {
	started = time.Now()
	wordlist := flag.String("w", "", "(Required) wordlist file path")
	hash := flag.String("h", "", "(Required) hash string that need to be found")
	core := flag.Int("c", -1, "(Optional) number of cpu core,  -1 = all core")
	thread := flag.Int("t", 50, "(Optional) number of concurrent thread")

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

	passes := hel.StrToArr(hel.FileStrMust(*wordlist), "\n")

	hashByte := []byte(*hash)

	var wg sync.WaitGroup
	var c = make(chan int, *thread)
	var checked = 0

	hel.Pl("Starting, total passwords to check", len(passes))

	for i, p := range passes {

		wg.Add(1)

		go func(hashByte []byte, pass []byte, i int) {

			c <- i

			if bcrypt.CompareHashAndPassword(hashByte, pass) == nil {
				fmt.Printf("\n\nFound pass `" + string(pass) + "`\n\n")
				done()
			}

			checked++

			fmt.Printf("\rChecked - %d ", checked)

			<-c
			wg.Done()

		}(hashByte, []byte(p), i)
	}
	wg.Wait()
	close(c)
	done()
}
func done() {
	hel.Pl("Done in:", time.Since(started))
	os.Exit(0)
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
