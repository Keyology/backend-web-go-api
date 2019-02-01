package main


import (
	"encoding/json"
	"io/ioutil"
	"github.com/labstack/echo"
	"net/http"
	"fmt"
)



//Joke represents a random bad joke sent


type Joke []struct{
	
	ID string `json:"id"`
	Type string `json:"type"`
	Setup string `json:"setup"`
	Punchline string `json:"punchline"`
}



func main(){
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello Keoni!")
	})

	e.GET("/joke", func(c echo.Context) error{

		//create request
		req, err := http.NewRequest("GET", "https://official-joke-api.appspot.com/random_joke", nil)

		// handle error
		if err != nil{
			fmt.Println("Error is req:", err)
		}

		// create a Client
		client := &http.Client{}

		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("error in sending req", err)
		}

		// Response data is available in response.Body; read until error or EOF.
		body, readErr := ioutil.ReadAll(resp.Body)
		if readErr != nil {
			
			fmt.Println(readErr)
	}
		//create a new joke and store data 
		joke := &Joke{} 

		var resObj req

		parseErr := json.Unmarshal(body, &joke)

		if parseErr != nil {
			fmt.Println(parseErr)
			
		}
		fmt.Println(joke)
		return c.JSON(http.StatusOK, joke)
	})

	
	e.Logger.Fatal(e.Start(":1323"))
}