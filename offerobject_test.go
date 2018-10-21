package googlepasses

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/sonda2208/googlepasses-go-client/walletobject"
	"github.com/stretchr/testify/assert"

	"github.com/gorilla/mux"
)

const (
	sampleOfferObjectID   = "1114132711145979111.TestOfferObject.1"
	sampleOfferObjectData = `
	{
		"kind": "walletobjects#offerObject",
		"id": "1114132711145979111.TestOfferObject.1",
		"classId": "1114132711145979111.TestOfferClass.1",
		"version": "1",
		"state": "expired",
		"barcode": {
			"kind": "walletobjects#barcode",
			"type": "qrCode",
			"value": "hello world",
			"alternateText": "Alternate text"
		},
		"messages": [
			{
				"kind": "walletobjects#walletObjectMessage",
				"body": "This is a notification message from object.",
				"displayInterval": {
					"kind": "walletobjects#timeInterval",
					"start": {
						"date": "2018-10-17T05:39:24.009Z"
					}
				},
				"messageType": "expirationNotification"
			},
			{
				"kind": "walletobjects#walletObjectMessage",
				"header": "Custom offer expiration notification time",
				"body": "This is a new notification message from object.",
				"displayInterval": {
					"kind": "walletobjects#timeInterval",
					"start": {
						"date": "2018-10-17T06:07:00.000Z"
					}
				},
				"messageType": "expirationNotification"
			},
			{
				"kind": "walletobjects#walletObjectMessage",
				"header": "Custom offer expiration notification time",
				"body": "This is a notification message from object.",
				"displayInterval": {
					"kind": "walletobjects#timeInterval",
					"start": {
						"date": "2018-10-17T05:43:51.362Z"
					}
				},
				"messageType": "expirationNotification"
			},
			{
				"kind": "walletobjects#walletObjectMessage",
				"header": "The Coffee Shop",
				"body": "This is a message",
				"displayInterval": {
					"kind": "walletobjects#timeInterval",
					"start": {
						"date": "2018-10-17T05:14:24.244Z"
					}
				},
				"messageType": "expirationNotification"
			},
			{
				"kind": "walletobjects#walletObjectMessage",
				"header": "Notification",
				"body": "This is a new notification message from object.",
				"displayInterval": {
					"kind": "walletobjects#timeInterval",
					"start": {
						"date": "2018-10-17T07:46:00.000Z"
					}
				},
				"messageType": "expirationNotification"
			},
			{
				"kind": "walletobjects#walletObjectMessage",
				"header": "Custom offer expiration notification time",
				"body": "This is a notification message from object.",
				"displayInterval": {
					"kind": "walletobjects#timeInterval",
					"start": {
						"date": "2018-10-17T05:44:24.009Z"
					}
				},
				"messageType": "expirationNotification"
			},
			{
				"kind": "walletobjects#walletObjectMessage",
				"header": "Custom offer expiration notification time",
				"body": "This is a new notification message from object.",
				"displayInterval": {
					"kind": "walletobjects#timeInterval",
					"start": {
						"date": "2018-10-17T06:09:49.616Z"
					}
				},
				"messageType": "text"
			},
			{
				"kind": "walletobjects#walletObjectMessage",
				"header": "The Coffee Shop",
				"body": "This is a message",
				"displayInterval": {
					"kind": "walletobjects#timeInterval",
					"start": {
						"date": "2018-10-17T05:12:28.174Z"
					}
				},
				"messageType": "expirationNotification"
			},
			{
				"kind": "walletobjects#walletObjectMessage",
				"body": "This is a notification message from object.",
				"displayInterval": {
					"kind": "walletobjects#timeInterval",
					"start": {
						"date": "2018-10-17T05:30:24.009Z"
					}
				},
				"messageType": "expirationNotification"
			},
			{
				"kind": "walletobjects#walletObjectMessage",
				"body": "This is a notification message from object.",
				"displayInterval": {
					"kind": "walletobjects#timeInterval",
					"start": {
						"date": "2018-10-17T05:28:28.112Z"
					}
				},
				"messageType": "expirationNotification"
			},
			{
				"kind": "walletobjects#walletObjectMessage",
				"body": "This is a notification message from object.",
				"displayInterval": {
					"kind": "walletobjects#timeInterval",
					"start": {
						"date": "2018-10-17T05:37:24.009Z"
					}
				},
				"messageType": "expirationNotification"
			},
			{
				"kind": "walletobjects#walletObjectMessage",
				"header": "Custom offer expiration notification time",
				"body": "This is a new notification message from object.",
				"displayInterval": {
					"kind": "walletobjects#timeInterval",
					"start": {
						"date": "2018-10-17T06:09:49.617Z"
					}
				},
				"messageType": "expirationNotification"
			},
			{
				"kind": "walletobjects#walletObjectMessage",
				"header": "Custom offer expiration notification time",
				"body": "This is a new notification message from object.",
				"displayInterval": {
					"kind": "walletobjects#timeInterval",
					"start": {
						"date": "2018-10-17T06:09:49.617Z"
					}
				},
				"messageType": "expirationNotification"
			},
			{
				"kind": "walletobjects#walletObjectMessage",
				"header": "Custom offer expiration notification time",
				"body": "This is a new notification message from object.",
				"displayInterval": {
					"kind": "walletobjects#timeInterval",
					"start": {
						"date": "2018-10-17T05:49:00.000Z"
					}
				},
				"messageType": "expirationNotification"
			},
			{
				"kind": "walletobjects#walletObjectMessage",
				"header": "Custom offer expiration notification time",
				"body": "This is a new notification message from object.",
				"displayInterval": {
					"kind": "walletobjects#timeInterval",
					"start": {
						"date": "2018-10-17T06:02:00.000Z"
					}
				},
				"messageType": "expirationNotification"
			}
		],
		"validTimeInterval": {
			"kind": "walletobjects#timeInterval",
			"end": {
				"date": "2018-10-17T20:10:00.000Z"
			}
		},
		"hasUsers": true,
		"disableExpirationNotification": false,
		"classReference": {
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
		}
	}`
	sampleOfferObjectsData = `
	{
		"pagination": {
			"kind": "walletobjects#pagination",
			"resultsPerPage": 7
		},
		"resources": [
			{
				"kind": "walletobjects#offerObject",
				"id": "1114132711145979111.TestOfferObject.1",
				"classId": "1114132711145979111.TestOfferClass.1",
				"version": "1",
				"state": "expired",
				"barcode": {
					"kind": "walletobjects#barcode",
					"type": "qrCode",
					"value": "hello world",
					"alternateText": "Alternate text"
				},
				"messages": [
					{
						"kind": "walletobjects#walletObjectMessage",
						"body": "This is a notification message from object.",
						"displayInterval": {
							"kind": "walletobjects#timeInterval",
							"start": {
								"date": "2018-10-17T05:39:24.009Z"
							}
						},
						"messageType": "expirationNotification"
					},
					{
						"kind": "walletobjects#walletObjectMessage",
						"header": "Custom offer expiration notification time",
						"body": "This is a new notification message from object.",
						"displayInterval": {
							"kind": "walletobjects#timeInterval",
							"start": {
								"date": "2018-10-17T06:07:00.000Z"
							}
						},
						"messageType": "expirationNotification"
					},
					{
						"kind": "walletobjects#walletObjectMessage",
						"header": "Custom offer expiration notification time",
						"body": "This is a notification message from object.",
						"displayInterval": {
							"kind": "walletobjects#timeInterval",
							"start": {
								"date": "2018-10-17T05:43:51.362Z"
							}
						},
						"messageType": "expirationNotification"
					},
					{
						"kind": "walletobjects#walletObjectMessage",
						"header": "The Coffee Shop",
						"body": "This is a message",
						"displayInterval": {
							"kind": "walletobjects#timeInterval",
							"start": {
								"date": "2018-10-17T05:14:24.244Z"
							}
						},
						"messageType": "expirationNotification"
					},
					{
						"kind": "walletobjects#walletObjectMessage",
						"header": "Notification",
						"body": "This is a new notification message from object.",
						"displayInterval": {
							"kind": "walletobjects#timeInterval",
							"start": {
								"date": "2018-10-17T07:46:00.000Z"
							}
						},
						"messageType": "expirationNotification"
					},
					{
						"kind": "walletobjects#walletObjectMessage",
						"header": "Custom offer expiration notification time",
						"body": "This is a notification message from object.",
						"displayInterval": {
							"kind": "walletobjects#timeInterval",
							"start": {
								"date": "2018-10-17T05:44:24.009Z"
							}
						},
						"messageType": "expirationNotification"
					},
					{
						"kind": "walletobjects#walletObjectMessage",
						"header": "Custom offer expiration notification time",
						"body": "This is a new notification message from object.",
						"displayInterval": {
							"kind": "walletobjects#timeInterval",
							"start": {
								"date": "2018-10-17T06:09:49.616Z"
							}
						},
						"messageType": "text"
					},
					{
						"kind": "walletobjects#walletObjectMessage",
						"header": "The Coffee Shop",
						"body": "This is a message",
						"displayInterval": {
							"kind": "walletobjects#timeInterval",
							"start": {
								"date": "2018-10-17T05:12:28.174Z"
							}
						},
						"messageType": "expirationNotification"
					},
					{
						"kind": "walletobjects#walletObjectMessage",
						"body": "This is a notification message from object.",
						"displayInterval": {
							"kind": "walletobjects#timeInterval",
							"start": {
								"date": "2018-10-17T05:30:24.009Z"
							}
						},
						"messageType": "expirationNotification"
					},
					{
						"kind": "walletobjects#walletObjectMessage",
						"body": "This is a notification message from object.",
						"displayInterval": {
							"kind": "walletobjects#timeInterval",
							"start": {
								"date": "2018-10-17T05:28:28.112Z"
							}
						},
						"messageType": "expirationNotification"
					},
					{
						"kind": "walletobjects#walletObjectMessage",
						"body": "This is a notification message from object.",
						"displayInterval": {
							"kind": "walletobjects#timeInterval",
							"start": {
								"date": "2018-10-17T05:37:24.009Z"
							}
						},
						"messageType": "expirationNotification"
					},
					{
						"kind": "walletobjects#walletObjectMessage",
						"header": "Custom offer expiration notification time",
						"body": "This is a new notification message from object.",
						"displayInterval": {
							"kind": "walletobjects#timeInterval",
							"start": {
								"date": "2018-10-17T06:09:49.617Z"
							}
						},
						"messageType": "expirationNotification"
					},
					{
						"kind": "walletobjects#walletObjectMessage",
						"header": "Custom offer expiration notification time",
						"body": "This is a new notification message from object.",
						"displayInterval": {
							"kind": "walletobjects#timeInterval",
							"start": {
								"date": "2018-10-17T06:09:49.617Z"
							}
						},
						"messageType": "expirationNotification"
					},
					{
						"kind": "walletobjects#walletObjectMessage",
						"header": "Custom offer expiration notification time",
						"body": "This is a new notification message from object.",
						"displayInterval": {
							"kind": "walletobjects#timeInterval",
							"start": {
								"date": "2018-10-17T05:49:00.000Z"
							}
						},
						"messageType": "expirationNotification"
					},
					{
						"kind": "walletobjects#walletObjectMessage",
						"header": "Custom offer expiration notification time",
						"body": "This is a new notification message from object.",
						"displayInterval": {
							"kind": "walletobjects#timeInterval",
							"start": {
								"date": "2018-10-17T06:02:00.000Z"
							}
						},
						"messageType": "expirationNotification"
					}
				],
				"validTimeInterval": {
					"kind": "walletobjects#timeInterval",
					"end": {
						"date": "2018-10-17T20:10:00.000Z"
					}
				},
				"hasUsers": true,
				"disableExpirationNotification": false,
				"classReference": {
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
				}
			},
			{
				"kind": "walletobjects#offerObject",
				"id": "1114132711145979111.TestOfferObject.2",
				"classId": "1114132711145979111.TestOfferClass.1",
				"version": "1",
				"state": "expired",
				"barcode": {
					"kind": "walletobjects#barcode",
					"type": "qrCode",
					"value": "hello world",
					"alternateText": "TestOfferObject.2"
				},
				"messages": [
					{
						"kind": "walletobjects#walletObjectMessage",
						"body": "This is a new notification.",
						"displayInterval": {
							"kind": "walletobjects#timeInterval",
							"start": {
								"date": "2018-10-17T08:44:36.309Z"
							}
						},
						"messageType": "expirationNotification"
					},
					{
						"kind": "walletobjects#walletObjectMessage",
						"body": "This is a new notification.",
						"displayInterval": {
							"kind": "walletobjects#timeInterval",
							"start": {
								"date": "2018-10-17T08:39:00.000Z"
							}
						},
						"messageType": "expirationNotification"
					},
					{
						"kind": "walletobjects#walletObjectMessage",
						"header": "Message",
						"body": "This is a new notification.",
						"displayInterval": {
							"kind": "walletobjects#timeInterval",
							"start": {
								"date": "2018-10-17T09:14:41.415Z"
							}
						},
						"messageType": "text"
					},
					{
						"kind": "walletobjects#walletObjectMessage",
						"header": "Notification",
						"body": "This is a new notification message from object.",
						"displayInterval": {
							"kind": "walletobjects#timeInterval",
							"start": {
								"date": "2018-10-17T08:29:00.000Z"
							}
						},
						"messageType": "expirationNotification"
					},
					{
						"kind": "walletobjects#walletObjectMessage",
						"body": "This is a new notification.",
						"displayInterval": {
							"kind": "walletobjects#timeInterval",
							"start": {
								"date": "2018-10-17T08:43:33.351Z"
							}
						},
						"messageType": "expirationNotification"
					}
				],
				"validTimeInterval": {
					"kind": "walletobjects#timeInterval",
					"end": {
						"date": "2018-10-18T16:55:32.504Z"
					}
				},
				"hasUsers": false,
				"classReference": {
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
				}
			},
			{
				"kind": "walletobjects#offerObject",
				"id": "1114132711145979111.TestOfferObject.3",
				"classId": "1114132711145979111.TestOfferClass.1",
				"version": "1",
				"state": "active",
				"barcode": {
					"kind": "walletobjects#barcode",
					"type": "dataMatrix",
					"value": "hello world",
					"alternateText": "TestOfferObject.3"
				},
				"messages": [
					{
						"kind": "walletobjects#walletObjectMessage",
						"header": "Message",
						"body": "This is a new notification.",
						"displayInterval": {
							"kind": "walletobjects#timeInterval",
							"start": {
								"date": "2018-10-17T09:15:25.332Z"
							}
						},
						"messageType": "text"
					}
				],
				"hasUsers": false,
				"classReference": {
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
				}
			},
			{
				"kind": "walletobjects#offerObject",
				"id": "1114132711145979111.TestOfferObject.4",
				"classId": "1114132711145979111.TestOfferClass.1",
				"version": "1",
				"state": "expired",
				"barcode": {
					"kind": "walletobjects#barcode",
					"type": "dataMatrix",
					"value": "hello world",
					"alternateText": "TestOfferObject.4"
				},
				"messages": [
					{
						"kind": "walletobjects#walletObjectMessage",
						"header": "Message",
						"body": "This is a new notification.",
						"displayInterval": {
							"kind": "walletobjects#timeInterval",
							"start": {
								"date": "2018-10-17T09:19:54.736Z"
							}
						},
						"messageType": "text"
					},
					{
						"kind": "walletobjects#walletObjectMessage",
						"header": "Notification",
						"body": "This is a new notification.",
						"displayInterval": {
							"kind": "walletobjects#timeInterval",
							"start": {
								"date": "2018-10-17T09:31:09.139Z"
							}
						},
						"messageType": "text"
					}
				],
				"validTimeInterval": {
					"kind": "walletobjects#timeInterval",
					"end": {
						"date": "2018-10-17T17:32:50.691Z"
					}
				},
				"locations": [
					{
						"kind": "walletobjects#latLongPoint",
						"latitude": 10.799652,
						"longitude": 106.70529
					}
				],
				"hasUsers": false,
				"classReference": {
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
				}
			},
			{
				"kind": "walletobjects#offerObject",
				"id": "1114132711145979111.TestOfferObject.5",
				"classId": "1114132711145979111.TestOfferClass.1",
				"version": "1",
				"state": "active",
				"messages": [
					{
						"kind": "walletobjects#walletObjectMessage",
						"header": "Notification",
						"body": "This is a new notification.",
						"displayInterval": {
							"kind": "walletobjects#timeInterval"
						},
						"messageType": "text"
					}
				],
				"locations": [
					{
						"kind": "walletobjects#latLongPoint",
						"latitude": 10.799652,
						"longitude": 106.70529
					}
				],
				"hasUsers": false,
				"classReference": {
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
				}
			},
			{
				"kind": "walletobjects#offerObject",
				"id": "1114132711145979111.TestOfferObject.6",
				"classId": "1114132711145979111.TestOfferClass.1",
				"version": "1",
				"state": "active",
				"locations": [
					{
						"kind": "walletobjects#latLongPoint",
						"latitude": 10.799652,
						"longitude": 106.70529
					}
				],
				"hasUsers": false,
				"classReference": {
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
				}
			},
			{
				"kind": "walletobjects#offerObject",
				"id": "1114132711145979111.TestOfferObject.7",
				"classId": "1114132711145979111.TestOfferClass.1",
				"version": "1",
				"state": "active",
				"messages": [
					{
						"kind": "walletobjects#walletObjectMessage",
						"header": "Notification",
						"body": "This is a new notification.",
						"displayInterval": {
							"kind": "walletobjects#timeInterval"
						},
						"messageType": "text"
					}
				],
				"hasUsers": true,
				"classReference": {
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
				}
			}
		]
	}
	`
)

