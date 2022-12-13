package utils_test

import (
	"final-project-4/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewHttpSuccess(t *testing.T) {
	stringSuccess := utils.NewHttpSuccess("Message Test", "Test Success")

	assert.NotEmpty(t, stringSuccess.Message)
	assert.NotEmpty(t, stringSuccess.Data)

	testStruct := struct {
		username string
	}{
		username: "wahyu_test",
	}

	structSuccess := utils.NewHttpSuccess("Success", testStruct)

	assert.NotEmpty(t, structSuccess.Data.username)
	assert.Equal(t, structSuccess.Data, testStruct)
}
