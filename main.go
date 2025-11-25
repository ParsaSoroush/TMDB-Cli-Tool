package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
)

const apiBaseURL = "https://api.themoviedb.org/3/movie"

// Structs to parse TMDB JSON response
type MovieResults struct {
	Results []Movie `json:"results"`
}

type Movie struct {
	Title       string  `json:"title"`
	ReleaseDate string  `json:"release_date"`
	VoteAverage float64 `json:"vote_average"`
}

func main() {
	// CLI flag
	movieType := flag.String("type", "popular", "Type of movie: playing | popular | top | upcoming")
	flag.Parse()

	apiKey := os.Getenv("TMDB_API_KEY")
	if apiKey == "" {
		fmt.Println("❌ ERROR: Missing API key. Set environment variable TMDB_API_KEY")
		return
	}

	endpoint := getEndpoint(*movieType)
	if endpoint == "" {
		fmt.Println("❌ ERROR: Invalid movie type. Supported: playing, popular, top, upcoming")
		return
	}

	fullURL := fmt.Sprintf("%s?api_key=%s", endpoint, apiKey)

	resp, err := http.Get(fullURL)
	if err != nil {
		fmt.Println("❌ Network error:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		fmt.Printf("❌ API error: Status %d\n", resp.StatusCode)
		return
	}

	var movieData MovieResults
	if err := json.NewDecoder(resp.Body).Decode(&movieData); err != nil {
		fmt.Println("❌ JSON parsing error:", err)
		return
	}

	displayMovies(*movieType, movieData.Results)
}

func getEndpoint(movieType string) string {
	switch movieType {
	case "playing":
		return apiBaseURL + "/now_playing"
	case "popular":
		return apiBaseURL + "/popular"
	case "top":
		return apiBaseURL + "/top_rated"
	case "upcoming":
		return apiBaseURL + "/upcoming"
	default:
		return ""
	}
}

func displayMovies(category string, movies []Movie) {
	fmt.Printf("\n🎬 Showing results for: %s\n", category)
	fmt.Println("----------------------------------------------------")

	for _, m := range movies {
		fmt.Printf("📌  %s\n", m.Title)
		fmt.Printf("   📅 Release Date: %s\n", m.ReleaseDate)
		fmt.Printf("   ⭐ Rating: %.1f\n", m.VoteAverage)
		fmt.Println("----------------------------------------------------")
	}
}
