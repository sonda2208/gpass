package googlepasses

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/sonda2208/googlepasses-go-client/walletobject"
	"github.com/stretchr/testify/assert"
)

const (
	sampleLoyaltyClassID   = "1114132711145979111.TestLoyaltyClass.1"
	sampleLoyaltyClassData = `
	{
		"kind": "walletobjects#loyaltyClass",
		"id": "1114132711145979111.TestLoyaltyClass.1",
		"version": "1",
		"issuerName": "thecoffeeshop",
		"reviewStatus": "approved",
		"programName": "Loyalty Card",
		"programLogo": {
			"kind": "walletobjects#image",
			"sourceUri": {
				"kind": "walletobjects#uri",
				"uri": "http://farm8.staticflickr.com/7340/11177041185_a61a7f2139_o.jpg"
			}
		}
	}`
	sampleLoyaltyClassesData = `
	{
		"pagination": {
			"kind": "walletobjects#pagination",
			"resultsPerPage": 5
		},
		"resources": [
			{
				"kind": "walletobjects#loyaltyClass",
				"id": "1114132711145979111.LoyaltyClass",
				"version": "1",
				"issuerName": "Baconrista",
				"allowMultipleUsersPerObject": true,
				"reviewStatus": "approved",
				"infoModuleData": {
					"showLastUpdateTime": true
				},
				"multipleDevicesAndHoldersAllowedStatus": "multipleHolders",
				"programName": "Baconrista Rewards",
				"programLogo": {
					"kind": "walletobjects#image",
					"sourceUri": {
						"kind": "walletobjects#uri",
						"uri": "http://farm8.staticflickr.com/7340/11177041185_a61a7f2139_o.jpg"
					}
				}
			},
			{
				"kind": "walletobjects#loyaltyClass",
				"id": "1114132711145979111.LoyaltyClass01",
				"version": "1",
				"issuerName": "Baconrista",
				"messages": [
					{
						"kind": "walletobjects#walletObjectMessage",
						"header": "Welcome to Baconrista",
						"body": "Featuring our new bacon donuts.",
						"displayInterval": {
							"kind": "walletobjects#timeInterval",
							"start": {
								"date": "2018-10-12T07:59:44.436Z"
							}
						}
					}
				],
				"allowMultipleUsersPerObject": true,
				"locations": [
					{
						"kind": "walletobjects#latLongPoint",
						"latitude": 37.424354,
						"longitude": -122.09508869999999
					},
					{
						"kind": "walletobjects#latLongPoint",
						"latitude": 40.7406578,
						"longitude": -74.00208940000002
					},
					{
						"kind": "walletobjects#latLongPoint",
						"latitude": 37.422601,
						"longitude": -122.085286
					}
				],
				"reviewStatus": "approved",
				"infoModuleData": {
					"showLastUpdateTime": true
				},
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
						"header": "Rewards details",
						"body": "Welcome to Baconrista rewards.  Enjoy your rewards for being a loyal customer.  10 points for ever dollar spent.  Redeem your points for free coffee, bacon and more!"
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
				"programName": "Baconrista Rewards",
				"programLogo": {
					"kind": "walletobjects#image",
					"sourceUri": {
						"kind": "walletobjects#uri",
						"uri": "http://farm8.staticflickr.com/7340/11177041185_a61a7f2139_o.jpg"
					}
				},
				"accountNameLabel": "Member Name",
				"accountIdLabel": "Member Id",
				"rewardsTierLabel": "Tier",
				"rewardsTier": "Gold"
			},
			{
				"kind": "walletobjects#loyaltyClass",
				"id": "1114132711145979111.LoyaltyClass02",
				"version": "1",
				"issuerName": "Baconrista",
				"messages": [
					{
						"kind": "walletobjects#walletObjectMessage",
						"header": "Welcome to Baconrista",
						"body": "Featuring our new bacon donuts.",
						"displayInterval": {
							"kind": "walletobjects#timeInterval",
							"start": {
								"date": "2018-10-12T08:10:18.030Z"
							}
						}
					}
				],
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
				"infoModuleData": {
					"showLastUpdateTime": true
				},
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
						"header": "Rewards details",
						"body": "Welcome to Baconrista rewards.  Enjoy your rewards for being a loyal customer.  10 points for ever dollar spent.  Redeem your points for free coffee, bacon and more!"
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
				"programName": "Baconrista Rewards",
				"programLogo": {
					"kind": "walletobjects#image",
					"sourceUri": {
						"kind": "walletobjects#uri",
						"uri": "http://farm8.staticflickr.com/7340/11177041185_a61a7f2139_o.jpg"
					}
				},
				"accountNameLabel": "Member Name",
				"accountIdLabel": "Member Id",
				"rewardsTierLabel": "Tier",
				"rewardsTier": "Gold"
			},
			{
				"kind": "walletobjects#loyaltyClass",
				"id": "1114132711145979111.TestGiftCardClass.5577006791947779410",
				"version": "1",
				"issuerName": "thecoffeeshop",
				"reviewStatus": "approved",
				"programName": "Loyalty Card",
				"programLogo": {
					"kind": "walletobjects#image",
					"sourceUri": {
						"kind": "walletobjects#uri",
						"uri": "http://farm8.staticflickr.com/7340/11177041185_a61a7f2139_o.jpg"
					}
				}
			},
			{
				"kind": "walletobjects#loyaltyClass",
				"id": "1114132711145979111.TestLoyaltyClass.1",
				"version": "1",
				"issuerName": "thecoffeeshop",
				"reviewStatus": "approved",
				"programName": "Loyalty Card",
				"programLogo": {
					"kind": "walletobjects#image",
					"sourceUri": {
						"kind": "walletobjects#uri",
						"uri": "http://farm8.staticflickr.com/7340/11177041185_a61a7f2139_o.jpg"
					}
				}
			}
		]
	}`
)

