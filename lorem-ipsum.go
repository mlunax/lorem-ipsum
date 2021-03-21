package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/urfave/cli/v2"
)

var words = [...]string{"Lorem", "ipsum", "dolor", "sit", "amet", "consectetur", "adipiscing", "elit", "sed", "do", "eiusmod", "tempor", "incididunt",
	"ut", "labore", "et", "dolore", "magna", "aliqua", "Ut", "enim", "ad", "minim", "veniam", "quis", "nostrud", "exercitation", "ullamco", "laboris", "nisi",
	"ut", "aliquip", "ex", "ea", "commodo", "consequat", "Duis", "aute", "irure", "dolor", "in", "reprehenderit", "in", "voluptate", "velit", "esse", "cillum",
	"dolore", "eu", "fugiat", "nulla", "pariatur", "Excepteur", "sint", "occaecat", "cupidatat", "non", "proident", "sunt", "in", "culpa", "qui", "officia", "deserunt",
	"mollit", "anim", "id", "est", "laborum"}

func main() {
	app := &cli.App{
		Name:      "lorem-ipsum",
		Usage:     "lorem-ipsum is the cli tool to generate popular unit test sentences",
		UsageText: "lorem-ipsum [number of words]",
		Action: func(c *cli.Context) error {
			numWords := uint64(20)
			if c.NArg() >= 1 {
				argument := c.Args().Get(0)
				result, err := strconv.ParseUint(argument, 10, 64)
				if err != nil {
					log.Fatal(err)
					return nil
				}
				numWords = result
			}
			var prefix []string
			var text []string
			if numWords <= 2 {
				prefix = words[:numWords]
				text = append(prefix)
			} else {
				prefix = words[:2]
				text = append(prefix)
				numWords -= 2
				rg := rand.New(rand.NewSource(time.Now().UnixNano()))
				for i := uint64(0); i < numWords; i++ {
					r := rg.Intn(len(words))
					word := words[r]
					word = strings.ToLower(word)
					text = append(text, word)
				}
			}
			fmt.Println(strings.Join(text, " "))
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
