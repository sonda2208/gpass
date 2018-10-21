package googlepasses

import (
	"testing"
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
	sampleOfferObjectsData = ""
)

func TestGetOfferObject(t *testing.T) {

}

func TestListOfferObjects(t *testing.T) {

}

func TestInsertOfferObject(t *testing.T) {

}

func TestPatchOfferObject(t *testing.T) {

}

func TestUpdateOfferObject(t *testing.T) {

}
