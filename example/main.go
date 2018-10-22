package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	googlepasses "github.com/sonda2208/googlepasses-go-client"
	"github.com/sonda2208/googlepasses-go-client/walletobject"
	"golang.org/x/oauth2/google"
	"golang.org/x/oauth2/jwt"
)

type AppConfig struct {
	ServiceAccountEmail      string `envconfig:"SERVICE_ACCOUNT_EMAIL" required:"true"`
	ServiceAccountPrivateKey string `envconfig:"SERVICE_ACCOUNT_PRIVATE_KEY" required:"true"`
	ApplicationName          string `envconfig:"APP_NAME" required:"true"`
	IssuerID                 string `envconfig:"ISSUER_ID" required:"true"`
	GiftCardClassPrefix      string `envconfig:"GIFTCARD_CLASS_PREFIX" required:"true"`
	GiftCardObjectPrefix     string `envconfig:"GIFTCARD_OBJECT_PREFIX" required:"true"`
	LoyaltyClassPrefix       string `envconfig:"LOYALTY_CLASS_PREFIX" required:"true"`
	LoyaltyObjectPrefix      string `envconfig:"LOYALTY_OBJECT_PREFIX" required:"true"`
	OfferClassPrefix         string `envconfig:"OFFER_CLASS_PREFIX" required:"true"`
	OfferObjectPrefix        string `envconfig:"OFFER_OBJECT_PREFIX" required:"true"`
}

func loadConfig(prefixEnv string) *AppConfig {
	conf := &AppConfig{}

	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	err = envconfig.Process(prefixEnv, conf)
	if err != nil {
		log.Fatal(err)
	}

	return conf
}

func loadJWTConfig(keyPath string) *jwt.Config {
	jsonConf, err := ioutil.ReadFile(keyPath)
	if err != nil {
		log.Fatal(err)
	}

	jwtConfig, err := google.JWTConfigFromJSON(jsonConf, googlepasses.GooglePayAPIScope)
	if err != nil {
		log.Fatal(err)
	}

	return jwtConfig
}

func testOfferClass(conf *AppConfig, client *http.Client) {
	// create offer class service
	ocClient := googlepasses.NewOfferClassClient("", client)
	_ = ocClient

	// list offer classes
	res, err := ocClient.List(conf.IssuerID, 2, "")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(res.Pagination)

	// get offer class
	oc, err := ocClient.Get(conf.IssuerID + ".OfferClass02")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(oc)

	// insert new offer class
	noc := &walletobject.OfferClass{
		ID:                fmt.Sprintf("%s.%s.2", conf.IssuerID, conf.OfferClassPrefix),
		IssuerName:        "thecoffeeshop",
		ReviewStatus:      "underReview",
		Provider:          "thecoffeeshop",
		RedemptionChannel: "online",
		Title:             "20% off",
	}
	noc, err = ocClient.Insert(noc)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(noc)
}

func testLoyaltyClass(conf *AppConfig, client *http.Client) {
	lcClient := googlepasses.NewLoyaltyClassClient(googlepasses.GooglePayAPIBasePath, client)

	// insert new loyalty class
	lc := &walletobject.LoyaltyClass{
		ID:           fmt.Sprintf("%s.%s.1", conf.IssuerID, conf.LoyaltyClassPrefix),
		IssuerName:   "thecoffeeshop",
		ReviewStatus: "underReview",
		ProgramName:  "Loyalty Card",
		ProgramLogo: &walletobject.Image{
			Kind: "walletobjects#image",
			SourceURI: &walletobject.URI{
				Kind: "walletobjects#uri",
				URI:  "http://farm8.staticflickr.com/7340/11177041185_a61a7f2139_o.jpg",
			},
		},
	}
	nlc, err := lcClient.Insert(lc)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(nlc)
}

func testLoyaltyObject(conf *AppConfig, client *http.Client) {
	loClient := googlepasses.NewLoyaltyObjectClient(googlepasses.GooglePayAPIBasePath, client)

	// insert new loyalty object
	lo := &walletobject.LoyaltyObject{
		ID:      fmt.Sprintf("%s.%s.1", conf.IssuerID, conf.LoyaltyObjectPrefix),
		ClassID: fmt.Sprintf("%s.%s.1", conf.IssuerID, conf.LoyaltyClassPrefix),
		State:   "active",
	}
	nlo, err := loClient.Insert(lo)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(nlo)
}

func testGiftcardClass(conf *AppConfig, client *http.Client) {
	gcClient := googlepasses.NewGiftcardClassClient(googlepasses.GooglePayAPIBasePath, client)

	res, err := gcClient.List(conf.IssuerID, 0, "")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(res.Resources[0])
}

func testGiftcardObject(conf *AppConfig, client *http.Client) {
	gcClient := googlepasses.NewGiftcardObjectClient(googlepasses.GooglePayAPIBasePath, client)

	res, err := gcClient.List(fmt.Sprintf("%s.%s.1", conf.IssuerID, conf.GiftCardClassPrefix), 0, "")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(res.Resources[0])
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	conf := loadConfig("S2AP")
	jwtConfig := loadJWTConfig(conf.ServiceAccountPrivateKey)

	tk, _ := jwtConfig.TokenSource(context.TODO()).Token()
	log.Println(tk.AccessToken)

	// testOfferClass(conf, jwtConfig.Client(context.TODO()))
	// testLoyaltyClass(conf, jwtConfig.Client(context.TODO()))
	// testLoyaltyObject(conf, jwtConfig.Client(context.TODO()))
	testGiftcardClass(conf, jwtConfig.Client(context.TODO()))
	// testGiftcardObject(conf, jwtConfig.Client(context.TODO()))
}
