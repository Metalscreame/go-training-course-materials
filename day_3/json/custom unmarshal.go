package main

import (
	"encoding/json"
	"fmt"
)

// Json to Golang Struct https://mholt.github.io/json-to-go/

const data = `
{ 
	"sequence": 3977318850, 
	"bids": [ 
          [ "4625.78", "0.80766325", 3 ] 
        ], 
	"asks": [ 
          [ "4625.79", "3.0154341", 3 ] 
        ]
}`

// we can't work with such struct. how can we marshall this?
type Struct struct {
	Sequence int64           `json:"sequence"`
	Bids     [][]interface{} `json:"bids"`
	Asks     [][]interface{} `json:"asks"`
}

type Bid struct {
	Price     string
	Size      string
	NumOrders int
}

type OrderBook struct {
	Sequence int64 `json:"sequence"`
	Bids     []Bid `json:"bids"`
	Asks     []Bid `json:"asks"`
}

func (b *Bid) UnmarshalJSON(data []byte) error {
	var v []interface{}
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	var ok bool
	b.Price, ok  = v[0].(string)
	if !ok {
		return fmt.Errorf("order was not expected")
	}

	b.Size, ok = v[1].(string)
	b.NumOrders = int(v[2].(float64))

	return nil
}

func main() {

	var orders OrderBook
	err := json.Unmarshal([]byte(data), &orders)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(orders)
}
