package config

type (
    NzbGet struct {
        BaseDir string
    }

    Deluge struct {
        BaseDir string
        ApiUrl string
        ApiKey string
    }

    FileBot struct {
        DestDir string
        AmcLogs string
        AmcExclude string
    }
)
