package main

import (
	"encoding/json"
	"fmt"
)
type Bird struct {
	Species string
	Description string
}

func main() {

	birdJson := `{"species":"pigeon", "description":"likes to perch on rocks"}`
	birdJsonArray :=`[
		{"species":"pigeon", "description":"likes to perch on rocks"},
		{"species":"eagle","description":"bird of prey"}
	]`

var bird Bird
var birds []Bird

json.Unmarshal([]byte(birdJson), &bird)
fmt.Printf("Species: %s, Description: %s \n", bird.Species, bird.Description)

//decoding array json

json.Unmarshal([]byte(birdJsonArray), &birds)
fmt.Printf("Birds: %+v \n", birds)
}