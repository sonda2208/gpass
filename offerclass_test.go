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

func TestGetOfferClass(t *testing.T) {
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

	client := NewOfferClassClient(server.URL, http.DefaultClient)

	router.HandleFunc("/offerClass/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if r.Method != "GET" {
			w.WriteHeader(http.StatusNotFound)
			return
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

		oc := &walletobject.OfferClass{}
		err = json.Unmarshal([]byte(sampleOfferClassData), oc)
		assert.NoError(t, err)

		assert.EqualValues(t, oc, res)
	})

	t.Run("Failed to get offer class", func(t *testing.T) {
		_, err := client.Get("abc")
		assert.Error(t, err)
	})
}

func TestListOfferClasses(t *testing.T) {
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

	client := NewOfferClassClient(server.URL, http.DefaultClient)

	router.HandleFunc("/offerClass", func(w http.ResponseWriter, r *http.Request) {
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
		w.Write([]byte(sampleOfferClassesData))
	})

	t.Run("List offer classes successfully", func(t *testing.T) {
		res, err := client.List(sampleIssuerID, 0, "")
		assert.NoError(t, err)

		lqr := &walletobject.ListQueryResponse{}
		err = json.Unmarshal([]byte(sampleOfferClassesData), lqr)
		assert.NoError(t, err)

		assert.EqualValues(t, lqr, res)
	})

	t.Run("Failed to list offer classes", func(t *testing.T) {
		_, err := client.List("", 0, "")
		assert.Error(t, err)
	})
}

