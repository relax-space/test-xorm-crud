package main

import (
	"testing"

	"github.com/pangpanglabs/goutils/test"
)

func TestFruitCRUD(t *testing.T) {
	t.Run("Update_unique_code", func(t *testing.T) {
		id := 2
		f := &Fruit{
			Code:"red",
		}
		affectedRow, err := f.Update(id)
		test.Ok(t, err)
		test.Equals(t, int64(1), affectedRow)
	})

	t.Run("Update_not_id", func(t *testing.T) {
		id := 3
		f := &Fruit{
			Code:"yellow",
		}
		affectedRow, err := f.Update(id)
		test.Ok(t, err)
		test.Equals(t, int64(1), affectedRow)
	})

	t.Run("Update_same_member", func(t *testing.T) {
		id := 2
		f := &Fruit{
			Code:"blue",
		}
		affectedRow, err := f.Update(id)
		test.Ok(t, err)
		test.Equals(t, int64(1), affectedRow)
	})

	t.Run("insert_unique_code", func(t *testing.T) {
		f := &Fruit{
			Code:"red",
		}
		affectedRow, err := f.Create()
		test.Ok(t, err)
		test.Equals(t, int64(1), affectedRow)
	})

}
