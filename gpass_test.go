package gpass_test

import (
	"context"
	"log"
	"testing"

	"github.com/sonda2208/gpass"
	"github.com/stretchr/testify/require"
)

func TestGPass(t *testing.T) {
	ctx := context.Background()
	cli, err := gpass.NewClient(ctx, "/Users/sonda/Downloads/mmpt-233505-c59ca5545a08.json")
	require.NoError(t, err)

	oc := cli.Issuer(3304132759545979026).OfferClass("3304132759545979026.CreditPlusDemo")

	ocm, err := oc.Metadata(ctx)
	if err != nil {
		log.Fatal(err)
	}

	_ = ocm
	_, err = oc.Update(ctx, &gpass.OfferClassMetadataToUpdate{
		ReviewStatus: "UNDER_REVIEW",
		LocalizedShortTitle: &gpass.LocalizedString{
			DefaultValue: &gpass.TranslatedString{
				Language: "en",
				Value:    "Status: Processing...",
			},
		},
	})
	require.NoError(t, err)
}
