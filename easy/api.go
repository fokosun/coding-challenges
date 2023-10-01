package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

func main() {
	cities := []string{"Denver", "Seattle", "Austin", "Miami", "Houston", "Omaha", "Dallas", "Portland", "Boston", "Chicago"}

	for i := 0; i < len(cities); i++ {
		FinestFoodHighestRating(cities[i], 500)
	}
}

func FinestFoodHighestRating(city string, votes int) {
	resp, getErr := http.Get("https://jsonmock.hackerrank.com/api/food_outlets?city=" + city)

	if getErr != nil {
		panic(getErr)
	}

	defer resp.Body.Close()
	body, readBodyErr := io.ReadAll(resp.Body)

	if readBodyErr != nil {
		panic(readBodyErr)
	}

	var p Pages

	err := json.Unmarshal(body, &p)

	if err != nil {
		panic(err)
	}

	max_ := 0.00
	highestRated := ""
	highs := make(map[float64]string)

	if p.TotalPages > 0 {
		for i := 1; i <= p.TotalPages; i++ {
			resp, getErr := http.Get("https://jsonmock.hackerrank.com/api/food_outlets?city=" + city + "&page=" + strconv.Itoa(i))

			if getErr != nil {
				panic(getErr)
			}

			defer resp.Body.Close()
			body, readBodyErr := io.ReadAll(resp.Body)

			if readBodyErr != nil {
				panic(readBodyErr)
			}

			var p_ Pages
			err := json.Unmarshal(body, &p_)

			if err != nil {
				fmt.Println("some err?")
				panic(err)
			}

			for i := range p_.Data {
				userVotes := p_.Data[i].UserRating.Votes
				userRatings := p_.Data[i].UserRating.AverageRating
				FoodOutletName := p_.Data[i].Name

				if userVotes >= votes {
					if _, ok := highs[userRatings]; ok {
						fmt.Printf("The highest rated food outlet in %v is: %v, with an average rating of: %v\n", city, highestRated, max_)
						return
					}

					if userRatings >= max_ {
						max_ = userRatings
						highestRated = p_.Data[i].Name
						highs[userRatings] = FoodOutletName
					}
				}
			}
		}
	}
}

type Pages struct {
	Page       int    `json:"page"`
	TotalPages int    `json:"total_pages"`
	PerPage    int    `json:"per_page"`
	Total      int    `json:"total"`
	Data       []Data `json:"data"`
}

type Data struct {
	City          string `json:"city"`
	Name          string `json:"name"`
	EstimatedCost int    `json:"estimated_cost"`
	UserRating    Rating `json:"user_rating"`
}

type Rating struct {
	AverageRating float64 `json:"average_rating"`
	Votes         int     `json:"votes"`
}
