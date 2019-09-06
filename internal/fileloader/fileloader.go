package fileloader

import (
    "io/ioutil"
    "encoding/json"
    "goplex/internal/config"
)

func LoadConfig(filename string) (config.Configuration, error) {
  bytes, err := ioutil.ReadFile(filename)
  if err != nil {
    return config.Configuration{}, err
  }

  var c config.Configuration
  err = json.Unmarshal(bytes, &c)
  if err != nil {
    return config.Configuration{}, err
  }

  return c, err
}
