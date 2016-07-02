# go-thesaurus
Simple wrapper written in go for Big Huge Thesaurus API

###Basic Usage
```go
thesa := thesaurus.Configure("yourapikey")

//Receives a thesaurus.Response and error message
response, err := thesa.LookUp("word")

//Receives []byte response in json format, status code and error message
responseJSON, code, err := thesa.LookUpResponse("word", "json")

//Receives []byte response in xml format
responseXML, code, err := thesa.LookUpResponse("word", "xml")

//Receives []byte response in plain text format
responseXML, code, err := thesa.LookUpResponse("word", "")

```
