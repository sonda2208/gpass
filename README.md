# googlepasses-go-client
Google Pay API for Passes Client in Go

## Examples

See [examples](https://github.com/sonda2208/googlepasses-go-client/tree/master/example) directory for featured examples.

#### Load JWT configuration from JSON key file

```go
jsonConf, err := ioutil.ReadFile(keyPath)
if err != nil {
    return err
}

jwtConfig, err := google.JWTConfigFromJSON(jsonConf, googlepasses.GooglePayAPIScope)
if err != nil {
    return err
}
```

#### List offer classes

```go
// create offer class client
client := googlepasses.NewOfferClassClient(googlepasses.GooglePayAPIBasePath, jwtConfig.Client(context.TODO()))

// list offer classes
res, err := client.List(conf.IssuerID, 0, "")
if err != nil {
    return
}
```

#### Insert loyalty object

```go
// create loyalty object client
client := googlepasses.NewLoyaltyObjectClient(googlepasses.GooglePayAPIBasePath, jwtConfig.Client(context.TODO()))

// insert new loyalty object
lo := &walletobject.LoyaltyObject{
    ID:      "loyalty_object_id",
    ClassID: "loyalty_class_id",
    State:   "active",
}

nlo, err := client.Insert(lo)
if err != nil {
    return err
}
```