func TestGetLoyaltyClass(t *testing.T) {
	var (
		router *mux.Router
		server *httptest.Server
	)

	teardown := func() func() {
		router = mux.NewRouter()
		server = httptest.NewServer(router)
		return func() {
			server.Close()
		}
	}()
	defer teardown()

	client := NewLoyaltyClassClient(server.URL, http.DefaultClient)

	router.HandleFunc("/loyaltyClass/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if r.Method != "GET" {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		if mux.Vars(r)["id"] != sampleLoyaltyClassID {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(sampleLoyaltyClassData))
	})

	t.Run("Get loyalty class successfully", func(t *testing.T) {
		res, err := client.Get(sampleLoyaltyClassID)
		assert.NoError(t, err)

		lc := &walletobject.LoyaltyClass{}
		err = json.Unmarshal([]byte(sampleLoyaltyClassData), lc)
		assert.NoError(t, err)

		assert.EqualValues(t, lc, res)
	})

	t.Run("Failed to get loyalty class", func(t *testing.T) {
		_, err := client.Get("abc")
		assert.Error(t, err)
	})
}

func TestListLoyaltyClasses(t *testing.T) {
	var (
		router *mux.Router
		server *httptest.Server
	)

	teardown := func() func() {
		router = mux.NewRouter()
		server = httptest.NewServer(router)
		return func() {
			server.Close()
		}
	}()
	defer teardown()

	client := NewLoyaltyClassClient(server.URL, http.DefaultClient)

	router.HandleFunc("/loyaltyClass", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if r.Method != "GET" {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		if r.URL.Query().Get("issuerId") != sampleIssuerID {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(sampleLoyaltyClassesData))
	})

	t.Run("List loyalty classes successfully", func(t *testing.T) {
		res, err := client.List(sampleIssuerID, 0, "")
		assert.NoError(t, err)

		lqr := &walletobject.ListQueryResponse{}
		err = json.Unmarshal([]byte(sampleLoyaltyClassesData), lqr)
		assert.NoError(t, err)

		assert.EqualValues(t, lqr, res)
	})

	t.Run("Failed to list loyalty classes", func(t *testing.T) {
		_, err := client.List("", 0, "")
		assert.Error(t, err)
	})
}

