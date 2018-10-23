package googlepasses

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/sonda2208/googlepasses-go-client/walletobject"
	"github.com/stretchr/testify/assert"

	"github.com/gorilla/mux"
)

const (
	sampleGiftcardObjectID   = "1114132711145979111.TestGiftCardObject.1"
	sampleGiftcardObjectData = `
	{
		"kind": "walletobjects#giftCardObject",
		"id": "1114132711145979111.TestGiftCardObject.1",
		"classId": "1114132711145979111.TestGiftCardClass.1",
		"version": "1",
		"state": "active",
		"hasUsers": false,
		"classReference": {
			"kind": "walletobjects#giftCardClass",
			"id": "1114132711145979111.TestGiftCardClass.1",
			"version": "1",
			"issuerName": "thecoffeeshop",
			"allowMultipleUsersPerObject": true,
			"reviewStatus": "approved",
			"multipleDevicesAndHoldersAllowedStatus": "multipleHolders",
			"allowBarcodeRedemption": false
		},
		"cardNumber": "318"
	}`
	sampleGiftcardObjectsData = `
	{
		"pagination": {
			"kind": "walletobjects#pagination",
			"resultsPerPage": 5
		},
		"resources": [
			{
				"kind": "walletobjects#giftCardObject",
				"id": "1114132711145979111.TestGiftCardObject.100",
				"classId": "1114132711145979111.TestGiftCardClass.1",
				"version": "1",
				"state": "active",
				"barcode": {
					"kind": "walletobjects#barcode",
					"type": "qrCode",
					"value": "Hello world!",
					"alternateText": "Alternate text"
				},
				"messages": [
					{
						"kind": "walletobjects#walletObjectMessage",
						"header": "The Coffee Shop",
						"body": "Thank you!",
						"displayInterval": {
							"kind": "walletobjects#timeInterval",
							"start": {
								"date": "2018-10-17T04:47:47.411Z"
							}
						},
						"id": "1",
						"messageType": "text"
					},
					{
						"kind": "walletobjects#walletObjectMessage",
						"header": "thecoffeeshop",
						"body": "This is a message from our server",
						"id": "3",
						"messageType": "text"
					},
					{
						"kind": "walletobjects#walletObjectMessage",
						"header": "The Coffee Shop",
						"body": "This is a message",
						"displayInterval": {
							"kind": "walletobjects#timeInterval",
							"start": {
								"date": "2018-10-17T04:47:47.411Z"
							}
						},
						"id": "1",
						"messageType": "text"
					},
					{
						"kind": "walletobjects#walletObjectMessage",
						"header": "The Coffee Shop",
						"body": "This is a message",
						"displayInterval": {
							"kind": "walletobjects#timeInterval",
							"start": {
								"date": "2018-10-17T04:47:47.411Z"
							}
						},
						"id": "1",
						"messageType": "text"
					},
					{
						"kind": "walletobjects#walletObjectMessage",
						"header": "The Coffee Shop",
						"body": "This is a message",
						"displayInterval": {
							"kind": "walletobjects#timeInterval",
							"start": {
								"date": "2018-10-17T04:47:47.411Z"
							}
						},
						"id": "1",
						"messageType": "text"
					},
					{
						"kind": "walletobjects#walletObjectMessage",
						"header": "thecoffeeshop",
						"body": "thank you",
						"id": "2",
						"messageType": "text"
					},
					{
						"kind": "walletobjects#walletObjectMessage",
						"header": "thecoffeeshop",
						"body": "This is a message from our server",
						"id": "2",
						"messageType": "text"
					},
					{
						"kind": "walletobjects#walletObjectMessage",
						"header": "thecoffeeshop",
						"body": "thank you",
						"id": "2",
						"messageType": "text"
					}
				],
				"hasUsers": true,
				"classReference": {
					"kind": "walletobjects#giftCardClass",
					"id": "1114132711145979111.TestGiftCardClass.1",
					"version": "1",
					"issuerName": "thecoffeeshop",
					"allowMultipleUsersPerObject": true,
					"reviewStatus": "approved",
					"multipleDevicesAndHoldersAllowedStatus": "multipleHolders",
					"allowBarcodeRedemption": false
				},
				"cardNumber": "81"
			},
			{
				"kind": "walletobjects#giftCardObject",
				"id": "1114132711145979111.TestGiftCardObject.2",
				"classId": "1114132711145979111.TestGiftCardClass.1",
				"version": "1",
				"state": "active",
				"hasUsers": false,
				"classReference": {
					"kind": "walletobjects#giftCardClass",
					"id": "1114132711145979111.TestGiftCardClass.1",
					"version": "1",
					"issuerName": "thecoffeeshop",
					"allowMultipleUsersPerObject": true,
					"reviewStatus": "approved",
					"multipleDevicesAndHoldersAllowedStatus": "multipleHolders",
					"allowBarcodeRedemption": false
				},
				"cardNumber": "318"
			},
			{
				"kind": "walletobjects#giftCardObject",
				"id": "1114132711145979111.TestGiftCardObject.3",
				"classId": "1114132711145979111.TestGiftCardClass.1",
				"version": "1",
				"state": "active",
				"hasUsers": false,
				"classReference": {
					"kind": "walletobjects#giftCardClass",
					"id": "1114132711145979111.TestGiftCardClass.1",
					"version": "1",
					"issuerName": "thecoffeeshop",
					"allowMultipleUsersPerObject": true,
					"reviewStatus": "approved",
					"multipleDevicesAndHoldersAllowedStatus": "multipleHolders",
					"allowBarcodeRedemption": false
				},
				"cardNumber": "%!d(string=3)"
			},
			{
				"kind": "walletobjects#giftCardObject",
				"id": "1114132711145979111.TestGiftCardObject.81",
				"classId": "1114132711145979111.TestGiftCardClass.1",
				"version": "1",
				"state": "active",
				"hasUsers": false,
				"classReference": {
					"kind": "walletobjects#giftCardClass",
					"id": "1114132711145979111.TestGiftCardClass.1",
					"version": "1",
					"issuerName": "thecoffeeshop",
					"allowMultipleUsersPerObject": true,
					"reviewStatus": "approved",
					"multipleDevicesAndHoldersAllowedStatus": "multipleHolders",
					"allowBarcodeRedemption": false
				},
				"cardNumber": "81"
			},
			{
				"kind": "walletobjects#giftCardObject",
				"id": "1114132711145979111.TestGiftCardObject.887",
				"classId": "1114132711145979111.TestGiftCardClass.1",
				"version": "1",
				"state": "active",
				"hasUsers": false,
				"classReference": {
					"kind": "walletobjects#giftCardClass",
					"id": "1114132711145979111.TestGiftCardClass.1",
					"version": "1",
					"issuerName": "thecoffeeshop",
					"allowMultipleUsersPerObject": true,
					"reviewStatus": "approved",
					"multipleDevicesAndHoldersAllowedStatus": "multipleHolders",
					"allowBarcodeRedemption": false
				},
				"cardNumber": "887"
			}
		]
	}
	`
)

