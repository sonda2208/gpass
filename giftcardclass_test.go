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
	sampleGiftcardClassID   = "1114132711145979111.TestGiftCardClass.1"
	sampleGiftcardClassData = `
	{
		"kind": "walletobjects#giftCardClass",
		"id": "1114132711145979111.TestGiftCardClass.1",
		"version": "1",
		"issuerName": "thecoffeeshop",
		"allowMultipleUsersPerObject": true,
		"reviewStatus": "approved",
		"multipleDevicesAndHoldersAllowedStatus": "multipleHolders",
		"allowBarcodeRedemption": false
	}
	`
	sampleGiftcardClassesData = `
	{
		"pagination": {
			"kind": "walletobjects#pagination",
			"resultsPerPage": 5
		},
		"resources": [
			{
				"kind": "walletobjects#giftCardClass",
				"id": "1114132711145979111.GiftCardClass02",
				"version": "1",
				"issuerName": "Baconrista",
				"allowMultipleUsersPerObject": true,
				"locations": [
					{
						"kind": "walletobjects#latLongPoint",
						"latitude": 37.422601,
						"longitude": -122.085286
					}
				],
				"reviewStatus": "approved",
				"textModulesData": [
					{
						"header": "Where to Redeem",
						"body": "All US gift cards are redeemable in any US and Puerto Rico Baconrista retail locations, or online at Baconrista.com whereavailable, for merchandise or services."
					}
				],
				"linksModuleData": {
					"uris": [
						{
							"kind": "walletobjects#uri",
							"uri": "http://www.baconrista.com/",
							"description": "Baconrista"
						}
					]
				},
				"multipleDevicesAndHoldersAllowedStatus": "multipleHolders",
				"merchantName": "Baconrista",
				"programLogo": {
					"kind": "walletobjects#image",
					"sourceUri": {
						"kind": "walletobjects#uri",
						"uri": "http://farm8.staticflickr.com/7340/11177041185_a61a7f2139_o.jpg"
					}
				}
			},
			{
				"kind": "walletobjects#giftCardClass",
				"id": "1114132711145979111.TestGiftCard",
				"version": "1",
				"issuerName": "Son",
				"allowMultipleUsersPerObject": true,
				"reviewStatus": "approved",
				"multipleDevicesAndHoldersAllowedStatus": "multipleHolders",
				"allowBarcodeRedemption": false
			},
			{
				"kind": "walletobjects#giftCardClass",
				"id": "1114132711145979111.TestGiftCardClass.1",
				"version": "1",
				"issuerName": "thecoffeeshop",
				"allowMultipleUsersPerObject": true,
				"reviewStatus": "approved",
				"multipleDevicesAndHoldersAllowedStatus": "multipleHolders",
				"allowBarcodeRedemption": false
			},
			{
				"kind": "walletobjects#giftCardClass",
				"id": "1114132711145979111.TestGiftCardClass.2",
				"version": "1",
				"issuerName": "thecoffeeshop",
				"reviewStatus": "approved"
			},
			{
				"kind": "walletobjects#giftCardClass",
				"id": "1114132711145979111.TestGiftCardClass81",
				"version": "1",
				"issuerName": "thecoffeeshop",
				"allowMultipleUsersPerObject": true,
				"reviewStatus": "approved",
				"multipleDevicesAndHoldersAllowedStatus": "multipleHolders",
				"allowBarcodeRedemption": false
			}
		]
	}
	`
)

func TestGetGiftcardClass(t *testing.T) {
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

	client := NewGiftcardClassClient(server.URL, http.DefaultClient)

	router.HandleFunc("/giftCardClass/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if r.Method != "GET" {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		if mux.Vars(r)["id"] != sampleGiftcardClassID {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(sampleGiftcardClassData))
	})

	t.Run("Get giftcard class successfully", func(t *testing.T) {
		res, err := client.Get(sampleGiftcardClassID)
		assert.NoError(t, err)

		gc := &walletobject.GiftcardClass{}
		err = json.Unmarshal([]byte(sampleGiftcardClassData), gc)
		assert.NoError(t, err)

		assert.EqualValues(t, gc, res)
	})

	t.Run("Failed to get giftcard class", func(t *testing.T) {
		_, err := client.Get("???")
		assert.Error(t, err)
	})
}

func TestListGiftcardClasses(t *testing.T) {
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

	client := NewGiftcardClassClient(server.URL, http.DefaultClient)

	router.HandleFunc("/giftCardClass", func(w http.ResponseWriter, r *http.Request) {
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
		w.Write([]byte(sampleGiftcardClassesData))
	})

	t.Run("List giftcard classes successfully", func(t *testing.T) {
		res, err := client.List(sampleIssuerID, 0, "")
		assert.NoError(t, err)

		lqr := &walletobject.ListQueryResponse{}
		err = json.Unmarshal([]byte(sampleGiftcardClassesData), lqr)
		assert.NoError(t, err)

		assert.EqualValues(t, lqr, res)
	})

	t.Run("Failed to list giftcard classes", func(t *testing.T) {
		_, err := client.List("", 0, "")
		assert.Error(t, err)
	})
}

