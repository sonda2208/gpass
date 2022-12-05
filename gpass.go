package gpass

import (
	"context"
	"crypto/rsa"
	"encoding/json"
	"io/ioutil"
	"time"

	"google.golang.org/api/option"

	"github.com/missmp/jwt-go"

	"golang.org/x/oauth2/google"

	"github.com/Hutchison-Technologies/gpass/walletobjects"
)

type Client struct {
	ProjectID   string
	wos         *walletobjects.Service
	credentials credentials
}

func NewClient(ctx context.Context, credentialFile string) (*Client, error) {
	var jsonData []byte
	if credentialFile != "" {
		b, err := ioutil.ReadFile(credentialFile)
		if err != nil {
			return nil, err
		}

		jsonData = b
	} else {
		c, err := google.FindDefaultCredentials(ctx)
		if err != nil {
			return nil, err
		}

		jsonData = c.JSON
	}

	var cred credentials
	err := json.Unmarshal(jsonData, &cred)
	if err != nil {
		return nil, err
	}

	wos, err := walletobjects.NewService(ctx, option.WithCredentialsJSON(jsonData))
	if err != nil {
		return nil, err
	}

	c := &Client{
		ProjectID:   cred.ProjectID,
		wos:         wos,
		credentials: cred,
	}
	return c, nil
}

func (c *Client) Close() error {
	return nil
}

type credentials struct {
	Type         string `json:"type"`
	ClientEmail  string `json:"client_email"`
	PrivateKeyID string `json:"private_key_id"`
	PrivateKey   string `json:"private_key"`
	TokenURL     string `json:"token_uri"`
	ProjectID    string `json:"project_id"`
}

type JWT struct {
	token        *jwt.Token
	signKey      *rsa.PrivateKey
	offerObjects []*walletobjects.OfferObject
	origins      []string
}

func NewJWT(c *Client) (*JWT, error) {
	token := jwt.New(jwt.SigningMethodRS256)
	token.Claims["iss"] = c.credentials.ClientEmail
	token.Claims["aud"] = "google"
	token.Claims["typ"] = "savetoandroidpay"

	signKey, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(c.credentials.PrivateKey))
	if err != nil {
		return nil, err
	}

	return &JWT{
		token:   token,
		signKey: signKey,
	}, nil
}

func (j *JWT) AddOfferObject(oo *OfferObject) *JWT {
	j.offerObjects = append(j.offerObjects, &walletobjects.OfferObject{
		ClassId: oo.OfferClassID,
		Id:      oo.OfferObjectID,
	})
	return j
}

func (j *JWT) AddOrigin(origins ...string) *JWT {
	j.origins = append(j.origins, origins...)
	return j
}

func (j *JWT) Sign() (string, error) {
	j.token.Claims["iat"] = time.Now().UTC().Unix()
	j.token.Claims["payload"] = map[string]interface{}{
		"offerObjects": j.offerObjects,
	}
	j.token.Claims["origins"] = j.origins
	res, err := j.token.SignedString(j.signKey)
	if err != nil {
		return "", err
	}

	return res, nil
}