func TestGetGiftcardObject(t *testing.T) {
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

	client := NewGiftcardObjectClient(server.URL, http.DefaultClient)

	router.HandleFunc("/giftCardObject/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if r.Method != "GET" {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		if mux.Vars(r)["id"] != sampleGiftcardObjectID {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		lo := &walletobject.GiftcardObject{}
		if err := json.Unmarshal([]byte(sampleGiftcardObjectData), lo); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		jsonData, err := json.Marshal(lo)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(jsonData)
	})

	t.Run("Get giftcard object successfully", func(t *testing.T) {
		res, err := client.Get(sampleGiftcardObjectID)
		assert.NoError(t, err)

		lo := &walletobject.GiftcardObject{}
		err = json.Unmarshal([]byte(sampleGiftcardObjectData), lo)
		assert.NoError(t, err)

		assert.EqualValues(t, lo, res)
	})

	t.Run("Failed to get giftcard object", func(t *testing.T) {
		_, err := client.Get("???")
		assert.Error(t, err)
	})
}

func TestListGiftcardObjects(t *testing.T) {
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

	client := NewGiftcardObjectClient(server.URL, http.DefaultClient)

	router.HandleFunc("/giftCardObject", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if r.Method != "GET" {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		if r.URL.Query().Get("classId") != sampleGiftcardClassID {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(sampleGiftcardObjectsData))
	})

	t.Run("List giftcard objects successfully", func(t *testing.T) {
		res, err := client.List(sampleGiftcardClassID, 0, "")
		assert.NoError(t, err)

		lqr := &walletobject.ListQueryResponse{}
		err = json.Unmarshal([]byte(sampleGiftcardObjectsData), lqr)
		assert.NoError(t, err)

		assert.EqualValues(t, lqr, res)
	})

	t.Run("Failed to list giftcard objects", func(t *testing.T) {
		_, err := client.List("", 0, "")
		assert.Error(t, err)
	})
}

