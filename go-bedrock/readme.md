

## API Call from cedrock console

```json
{
  "modelId": "anthropic.claude-v1",
  "contentType": "application/json",
  "accept": "*/*",
  "body": "{\"prompt\":\"Human: \\n\\nHuman: \\n\\nCompare JSON to XML. Write the top 3 differences in markdown.\\n\\n\\nAssistant:\",\"max_tokens_to_sample\":1011,\"temperature\":0,\"top_k\":250,\"top_p\":0.999,\"stop_sequences\":[\"\\n\\nHuman:\"],\"anthropic_version\":\"bedrock-2023-05-31\"}"
}
```
The Body

// prompt := &ClaudeV1{
	// 	Prompt:            "Human: \n\nHuman: \n\nCompare JSON to XML. Write the top 3 differences in markdown.\n\n\nAssistant:",
	// 	MaxTokensToSample: 1011,
	// 	Temperature:       0,
	// 	TopK:              250,
	// 	TopP:              0.999,
	// 	StopSequences:     []string{"\n\nHuman:"},
	// 	AnthropicVersion:  "bedrock-2023-05-31",
	// }

take bedrock input json
E.g. claude v1.3

```json
{
  "modelId": "anthropic.claude-v1",
  "contentType": "application/json",
  "accept": "*/*",
  "body": {
    "prompt": "\n\nHuman: Hello world\n\nAssistant:",
    "max_tokens_to_sample": 300,
    "temperature": 0.5,
    "top_k": 250,
    "top_p": 1,
    "stop_sequences": [
      "\\n\\nHuman:"
    ],
    "anthropic_version": "bedrock-2023-05-31"
  }
}
```

Convert to go struct

```go
type ClaudeV1 struct {
	Prompt              string   `json:"prompt"`
	MaxTokensToSample   int      `json:"max_tokens_to_sample"`
	Temperature         float64  `json:"temperature"`
	TopK                int      `json:"top_k"`
	TopP                int      `json:"top_p"`
	StopSequences       []string `json:"stop_sequences"`
	AnthropicVersion    string   `json:"anthropic_version"`
}
```

## First response

```json
response: &{[123 34 99 111 109 112 108 101 116 105 111 110 34 58 34 32 72 101 114 101 32 97 114 101 32 116 104 101 32 116 111 112 32 51 32 100 105 102 102 101 114 101 110 99 101 115 32 98 101 116 119 101 101 110 32 74 83 79 78 32 97 110 100 32 88 77 76 58 92 110 92 110 49 46 32 74 83 79 78 32 105 115 32 108 101 115 115 32 118 101 114 98 111 115 101 32 116 104 97 110 32 88 77 76 46 32 74 83 79 78 32 100 111 101 115 32 110 111 116 32 104 97 118 101 32 101 110 100 32 116 97 103 115 44 32 97 110 100 32 110 97 109 101 115 32 100 111 32 110 111 116 32 110 101 101 100 32 116 111 32 98 101 32 101 110 99 108 111 115 101 100 32 105 110 32 113 117 111 116 101 115 46 32 84 104 105 115 32 109 97 107 101 115 32 74 83 79 78 32 109 111 114 101 32 99 111 109 112 97 99 116 32 97 110 100 32 102 97 115 116 101 114 32 116 111 32 112 97 114 115 101 32 116 104 97 110 32 88 77 76 46 92 110 92 110 50 46 32 74 83 79 78 32 115 117 112 112 111 114 116 115 32 97 114 114 97 121 115 44 32 88 77 76 32 100 111 101 115 32 110 111 116 32 104 97 118 101 32 97 32 110 97 116 105 118 101 32 97 114 114 97 121 32 116 121 112 101 46 32 84 111 32 114 101 112 114 101 115 101 110 116 32 97 114 114 97 121 115 32 105 110 32 88 77 76 44 32 121 111 117 32 110 101 101 100 32 116 111 32 101 110 99 97 112 115 117 108 97 116 101 32 109 117 108 116 105 112 108 101 32 101 108 101 109 101 110 116 115 32 105 110 32 97 32 112 97 114 101 110 116 32 110 111 100 101 46 32 92 110 92 110 51 46 32 74 83 79 78 32 105 115 32 115 105 109 112 108 101 114 32 116 111 32 112 97 114 115 101 32 116 104 97 110 32 88 77 76 46 32 84 104 101 32 74 83 79 78 32 102 111 114 109 97 116 32 105 115 32 115 105 109 112 108 101 114 44 32 119 105 116 104 32 108 101 115 115 32 114 117 108 101 115 32 99 111 109 112 97 114 101 100 32 116 111 32 88 77 76 46 32 84 104 105 115 32 109 97 107 101 115 32 74 83 79 78 32 101 97 115 105 101 114 32 116 111 32 112 97 114 115 101 32 117 115 105 110 103 32 99 111 100 101 44 32 97 110 100 32 74 83 79 78 32 112 97 114 115 101 114 115 32 116 101 110 100 32 116 111 32 98 101 32 102 97 115 116 101 114 32 116 104 97 110 32 88 77 76 32 112 97 114 115 101 114 115 46 34 44 34 115 116 111 112 95 114 101 97 115 111 110 34 58 34 115 116 111 112 95 115 101 113 117 101 110 99 101 34 125] 0x14000020580 {map[{}:-739675000 {}:0x1400005c058 {}:79d7d0be-0238-4b7d-9102-30f11b34ed2f {}:{13925976069079731064 5629651959 0x10278c580} {}:{0 63832950094 <nil>} {}:{[{<nil> false false {map[{}:-739675000 {}:0x1400005c058 {}:79d7d0be-0238-4b7d-9102-30f11b34ed2f {}:{13925976069079731064 5629651959 0x10278c580} {}:{0 63832950094 <nil>}]}}]}]} {}}
``

## Respone 200

```json
{
  "completion": " Hello! My name is Claude.",
  "stop_reason": "stop_sequence",
  "model": "claude-2"
}
```

Now this json as go struct

```go
type Response struct {
	Completion  string `json:"completion"`
	StopReason  string `json:"stop_reason"`
	Model       string `json:"model"`
}
```

## Response error

```json
{
  "error": {
    "type": "invalid_request_error",
    "message": "Invalid request"
  }
}
```

Now this json as go struct

```go
type ResponseError struct {
	Error struct {
		Type    string `json:"type"`
		Message string `json:"message"`
	} `json:"error"`
}
```

## Sources

https://docs.anthropic.com/claude/docs

Prompt Design
https://docs.anthropic.com/claude/docs/introduction-to-prompt-design


## Jurassic

Given json:


```json
 "{\"prompt\":\"Compare JSON to XML. Write the top 3 differences in markdown.\",\"maxTokens\":200,\"temperature\":0.7,\"topP\":1,\"stopSequences\":[],\"countPenalty\":{\"scale\":0},\"presencePenalty\":{\"scale\":0},\"frequencyPenalty\":{\"scale\":0}}"

