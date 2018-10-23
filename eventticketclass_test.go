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
	sampleEventTicketClassID   = "1114132711145979111.TestEventTicketClass.1"
	sampleEventTicketClassData = `
	{
		"kind": "walletobjects#eventTicketClass",
		"id": "1114132711145979111.TestEventTicketClass.1",
		"version": "1",
		"issuerName": "thecoffeeshop",
		"reviewStatus": "approved",
		"eventName": {
			"kind": "walletobjects#localizedString",
			"defaultValue": {
				"kind": "walletobjects#translatedString",
				"language": "en-US",
				"value": "Grand Opening"
			}
		}
	}`
	sampleEventTicketClassesData = `
	{
		"pagination": {
			"kind": "walletobjects#pagination",
			"resultsPerPage": 1
		},
		"resources": [
			{
				"kind": "walletobjects#eventTicketClass",
				"id": "1114132711145979111.TestEventTicketClass.1",
				"version": "1",
				"issuerName": "thecoffeeshop",
				"reviewStatus": "approved",
				"eventName": {
					"kind": "walletobjects#localizedString",
					"defaultValue": {
						"kind": "walletobjects#translatedString",
						"language": "en-US",
						"value": "Grand Opening"
					}
				}
			}
		]
	}`
)

func TestGetEventTicketClass(t *testing.T) {
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

	client := NewEventTicketClassClient(server.URL, http.DefaultClient)

	router.HandleFunc("/eventTicketClass/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if r.Method != "GET" {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		if mux.Vars(r)["id"] != sampleEventTicketClassID {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(sampleEventTicketClassData))
	})

	t.Run("Get event ticket class successfully", func(t *testing.T) {
		res, err := client.Get(sampleEventTicketClassID)
		assert.NoError(t, err)

		et := &walletobject.EventTicketClass{}
		err = json.Unmarshal([]byte(sampleEventTicketClassData), et)
		assert.NoError(t, err)

		assert.EqualValues(t, et, res)
	})

	t.Run("Failed to get event ticket class", func(t *testing.T) {
		_, err := client.Get("???")
		assert.Error(t, err)
	})
}

func TestListEventTicketClasses(t *testing.T) {
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

	client := NewEventTicketClassClient(server.URL, http.DefaultClient)

	router.HandleFunc("/eventTicketClass", func(w http.ResponseWriter, r *http.Request) {
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
		w.Write([]byte(sampleEventTicketClassesData))
	})

	t.Run("List event ticket classes successfully", func(t *testing.T) {
		res, err := client.List(sampleIssuerID, 0, "")
		assert.NoError(t, err)

		lqr := &walletobject.ListQueryResponse{}
		err = json.Unmarshal([]byte(sampleEventTicketClassesData), lqr)
		assert.NoError(t, err)

		assert.EqualValues(t, lqr, res)
	})

	t.Run("Failed to list event ticket classes", func(t *testing.T) {
		_, err := client.List("", 0, "")
		assert.Error(t, err)
	})
}

