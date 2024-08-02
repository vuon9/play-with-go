# json2json
[WIP] to query json and append to new json, not implemented as enough to use, just demostrate.


Example to see how it works:
```go
// main.go

func main() {
    srcContent := `{"person":{"first":"John","last":"Smith","books":["Go in Action","Go in Action 2"],"school":{"address":"First Street","city":"New York"}}}`
    mapKeys := []MapKey{
        {"person.first", "first"},
        {"person.last", "last"},
        {"person.books.[0]", "nana.books1"},
    }

    mOut, err := jmapper.ConvertToDestKeymap(srcContent, mapKeys)
    if err != nil {
        panic(err)
    }
    outContent, _ := json.Marshal(mOut)
}
```

Result:
```go
{"first":"John","last":"Smith","nana":{"books1":"Go in Action"}}
```

## Thanks to
* github.com/thedevsaddam/gojsonq/v2