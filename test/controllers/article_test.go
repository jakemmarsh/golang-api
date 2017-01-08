// +build test

package controllers_test

import (
  "testing"
  "github.com/jakemmarsh/viral-api/test_utils"
)

func TestArticleController(t *testing.T) {
  t.Run("Test1", func(t *testing.T) {
    testutils.Ok(t, nil)
  })
}
