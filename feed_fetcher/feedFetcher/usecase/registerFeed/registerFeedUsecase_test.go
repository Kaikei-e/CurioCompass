package registerFeed

import "testing"

func TestRegisterSingleFeed(t *testing.T) {

	t.Run("TestRegisterSingleFeed", func(t *testing.T) {
		err := RegisterSingleFeed()
		if err != nil {
			t.Errorf("Expected %v, but got %v", nil, err)
		}

	})
}
