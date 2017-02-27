# go-thesaurus
Super simple wrapper written in go for Big Huge Thesaurus API

###Basic Usage
```go
thesa := thesaurus.Configure("yourapikey")

//Receives a thesaurus.Response and error message
response, err := thesa.LookUp("word")
if err!=nil{
  panic(err)
}

fmt.Println(response)

```
