import boto3
import json

def main():
  """
  Example how ti use bedrock runtime
  """
  # See https://boto3.amazonaws.com/v1/documentation/api/latest/reference/services/bedrock-runtime.html
  # create client with region
  # Check model available in region
  client = boto3.client('bedrock-runtime', region_name='eu-central-1')
  modelId = 'anthropic.claude-instant-v1'
  accept = 'application/json'
  contentType = 'application/json'
  instruction = "Compare JSON to XML. Write the top 3 differences in markdown."

  prompt = f"Human: \n\n{instruction}\n\n\nAssistant:"

  body = {
  	"prompt": prompt,
  	"max_tokens_to_sample": 1024,
  	"temperature": 0,
  	"top_p": 250,
  	"top_k": 0.999,
  	"stop_sequences": ["\n\nHuman:"],
  	"anthropic_version": "bedrock-2023-05-31",
  }

  json_dump = json.dumps(body)
  print("JSON: ", json_dump)
  body = json.dumps({"prompt": prompt,
                 "max_tokens_to_sample":4096,
                 "temperature":0.5,
                 "top_k":250,
                 "top_p":0.5,
                 "stop_sequences":[]
                  })
  try:
      response = client.invoke_model(body=body, modelId=modelId, accept=accept, contentType=contentType)

  except Exception as e:
      print("Configuration error: ", e)

      try:
          response_error = json.loads(response.body)
      except:
          print("Error unmarshalling: ", e)

      print("ResponseError Message: ", response_error.get("Message"))
      print("ResponseError Type: ", response_error.get("Type"))

      raise e

  try:
      content = response['body'].read().decode('utf-8')
  except Exception as e:
      print("Error unmarshalling: ", e)

      raise e
  finally:
      print("Response: ", content)

if __name__ == '__main__':
    # execute only if run as the entry point into the program
    main()
