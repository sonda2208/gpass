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
	sampleLoyaltyObjectID   = "1114132711145979111.TestLoyaltyObject.1"
	sampleLoyaltyObjectData = `
	{
		"kind": "walletobjects#loyaltyObject",
		"id": "3304132759545979026.TestLoyaltyObject.1",
		"classId": "3304132759545979026.TestLoyaltyClass.1",
		"version": "1",
		"state": "active",
		"hasUsers": false,
		"classReference": {
			"kind": "walletobjects#loyaltyClass",
			"id": "3304132759545979026.TestLoyaltyClass.1",
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
	}`
	sampleLoyaltyObjectsData = `
	{
		"pagination": {
			"kind": "walletobjects#pagination",
			"resultsPerPage": 1
		},
		"resources": [
			{
				"kind": "walletobjects#loyaltyObject",
				"id": "3304132759545979026.TestLoyaltyObject.1",
				"classId": "3304132759545979026.TestLoyaltyClass.1",
				"version": "1",
				"state": "active",
				"hasUsers": false,
				"classReference": {
					"kind": "walletobjects#loyaltyClass",
					"id": "3304132759545979026.TestLoyaltyClass.1",
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
			}
		]
	}
	`
)

func TestGetLoyaltyObject(t *testing.T) {
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

	client := NewLoyaltyObjectClient(server.URL, http.DefaultClient)

	router.HandleFunc("/loyaltyObject/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if r.Method != "GET" {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		if mux.Vars(r)["id"] != sampleLoyaltyObjectID {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		lo := &walletobject.LoyaltyObject{}
		if err := json.Unmarshal([]byte(sampleLoyaltyObjectData), lo); err != nil {
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

	t.Run("Get loyalty object successfully", func(t *testing.T) {
		res, err := client.Get(sampleLoyaltyObjectID)
		assert.NoError(t, err)

		lo := &walletobject.LoyaltyObject{}
		err = json.Unmarshal([]byte(sampleLoyaltyObjectData), lo)
		assert.NoError(t, err)

		assert.EqualValues(t, lo, res)
	})

	t.Run("Failed to get loyalty object", func(t *testing.T) {
		_, err := client.Get("???")
		assert.Error(t, err)
	})
}

func TestListLoyaltyObjects(t *testing.T) {
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

	client := NewLoyaltyObjectClient(server.URL, http.DefaultClient)

	router.HandleFunc("/loyaltyObject", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if r.Method != "GET" {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		if r.URL.Query().Get("classId") != sampleLoyaltyClassID {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(sampleLoyaltyObjectsData))
	})

	t.Run("List loyalty objects successfully", func(t *testing.T) {
		res, err := client.List(sampleLoyaltyClassID, 0, "")
		assert.NoError(t, err)

		lqr := &walletobject.ListQueryResponse{}
		err = json.Unmarshal([]byte(sampleLoyaltyObjectsData), lqr)
		assert.NoError(t, err)

		assert.EqualValues(t, lqr, res)
	})

	t.Run("Failed to list loyalty objects", func(t *testing.T) {
		_, err := client.List("", 0, "")
		assert.Error(t, err)
	})
}

func TestModifyLinkedOfferObject(t *testing.T) {
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

	client := NewLoyaltyObjectClient(server.URL, http.DefaultClient)

	router.HandleFunc("/loyaltyObject/{id}/modifyLinkedOfferObjects", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if r.Method != "POST" {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		body := new(bytes.Buffer)
		body.ReadFrom(r.Body)

		ids := &walletobject.LinkedOfferObjectIds{}
		if err := json.Unmarshal(body.Bytes(), ids); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		lo := &walletobject.LoyaltyObject{}
		if err := json.Unmarshal([]byte(sampleLoyaltyObjectData), lo); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		tmpIds := make(map[string]bool)
		for _, id := range lo.LinkedOfferIds {
			tmpIds[id] = true
		}

		for _, id := range ids.AddLinkedOfferObjectIds {
			tmpIds[id] = true
		}

		for _, id := range ids.RemoveLinkedOfferObjectIds {
			delete(tmpIds, id)
		}

		lo.LinkedOfferIds = make([]string, 0)
		for k := range tmpIds {
			lo.LinkedOfferIds = append(lo.LinkedOfferIds, k)
		}

		jsonData, err := json.Marshal(lo)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(jsonData)
	})

	t.Run("Modify linked offer object successfully", func(t *testing.T) {
		ids := &walletobject.LinkedOfferObjectIds{
			AddLinkedOfferObjectIds:    []string{"1", "2"},
			RemoveLinkedOfferObjectIds: []string{"3", "4"},
		}
		res, err := client.ModifyLinkedOfferObject(sampleLoyaltyObjectID, ids)
		assert.NoError(t, err)
		assert.EqualValues(t, res.LinkedOfferIds, ids.AddLinkedOfferObjectIds)
	})

	t.Run("Failed to modify linked offer object", func(t *testing.T) {
		ids := &walletobject.LinkedOfferObjectIds{}
		_, err := client.ModifyLinkedOfferObject("???", ids)
		assert.Error(t, err)
	})
}

func TestInsertLoyaltyObject(t *testing.T) {
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

	client := NewLoyaltyObjectClient(server.URL, http.DefaultClient)

	router.HandleFunc("/loyaltyObject", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if r.Method != "POST" {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		body := new(bytes.Buffer)
		body.ReadFrom(r.Body)

		lo := &walletobject.LoyaltyObject{}
		err := json.Unmarshal(body.Bytes(), lo)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if lo.ClassID != sampleLoyaltyClassID ||
			lo.ID == "" ||
			lo.State == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		states := map[string]string{
			"active":    "",
			"completed": "",
			"expired":   "",
			"inactive":  "",
		}
		if _, ok := states[lo.State]; !ok {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(body.Bytes())
	})

	t.Run("Insert new loyalty object successfully", func(t *testing.T) {
		lo := &walletobject.LoyaltyObject{
			ID:      sampleLoyaltyObjectID,
			ClassID: sampleLoyaltyClassID,
			State:   "active",
		}

		res, err := client.Insert(lo)
		assert.NoError(t, err)

		assert.EqualValues(t, lo, res)
	})

	t.Run("Failed to insert new loyalty object", func(t *testing.T) {
		objects := []*walletobject.LoyaltyObject{
			&walletobject.LoyaltyObject{
				ID:      sampleLoyaltyObjectID,
				ClassID: sampleLoyaltyClassID,
			},
			&walletobject.LoyaltyObject{
				ID:      sampleLoyaltyObjectID,
				ClassID: sampleLoyaltyClassID,
				State:   "???",
			},
			&walletobject.LoyaltyObject{
				ID:      sampleLoyaltyObjectID,
				ClassID: "???",
				State:   "active",
			},
		}

		for _, c := range objects {
			_, err := client.Insert(c)
			assert.Error(t, err)
		}
	})
}

func TestPatchLoyaltyObject(t *testing.T) {
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

	client := NewLoyaltyObjectClient(server.URL, http.DefaultClient)

	router.HandleFunc("/loyaltyObject/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if r.Method != "PATCH" {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		if mux.Vars(r)["id"] != sampleLoyaltyObjectID {
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

		lo := &walletobject.LoyaltyObject{}
		err = json.Unmarshal([]byte(sampleLoyaltyObjectData), lo)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		respData, err := json.Marshal(lo)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(respData)
	})

	t.Run("Patch loyalty object successfully", func(t *testing.T) {
		payload := map[string]string{
			"accountName": "AcountName",
		}
		res, err := client.Patch(sampleLoyaltyObjectID, payload)
		assert.NoError(t, err)
		assert.NotNil(t, res)
	})

	t.Run("Failed to patch loyalty object", func(t *testing.T) {
		payload := map[string]string{
			"accountName": "AcountName",
		}
		_, err := client.Patch("???", payload)
		assert.Error(t, err)
	})
}

func TestUpdateLoyaltyObject(t *testing.T) {
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

	client := NewLoyaltyObjectClient(server.URL, http.DefaultClient)

	router.HandleFunc("/loyaltyObject/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if r.Method != "PUT" {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		if mux.Vars(r)["id"] != sampleLoyaltyObjectID {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		body := new(bytes.Buffer)
		body.ReadFrom(r.Body)

		lo := &walletobject.LoyaltyObject{}
		err := json.Unmarshal(body.Bytes(), lo)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if lo.ClassID != sampleLoyaltyClassID ||
			lo.ID == "" ||
			lo.State == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		states := map[string]string{
			"active":    "",
			"completed": "",
			"expired":   "",
			"inactive":  "",
		}
		if _, ok := states[lo.State]; !ok {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		respData, err := json.Marshal(lo)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(respData)
	})

	t.Run("Update loyalty object successfully", func(t *testing.T) {
		lo := &walletobject.LoyaltyObject{
			ID:      sampleLoyaltyObjectID,
			ClassID: sampleLoyaltyClassID,
			State:   "active",
		}

		res, err := client.Update(sampleLoyaltyObjectID, lo)
		assert.NoError(t, err)

		assert.EqualValues(t, lo, res)
	})

	t.Run("Failed to update loyalty object", func(t *testing.T) {
		objects := []*walletobject.LoyaltyObject{
			&walletobject.LoyaltyObject{
				ID:      sampleLoyaltyObjectID,
				ClassID: sampleLoyaltyClassID,
			},
			&walletobject.LoyaltyObject{
				ID:      sampleLoyaltyObjectID,
				ClassID: sampleLoyaltyClassID,
				State:   "???",
			},
			&walletobject.LoyaltyObject{
				ID:      sampleLoyaltyObjectID,
				ClassID: "???",
				State:   "active",
			},
		}

		for _, c := range objects {
			_, err := client.Update(c.ID, c)
			assert.Error(t, err)
		}
	})
}
