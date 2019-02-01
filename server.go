package main

import (
	"net/http"
	"github.com/labstack/echo"
	"fmt"
)
type Joke []struct{
	// Joke represents a random bad joke sent
	
	ID	string `json:"id"`
	Type string `json:"type"`
	Setup string `json:"setup"`
	Punchline string `json:"punchline"`

}


func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, world!")

	})

	e.GET("/joke", func(c echo.Context) error{
		// Build the request 
		req, err := http.NewRequest("GET","https://official-joke-api.appspot.com/random_joke",nil)
			if err != nil{
				fmt.Println("error is req: ", err)
			}
	})
	e.Logger.Fatal(e.Start(":1323"))
}

