package parameters

import (
  "github.com/arturoguerra/goautoplex/internal/models"
  "flag"
  "errors"
)

func stringInArray(str string, list []string) bool {
  for _, v := range list {
    if v == str {
      return true
    }
  }

  return false
}

func Parse() (*models.Config, error) {
  validModes := []string{"nzbget", "deluge"}
  modePtr := flag.String("mode", "[nzbget, deluge]", "string")

  nzbgetStatusPtr := flag.String("nzbget-status", "[SUCCESS, FAILURE]", "string")
  nzbgetCategoryPtr := flag.String("nzbget-category", "[movies, tv, anime]", "string")
  nzbgetDirectoryPtr := flag.String("nzbget-directory", "nzbget source dir", "string")

  delugeTorrentIdPtr := flag.String("deluge-torrentid", "torrent id string", "string")
  delugeTorrentNamePtr := flag.String("deluge-torrentname", "name of torrent", "string")
  delugeTorrentDirPtr := flag.String("deluge-torrentdir", "deluge source dir", "string")

  flag.Parse()

  validMode := stringInArray(*modePtr, validModes)

  if validMode {
    nzbGet := models.NzbGet{*nzbgetStatusPtr, *nzbgetCategoryPtr, *nzbgetDirectoryPtr}
    deluge := models.Deluge{*delugeTorrentIdPtr, *delugeTorrentNamePtr, *delugeTorrentDirPtr}
    config := &models.Config{*modePtr, &nzbGet, &deluge}
    return config, nil
  } else {
    err := errors.New("Invalid github.com/arturoguerra/goautoplex mode")
    return &models.Config{}, err
  }
}
