package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/your_module_path/persistence-service/model"
)

func TestMessageModel(t *testing.T) {
	msg := model.Message{
		UserID:  1,
		Message: "Hello World",
	}

	assert.Equal(t, 1, msg.UserID)
	assert.Equal(t, "Hello World", msg.Message)
}