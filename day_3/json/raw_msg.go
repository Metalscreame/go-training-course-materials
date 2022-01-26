package main

import (
	"encoding/json"
	"fmt"
)

/*
What if I have a large file and I want only a few fields and avoid decoding all of it?
json.RawMessage is a raw encoded JSON value. It can be used to delay JSON decoding.
We can unmarshal the JSON data into a map[string]json.RawMessage and pick out the fields we want.
E.g. In the JSON, below, I want to get the values of only two fields, accountid and paymentid and ignore the rest.
	*/

const data2 = `
{
     "accountid": "162",
     "name": "myclient",
     "paymentid": "A345345",
     "sequence": 3977318850,
     "paymentrow": [{
         "date": "2018 - Jul - 22",
         "amount ": 137.18
     }, {
         "date": "2018-Jul-22",
         "amount": 137.18
     }]
 }`

func main() {
	var objmap map[string]*json.RawMessage

	err := json.Unmarshal([]byte(data2), &objmap)
	if err != nil {
		fmt.Println(err)
	}

	var accountID, paymentID string
	err = json.Unmarshal(*objmap["accountid"], &accountID)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(accountID)
	err = json.Unmarshal(*objmap["paymentid"], &paymentID)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(paymentID)

}