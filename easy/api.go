package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

func main() {
	FinestFoodHighestRating("Austin", 1000)
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
				votes_ := p_.Data[i].UserRating.Votes
				rating_ := p_.Data[i].UserRating.AverageRating
				name_ := p_.Data[i].Name

				if votes_ >= votes {
					if _, ok := highs[rating_]; ok {
						fmt.Printf("Highest Rated in %v is: %v, with an average rating of: %v\n", city, highestRated, max_)
						return
					}

					if rating_ >= max_ {
						max_ = rating_
						highestRated = p_.Data[i].Name
						highs[rating_] = name_
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