```

Now this json as go struct

```go
type Jurassic struct {
    Prompt           string   `json:"prompt"`
    MaxTokens        int      `json:"maxTokens"`
    Temperature      float64  `json:"temperature"`
    TopP             float64  `json:"topP"`
    StopSequences    []string `json:"stopSequences"`
    CountPenalty     struct {
        Scale int `json:"scale"`
    } `json:"countPenalty"`
    PresencePenalty struct {
        Scale int `json:"scale"`
    } `json:"presencePenalty"`
    FrequencyPenalty struct {
        Scale int `json:"scale"`
    } `json:"frequencyPenalty"`
}
```

Jurassic response

```json
{
  "id": "7921a78e-d905-c9df-27e3-88e4831e3c3b",
  "prompt": {
    "text": "I will", // Prompt <---
    "tokens": [
      {
        "generatedToken": {
          "token": "▁I▁will",
          "logprob": -9.654844284057617,
          "raw_logprob": -9.654844284057617
        },
        "topTokens": null,
        "textRange": {
          "start": 0,
          "end": 6
        }
      }
    ]
  },
  "completions": [
    {
      "data": {
        "text": " complete this", // Completion <---
        "tokens": [
          {
            "generatedToken": {
              "token": "▁complete",
              "logprob": -6.360593795776367,
              "raw_logprob": -6.491568565368652
            },
            "topTokens": null,
            "textRange": {
              "start": 0,
              "end": 9
            }
          },
          {
            "generatedToken": {
              "token": "▁this",
              "logprob": -2.253995180130005,
              "raw_logprob": -2.886030673980713
            },
            "topTokens": null,
            "textRange": {
              "start": 9,
              "end": 14
            }
          }
        ]
      },
      "finishReason": {
        "reason": "length",
        "length": 2
      }
    }
  ]
}
```

Now this json as go struct

```go
type JurassicResponse struct {
  ID          string `json:"id"`
  Prompt      struct {
    Text   string `json:"text"`
    Tokens []struct {
      GeneratedToken struct {
        Token     string  `json:"token"`
        Logprob   float64 `json:"logprob"`
        RawLogprob float64 `json:"raw_logprob"`
      } `json:"generatedToken"`
      TopTokens  interface{} `json:"topTokens"`
      TextRange  struct {
        Start int `json:"start"`
        End   int `json:"end"`
      } `json:"textRange"`
    } `json:"tokens"`
  } `json:"prompt"`
  Completions []struct {
    Data struct {
      Text   string `json:"text"`
      Tokens []struct {
        GeneratedToken struct {
          Token     string  `json:"token"`
          Logprob   float64 `json:"logprob"`
          RawLogprob float64 `json:"raw_logprob"`
        } `json:"generatedToken"`
        TopTokens interface{} `json:"topTokens"`
        TextRange struct {
          Start int `json:"start"`
          End   int `json:"end"`
        } `json:"textRange"`
      } `json:"tokens"`
    } `json:"data"`
    FinishReason struct {
      Reason string `json:"reason"`
      Length int    `json:"length"`
    } `json:"finishReason"`
  } `json:"completions"`
}
```

With Python SDK:

https://docs.ai21.com/reference/python-sdk
print(response.summaries[0].text)
