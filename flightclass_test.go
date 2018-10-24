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
	sampleFlightClassID   = "1114132711145979111.TestFlightClass.1"
	sampleFlightClassData = `
	{
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
	}`
	sampleFlightClassesData = `
	{
		"pagination": {
			"kind": "walletobjects#pagination",
			"resultsPerPage": 1
		},
		"resources": [
			{
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
			}
		]
	}`
)

func TestGetFlightClass(t *testing.T) {
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

	client := NewFlightClassClient(server.URL, http.DefaultClient)

	router.HandleFunc("/flightClass/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if r.Method != "GET" {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		if mux.Vars(r)["id"] != sampleFlightClassID {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(sampleFlightClassData))
	})

	t.Run("Get flight class successfully", func(t *testing.T) {
		res, err := client.Get(sampleFlightClassID)
		assert.NoError(t, err)

		fc := &walletobject.FlightClass{}
		err = json.Unmarshal([]byte(sampleFlightClassData), fc)
		assert.NoError(t, err)

		assert.EqualValues(t, fc, res)
	})

	t.Run("Failed to get flight class", func(t *testing.T) {
		_, err := client.Get("???")
		assert.Error(t, err)
	})
}

func TestListFlightClasses(t *testing.T) {
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

	client := NewFlightClassClient(server.URL, http.DefaultClient)

	router.HandleFunc("/flightClass", func(w http.ResponseWriter, r *http.Request) {
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
		w.Write([]byte(sampleFlightClassesData))
	})

	t.Run("List flight classes successfully", func(t *testing.T) {
		res, err := client.List(sampleIssuerID, 0, "")
		assert.NoError(t, err)

		lqr := &walletobject.ListQueryResponse{}
		err = json.Unmarshal([]byte(sampleFlightClassesData), lqr)
		assert.NoError(t, err)

		assert.EqualValues(t, lqr, res)
	})

	t.Run("Failed to list flight classes", func(t *testing.T) {
		_, err := client.List("", 0, "")
		assert.Error(t, err)
	})
}

