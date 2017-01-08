// +build test

package main_test

import (
  "testing"
  "github.com/jakemmarsh/viral-api/test_utils"
)

func TestMain(t *testing.T) {
  t.Run("Test1", func(t *testing.T) {
    testutils.Ok(t, nil)
  })
}
