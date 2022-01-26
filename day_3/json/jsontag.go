package main

import (
	"encoding/json"
	"fmt"
	"time"
)

/*
The "omitempty" option specifies that the field should be omitted from the encoding if the field has an empty value,
defined as false, 0, a nil pointer, a nil interface value, and any empty array, slice, map, or string.

As a special case, if the field tag is "-", the field is always omitted. Note that a field with name "-"
can still be generated using the tag "-,".


The "string" option signals that a field is stored as JSON inside a JSON-encoded string.
It applies only to fields of string, floating point, integer, or boolean types.
This extra level of encoding is sometimes used when communicating with JavaScript programs:

Int64String int64 `json:",string"`

Pointer values encode as the value pointed to. A nil pointer encodes as the null JSON value.
Interface values encode as the value contained in the interface. A nil interface value encodes as the null JSON value.
Channel, complex, and function values cannot be encoded in JSON.
Attempting to encode such a value causes Marshal to return an UnsupportedTypeError.

*/

type User struct {
	Name      string    `json:"name"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	ExpiresAt time.Time `json:"expiresAt,omitempty"` // contains default value from .String() , not nil will be shown
	ExpiresAt2 *time.Time `json:"expiresAt2,omitempty"` // contains default value from .String() , not nil will be shown omited as nil

	Pointer  *string // not omited - null
	EmptyOmitted string `json:"EmptyOmitted,omitempty"` // omited as empty
}

func main() {
	u := User{
		Name:      "123",
		Password:  "123",
		CreatedAt: time.Now(),
		Pointer:   nil,
		EmptyOmitted:  "",
	}
	b, _ := json.Marshal(u)
	fmt.Println(string(b))
}
