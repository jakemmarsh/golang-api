// +build test

package utils_test

import (
  "os"
  "testing"
  "github.com/jakemmarsh/viral-api/utils"
  "github.com/jakemmarsh/viral-api/test_utils"
)

func TestGetEnvVariable(t *testing.T) {
  t.Run("GetEnvVariable when variable exists", func(t *testing.T) {
    os.Setenv("FOO", "1")
    testutils.Equals(t, utils.GetEnvVariable("FOO", "5"), "1")
  })

  t.Run("GetEnvVariable when variable does not exist", func(t *testing.T) {
    os.Setenv("FOO", "")
    testutils.Equals(t, utils.GetEnvVariable("FOO", "5"), "5")
  })
}
