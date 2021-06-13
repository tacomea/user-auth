package repository

import (
	"testing"
	"userCreation/domain"

	"github.com/stretchr/testify/require"
)

func TestSyncMapUserRepository(t *testing.T) {
	dbRepo := NewSyncMapUserRepository()

	email := "test@example.com"
	testData := domain.User{
		Email:    email,
		Password: []byte("password"),
	}

	t.Run("Create / Check Test", func(t *testing.T) {
		dbRepo.Create(testData)
		r1, _ := dbRepo.Check(email)
		require.Equal(t, r1, testData)
	})

	t.Run("Delete Test", func(t *testing.T) {
		dbRepo.Delete(email)
		r2, _ := dbRepo.Check(email)
		require.Empty(t, r2)
	})
}

func TestSyncMapSessionRepository(t *testing.T) {
	dbRepo := NewSyncMapSessionRepository()

	id := "1"
	testData := domain.Session{
		ID:    id,
		Email: "test@example.com",
	}

	t.Run("Create / Check Test", func(t *testing.T) {
		dbRepo.Store(testData)
		r1, _ := dbRepo.Load(id)
		require.Equal(t, r1, testData)
	})

	t.Run("Delete Test", func(t *testing.T) {
		dbRepo.Delete(id)
		r2, _ := dbRepo.Load(id)
		require.Empty(t, r2)
	})
}
