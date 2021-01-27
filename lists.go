package eiscp

var resolutions = map[string]string{
	"00": "through",
	"01": "auto",
	"02": "480p",
	"03": "720p",
	"04": "1080i",
	"05": "1080p",
	"06": "source",
	"07": "[1080p, 24fs]",
	"08": "4k-upscaling",
	"13": "1680x720p",
	"15": "2560x1080p",
}

var vwm = map[string]string{
	"00": "auto",
	"01": "4-3",
	"02": "full",
	"03": "zoom",
	"04": "wide zoom",
	"05": "smart zoom",
	"UP": "up",
}

var ListeningModes = map[string]string{
	"00":     "stereo",
	"01":     "direct",
	"02":     "surround",
	"03":     "[film, game-rpg]",
	"04":     "thx",
	"05":     "[action, game-action]",
	"06":     "[musical, game-rock]",
	"07":     "mono-movie",
	"08":     "orchestra",
	"09":     "unplugged",
	"0A":     "studio-mix",
	"0B":     "tv-logic",
	"0C":     "all-ch-stereo",
	"0D":     "theater-dimensional",
	"0E":     "enhanced-7, enhance, game-sports",
	"0F":     "mono",
	"11":     "pure-audio",
	"12":     "multiplex",
	"13":     "full-mono",
	"14":     "[dolby-virtual, surround-enhanced]",
	"15":     "dts-surround-sensation",
	"16":     "audyssey-dsx",
	"1F":     "whole house",
	"23":     "stage",
	"25":     "action",
	"26":     "music",
	"2E":     "sports",
	"40":     "straight-decode",
	"41":     "dolby-ex",
	"42":     "thx-cinema",
	"43":     "thx-surround-ex",
	"44":     "thx-music",
	"45":     "thx-games",
	"50":     "[thx-u2, s1, i, s-cinema, cinema2]",
	"51":     "[thx-musicmode, thx-u2, s2, i, s-music]",
	"52":     "[thx-games, thx-u2, s2, i, s-games]",
	"80":     "[plii, pliix-movie, dolby-atmos, dolby-surround]",
	"81":     "[plii, pliix-music]",
	"82":     "[neo-6-cinema, neo-x-cinema, dts-x, neural-x]",
	"83":     "[neo-6-music, neo-x-music]",
	"84":     "[plii, pliix-thx-cinema, dolby-surround-thx-cinema]",
	"85":     "[neo-6, neo-x-thx-cinema, dts-neural-x-thx-cinema]",
	"86":     "[plii, pliix-game]",
	"87":     "neural-surr",
	"88":     "[neural-thx, nexural-surround]",
	"89":     "[plii, pliix-thx-games, dolby-surround-thx-games]",
	"8A":     "[neo-6, neo-x-thx-games, dts-neural-x-thx-games]",
	"8B":     "[plii, pliix-thx-music, dolby-surround-thx-music]",
	"8C":     "[neo-6, neo-x-thx-music, dts-neural-x-thx-music]",
	"8D":     "neural-thx-cinema",
	"8E":     "neural-thx-music",
	"8F":     "neural-thx-games",
	"90":     "pliiz-height",
	"91":     "neo-6-cinema-dts-surround-sensation",
	"92":     "neo-6-music-dts-surround-sensation",
	"93":     "neural-digital-music",
	"94":     "pliiz-height-thx-cinema",
	"95":     "pliiz-height-thx-music",
	"96":     "pliiz-height-thx-games",
	"97":     "[pliiz-height-thx-u2, s2-cinema]",
	"98":     "[pliiz-height-thx-u2, s2-music]",
	"99":     "[pliiz-height-thx-u2, s2-games]",
	"9A":     "neo-x-game",
	"A0":     "[pliix, plii-movie-audyssey-dsx]",
	"A1":     "[pliix, plii-music-audyssey-dsx]",
	"A2":     "[pliix, plii-game-audyssey-dsx]",
	"A3":     "neo-6-cinema-audyssey-dsx",
	"A4":     "neo-6-music-audyssey-dsx",
	"A5":     "neural-surround-audyssey-dsx",
	"A6":     "neural-digital-music-audyssey-dsx",
	"A7":     "dolby-ex-audyssey-dsx",
	"FF":     "auto-surround",
	"MOVIE":  "movie",
	"MUSIC":  "music",
	"GAME":   "game",
	"THX":    "thx",
	"AUTO":   "auto",
	"SURR":   "surr",
	"STEREO": "stereo",
}

type NLT struct {
	ServiceType NetSource
	UIType      string // 1 - int // 0 : List, 1 : Menu, 2 : Playback, 3 : Popup, 4 : Keyboard, 5 : Menu
	LayerType   string // 1 - int // 0 : NET TOP, 1 : Service Top,DLNA/USB/iPod Top, 2 : under 2nd Layer
	CurrentPos  string // 4 - hex
	NumItems    string // 4 - hex
	NumLayers   string // 2 - hex
	Reserved    string // 2 - unused
	IconL       NetSource
	IconR       NetSource
	Status      string // 2 -- hex -- lookup table // 00 : None, 01 : Connecting, 02 : Acquiring License, 03 : Buffering 04 : Cannot Play, 05 : Searching, 06 : Profile update, 07 : Operation disabled 08 : Server Start-up, 09 : Song rated as Favorite, 0A : Song banned from station, 0B : Authentication Failed, 0C : Spotify Paused(max 1 device), 0D : Track Not Available, 0E : Cannot Skip
	Title       string // the rest
}

type NLS struct {
	InfoType string // (A : ASCII letter, C : Cursor Info, U : Unicode letter)
	LineInfo string // (0-9 : 1st to 10th Line)
	Property string // varies based on context
	Line     string
}

type NetworkStatus struct {
	Source string
	Front  string
	Rear   string
}

type NetworkPlayStatus struct {
	State   string
	Repeat  string
	Shuffle string
}

type NetworkMenuStatus struct {
	Menu               bool
	PositiveButtonIcon bool
	NegativeButtonIcon bool
	SeekTime           bool
	ElapsedTimeMode    int
	Service            string
	ServiceName        string
}