func TestInsertLoyaltyClass(t *testing.T) {
	var (
		router *mux.Router
		server *httptest.Server
	)

	teardown := func() func() {
		router = mux.NewRouter()
		server = httptest.NewServer(router)
		return func() {
			server.Close()
		}
	}()
	defer teardown()

	client := NewLoyaltyClassClient(server.URL, http.DefaultClient)

	router.HandleFunc("/loyaltyClass", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if r.Method != "POST" {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		body := new(bytes.Buffer)
		body.ReadFrom(r.Body)

		lc := &walletobject.LoyaltyClass{}
		err := json.Unmarshal(body.Bytes(), lc)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if lc.ID != sampleLoyaltyClassID ||
			lc.IssuerName == "" ||
			lc.ProgramLogo.SourceURI.URI == "" ||
			lc.ProgramName == "" ||
			lc.ReviewStatus == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		statuses := map[string]string{
			"approved":    "",
			"draft":       "",
			"rejected":    "",
			"underReview": "",
		}
		if _, ok := statuses[lc.ReviewStatus]; !ok {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(body.Bytes())
	})

	t.Run("Insert new loyalty class successfully", func(t *testing.T) {
		lc := &walletobject.LoyaltyClass{
			ID:           sampleLoyaltyClassID,
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

		res, err := client.Insert(lc)
		assert.NoError(t, err)

		assert.EqualValues(t, lc, res)
	})

	t.Run("Failed to insert new loyalty class", func(t *testing.T) {
		classes := []*walletobject.LoyaltyClass{
			&walletobject.LoyaltyClass{
				ID:           "???",
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
			},
			&walletobject.LoyaltyClass{
				ID:           sampleLoyaltyClassID,
				ReviewStatus: "underReview",
				ProgramName:  "Loyalty Card",
				ProgramLogo: &walletobject.Image{
					Kind: "walletobjects#image",
					SourceURI: &walletobject.URI{
						Kind: "walletobjects#uri",
						URI:  "http://farm8.staticflickr.com/7340/11177041185_a61a7f2139_o.jpg",
					},
				},
			},
			&walletobject.LoyaltyClass{
				ID:           sampleLoyaltyClassID,
				IssuerName:   "thecoffeeshop",
				ReviewStatus: "???",
				ProgramName:  "Loyalty Card",
				ProgramLogo: &walletobject.Image{
					Kind: "walletobjects#image",
					SourceURI: &walletobject.URI{
						Kind: "walletobjects#uri",
						URI:  "http://farm8.staticflickr.com/7340/11177041185_a61a7f2139_o.jpg",
					},
				},
			},
		}

		for _, c := range classes {
			_, err := client.Insert(c)
			assert.Error(t, err)
		}
	})
}

func TestPatchLoyaltyClass(t *testing.T) {
	var (
		router *mux.Router
		server *httptest.Server
	)

	teardown := func() func() {
		router = mux.NewRouter()
		server = httptest.NewServer(router)
		return func() {
			server.Close()
		}
	}()
	defer teardown()

	client := NewLoyaltyClassClient(server.URL, http.DefaultClient)

	router.HandleFunc("/loyaltyClass/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if r.Method != "PATCH" {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		if mux.Vars(r)["id"] != sampleLoyaltyClassID {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		body := new(bytes.Buffer)
		body.ReadFrom(r.Body)

		bodyValues := make(map[string]string)
		err := json.Unmarshal(body.Bytes(), &bodyValues)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		lc := &walletobject.LoyaltyClass{}
		err = json.Unmarshal([]byte(sampleLoyaltyClassData), lc)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		reviewStatusValues := map[string]string{
			"approved":    "",
			"draft":       "",
			"rejected":    "",
			"underReview": "",
		}
		if _, ok := reviewStatusValues[bodyValues["reviewStatus"]]; !ok {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		respData, err := json.Marshal(lc)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(respData)
	})

	t.Run("Patch loyalty class successfully", func(t *testing.T) {
		payload := map[string]string{
			"reviewStatus":   "underReview",
			"enableSmartTap": "true",
		}
		res, err := client.Patch(sampleLoyaltyClassID, payload)
		assert.NoError(t, err)
		assert.NotNil(t, res)
	})

	t.Run("Failed to patch loyalty class", func(t *testing.T) {
		payloads := []map[string]string{
			map[string]string{
				"enableSmartTap": "true",
			},
			map[string]string{
				"reviewStatus":   "???",
				"enableSmartTap": "true",
			},
		}

		for _, p := range payloads {
			_, err := client.Patch(sampleLoyaltyClassID, p)
			assert.Error(t, err)
		}

		payload := map[string]string{
			"reviewStatus": "underReview",
			"finePrint":    "0% off any t-shirt at Adam's Apparel.",
		}
		_, err := client.Patch("???", payload)
		assert.Error(t, err)
	})
}

func TestUpdateLoyaltyClass(t *testing.T) {
	var (
		router *mux.Router
		server *httptest.Server
	)

	teardown := func() func() {
		router = mux.NewRouter()
		server = httptest.NewServer(router)
		return func() {
			server.Close()
		}
	}()
	defer teardown()

	client := NewLoyaltyClassClient(server.URL, http.DefaultClient)

	router.HandleFunc("/loyaltyClass/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if r.Method != "PUT" {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		if mux.Vars(r)["id"] != sampleLoyaltyClassID {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		body := new(bytes.Buffer)
		body.ReadFrom(r.Body)

		lc := &walletobject.LoyaltyClass{}
		err := json.Unmarshal(body.Bytes(), lc)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if lc.ID != sampleLoyaltyClassID ||
			lc.IssuerName == "" ||
			lc.ProgramLogo.SourceURI.URI == "" ||
			lc.ProgramName == "" ||
			lc.ReviewStatus == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		statuses := map[string]string{
			"approved":    "",
			"draft":       "",
			"rejected":    "",
			"underReview": "",
		}
		if _, ok := statuses[lc.ReviewStatus]; !ok {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		respData, err := json.Marshal(lc)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(respData)
	})

	t.Run("Update loyalty class successfully", func(t *testing.T) {
		lc := &walletobject.LoyaltyClass{
			ID:           sampleLoyaltyClassID,
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

		res, err := client.Update(sampleLoyaltyClassID, lc)
		assert.NoError(t, err)

		assert.EqualValues(t, lc, res)
	})

	t.Run("Failed to update loyalty class", func(t *testing.T) {
		classes := []*walletobject.LoyaltyClass{
			&walletobject.LoyaltyClass{
				ID:           "???",
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
			},
			&walletobject.LoyaltyClass{
				ID:          sampleLoyaltyClassID,
				IssuerName:  "thecoffeeshop",
				ProgramName: "Loyalty Card",
				ProgramLogo: &walletobject.Image{
					Kind: "walletobjects#image",
					SourceURI: &walletobject.URI{
						Kind: "walletobjects#uri",
						URI:  "http://farm8.staticflickr.com/7340/11177041185_a61a7f2139_o.jpg",
					},
				},
			},
			&walletobject.LoyaltyClass{
				ID:           sampleLoyaltyClassID,
				IssuerName:   "thecoffeeshop",
				ReviewStatus: "???",
				ProgramName:  "Loyalty Card",
				ProgramLogo: &walletobject.Image{
					Kind: "walletobjects#image",
					SourceURI: &walletobject.URI{
						Kind: "walletobjects#uri",
						URI:  "http://farm8.staticflickr.com/7340/11177041185_a61a7f2139_o.jpg",
					},
				},
			},
		}

		for _, c := range classes {
			_, err := client.Update(c.ID, c)
			assert.Error(t, err)
		}
	})
}
