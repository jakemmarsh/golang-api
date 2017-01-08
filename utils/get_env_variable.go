package utils

import "os"

func GetEnvVariable(key string, defaultValue string) string {
  var value string

  if os.Getenv("PORT") != "" {
    value = os.Getenv("PORT")
  } else {
    value = defaultValue
  }

  return value
}
