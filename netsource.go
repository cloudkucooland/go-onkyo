package eiscp

// Source name of input channel
type NetSource string

// Sample sources
const (
	NetSrcDLNA        NetSource = "00"
	NetSrcFavorite              = "01"
	NetSrcVTuner                = "02"
	NetSrcSiriusXM              = "03"
	NetSrcPandora               = "04"
	NetSrcRhapsody              = "05"
	NetSrcLastFM                = "06"
	NetSrcNapster               = "07"
	NetSrcSlacker               = "08"
	NetSrcMediafly              = "09"
	NetSrcSpotify               = "0A"
	NetSrcAUPEO                 = "0B"
	NetSrcRadiko                = "0C"
	NetSrcEOnkyo                = "0D"
	NetSrcTuneIn                = "0E"
	NetSrcMP3Tunes              = "0F"
	NetSrcSimfy                 = "10"
	NetSrcHomeMedia             = "11"
	NetSrcDeezer                = "12"
	NetSrciHeartRadio           = "13"
	NetSrcUnknown               = "FF"
)

// SourceByName - map channel name to source enum const
var NetSourceByName = map[string]NetSource{
	"DLNA":        NetSrcDLNA,
	"Favorite":    NetSrcFavorite,
	"vTuner":      NetSrcVTuner,
	"SiriusXM":    NetSrcSiriusXM,
	"Pandora":     NetSrcPandora,
	"Rhapsody":    NetSrcRhapsody,
	"Last.fm":     NetSrcLastFM,
	"Napster":     NetSrcNapster,
	"Slacker":     NetSrcSlacker,
	"Mediafly":    NetSrcMediafly,
	"Spotify":     NetSrcSpotify,
	"AUPEO!":      NetSrcAUPEO,
	"radiko":      NetSrcRadiko,
	"e-Onkyo":     NetSrcEOnkyo,
	"TuneIn":      NetSrcTuneIn,
	"mp3Tunes":    NetSrcMP3Tunes,
	"Simfy":       NetSrcSimfy,
	"HomeMedia":   NetSrcHomeMedia,
	"Deezer":      NetSrcDeezer,
	"iHeartRadio": NetSrciHeartRadio,
	"unknown":     NetSrcUnknown,
}

var NetSourceToName = map[NetSource]string{
	NetSrcDLNA:        "DLNA",
	NetSrcFavorite:    "Favorite",
	NetSrcVTuner:      "vTuner",
	NetSrcSiriusXM:    "SiriusXM",
	NetSrcPandora:     "Pandora",
	NetSrcRhapsody:    "Rhapsody",
	NetSrcLastFM:      "Last.fm",
	NetSrcNapster:     "Napster",
	NetSrcSlacker:     "Slacker",
	NetSrcMediafly:    "Mediafly",
	NetSrcSpotify:     "Spotify",
	NetSrcAUPEO:       "AUPEO!",
	NetSrcRadiko:      "radiko",
	NetSrcEOnkyo:      "e-Onkyo",
	NetSrcTuneIn:      "TuneIn",
	NetSrcMP3Tunes:    "mp3Tunes",
	NetSrcSimfy:       "Simfy",
	NetSrcHomeMedia:   "HomeMedia",
	NetSrcDeezer:      "Deezer",
	NetSrciHeartRadio: "iHeartRadio",
	NetSrcUnknown:     "unknown",
}