func TestInsertGiftcardClass(t *testing.T) {
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

	client := NewGiftcardClassClient(server.URL, http.DefaultClient)

	router.HandleFunc("/giftCardClass", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if r.Method != "POST" {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		body := new(bytes.Buffer)
		body.ReadFrom(r.Body)

		gc := &walletobject.GiftcardClass{}
		err := json.Unmarshal(body.Bytes(), gc)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if gc.ID != sampleGiftcardClassID ||
			gc.IssuerName == "" ||
			gc.ReviewStatus == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		statuses := map[string]string{
			"approved":    "",
			"draft":       "",
			"rejected":    "",
			"underReview": "",
		}
		if _, ok := statuses[gc.ReviewStatus]; !ok {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(body.Bytes())
	})

	t.Run("Insert new giftcard class successfully", func(t *testing.T) {
		gc := &walletobject.GiftcardClass{
			ID:           sampleGiftcardClassID,
			IssuerName:   "thecoffeeshop",
			ReviewStatus: "underReview",
		}

		res, err := client.Insert(gc)
		assert.NoError(t, err)

		assert.EqualValues(t, gc, res)
	})

	t.Run("Failed to insert new giftcard class", func(t *testing.T) {
		classes := []*walletobject.GiftcardClass{
			&walletobject.GiftcardClass{
				ID:           "???",
				IssuerName:   "thecoffeeshop",
				ReviewStatus: "underReview",
			},
			&walletobject.GiftcardClass{
				ID:           sampleGiftcardClassID,
				ReviewStatus: "underReview",
			},
			&walletobject.GiftcardClass{
				ID:           sampleGiftcardClassID,
				IssuerName:   "thecoffeeshop",
				ReviewStatus: "???",
			},
		}

		for _, c := range classes {
			_, err := client.Insert(c)
			assert.Error(t, err)
		}
	})
}

func TestPatchGiftcardClass(t *testing.T) {
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

	client := NewGiftcardClassClient(server.URL, http.DefaultClient)

	router.HandleFunc("/giftCardClass/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if r.Method != "PATCH" {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		if mux.Vars(r)["id"] != sampleGiftcardClassID {
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

		lc := &walletobject.GiftcardClass{}
		err = json.Unmarshal([]byte(sampleGiftcardClassData), lc)
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

	t.Run("Patch giftcard class successfully", func(t *testing.T) {
		payload := map[string]string{
			"reviewStatus":           "underReview",
			"allowBarcodeRedemption": "true",
		}
		res, err := client.Patch(sampleGiftcardClassID, payload)
		assert.NoError(t, err)
		assert.NotNil(t, res)
	})

	t.Run("Failed to patch giftcard class", func(t *testing.T) {
		payloads := []map[string]string{
			map[string]string{
				"allowBarcodeRedemption": "true",
			},
			map[string]string{
				"reviewStatus":           "???",
				"allowBarcodeRedemption": "true",
			},
		}

		for _, p := range payloads {
			_, err := client.Patch(sampleGiftcardClassID, p)
			assert.Error(t, err)
		}

		payload := map[string]string{
			"reviewStatus":           "underReview",
			"allowBarcodeRedemption": "true",
		}
		_, err := client.Patch("???", payload)
		assert.Error(t, err)
	})
}

func TestUpdateGiftcardClass(t *testing.T) {
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

	client := NewGiftcardClassClient(server.URL, http.DefaultClient)

	router.HandleFunc("/giftCardClass/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if r.Method != "PUT" {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		if mux.Vars(r)["id"] != sampleGiftcardClassID {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		body := new(bytes.Buffer)
		body.ReadFrom(r.Body)

		gc := &walletobject.GiftcardClass{}
		err := json.Unmarshal(body.Bytes(), gc)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if gc.ID != sampleGiftcardClassID ||
			gc.IssuerName == "" ||
			gc.ReviewStatus == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		statuses := map[string]string{
			"approved":    "",
			"draft":       "",
			"rejected":    "",
			"underReview": "",
		}
		if _, ok := statuses[gc.ReviewStatus]; !ok {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		respData, err := json.Marshal(gc)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(respData)
	})

	t.Run("Update giftcard class successfully", func(t *testing.T) {
		gc := &walletobject.GiftcardClass{
			ID:           sampleGiftcardClassID,
			IssuerName:   "thecoffeeshop",
			ReviewStatus: "underReview",
		}

		res, err := client.Update(sampleGiftcardClassID, gc)
		assert.NoError(t, err)

		assert.EqualValues(t, gc, res)
	})

	t.Run("Failed to update giftcard class", func(t *testing.T) {
		classes := []*walletobject.GiftcardClass{
			&walletobject.GiftcardClass{
				ID:           "???",
				IssuerName:   "thecoffeeshop",
				ReviewStatus: "underReview",
			},
			&walletobject.GiftcardClass{
				ID:           sampleGiftcardClassID,
				ReviewStatus: "underReview",
			},
			&walletobject.GiftcardClass{
				ID:           sampleGiftcardClassID,
				IssuerName:   "thecoffeeshop",
				ReviewStatus: "???",
			},
		}

		for _, c := range classes {
			_, err := client.Update(c.ID, c)
			assert.Error(t, err)
		}
	})
}
