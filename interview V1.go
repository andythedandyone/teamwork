// package customerimporter reads from the given customers.csv file and returns a
// sorted (data structure of your choice) of email domains along with the number
// of customers with e-mail addresses for each domain.  Any errors should be
// logged (or handled). Performance matters (this is only ~3k lines, but *could*
// be 1m lines or run on a small machine).
package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strings"
)

// Client : used to store final data
type Client struct {
	domain string
	amount int
}

func main() {

	csFile, error := os.Open("customers.csv")
	if error != nil {
		fmt.Println("there was an error", error)
		os.Exit(1)
	}
	defer csFile.Close()

	csvFile := csv.NewReader(csFile)

	var clientList []Client
	var clientSingle Client

	for {
		data, err := csvFile.Read()
		if data == nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}

		splitEmail := strings.Split(data[2], "@")
		domain := strings.Join(splitEmail[1:], "")

		if domain == "" {
			continue
		} else {
			if len(clientList) == 0 {
				clientSingle.domain = domain
				clientSingle.amount++
				clientList = append(clientList, clientSingle)
			} else {
				for i, thisClient := range clientList {
					if thisClient.domain != domain && len(clientList)-1 == i {
						var x Client
						x.domain = domain
						x.amount++
						clientList = append(clientList, x)
					} else if thisClient.domain == domain {
						clientList[i].amount++
						break
					}
				}
			}
		}
	}

	sort.Slice(clientList, func(i, j int) bool {
		return clientList[i].domain < clientList[j].domain
	})

	// Un-comment the loop below to see the output on a sorted way.
	// Without printing to screen you see how fast the function loops through

	// for i, data := range clientList {
	// 	fmt.Println(" ---> ", i, data)
	// }
}
