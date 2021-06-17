package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sort"
)

var jsonStr = []byte(`
{
    "things": [
        {
            "name": "Alice",
            "age": 37
        },
        {
            "city": "Ipoh",
            "country": "Malaysia"
        },
        {
            "name": "Bob",
            "age": 36
        },
        {
            "city": "Northampton",
            "country": "England"
        },
 		{
            "name": "Albert",
            "age": 3
        },
		{
            "city": "Dnipro",
            "country": "Ukraine"
        },
		{
            "name": "Roman",
            "age": 32
        },
		{
            "city": "New York City",
            "country": "US"
        }
    ]
}`)

type HumanDecoder interface {
	Decode(data []byte) ([]Person, []Place)
	Sort(dataToSort interface{})

	Print(interface{})
	Printlen(persons []Person, places []Place)
}

type Logger interface {
	Println(v ...interface{})
	Fatalf(format string, v ...interface{})
}

type Service struct {
	log Logger
}

func NewService(log Logger) *Service {
	return &Service{log: log}
}

func (s *Service) Decode(data []byte) ([]Person, []Place) {
	return solutionD(data)
}

func (s *Service) Sort(dataToSort interface{}) {
	switch v := dataToSort.(type) {
	case []Person:
		sort.Sort(PersonByLength(v))
	case []Place:
		sort.Sort(PlacesByLength(v))
	default:
		s.log.Fatalf("can't parse type of %v", v)
	}
}
func (s *Service) Print(d interface{}) {
	s.log.Println(d)
}
func (s *Service) Printlen(persons []Person, places []Place) {
	s.log.Println(fmt.Sprintf("%d %d", len(persons), len(places)))
}

func main() {
	logger := log.New(os.Stdout, "INFO: ", 0)
	srv := NewService(logger)

	personsB, placesB := srv.Decode(jsonStr)

	srv.Printlen(personsB, placesB)

	srv.Sort(personsB)
	srv.Sort(placesB)

	srv.Print(personsB)
	srv.Print(placesB)

	personsC, placesC := solutionC(jsonStr)
	fmt.Printf("%d %d\n", len(personsC), len(placesC))

	sort.Sort(PersonByLength(personsC))
	sort.Sort(PlacesByLength(placesC))
	fmt.Println(personsC, placesC)

	persons, places := solutionD(jsonStr)
	fmt.Printf("%d %d\n", len(persons), len(places))

	sort.Sort(PersonByLength(persons))
	sort.Sort(PlacesByLength(places))
	fmt.Println(persons, places)
}

type Person struct {
	Name string
	Age  int
}

type Place struct {
	City    string
	Country string
}

type Mixed struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	City    string `json:"city"`
	Country string `json:"country"`
}

func solutionB(jsonStr []byte) ([]Person, []Place) {
	persons := []Person{}
	places := []Place{}
	var data map[string][]Mixed
	err := json.Unmarshal(jsonStr, &data)
	if err != nil {
		fmt.Println(err)
		return persons, places
	}

	for i := range data["things"] {
		item := data["things"][i]
		if item.Name != "" {
			persons = append(persons, Person{item.Name, item.Age})
		} else {
			places = append(places, Place{item.City, item.Country})
		}

	}
	return persons, places
}

//SolutionC
func solutionC(jsonStr []byte) ([]Person, []Place) {
	people := []Person{}
	places := []Place{}
	var data map[string][]json.RawMessage
	err := json.Unmarshal(jsonStr, &data)
	if err != nil {
		fmt.Println(err)
		return people, places
	}

	for _, thing := range data["things"] {
		people = addPersonC(thing, people)
		places = addPlaceC(thing, places)
	}
	return people, places
}

func addPersonC(thing json.RawMessage, people []Person) []Person {
	person := Person{}
	if err := json.Unmarshal(thing, &person); err != nil {
		fmt.Println(err)
	} else {
		if person != *new(Person) {
			people = append(people, person)
		}
	}

	return people
}

func addPlaceC(thing json.RawMessage, places []Place) []Place {
	place := Place{}
	if err := json.Unmarshal(thing, &place); err != nil {
		fmt.Println(err)
	} else {
		if place != *new(Place) {
			places = append(places, place)
		}
	}

	return places
}

func (p *Person) unmarshal(raw json.RawMessage) bool {
	if err := json.Unmarshal(raw, p); err != nil {
		return false
	}
	return p.Name != "" && p.Age != 0
}

func (p *Place) unmarshal(raw json.RawMessage) bool {
	if err := json.Unmarshal(raw, p); err != nil {
		return false
	}
	return p.City != "" && p.Country != ""
}

func solutionD(data []byte) ([]Person, []Place) {
	var doc struct {
		Things []json.RawMessage `json:"things"`
	}
	if err := json.Unmarshal(data, &doc); err != nil {
		log.Fatal(err)
	}

	people, places := make([]Person, 0, len(doc.Things)), make([]Place, 0, len(doc.Things))
	for _, thing := range doc.Things {
		peep := &Person{}
		if peep.unmarshal(thing) {
			people = append(people, *peep)
			continue
		}

		spot := &Place{}
		if spot.unmarshal(thing) {
			places = append(places, *spot)
			continue
		}

		fmt.Printf("unable to unmarshal: %v\n", thing)
	}

	return people, places
}

type PersonByLength []Person

func (s PersonByLength) Len() int           { return len(s) }
func (s PersonByLength) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s PersonByLength) Less(i, j int) bool { return len(s[i].Name) < len(s[j].Name) }

type PlacesByLength []Place

func (s PlacesByLength) Len() int           { return len(s) }
func (s PlacesByLength) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s PlacesByLength) Less(i, j int) bool { return len(s[i].City) < len(s[j].City) }
