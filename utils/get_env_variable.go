package utils

import (
  "os"
)

func GetEnvVariable(key string, defaultValue string) string {
  var value string

  if os.Getenv(key) != "" {
    value = os.Getenv(key)
  } else {
    value = defaultValue
  }

  return value
}
