package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// Movie simple movie info
type Movie struct {
	Title string
	// Tags for json
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"`
	Actors []string
}

func main() {
	var movies = []Movie{
		{
			Title:  "The Shawshank Redemption",
			Year:   1994,
			Color:  true,
			Actors: []string{"Tim Robbins", "Morgan Freeman", "Bob Gunton", "William Sadler"},
		},
		{
			Title:  "Casablanca",
			Year:   1942,
			Actors: []string{"Humphrey Bogart", "Ingrid Nergman"},
		},
	}
	fmt.Println(movies)
	data, err := json.Marshal(movies)
	if err != nil {
		log.Fatalf("Json Marshaling failed:%s ", err)
	}
	fmt.Printf("%s \n", data)
	data, err = json.MarshalIndent(movies, "a", "bb")
	if err != nil {
		log.Fatalf("Json Marshaling failed:%s ", err)
	}
	fmt.Printf("%s \n", data)

}
