package commands

import (
	"net/http"
	"log"
	"io"
	"fmt"
	"encoding/json"
)

type Config struct {
	Next string
	Previous *string
}

type LocationArea struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func Map(cfg *Config) error {
	if cfg == nil {
		return fmt.Errorf("nil config")
	}

	url := "https://pokeapi.co/api/v2/location-area/"
	if cfg.Next != "" {
    	url = cfg.Next
	}


	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}
	
	var locAreas LocationArea
	if err := json.Unmarshal(body, &locAreas); err != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err)
	}
	
	for _, locArea := range locAreas.Results {
		fmt.Printf("%s", locArea.Name)
		fmt.Println()
	}

	cfg.Next = locAreas.Next
	cfg.Previous = locAreas.Previous
	
	return nil
}

func Mapb(cfg *Config) error {

	if cfg == nil {
		return fmt.Errorf("nil config")
	}
	
	if cfg.Previous == nil {
		fmt.Println("You're on the first page")
		return nil
	}
	url := *cfg.Previous

	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}
	
	var locAreas LocationArea
	if err := json.Unmarshal(body, &locAreas); err != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err)
	}
	
	for _, locArea := range locAreas.Results {
		fmt.Printf("%s", locArea.Name)
		fmt.Println()
	}

	cfg.Next = locAreas.Next
	cfg.Previous = locAreas.Previous
	
	return nil
}
