package pii_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/ThreeDotsLabs/pii"
)

func TestMaskingAnonymizer(t *testing.T) {
	ctx := context.Background()

	a := pii.NewStructAnonymizer[string, testStruct](pii.MaskingAnonymizer[string]{})

	s := testStruct{
		FirstName: "John",
		LastName:  "Doe",
		Company:   "ThreeDotsLabs",
	}

	anonymized, err := a.Anonymize(ctx, "***", s)
	require.NoError(t, err)

	assert.Equal(t, "***", anonymized.FirstName)
	assert.Equal(t, "***", anonymized.LastName)
	assert.Equal(t, "ThreeDotsLabs", anonymized.Company)

	deanonymized, err := a.Deanonymize(ctx, "", anonymized)
	require.NoError(t, err)

	assert.Equal(t, "***", deanonymized.FirstName)
	assert.Equal(t, "***", deanonymized.LastName)
	assert.Equal(t, "ThreeDotsLabs", deanonymized.Company)
}