func TestInsertFlightClass(t *testing.T) {
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

	client := NewFlightClassClient(server.URL, http.DefaultClient)

	router.HandleFunc("/flightClass", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if r.Method != "POST" {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		body := new(bytes.Buffer)
		body.ReadFrom(r.Body)

		fc := &walletobject.FlightClass{}
		err := json.Unmarshal(body.Bytes(), fc)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if fc.ID != sampleFlightClassID ||
			fc.IssuerName == "" ||
			fc.ReviewStatus == "" ||
			fc.LocalScheduledDepartureDateTime == "" ||
			fc.Destination == nil ||
			fc.Origin == nil ||
			fc.FlightHeader == nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		reviewStatusValues := map[string]string{
			"approved":    "",
			"draft":       "",
			"rejected":    "",
			"underReview": "",
		}
		if _, ok := reviewStatusValues[fc.ReviewStatus]; !ok {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(body.Bytes())
	})

	t.Run("Insert new flight class successfully", func(t *testing.T) {
		fc := &walletobject.FlightClass{
			ID:                              sampleFlightClassID,
			IssuerName:                      "FancyAirline",
			LocalScheduledDepartureDateTime: "2027-03-05T06:30:00",
			ReviewStatus:                    "underReview",
			Origin: &walletobject.AirportInfo{
				AirportIataCode: "SGN",
				Terminal:        "TSN",
			},
			Destination: &walletobject.AirportInfo{
				AirportIataCode: "HAN",
				Terminal:        "NB",
			},
			FlightHeader: &walletobject.FlightHeader{
				Carrier: &walletobject.FlightCarrier{
					CarrierIataCode: "LX",
					AirlineLogo: &walletobject.Image{
						SourceURI: &walletobject.URI{
							URI: "https://cdn.logojoy.com/wp-content/uploads/2018/05/30142202/1_big-768x591.jpg",
						},
					},
				},
				FlightNumber:          "113",
				OperatingFlightNumber: "113",
			},
		}

		res, err := client.Insert(fc)
		assert.NoError(t, err)

		assert.EqualValues(t, fc, res)
	})

	t.Run("Failed to insert new flight class", func(t *testing.T) {
		classes := []*walletobject.FlightClass{
			&walletobject.FlightClass{
				ID:                              "???",
				IssuerName:                      "FancyAirline",
				LocalScheduledDepartureDateTime: "2027-03-05T06:30:00",
				ReviewStatus:                    "underReview",
				Origin: &walletobject.AirportInfo{
					AirportIataCode: "SGN",
					Terminal:        "TSN",
				},
				Destination: &walletobject.AirportInfo{
					AirportIataCode: "HAN",
					Terminal:        "NB",
				},
				FlightHeader: &walletobject.FlightHeader{
					Carrier: &walletobject.FlightCarrier{
						CarrierIataCode: "LX",
						AirlineLogo: &walletobject.Image{
							SourceURI: &walletobject.URI{
								URI: "https://cdn.logojoy.com/wp-content/uploads/2018/05/30142202/1_big-768x591.jpg",
							},
						},
					},
					FlightNumber:          "113",
					OperatingFlightNumber: "113",
				},
			},
			&walletobject.FlightClass{
				ID:                              sampleFlightClassID,
				IssuerName:                      "FancyAirline",
				LocalScheduledDepartureDateTime: "2027-03-05T06:30:00",
				ReviewStatus:                    "underReview",
				Destination: &walletobject.AirportInfo{
					AirportIataCode: "HAN",
					Terminal:        "NB",
				},
				FlightHeader: &walletobject.FlightHeader{
					Carrier: &walletobject.FlightCarrier{
						CarrierIataCode: "LX",
						AirlineLogo: &walletobject.Image{
							SourceURI: &walletobject.URI{
								URI: "https://cdn.logojoy.com/wp-content/uploads/2018/05/30142202/1_big-768x591.jpg",
							},
						},
					},
					FlightNumber:          "113",
					OperatingFlightNumber: "113",
				},
			},
			&walletobject.FlightClass{
				ID:                              sampleFlightClassID,
				IssuerName:                      "FancyAirline",
				LocalScheduledDepartureDateTime: "2027-03-05T06:30:00",
				ReviewStatus:                    "???",
				Origin: &walletobject.AirportInfo{
					AirportIataCode: "SGN",
					Terminal:        "TSN",
				},
				Destination: &walletobject.AirportInfo{
					AirportIataCode: "HAN",
					Terminal:        "NB",
				},
				FlightHeader: &walletobject.FlightHeader{
					Carrier: &walletobject.FlightCarrier{
						CarrierIataCode: "LX",
						AirlineLogo: &walletobject.Image{
							SourceURI: &walletobject.URI{
								URI: "https://cdn.logojoy.com/wp-content/uploads/2018/05/30142202/1_big-768x591.jpg",
							},
						},
					},
					FlightNumber:          "113",
					OperatingFlightNumber: "113",
				},
			},
		}

		for _, c := range classes {
			_, err := client.Insert(c)
			assert.Error(t, err)
		}
	})
}

func TestPatchFlightClass(t *testing.T) {
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

	client := NewFlightClassClient(server.URL, http.DefaultClient)

	router.HandleFunc("/flightClass/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if r.Method != "PATCH" {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		if mux.Vars(r)["id"] != sampleFlightClassID {
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

		fc := &walletobject.FlightClass{}
		err = json.Unmarshal([]byte(sampleFlightClassData), fc)
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

		fc.ReviewStatus = bodyValues["reviewStatus"]

		respData, err := json.Marshal(fc)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(respData)
	})

	t.Run("Patch flight class successfully", func(t *testing.T) {
		payload := map[string]string{
			"reviewStatus":                "underReview",
			"allowMultipleUsersPerObject": "true",
		}
		res, err := client.Patch(sampleFlightClassID, payload)
		assert.NoError(t, err)
		assert.NotNil(t, res)
	})

	t.Run("Failed to patch flight class", func(t *testing.T) {
		payloads := []map[string]string{
			map[string]string{
				"allowMultipleUsersPerObject": "true",
			},
			map[string]string{
				"reviewStatus": "???",
			},
		}

		for _, p := range payloads {
			_, err := client.Patch(sampleFlightClassID, p)
			assert.Error(t, err)
		}

		payload := map[string]string{
			"reviewStatus": "underReview",
		}
		_, err := client.Patch("???", payload)
		assert.Error(t, err)
	})
}