func TestInsertGiftcardObject(t *testing.T) {
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

	client := NewGiftcardObjectClient(server.URL, http.DefaultClient)

	router.HandleFunc("/giftCardObject", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if r.Method != "POST" {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		body := new(bytes.Buffer)
		body.ReadFrom(r.Body)

		gc := &walletobject.GiftcardObject{}
		err := json.Unmarshal(body.Bytes(), gc)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if gc.ClassID != sampleGiftcardClassID ||
			gc.CardNumber == "" ||
			gc.ID == "" ||
			gc.State == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		states := map[string]string{
			"active":    "",
			"completed": "",
			"expired":   "",
			"inactive":  "",
		}
		if _, ok := states[gc.State]; !ok {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(body.Bytes())
	})

	t.Run("Insert new giftcard object successfully", func(t *testing.T) {
		gc := &walletobject.GiftcardObject{
			ID:         sampleGiftcardObjectID,
			ClassID:    sampleGiftcardClassID,
			CardNumber: "11",
			State:      "active",
		}

		res, err := client.Insert(gc)
		assert.NoError(t, err)

		assert.EqualValues(t, gc, res)
	})

	t.Run("Failed to insert new giftcard object", func(t *testing.T) {
		objects := []*walletobject.GiftcardObject{
			&walletobject.GiftcardObject{
				ID:         sampleGiftcardObjectID,
				ClassID:    "???",
				CardNumber: "11",
				State:      "active",
			},
			&walletobject.GiftcardObject{
				ID:      sampleGiftcardObjectID,
				ClassID: sampleGiftcardClassID,
				State:   "active",
			},
			&walletobject.GiftcardObject{
				ID:         sampleGiftcardObjectID,
				ClassID:    sampleGiftcardClassID,
				CardNumber: "11",
				State:      "???",
			},
		}

		for _, c := range objects {
			_, err := client.Insert(c)
			assert.Error(t, err)
		}
	})
}

func TestPatchGiftcardObject(t *testing.T) {
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

	client := NewGiftcardObjectClient(server.URL, http.DefaultClient)

	router.HandleFunc("/giftCardObject/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if r.Method != "PATCH" {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		if mux.Vars(r)["id"] != sampleGiftcardObjectID {
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

		gc := &walletobject.GiftcardObject{}
		err = json.Unmarshal([]byte(sampleGiftcardObjectData), gc)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
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

	t.Run("Patch giftcard object successfully", func(t *testing.T) {
		payload := map[string]string{
			"pin": "1111",
		}
		res, err := client.Patch(sampleGiftcardObjectID, payload)
		assert.NoError(t, err)
		assert.NotNil(t, res)
	})

	t.Run("Failed to patch giftcard object", func(t *testing.T) {
		payload := map[string]string{
			"pin": "1111",
		}
		_, err := client.Patch("???", payload)
		assert.Error(t, err)
	})
}

func TestUpdateGiftcardObject(t *testing.T) {
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

	client := NewGiftcardObjectClient(server.URL, http.DefaultClient)

	router.HandleFunc("/giftCardObject/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if r.Method != "PUT" {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		if mux.Vars(r)["id"] != sampleGiftcardObjectID {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		body := new(bytes.Buffer)
		body.ReadFrom(r.Body)

		gc := &walletobject.GiftcardObject{}
		err := json.Unmarshal(body.Bytes(), gc)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if gc.ClassID != sampleGiftcardClassID ||
			gc.CardNumber == "" ||
			gc.ID == "" ||
			gc.State == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		states := map[string]string{
			"active":    "",
			"completed": "",
			"expired":   "",
			"inactive":  "",
		}
		if _, ok := states[gc.State]; !ok {
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

	t.Run("Update giftcard object successfully", func(t *testing.T) {
		gc := &walletobject.GiftcardObject{
			ID:         sampleGiftcardObjectID,
			ClassID:    sampleGiftcardClassID,
			CardNumber: "11",
			State:      "active",
		}

		res, err := client.Update(sampleGiftcardObjectID, gc)
		assert.NoError(t, err)

		assert.EqualValues(t, gc, res)
	})

	t.Run("Failed to update giftcard object", func(t *testing.T) {
		objects := []*walletobject.GiftcardObject{
			&walletobject.GiftcardObject{
				ID:         sampleGiftcardObjectID,
				ClassID:    "???",
				CardNumber: "11",
				State:      "active",
			},
			&walletobject.GiftcardObject{
				ID:      sampleGiftcardObjectID,
				ClassID: sampleGiftcardClassID,
				State:   "active",
			},
			&walletobject.GiftcardObject{
				ID:         sampleGiftcardObjectID,
				ClassID:    sampleGiftcardClassID,
				CardNumber: "11",
				State:      "???",
			},
		}

		for _, c := range objects {
			_, err := client.Update(c.ID, c)
			assert.Error(t, err)
		}
	})
}
