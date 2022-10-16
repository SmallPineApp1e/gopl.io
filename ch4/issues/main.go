// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 112.
//!+

// Issues prints a table of GitHub issues matching the search terms.
package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"gopl.io/ch4/github"
)

const (
	IN_ONE_MONTH  = "in_one_month"
	IN_ONE_YEAR   = "in_one_year"
	OVER_ONE_YEAR = "over_one_year"
	OVER_TWO_YEAR = "over_two_year"
)

const (
	MONTH    = 24 * 60 * 60 * 30
	YEAR     = MONTH * 12
	TWO_YEAR = YEAR * 2
)

// !+
func main() {
	practice410()
}

func example() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}
}

func practice410() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	now := time.Now().UTC().UnixMilli()
	classification := make(map[string][]*github.Issue)
	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		create_time := item.CreatedAt.UTC().UnixMilli()
		t := IN_ONE_MONTH
		duration := (now - create_time) / 1000
		if duration > MONTH && duration < YEAR {
			t = IN_ONE_YEAR
		} else if duration > YEAR && duration < TWO_YEAR {
			t = OVER_ONE_YEAR
		} else if duration > TWO_YEAR {
			t = OVER_TWO_YEAR
		}
		vals, ok := classification[t]
		if !ok {
			classification[t] = []*github.Issue{}
			vals = classification[t]
		}
		vals = append(vals, item)
		classification[t] = vals
	}
	for k, vals := range classification {
		fmt.Printf("type: %-5v\n", k)
		for _, item := range vals {
			fmt.Printf("#%-5d %9.9s %.55s %+6s\n",
				item.Number, item.User.Login, item.Title, item.CreatedAt.UTC())
		}
	}
}

//!-

/*
//!+textoutput
$ go build gopl.io/ch4/issues
$ ./issues repo:golang/go is:open json decoder
13 issues:
#5680    eaigner encoding/json: set key converter on en/decoder
#6050  gopherbot encoding/json: provide tokenizer
#8658  gopherbot encoding/json: use bufio
#8462  kortschak encoding/json: UnmarshalText confuses json.Unmarshal
#5901        rsc encoding/json: allow override type marshaling
#9812  klauspost encoding/json: string tag not symmetric
#7872  extempora encoding/json: Encoder internally buffers full output
#9650    cespare encoding/json: Decoding gives errPhase when unmarshalin
#6716  gopherbot encoding/json: include field name in unmarshal error me
#6901  lukescott encoding/json, encoding/xml: option to treat unknown fi
#6384    joeshaw encoding/json: encode precise floating point integers u
#6647    btracey x/tools/cmd/godoc: display type kind of each named type
#4237  gjemiller encoding/base64: URLEncoding padding is optional
//!-textoutput
*/
