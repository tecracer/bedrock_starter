package main

import (
	"bedrock"
	"fmt"
)

func main() {
	// Call Claude via bedrock
	fmt.Printf("Calling Claude...\n")
	response_claude, err := bedrock.CallClaude(bedrock.Client, "Compare JSON to XML. Write the top 3 differences in markdown.")
	if err != nil {
		fmt.Printf("Error calling Claude: %v\n", err)
	} else {
		fmt.Printf("Response:\n%s\n", *response_claude)
	}

}
