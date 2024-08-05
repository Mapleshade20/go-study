package main

import (
	"log"
	"net/http"
	"strconv"

	"example.com/lissajous"
)

func main() {
	http.HandleFunc("/", handler) // each request calls handler
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	params := struct {
		cycles  int
		res     float64
		size    int
		nframes int
		delay   int
	}{
		cycles:  5,     // number of complete x oscillator revolutions
		res:     0.001, // angular resolution
		size:    100,   // image canvas covers [-size..+size]
		nframes: 64,    // number of animation frames
		delay:   8,     // delay between frames (in 10ms units)
	}

	queryParams := map[string]interface{}{
		"cycles":  &params.cycles,
		"res":     &params.res,
		"size":    &params.size,
		"nframes": &params.nframes,
		"delay":   &params.delay,
	}

	for key, ptr := range queryParams {
		if value := r.URL.Query().Get(key); value != "" {
			switch p := ptr.(type) {
			case *int:
				if v, err := strconv.Atoi(value); err == nil && v > 0 {
					*p = v
				} else {
					log.Printf("%s: parsing error", key)
				}
			case *float64:
				if v, err := strconv.ParseFloat(value, 64); err == nil && v > 0.0 {
					*p = v
				} else {
					log.Printf("%s: parsing error", key)
				}
			}
		}
	}

	lissajous.Lissajous(w, params.cycles, params.res, params.size, params.nframes, params.delay)
}
