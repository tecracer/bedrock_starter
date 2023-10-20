async function rockbed() {
  const { BedrockRuntimeClient, InvokeModelCommand } = require("@aws-sdk/client-bedrock-runtime");

  const config = { region: "eu-central-1" };
  const client = new BedrockRuntimeClient(config);

  let instruction = "Compare JSON to XML. Write the top 3 differences in markdown."
  let complete_prompt = `Human: \n\n${instruction}\n\n\nAssistant:`;

  const body = {
    "prompt": complete_prompt,
    "max_tokens_to_sample": 1024,
    "temperature": 0,
    "top_k": 250,
    "top_p": 0.999,
    "stop_sequences": ["\n\nHuman:"],
    "anthropic_version": "bedrock-2023-05-31",
  }
  const buffer = Buffer.from(JSON.stringify(body));

  modelId = 'anthropic.claude-instant-v1'
  accept = '*/*'
  contentType = 'application/json'
  instruction = "Compare JSON to XML. Write the top 3 differences in markdown."

  const input = {
    body: buffer,
    contentType,
    accept,
    modelId
  };

  const command = new InvokeModelCommand(input);
  try {
    const response = await client.send(command);
    console.log("Response:\n");
    let decoder = new TextDecoder("utf-8");
    let content = JSON.parse(decoder.decode(response.body));
    console.log(content.completion);
  } catch (error) {
    console.error(error);
  }
}

rockbed()
  .then(() => console.log('Done'))
  .catch(err => console.error(err));
