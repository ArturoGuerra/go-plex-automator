package main

import (
  "os"
  "utils"
)
/*
mode = nzbget, deluge
nzbget:
  NZBPP_TOTALSTATUS
  NZBPP_CATEGORY
  NZBPP_DIRECTORY

deluge:
  TORRENT_ID
  TORRENT_NAME
  TORRENT_DIR
  IMPORTANT=Requires deluge web api to clean finished torrents to prevent garbage collection
*/

type Handler struct {
  Args *utils.Args
)

func New(h Handler) *Handler {
  return &Handler{}
}

func stringInArray(str string, list []string) bool {
  for _, v := range list {
    if v == str {
      return true
    }
  }

  return false
}

func Parse(h *Handler) (*Args, error) {
  validModes := []string{"nzbget", "deluge"}
  modePtr := flag.String("mode", "[nzbget, deluge]", "string")

  nzbgetStatusPtr := flag.String("nzbget-status", "[SUCCESS, FAILURE]", "string")
  nzbgetCategoryPtr := flag.String("nzbget-Category", "[movies, tv, anime]", "string")
  nzbgetDirectoryPtr := flag.String("nzbget-directory", "nzbget source dir", "string")

  delugeTorrentIdPtr := flag.String("deluge-torrentid", "torrent id string", "string")
  delugeTorrentNamePtr := flag.String("deluge-torrentname", "name of torrent", "string")
  delugeTorrentDirPtr := flag.String("deluge-torrentdir", "deluge source dir", "string")

  flag.Parse()

  validMode := stringInArray(*modePtr, validModes)

  if validMode {
    nzbGet := utils.NzbGet{*nzbgetStatusPtr, *nzbgetCategoryPtr, *nzbgetDirectoryPtr}
    deluge := utils.Deluge{*delugeTorrentIdPtr, *delugeTorrentNamePtr, *delugeTorrentDirPtr}
    args := &utils.Args{*modePtr, &nzbGet, &deluge}
    h.Args = args
    return args, nil
  }

  err := errors.New("Invalid mode")
  h.Args = &utils.Args{}
  return &utils.Args{}, err

}
