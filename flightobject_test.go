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
	sampleFlightObjectID   = "1114132711145979111.TestFlightObject.1"
	sampleFlightObjectData = `
	{
		"kind": "walletobjects#flightObject",
		"id": "1114132711145979111.TestFlightObject.1",
		"classId": "1114132711145979111.TestFlightClass.1",
		"version": "1",
		"state": "active",
		"hasUsers": false,
		"classReference": {
			"kind": "walletobjects#flightClass",
			"id": "1114132711145979111.TestFlightClass.1",
			"version": "1",
			"issuerName": "FancyAirline",
			"reviewStatus": "approved",
			"localScheduledDepartureDateTime": "2027-03-05T06:30",
			"flightHeader": {
				"kind": "walletobjects#flightHeader",
				"carrier": {
					"kind": "walletobjects#flightCarrier",
					"carrierIataCode": "LX",
					"airlineLogo": {
						"kind": "walletobjects#image",
						"sourceUri": {
							"kind": "walletobjects#uri",
							"uri": "https://cdn.logojoy.com/wp-content/uploads/2018/05/30142202/1_big-768x591.jpg"
						}
					}
				},
				"flightNumber": "113",
				"operatingFlightNumber": "113"
			},
			"origin": {
				"kind": "walletobjects#airportInfo",
				"airportIataCode": "SGN",
				"terminal": "TSN"
			},
			"destination": {
				"kind": "walletobjects#airportInfo",
				"airportIataCode": "HAN",
				"terminal": "NB"
			}
		},
		"passengerName": "Charles Xavier",
		"reservationInfo": {
			"kind": "walletobjects#reservationInfo",
			"confirmationCode": "xmen"
		}
	}`
	sampleFlightObjectsData = `
	{
		"pagination": {
			"kind": "walletobjects#pagination",
			"resultsPerPage": 1
		},
		"resources": [
			{
				"kind": "walletobjects#flightObject",
				"id": "1114132711145979111.TestFlightObject.1",
				"classId": "1114132711145979111.TestFlightClass.1",
				"version": "1",
				"state": "active",
				"hasUsers": false,
				"classReference": {
					"kind": "walletobjects#flightClass",
					"id": "1114132711145979111.TestFlightClass.1",
					"version": "1",
					"issuerName": "FancyAirline",
					"reviewStatus": "approved",
					"localScheduledDepartureDateTime": "2027-03-05T06:30",
					"flightHeader": {
						"kind": "walletobjects#flightHeader",
						"carrier": {
							"kind": "walletobjects#flightCarrier",
							"carrierIataCode": "LX",
							"airlineLogo": {
								"kind": "walletobjects#image",
								"sourceUri": {
									"kind": "walletobjects#uri",
									"uri": "https://cdn.logojoy.com/wp-content/uploads/2018/05/30142202/1_big-768x591.jpg"
								}
							}
						},
						"flightNumber": "113",
						"operatingFlightNumber": "113"
					},
					"origin": {
						"kind": "walletobjects#airportInfo",
						"airportIataCode": "SGN",
						"terminal": "TSN"
					},
					"destination": {
						"kind": "walletobjects#airportInfo",
						"airportIataCode": "HAN",
						"terminal": "NB"
					}
				},
				"passengerName": "Charles Xavier",
				"reservationInfo": {
					"kind": "walletobjects#reservationInfo",
					"confirmationCode": "xmen"
				}
			}
		]
	}
	`
)

func TestGetFlightObject(t *testing.T) {
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

	client := NewFlightObjectClient(server.URL, http.DefaultClient)

	router.HandleFunc("/flightObject/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if r.Method != "GET" {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		if mux.Vars(r)["id"] != sampleFlightObjectID {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		f := &walletobject.FlightObject{}
		if err := json.Unmarshal([]byte(sampleFlightObjectData), f); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		jsonData, err := json.Marshal(f)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(jsonData)
	})

	t.Run("Get flight object successfully", func(t *testing.T) {
		res, err := client.Get(sampleFlightObjectID)
		assert.NoError(t, err)

		f := &walletobject.FlightObject{}
		err = json.Unmarshal([]byte(sampleFlightObjectData), f)
		assert.NoError(t, err)

		assert.EqualValues(t, f, res)
	})

	t.Run("Failed to get flight object", func(t *testing.T) {
		_, err := client.Get("???")
		assert.Error(t, err)
	})
}

func TestListFlightObjects(t *testing.T) {
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

	client := NewFlightObjectClient(server.URL, http.DefaultClient)

	router.HandleFunc("/flightObject", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if r.Method != "GET" {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		if r.URL.Query().Get("classId") != sampleFlightClassID {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(sampleFlightObjectsData))
	})

	t.Run("List flight objects successfully", func(t *testing.T) {
		res, err := client.List(sampleFlightClassID, 0, "")
		assert.NoError(t, err)

		lqr := &walletobject.ListQueryResponse{}
		err = json.Unmarshal([]byte(sampleFlightObjectsData), lqr)
		assert.NoError(t, err)

		assert.EqualValues(t, lqr, res)
	})

	t.Run("Failed to list flight objects", func(t *testing.T) {
		_, err := client.List("", 0, "")
		assert.Error(t, err)
	})
}

