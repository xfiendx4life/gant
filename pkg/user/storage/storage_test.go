package storage_test

import (
	"context"
	"log"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/stretchr/testify/assert"
	"github.com/xfiendx4life/gant/pkg/models"
	"github.com/xfiendx4life/gant/pkg/user/storage"
)

var testUser = models.User{
	Name:     "test",
	Email:    "test@mail.ru",
	Password: "123456",
}

func getCleaner() *pgxpool.Pool {
	pool, err := pgxpool.Connect(context.Background(),
		"postgresql://localhost:5432/test_diagram?user=test&password=123")
	if err != nil {
		log.Fatal("can't connect to db")
	}
	return pool
}

var cleaner = getCleaner()

func clean() {
	cleaner.Exec(context.Background(), "DELETE FROM users")
}

func TestCreate(t *testing.T) {
	st, err := storage.New(context.Background(),
		"postgresql://localhost:5432/test_diagram?user=test&password=123")
	defer clean()
	assert.NoError(t, err)
	id, err := st.Create(&testUser)
	assert.NoError(t, err)
	log.Println(id)
	assert.NotNil(t, id)
}

func TestCreateWithContext(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	st, err := storage.New(ctx,
		"postgresql://localhost:5432/test_diagram?user=test&password=123")
	defer clean()
	assert.NoError(t, err)
	cancel()
	_, err = st.Create(&testUser)
	log.Println(err)
	assert.Error(t, err)
}

func TestGet(t *testing.T) {
	st, err := storage.New(context.Background(),
		"postgresql://localhost:5432/test_diagram?user=test&password=123")
	defer clean()
	assert.NoError(t, err)
	cleaner.QueryRow(context.Background(),
		"INSERT INTO users (name, email, password) VALUES ($1, $2, $3)",
		testUser.Name, testUser.Email, testUser.Password)
	user, err := st.Get(testUser.Email)
	assert.NoError(t, err)
	assert.Equal(t, testUser.Name, user.Name)
	assert.Equal(t, testUser.Email, user.Email)
	assert.Equal(t, testUser.Password, user.Password)
}
