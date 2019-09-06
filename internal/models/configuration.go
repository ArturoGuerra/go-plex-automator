package models

type (
  NzbGet struct {
    NzbppTotalStatus string // NZBAPP_TOTALSTATUS
    NzbppCategory string // NZBPP_CATEGORY
    NzbppDirectory string // NZBPP_DIRECTORY
  }

  Deluge struct {
    TorrentId string // TORRENT_ID
    TorrentName string // TORRENT_NAME
    TorrentDir string // TORRENT_DIR
  }

  Config struct {
    Mode string
    NzbGet *NzbGet
    Deluge *Deluge
  }
)
