package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"fmt"

	"github.com/labstack/echo"
)

//Joke represents a random bad joke sent
type Joke struct {
	ID        string `json:"id"`
	Type      string `json:"type"`
	Setup     string `json:"setup"`
	Punchline string `json:"punchline"`
}

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
	

	//call the api

	response, err := http.Get("https://official-joke-api.appspot.com/random_joke")

	fmt.Println("this is response", response)
	fmt.Println("this is response err:", err)



	if err != nil{
		fmt.Println("This is a req error", err)

	}

	

	// get the response
	data, rederr := ioutil.ReadAll(response.Body)
	fmt.Println("this is data ", data)
	fmt.Println("this is rederr:", rederr)

	if rederr != nil{
		fmt.Println("This is a read error:", rederr)
	}

	// Instantiate new Joke

	var joke Joke

	fmt.Print("this is joke variable", joke)

	// Unmarshal data

	json.Unmarshal(data, &joke)

	fmt.Println("this is unmarshaling json data",json.Unmarshal(data, &joke))

	

		return c.JSON(http.StatusOK, joke)

	})

	e.Logger.Fatal(e.Start(":4000"))

}
