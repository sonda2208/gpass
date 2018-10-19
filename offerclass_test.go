package googlepasses

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

var (
	router *mux.Router
	server *httptest.Server
	client *OfferClassClient
)

const (
	sampleIssuerID       = "1114132711145979111"
	sampleOfferClassID   = "1114132711145979111.TestOfferClass.1"
	sampleOfferClassData = `
	{
		"kind": "walletobjects#offerClass",
		"id": "1114132711145979111.TestOfferClass.1",
		"version": "1",
		"issuerName": "thecoffeeshop",
		"allowMultipleUsersPerObject": true,
		"reviewStatus": "approved",
		"multipleDevicesAndHoldersAllowedStatus": "multipleHolders",
		"title": "20% off",
		"redemptionChannel": "online",
		"provider": "thecoffeeshop"
	}`
	sampleOfferClassesData = `
	{
		"pagination": {
			"kind": "walletobjects#pagination",
			"resultsPerPage": 4
		},
		"resources": [
			{
				"kind": "walletobjects#offerClass",
				"id": "1114132711145979111.OfferClass02",
				"version": "1",
				"issuerName": "Baconrista Coffee",
				"messages": [
					{
						"kind": "walletobjects#walletObjectMessage",
						"header": "Message",
						"body": "This is a message body",
						"id": "1"
					}
				],
				"allowMultipleUsersPerObject": true,
				"locations": [
					{
						"kind": "walletobjects#latLongPoint",
						"latitude": 37.422601,
						"longitude": -122.085286
					},
					{
						"kind": "walletobjects#latLongPoint",
						"latitude": 40.7406578,
						"longitude": -74.00208940000002
					},
					{
						"kind": "walletobjects#latLongPoint",
						"latitude": 37.424354,
						"longitude": -122.09508869999999
					}
				],
				"reviewStatus": "approved",
				"imageModulesData": [
					{
						"mainImage": {
							"kind": "walletobjects#image",
							"sourceUri": {
								"kind": "walletobjects#uri",
								"uri": "http://farm4.staticflickr.com/3738/12440799783_3dc3c20606_b.jpg"
							}
						}
					}
				],
				"textModulesData": [
					{
						"header": "Details",
						"body": "20% off one cup of coffee at all Baconrista Coffee locations.  Only one can be used per visit."
					},
					{
						"header": "About Baconrista",
						"body": "Since 2013, Baconrista Coffee has been committed to making high quality bacon coffee. Visit us in our stores or online at www.baconrista.com"
					}
				],
				"linksModuleData": {
					"uris": [
						{
							"kind": "walletobjects#uri",
							"uri": "http://maps.google.com/maps?q=google",
							"description": "Nearby Locations"
						},
						{
							"kind": "walletobjects#uri",
							"uri": "tel:6505555555",
							"description": "Call Customer Service"
						}
					]
				},
				"multipleDevicesAndHoldersAllowedStatus": "multipleHolders",
				"title": "20% off one bacon fat latte",
				"redemptionChannel": "both",
				"provider": "Baconrista Deals",
				"titleImage": {
					"kind": "walletobjects#image",
					"sourceUri": {
						"kind": "walletobjects#uri",
						"uri": "http://farm4.staticflickr.com/3723/11177041115_6e6a3b6f49_o.jpg"
					}
				}
			},
			{
				"kind": "walletobjects#offerClass",
				"id": "1114132711145979111.OfferClassMonster3",
				"version": "1",
				"issuerName": "Baconrista Coffee",
				"allowMultipleUsersPerObject": true,
				"locations": [
					{
						"kind": "walletobjects#latLongPoint",
						"latitude": 37.424354,
						"longitude": -122.09508869999999
					},
					{
						"kind": "walletobjects#latLongPoint",
						"latitude": 37.422601,
						"longitude": -122.085286
					},
					{
						"kind": "walletobjects#latLongPoint",
						"latitude": 40.7406578,
						"longitude": -74.00208940000002
					}
				],
				"reviewStatus": "approved",
				"imageModulesData": [
					{
						"mainImage": {
							"kind": "walletobjects#image",
							"sourceUri": {
								"kind": "walletobjects#uri",
								"uri": "http://farm4.staticflickr.com/3738/12440799783_3dc3c20606_b.jpg"
							}
						}
					}
				],
				"textModulesData": [
					{
						"header": "Details",
						"body": "20% off one cup of coffee at all Baconrista Coffee locations.  Only one can be used per visit."
					},
					{
						"header": "About Baconrista",
						"body": "Since 2013, Baconrista Coffee has been committed to making high quality bacon coffee. Visit us in our stores or online at www.baconrista.com"
					}
				],
				"linksModuleData": {
					"uris": [
						{
							"kind": "walletobjects#uri",
							"uri": "http://maps.google.com/maps?q=google",
							"description": "Nearby Locations"
						},
						{
							"kind": "walletobjects#uri",
							"uri": "tel:6505555555",
							"description": "Call Customer Service"
						}
					]
				},
				"multipleDevicesAndHoldersAllowedStatus": "multipleHolders",
				"title": "20% off one bacon fat latte",
				"redemptionChannel": "both",
				"provider": "Baconrista Deals",
				"titleImage": {
					"kind": "walletobjects#image",
					"sourceUri": {
						"kind": "walletobjects#uri",
						"uri": "http://farm4.staticflickr.com/3723/11177041115_6e6a3b6f49_o.jpg"
					}
				}
			},
			{
				"kind": "walletobjects#offerClass",
				"id": "1114132711145979111.TestOfferClass.1",
				"version": "1",
				"issuerName": "thecoffeeshop",
				"allowMultipleUsersPerObject": true,
				"reviewStatus": "approved",
				"multipleDevicesAndHoldersAllowedStatus": "multipleHolders",
				"title": "20% off",
				"redemptionChannel": "online",
				"provider": "thecoffeeshop"
			},
			{
				"kind": "walletobjects#offerClass",
				"id": "1114132711145979111.TestOfferClass.2",
				"version": "1",
				"issuerName": "thecoffeeshop",
				"reviewStatus": "approved",
				"title": "20% off",
				"redemptionChannel": "online",
				"provider": "thecoffeeshop"
			}
		]
	}`
)