func TestInsertOfferClass(t *testing.T) {
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

	client := NewOfferClassClient(server.URL, http.DefaultClient)

	router.HandleFunc("/offerClass", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if r.Method != "POST" {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		body := new(bytes.Buffer)
		body.ReadFrom(r.Body)

		oc := &walletobject.OfferClass{}
		err := json.Unmarshal(body.Bytes(), oc)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if oc.ID != sampleOfferClassID ||
			oc.IssuerName == "" ||
			oc.Provider == "" ||
			oc.Title == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		redemptionChannelValues := map[string]string{
			"both":                    "",
			"instore":                 "",
			"online":                  "",
			"temporaryPriceReduction": "",
		}
		if _, ok := redemptionChannelValues[oc.RedemptionChannel]; !ok {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		reviewStatusValues := map[string]string{
			"approved":    "",
			"draft":       "",
			"rejected":    "",
			"underReview": "",
		}
		if _, ok := reviewStatusValues[oc.ReviewStatus]; !ok {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(body.Bytes())
	})

	t.Run("Insert new offer class successfully", func(t *testing.T) {
		oc := &walletobject.OfferClass{
			ID:                sampleOfferClassID,
			IssuerName:        "thecoffeeshop",
			ReviewStatus:      "underReview",
			Provider:          "thecoffeeshop",
			RedemptionChannel: "online",
			Title:             "20% off",
		}

		res, err := client.Insert(oc)
		assert.NoError(t, err)

		assert.EqualValues(t, oc, res)
	})

	t.Run("Failed to insert new offer class", func(t *testing.T) {
		classes := []*walletobject.OfferClass{
			&walletobject.OfferClass{
				ID:         sampleOfferClassID,
				IssuerName: "thecoffeeshop",
			},
			&walletobject.OfferClass{
				ID:                sampleOfferClassID,
				IssuerName:        "thecoffeeshop",
				ReviewStatus:      "???",
				Provider:          "thecoffeeshop",
				RedemptionChannel: "online",
				Title:             "20% off",
			},
			&walletobject.OfferClass{
				ID:                sampleOfferClassID,
				IssuerName:        "thecoffeeshop",
				ReviewStatus:      "underReview",
				Provider:          "thecoffeeshop",
				RedemptionChannel: "???",
				Title:             "20% off",
			},
		}

		for _, c := range classes {
			_, err := client.Insert(c)
			assert.Error(t, err)
		}
	})
}

func TestPatchOfferClass(t *testing.T) {
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

	client := NewOfferClassClient(server.URL, http.DefaultClient)

	router.HandleFunc("/offerClass/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if r.Method != "PATCH" {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		if mux.Vars(r)["id"] != sampleOfferClassID {
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

		oc := &walletobject.OfferClass{}
		err = json.Unmarshal([]byte(sampleOfferClassData), oc)
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

		oc.FinePrint = bodyValues["finePrint"]
		oc.ReviewStatus = bodyValues["reviewStatus"]

		respData, err := json.Marshal(oc)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(respData)
	})

	t.Run("Patch offer class successfully", func(t *testing.T) {
		payload := map[string]string{
			"reviewStatus": "underReview",
			"finePrint":    "0% off any t-shirt at Adam's Apparel.",
		}
		res, err := client.Patch(sampleOfferClassID, payload)
		assert.NoError(t, err)
		assert.NotNil(t, res)
	})

	t.Run("Failed to patch offer class", func(t *testing.T) {
		payloads := []map[string]string{
			map[string]string{
				"finePrint": "0% off any t-shirt at Adam's Apparel.",
			},
			map[string]string{
				"reviewStatus": "???",
				"finePrint":    "0% off any t-shirt at Adam's Apparel.",
			},
		}

		for _, p := range payloads {
			_, err := client.Patch(sampleOfferClassID, p)
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

func TestUpdateOfferClass(t *testing.T) {
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

	client := NewOfferClassClient(server.URL, http.DefaultClient)

	router.HandleFunc("/offerClass/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if r.Method != "PUT" {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		if mux.Vars(r)["id"] != sampleOfferClassID {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		body := new(bytes.Buffer)
		body.ReadFrom(r.Body)

		noc := &walletobject.OfferClass{}
		err := json.Unmarshal(body.Bytes(), noc)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if noc.ID != sampleOfferClassID ||
			noc.IssuerName == "" ||
			noc.Provider == "" ||
			noc.Title == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		redemptionChannelValues := map[string]string{
			"both":                    "",
			"instore":                 "",
			"online":                  "",
			"temporaryPriceReduction": "",
		}
		if _, ok := redemptionChannelValues[noc.RedemptionChannel]; !ok {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		reviewStatusValues := map[string]string{
			"approved":    "",
			"draft":       "",
			"rejected":    "",
			"underReview": "",
		}
		if _, ok := reviewStatusValues[noc.ReviewStatus]; !ok {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		respData, err := json.Marshal(noc)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(respData)
	})

	t.Run("Update offer class successfully", func(t *testing.T) {
		oc := &walletobject.OfferClass{
			ID:                sampleOfferClassID,
			IssuerName:        "thecoffeeshop",
			ReviewStatus:      "underReview",
			Provider:          "thecoffeeshop",
			RedemptionChannel: "online",
			Title:             "20% off",
		}

		res, err := client.Update(sampleOfferClassID, oc)
		assert.NoError(t, err)

		assert.EqualValues(t, oc, res)
	})

	t.Run("Failed to update offer class", func(t *testing.T) {
		classes := []*walletobject.OfferClass{
			&walletobject.OfferClass{
				ID:                "???",
				IssuerName:        "thecoffeeshop",
				ReviewStatus:      "underReview",
				Provider:          "thecoffeeshop",
				RedemptionChannel: "online",
				Title:             "20% off",
			},
			&walletobject.OfferClass{
				ID:                sampleOfferClassID,
				IssuerName:        "thecoffeeshop",
				ReviewStatus:      "underReview",
				Provider:          "thecoffeeshop",
				RedemptionChannel: "online",
			},
			&walletobject.OfferClass{
				ID:                sampleOfferClassID,
				IssuerName:        "thecoffeeshop",
				ReviewStatus:      "???",
				Provider:          "thecoffeeshop",
				RedemptionChannel: "online",
				Title:             "20% off",
			},
			&walletobject.OfferClass{
				ID:                sampleOfferClassID,
				IssuerName:        "thecoffeeshop",
				ReviewStatus:      "underReview",
				Provider:          "thecoffeeshop",
				RedemptionChannel: "???",
				Title:             "20% off",
			},
		}

		for _, c := range classes {
			_, err := client.Update(c.ID, c)
			assert.Error(t, err)
		}
	})
}
