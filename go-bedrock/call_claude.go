package bedrock

//import amazon bedrock runtime go v2
import (
	claude "bedrock/claude"
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"golang.org/x/exp/slog"

	"github.com/aws/aws-sdk-go-v2/service/bedrockruntime"
)

func CallClaude(client *bedrockruntime.Client, prompt string) (*string, error) {

	completePrompt := fmt.Sprintf("Human: \n\n%s\n\n\nAssistant:", prompt)

	body := &claude.ClaudeV1{
		Prompt:            completePrompt,
		MaxTokensToSample: 1024,
		Temperature:       0,
		TopK:              250,
		TopP:              0.999,
		StopSequences:     []string{"\n\nHuman:"},
		AnthropicVersion:  "bedrock-2023-05-31",
	}

	byteSlice, err := json.Marshal(body)
	if err != nil {
		slog.Error("Error unmarshalling: ", "Error", err)
		return nil, err
	}

	// call the bedrock runtime
	response, err := client.InvokeModel(context.TODO(), &bedrockruntime.InvokeModelInput{
		ModelId:     aws.String(string(claude.ModelIdClaudeInstantV1)),
		ContentType: aws.String("application/json"),
		Accept:      aws.String("*/*"),
		Body:        byteSlice,
	})

	if err != nil {
		slog.Error("configuration error, " + err.Error())
		var responseError claude.ResponseError
		err = json.Unmarshal(response.Body, &responseError)
		if err != nil {
			slog.Error("Error unmarshalling: ", "Error", err)
		}
		fmt.Printf("responseError Message: %v\n", responseError.Error.Message)
		fmt.Printf("responseError Type: %v\n", responseError.Error.Type)
		return nil, err

	}
	// https://pkg.go.dev/github.com/aws/aws-sdk-go-v2/service/bedrockruntime@v1.1.2#InvokeModelOutput
	// fmt.Printf("response: %v\n", string(response.Body))

	// response: {"completion":" Here are the top 3 differences between JSON and XML:\n\n1. JSON is less verbose than XML. JSON does not have end tags, and names do not need to be enclosed in quotes. This makes JSON more compact and faster to parse than XML.\n\n2. JSON supports arrays, XML does not have a native array type. To represent arrays in XML, you need to encapsulate multiple elements in a parent node. \n\n3. JSON is simpler to parse than XML. The JSON format is simpler, with less rules compared to XML. This makes JSON easier to parse using code, and JSON parsers tend to be faster than XML parsers.","stop_reason":"stop_sequence"}
	// https://docs.anthropic.com/claude/reference/complete_post
	var responseStruct claude.Response
	err = json.Unmarshal(response.Body, &responseStruct)
	if err != nil {
		slog.Error("Error unmarshalling: ", "Error", err)
		return nil, err
	}

	return &responseStruct.Completion, nil
}
