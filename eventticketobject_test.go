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
	sampleEventTicketObjectID   = "1114132711145979111.TestEventTicketObject.1"
	sampleEventTicketObjectData = `
	{
		"kind": "walletobjects#eventTicketObject",
		"id": "1114132711145979111.TestEventTicketObject.1",
		"classId": "1114132711145979111.TestEventTicketClass.1",
		"version": "1",
		"state": "active",
		"hasUsers": false,
		"classReference": {
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
	}`
	sampleEventTicketObjectsData = `
	{
		"pagination": {
			"kind": "walletobjects#pagination",
			"resultsPerPage": 1
		},
		"resources": [
			{
				"kind": "walletobjects#eventTicketObject",
				"id": "1114132711145979111.TestEventTicketObject.1",
				"classId": "1114132711145979111.TestEventTicketClass.1",
				"version": "1",
				"state": "active",
				"hasUsers": false,
				"classReference": {
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
			}
		]
	}
	`
)

func TestGetEventTicketObject(t *testing.T) {
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

	client := NewEventTicketObjectClient(server.URL, http.DefaultClient)

	router.HandleFunc("/eventTicketObject/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if r.Method != "GET" {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		if mux.Vars(r)["id"] != sampleEventTicketObjectID {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		lo := &walletobject.EventTicketObject{}
		if err := json.Unmarshal([]byte(sampleEventTicketObjectData), lo); err != nil {
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

	t.Run("Get event ticket object successfully", func(t *testing.T) {
		res, err := client.Get(sampleEventTicketObjectID)
		assert.NoError(t, err)

		lo := &walletobject.EventTicketObject{}
		err = json.Unmarshal([]byte(sampleEventTicketObjectData), lo)
		assert.NoError(t, err)

		assert.EqualValues(t, lo, res)
	})

	t.Run("Failed to get event ticket object", func(t *testing.T) {
		_, err := client.Get("???")
		assert.Error(t, err)
	})
}

func TestListEventTicketObjects(t *testing.T) {
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

	client := NewEventTicketObjectClient(server.URL, http.DefaultClient)

	router.HandleFunc("/eventTicketObject", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if r.Method != "GET" {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		if r.URL.Query().Get("classId") != sampleEventTicketClassID {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(sampleEventTicketObjectsData))
	})

	t.Run("List event ticket objects successfully", func(t *testing.T) {
		res, err := client.List(sampleEventTicketClassID, 0, "")
		assert.NoError(t, err)

		lqr := &walletobject.ListQueryResponse{}
		err = json.Unmarshal([]byte(sampleEventTicketObjectsData), lqr)
		assert.NoError(t, err)

		assert.EqualValues(t, lqr, res)
	})

	t.Run("Failed to list event ticket objects", func(t *testing.T) {
		_, err := client.List("", 0, "")
		assert.Error(t, err)
	})
}

func TestInsertEventTicketObject(t *testing.T) {
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

	client := NewEventTicketObjectClient(server.URL, http.DefaultClient)

	router.HandleFunc("/eventTicketObject", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if r.Method != "POST" {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		body := new(bytes.Buffer)
		body.ReadFrom(r.Body)

		et := &walletobject.EventTicketObject{}
		err := json.Unmarshal(body.Bytes(), et)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if et.ClassID != sampleEventTicketClassID ||
			et.ID == "" ||
			et.State == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		states := map[string]string{
			"active":    "",
			"completed": "",
			"expired":   "",
			"inactive":  "",
		}
		if _, ok := states[et.State]; !ok {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(body.Bytes())
	})

	t.Run("Insert new event ticket object successfully", func(t *testing.T) {
		et := &walletobject.EventTicketObject{
			ID:      sampleEventTicketObjectID,
			ClassID: sampleEventTicketClassID,
			State:   "active",
		}

		res, err := client.Insert(et)
		assert.NoError(t, err)

		assert.EqualValues(t, et, res)
	})

	t.Run("Failed to insert new event ticket object", func(t *testing.T) {
		objects := []*walletobject.EventTicketObject{}

		for _, c := range objects {
			_, err := client.Insert(c)
			assert.Error(t, err)
		}
	})
}

func TestPatchEventTicketObject(t *testing.T) {
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

	client := NewEventTicketObjectClient(server.URL, http.DefaultClient)

	router.HandleFunc("/eventTicketObject/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if r.Method != "PATCH" {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		if mux.Vars(r)["id"] != sampleEventTicketObjectID {
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

		et := &walletobject.EventTicketObject{}
		err = json.Unmarshal([]byte(sampleEventTicketObjectData), et)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
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

	t.Run("Patch event ticket object successfully", func(t *testing.T) {
		payload := map[string]string{
			"ticketHolderName": "David",
		}
		res, err := client.Patch(sampleEventTicketObjectID, payload)
		assert.NoError(t, err)
		assert.NotNil(t, res)
	})

	t.Run("Failed to patch event ticket object", func(t *testing.T) {
		payload := map[string]string{
			"ticketHolderName": "David",
		}
		_, err := client.Patch("???", payload)
		assert.Error(t, err)
	})
}

func TestUpdateEventTicketObject(t *testing.T) {
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

	client := NewEventTicketObjectClient(server.URL, http.DefaultClient)

	router.HandleFunc("/eventTicketObject/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if r.Method != "PUT" {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		if mux.Vars(r)["id"] != sampleEventTicketObjectID {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		body := new(bytes.Buffer)
		body.ReadFrom(r.Body)

		et := &walletobject.EventTicketObject{}
		err := json.Unmarshal(body.Bytes(), et)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if et.ClassID != sampleEventTicketClassID ||
			et.ID == "" ||
			et.State == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		states := map[string]string{
			"active":    "",
			"completed": "",
			"expired":   "",
			"inactive":  "",
		}
		if _, ok := states[et.State]; !ok {
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

	t.Run("Update event ticket object successfully", func(t *testing.T) {
		et := &walletobject.EventTicketObject{
			ID:      sampleEventTicketObjectID,
			ClassID: sampleEventTicketClassID,
			State:   "active",
		}

		res, err := client.Update(sampleEventTicketObjectID, et)
		assert.NoError(t, err)

		assert.EqualValues(t, et, res)
	})

	t.Run("Failed to update event ticket object", func(t *testing.T) {
		objects := []*walletobject.EventTicketObject{
			&walletobject.EventTicketObject{
				ID:      "???",
				ClassID: sampleEventTicketClassID,
				State:   "active",
			},
			&walletobject.EventTicketObject{
				ID:    sampleEventTicketObjectID,
				State: "active",
			},
			&walletobject.EventTicketObject{
				ID:      sampleEventTicketObjectID,
				ClassID: sampleEventTicketClassID,
				State:   "???",
			},
		}

		for _, c := range objects {
			_, err := client.Update(c.ID, c)
			assert.Error(t, err)
		}
	})
}
