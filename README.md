# Multi-Modal Medical Diagnosis System 

This repo serves as the server for the RAG model, it interacts with 
the model through REST API. 

## Instruction 

Add dependencies
```
go mod tidy
``` 
Config ```application-dev.yaml``` under ```common/config/``` to adjust server settings based on need. 

Run the server 
```
go run main.go 
// or 
go build 
go run main 
```

## API 
```go
/rag        // homepage
/rag/query  // send query request to RAG 
/rag/upload // upload and save file 
```