func setup() func() {
	router = mux.NewRouter()
	server = httptest.NewServer(router)
	client = NewOfferClassClient(server.URL, http.DefaultClient)

	return func() {
		server.Close()
	}
}

func TestGetOfferClass(t *testing.T) {
	teardown := setup()
	defer teardown()

	router.HandleFunc("/offerClass/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if r.Method != "GET" {
			w.WriteHeader(http.StatusNotFound)
		}

		if mux.Vars(r)["id"] != sampleOfferClassID {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(sampleOfferClassData))
	})

	t.Run("Get offer class successfully", func(t *testing.T) {
		res, err := client.Get(sampleOfferClassID)
		assert.NoError(t, err)

		jsonData, err := json.Marshal(res)
		assert.NoError(t, err)

		assert.JSONEq(t, sampleOfferClassData, string(jsonData))
	})

	t.Run("Failed to get offer class", func(t *testing.T) {
		_, err := client.Get("abc")
		assert.Error(t, err)
	})
}

func TestListOfferClasses(t *testing.T) {
	teardown := setup()
	defer teardown()

	router.HandleFunc("/offerClass", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if r.Method != "GET" {
			w.WriteHeader(http.StatusNotFound)
		}

		if r.URL.Query().Get("issuerId") != sampleIssuerID {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(sampleOfferClassesData))
	})

	t.Run("List offer classes successfully", func(t *testing.T) {
		res, err := client.List(sampleIssuerID, 0, "")
		assert.NoError(t, err)

		jsonData, err := json.Marshal(res)
		assert.NoError(t, err)

		assert.JSONEq(t, sampleOfferClassesData, string(jsonData))
	})

	t.Run("Failed to list offer classes", func(t *testing.T) {
		_, err := client.List("", 0, "")
		assert.Error(t, err)
	})
}
