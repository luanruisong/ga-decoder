## decoder for google auth 


### useage

```shell
go get github.com/luanruisong/ga-decoder
```

```go
func main(){
    url := `otpauth-migration://offline?data={data str}`
    res, err := ga_decoder.DecoderMigrationPayload(url)
    if err != nil {
        panic(err)
    }
    
    fmt.Println(res)
    
    
    for _, v := range res.OtpParameters {
        fmt.Println(ga_decoder.DecodeSecret(v.Secret))
    }
}
```