func TestUpdateFlightClass(t *testing.T) {
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

	client := NewFlightClassClient(server.URL, http.DefaultClient)

	router.HandleFunc("/flightClass/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if r.Method != "PUT" {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		if mux.Vars(r)["id"] != sampleFlightClassID {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		body := new(bytes.Buffer)
		body.ReadFrom(r.Body)

		fc := &walletobject.FlightClass{}
		err := json.Unmarshal(body.Bytes(), fc)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if fc.ID != sampleFlightClassID ||
			fc.IssuerName == "" ||
			fc.ReviewStatus == "" ||
			fc.LocalScheduledDepartureDateTime == "" ||
			fc.Destination == nil ||
			fc.Origin == nil ||
			fc.FlightHeader == nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		reviewStatusValues := map[string]string{
			"approved":    "",
			"draft":       "",
			"rejected":    "",
			"underReview": "",
		}
		if _, ok := reviewStatusValues[fc.ReviewStatus]; !ok {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		respData, err := json.Marshal(fc)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(respData)
	})

	t.Run("Update flight class successfully", func(t *testing.T) {
		fc := &walletobject.FlightClass{
			ID:                              sampleFlightClassID,
			IssuerName:                      "FancyAirline",
			LocalScheduledDepartureDateTime: "2027-03-05T06:30:00",
			ReviewStatus:                    "underReview",
			Origin: &walletobject.AirportInfo{
				AirportIataCode: "SGN",
				Terminal:        "TSN",
			},
			Destination: &walletobject.AirportInfo{
				AirportIataCode: "HAN",
				Terminal:        "NB",
			},
			FlightHeader: &walletobject.FlightHeader{
				Carrier: &walletobject.FlightCarrier{
					CarrierIataCode: "LX",
					AirlineLogo: &walletobject.Image{
						SourceURI: &walletobject.URI{
							URI: "https://cdn.logojoy.com/wp-content/uploads/2018/05/30142202/1_big-768x591.jpg",
						},
					},
				},
				FlightNumber:          "113",
				OperatingFlightNumber: "113",
			},
		}

		res, err := client.Update(sampleFlightClassID, fc)
		assert.NoError(t, err)

		assert.EqualValues(t, fc, res)
	})

	t.Run("Failed to update flight class", func(t *testing.T) {
		classes := []*walletobject.FlightClass{
			&walletobject.FlightClass{
				ID:                              "???",
				IssuerName:                      "FancyAirline",
				LocalScheduledDepartureDateTime: "2027-03-05T06:30:00",
				ReviewStatus:                    "underReview",
				Origin: &walletobject.AirportInfo{
					AirportIataCode: "SGN",
					Terminal:        "TSN",
				},
				Destination: &walletobject.AirportInfo{
					AirportIataCode: "HAN",
					Terminal:        "NB",
				},
				FlightHeader: &walletobject.FlightHeader{
					Carrier: &walletobject.FlightCarrier{
						CarrierIataCode: "LX",
						AirlineLogo: &walletobject.Image{
							SourceURI: &walletobject.URI{
								URI: "https://cdn.logojoy.com/wp-content/uploads/2018/05/30142202/1_big-768x591.jpg",
							},
						},
					},
					FlightNumber:          "113",
					OperatingFlightNumber: "113",
				},
			},
			&walletobject.FlightClass{
				ID:                              sampleFlightClassID,
				IssuerName:                      "FancyAirline",
				LocalScheduledDepartureDateTime: "2027-03-05T06:30:00",
				ReviewStatus:                    "underReview",
				Destination: &walletobject.AirportInfo{
					AirportIataCode: "HAN",
					Terminal:        "NB",
				},
				FlightHeader: &walletobject.FlightHeader{
					Carrier: &walletobject.FlightCarrier{
						CarrierIataCode: "LX",
						AirlineLogo: &walletobject.Image{
							SourceURI: &walletobject.URI{
								URI: "https://cdn.logojoy.com/wp-content/uploads/2018/05/30142202/1_big-768x591.jpg",
							},
						},
					},
					FlightNumber:          "113",
					OperatingFlightNumber: "113",
				},
			},
			&walletobject.FlightClass{
				ID:                              sampleFlightClassID,
				IssuerName:                      "FancyAirline",
				LocalScheduledDepartureDateTime: "2027-03-05T06:30:00",
				ReviewStatus:                    "???",
				Origin: &walletobject.AirportInfo{
					AirportIataCode: "SGN",
					Terminal:        "TSN",
				},
				Destination: &walletobject.AirportInfo{
					AirportIataCode: "HAN",
					Terminal:        "NB",
				},
				FlightHeader: &walletobject.FlightHeader{
					Carrier: &walletobject.FlightCarrier{
						CarrierIataCode: "LX",
						AirlineLogo: &walletobject.Image{
							SourceURI: &walletobject.URI{
								URI: "https://cdn.logojoy.com/wp-content/uploads/2018/05/30142202/1_big-768x591.jpg",
							},
						},
					},
					FlightNumber:          "113",
					OperatingFlightNumber: "113",
				},
			},
		}

		for _, c := range classes {
			_, err := client.Update(c.ID, c)
			assert.Error(t, err)
		}
	})
}