func TestGetOfferObject(t *testing.T) {
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

	client := NewOfferObjectClient(server.URL, http.DefaultClient)

	router.HandleFunc("/offerObject/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if r.Method != "GET" {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		if mux.Vars(r)["id"] != sampleOfferObjectID {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		oo := &walletobject.OfferObject{}
		if err := json.Unmarshal([]byte(sampleOfferObjectData), oo); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		jsonData, err := json.Marshal(oo)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(jsonData)
	})

	t.Run("Get offer object successfully", func(t *testing.T) {
		res, err := client.Get(sampleOfferObjectID)
		assert.NoError(t, err)

		oo := &walletobject.OfferObject{}
		err = json.Unmarshal([]byte(sampleOfferObjectData), oo)
		assert.NoError(t, err)

		assert.EqualValues(t, oo, res)
	})

	t.Run("Failed to get offer object", func(t *testing.T) {
		_, err := client.Get("???")
		assert.Error(t, err)
	})
}

func TestListOfferObjects(t *testing.T) {
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

	client := NewOfferObjectClient(server.URL, http.DefaultClient)

	router.HandleFunc("/offerObject", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if r.Method != "GET" {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		if r.URL.Query().Get("classId") != sampleOfferClassID {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(sampleOfferObjectsData))
	})

	t.Run("List offer objects successfully", func(t *testing.T) {
		res, err := client.List(sampleOfferClassID, 0, "")
		assert.NoError(t, err)

		lqr := &walletobject.ListQueryResponse{}
		err = json.Unmarshal([]byte(sampleOfferObjectsData), lqr)
		assert.NoError(t, err)

		assert.EqualValues(t, lqr, res)
	})

	t.Run("Failed to list offer objects", func(t *testing.T) {
		_, err := client.List("", 0, "")
		assert.Error(t, err)
	})
}