func TestInsertEventTicketClass(t *testing.T) {
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

	client := NewEventTicketClassClient(server.URL, http.DefaultClient)

	router.HandleFunc("/eventTicketClass", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if r.Method != "POST" {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		body := new(bytes.Buffer)
		body.ReadFrom(r.Body)

		et := &walletobject.EventTicketClass{}
		err := json.Unmarshal(body.Bytes(), et)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if et.ID != sampleEventTicketClassID ||
			et.IssuerName == "" ||
			et.ReviewStatus == "" ||
			et.EventName == nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		reviewStatusValues := map[string]string{
			"approved":    "",
			"draft":       "",
			"rejected":    "",
			"underReview": "",
		}
		if _, ok := reviewStatusValues[et.ReviewStatus]; !ok {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(body.Bytes())
	})

	t.Run("Insert new event ticket class successfully", func(t *testing.T) {
		et := &walletobject.EventTicketClass{
			ID:           sampleEventTicketClassID,
			IssuerName:   "thecoffeeshop",
			ReviewStatus: "underReview",
			EventName:    &walletobject.LocalizedString{},
		}

		res, err := client.Insert(et)
		assert.NoError(t, err)

		assert.EqualValues(t, et, res)
	})

	t.Run("Failed to insert new event ticket class", func(t *testing.T) {
		classes := []*walletobject.EventTicketClass{
			&walletobject.EventTicketClass{
				ID:           "???",
				IssuerName:   "thecoffeeshop",
				ReviewStatus: "underReview",
				EventName:    &walletobject.LocalizedString{},
			},
			&walletobject.EventTicketClass{
				ID:           sampleEventTicketClassID,
				ReviewStatus: "underReview",
				EventName:    &walletobject.LocalizedString{},
			},
			&walletobject.EventTicketClass{
				ID:           sampleEventTicketClassID,
				IssuerName:   "thecoffeeshop",
				ReviewStatus: "???",
				EventName:    &walletobject.LocalizedString{},
			},
			&walletobject.EventTicketClass{
				ID:           sampleEventTicketClassID,
				IssuerName:   "thecoffeeshop",
				ReviewStatus: "underReview",
			},
		}

		for _, c := range classes {
			_, err := client.Insert(c)
			assert.Error(t, err)
		}
	})
}

func TestPatchEventTicketClass(t *testing.T) {
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

	client := NewEventTicketClassClient(server.URL, http.DefaultClient)

	router.HandleFunc("/eventTicketClass/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if r.Method != "PATCH" {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		if mux.Vars(r)["id"] != sampleEventTicketClassID {
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

		et := &walletobject.EventTicketClass{}
		err = json.Unmarshal([]byte(sampleEventTicketClassData), et)
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

		et.ReviewStatus = bodyValues["reviewStatus"]

		respData, err := json.Marshal(et)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(respData)
	})

	t.Run("Patch event ticket class successfully", func(t *testing.T) {
		payload := map[string]string{
			"reviewStatus":                "underReview",
			"allowMultipleUsersPerObject": "true",
		}
		res, err := client.Patch(sampleEventTicketClassID, payload)
		assert.NoError(t, err)
		assert.NotNil(t, res)
	})

	t.Run("Failed to patch event ticket class", func(t *testing.T) {
		payloads := []map[string]string{
			map[string]string{
				"allowMultipleUsersPerObject": "true",
			},
			map[string]string{
				"reviewStatus":                "???",
				"allowMultipleUsersPerObject": "true",
			},
		}

		for _, p := range payloads {
			_, err := client.Patch(sampleEventTicketClassID, p)
			assert.Error(t, err)
		}

		payload := map[string]string{
			"reviewStatus":                "underReview",
			"allowMultipleUsersPerObject": "true",
		}
		_, err := client.Patch("???", payload)
		assert.Error(t, err)
	})
}

func TestUpdateEventTicketClass(t *testing.T) {
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

	client := NewEventTicketClassClient(server.URL, http.DefaultClient)

	router.HandleFunc("/eventTicketClass/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if r.Method != "PUT" {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		if mux.Vars(r)["id"] != sampleEventTicketClassID {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		body := new(bytes.Buffer)
		body.ReadFrom(r.Body)

		et := &walletobject.EventTicketClass{}
		err := json.Unmarshal(body.Bytes(), et)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if et.ID != sampleEventTicketClassID ||
			et.IssuerName == "" ||
			et.ReviewStatus == "" ||
			et.EventName == nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		reviewStatusValues := map[string]string{
			"approved":    "",
			"draft":       "",
			"rejected":    "",
			"underReview": "",
		}
		if _, ok := reviewStatusValues[et.ReviewStatus]; !ok {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		respData, err := json.Marshal(et)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(respData)
	})

	t.Run("Update event ticket class successfully", func(t *testing.T) {
		et := &walletobject.EventTicketClass{
			ID:           sampleEventTicketClassID,
			IssuerName:   "thecoffeeshop",
			ReviewStatus: "underReview",
			EventName:    &walletobject.LocalizedString{},
		}

		res, err := client.Update(sampleEventTicketClassID, et)
		assert.NoError(t, err)

		assert.EqualValues(t, et, res)
	})

	t.Run("Failed to update event ticket class", func(t *testing.T) {
		classes := []*walletobject.EventTicketClass{
			&walletobject.EventTicketClass{
				ID:           "???",
				IssuerName:   "thecoffeeshop",
				ReviewStatus: "underReview",
				EventName:    &walletobject.LocalizedString{},
			},
			&walletobject.EventTicketClass{
				ID:           sampleEventTicketClassID,
				IssuerName:   "thecoffeeshop",
				ReviewStatus: "???",
				EventName:    &walletobject.LocalizedString{},
			},
			&walletobject.EventTicketClass{
				ID:           sampleEventTicketClassID,
				IssuerName:   "thecoffeeshop",
				ReviewStatus: "underReview",
			},
		}

		for _, c := range classes {
			_, err := client.Update(c.ID, c)
			assert.Error(t, err)
		}
	})
}
