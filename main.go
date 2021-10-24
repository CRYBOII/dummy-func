package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"path/filepath"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func main() {
	numbPtr := flag.Int("len", 5, "how many functions you want to generate ?")
	flag.Parse()

	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	fucns := "package main"

	for i := 0; i < *numbPtr; i++ {
		fucns += fmt.Sprintf(`
func  %s() {
            
}
            `, RandStringRunes(10))

	}

	fmt.Println(wd)
	err = ioutil.WriteFile(filepath.Join(wd, filepath.Base(fmt.Sprintf("%s.go", RandStringRunes(7)))), []byte(fucns), 0755)
	if err != nil {
		fmt.Printf("Unable to write file: %v", err)
	}

}
