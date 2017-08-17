package postgres

import (
	"reflect"
	"testing"
	"time"

	randomdata "github.com/Pallinder/go-randomdata"
	uuid "github.com/satori/go.uuid"
	"github.com/tthanh/yoblog"
)

func Test_Account(t *testing.T) {
	now := time.Now()

	account := &yoblog.Account{
		ID:        uuid.NewV4().String(),
		Email:     randomdata.Email(),
		Name:      randomdata.FirstName(randomdata.Female),
		CreatedAt: now.Unix(),
		UpdatedAt: now.Unix(),
	}

	t.Run("Create", func(t *testing.T) {
		accountID, err := accountStore.Create(account)
		if err != nil {
			t.Fatal(err)
		}

		account.ID = accountID
	})

	t.Run("GetByID", func(t *testing.T) {
		t.Run("Valid", func(t *testing.T) {
			_account, err := accountStore.GetByID(account.ID)
			if err != nil {
				t.Fatal(err)
			}

			if !reflect.DeepEqual(*account, _account) {
				t.Fatalf("Mismatch: %v != %v", account, _account)
			}
		})

		t.Run("Invalid", func(t *testing.T) {
			_, err := accountStore.GetByID(uuid.NewV4().String())
			if err == nil {
				t.Fatal("Expected error")
			}
		})
	})

	t.Run("Delete", func(t *testing.T) {
		err := accountStore.Delete(account.ID)
		if err != nil {
			t.Fatal(err)
		}
	})
}
