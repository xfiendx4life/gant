package storage_test

import (
	"context"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xfiendx4life/gant/pkg/models"
	"github.com/xfiendx4life/gant/pkg/user/storage"
)

var testUser = models.User{
	Name:     "test",
	Email:    "test@mail.ru",
	Password: "123456",
}

func TestCreate(t *testing.T) {
	st, err := storage.New(context.Background(),
		"postgresql://localhost:5432/test_diagram?user=test&password=123")
	assert.NoError(t, err)

	id, err := st.Create(&testUser)
	assert.NoError(t, err)
	log.Println(id)
	assert.NotNil(t, id)
}
