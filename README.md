## A simple conversational Go Gemini-ai chat 

It keeps converastion history for a session

### Install required packages, run:

```bash
go mod tidy
```

To get started set environment variable, for example:

```bash
export API_KEY=your_api_key
```

or pass API_KEY as flag when running go:
```bash
go run main.go -api-key=your_api_key
```

get api key here : [https://aistudio.google.com/app/apikey](https://aistudio.google.com/app/apikey)


googles examples to get started with the Gemini API in Go applications [https://ai.google.dev/tutorials/go_quickstart](https://ai.google.dev/tutorials/go_quickstart)



### Run chat:

```bash
go run main.go
```