func TestInsertOfferObject(t *testing.T) {
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

	client := NewOfferObjectClient(server.URL, http.DefaultClient)

	router.HandleFunc("/offerObject", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if r.Method != "POST" {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		body := new(bytes.Buffer)
		body.ReadFrom(r.Body)

		oo := &walletobject.OfferObject{}
		err := json.Unmarshal(body.Bytes(), oo)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if oo.ClassID != sampleOfferClassID ||
			oo.ID == "" ||
			oo.State == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		states := map[string]string{
			"active":          "",
			"completed":       "",
			"rejecexpiredted": "",
			"inactive":        "",
		}
		if _, ok := states[oo.State]; !ok {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(body.Bytes())
	})

	t.Run("Insert new offer object successfully", func(t *testing.T) {
		oo := &walletobject.OfferObject{
			ID:      "1114132711145979111.TestOfferObject.1",
			ClassID: "1114132711145979111.TestOfferClass.1",
			State:   "active",
		}

		res, err := client.Insert(oo)
		assert.NoError(t, err)

		assert.EqualValues(t, oo, res)
	})

	t.Run("Failed to insert new offer object", func(t *testing.T) {
		objects := []*walletobject.OfferObject{
			&walletobject.OfferObject{
				ID:      sampleOfferObjectID,
				ClassID: "1114132711145979111.TestOfferClass.1",
			},
			&walletobject.OfferObject{
				ID:      sampleOfferObjectID,
				ClassID: "1114132711145979111.TestOfferClass.1",
				State:   "???",
			},
			&walletobject.OfferObject{
				ID:      sampleOfferObjectID,
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

func TestPatchOfferObject(t *testing.T) {
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

	client := NewOfferObjectClient(server.URL, http.DefaultClient)

	router.HandleFunc("/offerObject/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if r.Method != "PATCH" {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		if mux.Vars(r)["id"] != sampleOfferObjectID {
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

		oo := &walletobject.OfferObject{}
		err = json.Unmarshal([]byte(sampleOfferObjectData), oo)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		oo.DisableExpirationNotification, _ = strconv.ParseBool(bodyValues["disableExpirationNotification"])

		respData, err := json.Marshal(oo)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(respData)
	})

	t.Run("Patch offer object successfully", func(t *testing.T) {
		payload := map[string]string{
			"disableExpirationNotification": "true",
		}
		res, err := client.Patch(sampleOfferObjectID, payload)
		assert.NoError(t, err)
		assert.NotNil(t, res)
	})

	t.Run("Failed to patch offer object", func(t *testing.T) {
		payload := map[string]string{
			"disableExpirationNotification": "true",
		}
		_, err := client.Patch("???", payload)
		assert.Error(t, err)
	})
}

func TestUpdateOfferObject(t *testing.T) {
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

	client := NewOfferObjectClient(server.URL, http.DefaultClient)

	router.HandleFunc("/offerObject/{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		if r.Method != "PUT" {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		if mux.Vars(r)["id"] != sampleOfferObjectID {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		body := new(bytes.Buffer)
		body.ReadFrom(r.Body)

		noo := &walletobject.OfferObject{}
		err := json.Unmarshal(body.Bytes(), noo)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if noo.ClassID != sampleOfferClassID ||
			noo.ID == "" ||
			noo.State == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		states := map[string]string{
			"active":    "",
			"completed": "",
			"expired":   "",
			"inactive":  "",
		}
		if _, ok := states[noo.State]; !ok {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		respData, err := json.Marshal(noo)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(respData)
	})

	t.Run("Update offer object successfully", func(t *testing.T) {
		oo := &walletobject.OfferObject{
			ID:      "1114132711145979111.TestOfferObject.1",
			ClassID: "1114132711145979111.TestOfferClass.1",
			State:   "active",
		}

		res, err := client.Update(sampleOfferObjectID, oo)
		assert.NoError(t, err)

		assert.EqualValues(t, oo, res)
	})

	t.Run("Failed to update offer object", func(t *testing.T) {
		objects := []*walletobject.OfferObject{
			&walletobject.OfferObject{
				ID:      sampleOfferObjectID,
				ClassID: "1114132711145979111.TestOfferClass.1",
			},
			&walletobject.OfferObject{
				ID:      sampleOfferObjectID,
				ClassID: "1114132711145979111.TestOfferClass.1",
				State:   "???",
			},
			&walletobject.OfferObject{
				ID:      sampleOfferObjectID,
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
