package config

type (
  FileBot struct {
    Amc string `json:"amc"`
    Logs string `json:"logs"`
  }

  CouchPotato struct {
    Url string `json:"url"`
    Api string `json:"api"`
  }

  SickRage struct {
    Url string `json:"url"`
    Api string `json:"api"`
  }

  Deluge struct {
    Api string `json:"api"`
    Password string `json:"password"`
    TorrentsDir string `json:"torrents_dir"`
  }

  NzbGet struct {
    BaseDir string `json:"base_dir"`
  }

  Configuration struct {
    Plex string `json:"plex"`
    NzbGet NzbGet `json:"nzbget"`
    Deluge Deluge `json:"deluge"`
    FileBot FileBot `json:"filebot"`
    SickRage SickRage `json:"sickrage"`
    CouchPotato CouchPotato `json:"couchpotato"`
  }
)