func TestInsertFlightObject(t *testing.T) {
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

	client := NewFlightObjectClient(server.URL, http.DefaultClient)

	router.HandleFunc("/flightObject", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if r.Method != "POST" {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		body := new(bytes.Buffer)
		body.ReadFrom(r.Body)

		f := &walletobject.FlightObject{}
		err := json.Unmarshal(body.Bytes(), f)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if f.ClassID != sampleFlightClassID ||
			f.ID == "" ||
			f.State == "" ||
			f.PassengerName == "" ||
			f.ReservationInfo == nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		states := map[string]string{
			"active":    "",
			"completed": "",
			"expired":   "",
			"inactive":  "",
		}
		if _, ok := states[f.State]; !ok {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(body.Bytes())
	})

	t.Run("Insert new flight object successfully", func(t *testing.T) {
		f := &walletobject.FlightObject{
			ID:            sampleFlightObjectID,
			ClassID:       sampleFlightClassID,
			State:         "active",
			PassengerName: "Charles Xavier",
			ReservationInfo: &walletobject.ReservationInfo{
				ConfirmationCode: "xmen",
			},
		}

		res, err := client.Insert(f)
		assert.NoError(t, err)

		assert.EqualValues(t, f, res)
	})

	t.Run("Failed to insert new flight object", func(t *testing.T) {
		objects := []*walletobject.FlightObject{
			&walletobject.FlightObject{
				ID:            sampleFlightObjectID,
				ClassID:       "???",
				State:         "active",
				PassengerName: "Charles Xavier",
				ReservationInfo: &walletobject.ReservationInfo{
					ConfirmationCode: "xmen",
				},
			},
			&walletobject.FlightObject{
				ID:            sampleFlightObjectID,
				ClassID:       sampleFlightClassID,
				State:         "active",
				PassengerName: "Charles Xavier",
			},
			&walletobject.FlightObject{
				ID:            sampleFlightObjectID,
				ClassID:       sampleFlightClassID,
				State:         "???",
				PassengerName: "Charles Xavier",
				ReservationInfo: &walletobject.ReservationInfo{
					ConfirmationCode: "xmen",
				},
			},
		}

		for _, o := range objects {
			_, err := client.Insert(o)
			assert.Error(t, err)
		}
	})
}

func TestPatchFlightObject(t *testing.T) {
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

	client := NewFlightObjectClient(server.URL, http.DefaultClient)

	router.HandleFunc("/flightObject/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if r.Method != "PATCH" {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		if mux.Vars(r)["id"] != sampleFlightObjectID {
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

		f := &walletobject.FlightObject{}
		err = json.Unmarshal([]byte(sampleFlightObjectData), f)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		respData, err := json.Marshal(f)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(respData)
	})

	t.Run("Patch flight object successfully", func(t *testing.T) {
		payload := map[string]string{
			"smartTapRedemptionValue": "1111",
		}
		res, err := client.Patch(sampleFlightObjectID, payload)
		assert.NoError(t, err)
		assert.NotNil(t, res)
	})

	t.Run("Failed to patch flight object", func(t *testing.T) {
		payload := map[string]string{
			"smartTapRedemptionValue": "1111",
		}
		_, err := client.Patch("???", payload)
		assert.Error(t, err)
	})
}

func TestUpdateFlightObject(t *testing.T) {
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

	client := NewFlightObjectClient(server.URL, http.DefaultClient)

	router.HandleFunc("/flightObject/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if r.Method != "PUT" {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		if mux.Vars(r)["id"] != sampleFlightObjectID {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		body := new(bytes.Buffer)
		body.ReadFrom(r.Body)

		f := &walletobject.FlightObject{}
		err := json.Unmarshal(body.Bytes(), f)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if f.ClassID != sampleFlightClassID ||
			f.ID == "" ||
			f.State == "" ||
			f.PassengerName == "" ||
			f.ReservationInfo == nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		states := map[string]string{
			"active":    "",
			"completed": "",
			"expired":   "",
			"inactive":  "",
		}
		if _, ok := states[f.State]; !ok {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		respData, err := json.Marshal(f)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(respData)
	})

	t.Run("Update flight object successfully", func(t *testing.T) {
		f := &walletobject.FlightObject{
			ID:            sampleFlightObjectID,
			ClassID:       sampleFlightClassID,
			State:         "active",
			PassengerName: "Charles Xavier",
			ReservationInfo: &walletobject.ReservationInfo{
				ConfirmationCode: "xmen",
			},
		}

		res, err := client.Update(sampleFlightObjectID, f)
		assert.NoError(t, err)

		assert.EqualValues(t, f, res)
	})

	t.Run("Failed to update flight object", func(t *testing.T) {
		objects := []*walletobject.FlightObject{
			&walletobject.FlightObject{
				ID:            sampleFlightObjectID,
				ClassID:       "???",
				State:         "active",
				PassengerName: "Charles Xavier",
				ReservationInfo: &walletobject.ReservationInfo{
					ConfirmationCode: "xmen",
				},
			},
			&walletobject.FlightObject{
				ID:            sampleFlightObjectID,
				ClassID:       sampleFlightClassID,
				State:         "active",
				PassengerName: "Charles Xavier",
			},
			&walletobject.FlightObject{
				ID:            sampleFlightObjectID,
				ClassID:       sampleFlightClassID,
				State:         "???",
				PassengerName: "Charles Xavier",
				ReservationInfo: &walletobject.ReservationInfo{
					ConfirmationCode: "xmen",
				},
			},
		}

		for _, o := range objects {
			_, err := client.Update(o.ID, o)
			assert.Error(t, err)
		}
	})
}
