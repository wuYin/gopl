package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

const (
	API_URL = "http://www.omdbapi.com?apikey="
	API_KEY = "3e356e0a"
)

type Movie struct {
	Title  string
	Year   string
	Poster string
}

//  go run 4.13.go r
func main() {
	flag.Parse()
	if flag.NArg() == 0 {
		fmt.Println("input movie title")
		return
	}
	movie, err := search(os.Args[1])
	if err != nil {
		log.Fatalf("search error: %v", err)
	}
	emptyM := Movie{}
	if movie == emptyM {
		log.Fatalf("search no result")
	}

	err = downloadPoster(movie)
	if err != nil {
		log.Fatalf("download error: %v", err)
	}
	fmt.Printf("%+v\n", movie)
}

func search(name string) (m Movie, err error) {
	url := fmt.Sprintf("%s%s&t=%s", API_URL, API_KEY, name)
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("resp status code: %v", resp.Status)
		return
	}
	err = json.NewDecoder(resp.Body).Decode(&m)
	if err != nil {
		return
	}
	return
}

func downloadPoster(m Movie) error {
	resp, err := http.Get(m.Poster)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("resp status code: %v", resp.Status)
	}
	poster := m.Title + filepath.Ext(m.Poster)
	f, err := os.Create(poster)
	if err != nil {
		return err
	}
	defer f.Close()
	w := bufio.NewWriter(f)
	_, err = w.ReadFrom(resp.Body)
	if err != nil {
		return err
	}
	err = w.Flush()
	if err != nil {
		return err
	}
	return nil
}
