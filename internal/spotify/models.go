package spotify

// SpotifyResponse represents the JSON response received from the Spotify API.
type SpotifyResponse struct {
	Device               Device  `json:"device"`
	RepeatState          string  `json:"repeat_state"`
	ShuffleState         bool    `json:"shuffle_state"`
	Context              Context `json:"context"`
	Timestamp            int64   `json:"timestamp"`
	ProgressMS           int64   `json:"progress_ms"`
	IsPlaying            bool    `json:"is_playing"`
	Item                 Item    `json:"item"`
	CurrentlyPlayingType string  `json:"currently_playing_type"`
	Actions              Actions `json:"actions"`
}

// Device represents the device information.
type Device struct {
	ID            string `json:"id"`
	IsActive      bool   `json:"is_active"`
	IsPrivate     bool   `json:"is_private_session"`
	IsRestricted  bool   `json:"is_restricted"`
	Name          string `json:"name"`
	Type          string `json:"type"`
	VolumePercent int    `json:"volume_percent"`
}

// Context represents the context information.
type Context struct {
	Type        string      `json:"type"`
	Href        string      `json:"href"`
	ExternalURL ExternalURL `json:"external_urls"`
	URI         string      `json:"uri"`
}

// ExternalURL represents the external URL information.
type ExternalURL struct {
	Spotify string `json:"spotify"`
}

// Item represents the item information (e.g., a track).
type Item struct {
	Album            Album        `json:"album"`
	Artists          []Artist     `json:"artists"`
	AvailableMarkets []string     `json:"available_markets"`
	DiscNumber       int          `json:"disc_number"`
	DurationMS       int64        `json:"duration_ms"`
	Explicit         bool         `json:"explicit"`
	ExternalIDs      ExternalIDs  `json:"external_ids"`
	ExternalURLs     ExternalURL  `json:"external_urls"`
	Href             string       `json:"href"`
	ID               string       `json:"id"`
	IsPlayable       bool         `json:"is_playable"`
	LinkedFrom       LinkedFrom   `json:"linked_from"`
	Restrictions     Restrictions `json:"restrictions"`
	Name             string       `json:"name"`
	Popularity       int          `json:"popularity"`
	PreviewURL       string       `json:"preview_url"`
	TrackNumber      int          `json:"track_number"`
	Type             string       `json:"type"`
	URI              string       `json:"uri"`
	IsLocal          bool         `json:"is_local"`
}

// Album represents the album information.
type Album struct {
	AlbumType            string       `json:"album_type"`
	TotalTracks          int          `json:"total_tracks"`
	AvailableMarkets     []string     `json:"available_markets"`
	ExternalURLs         ExternalURL  `json:"external_urls"`
	Href                 string       `json:"href"`
	ID                   string       `json:"id"`
	Images               []Image      `json:"images"`
	Name                 string       `json:"name"`
	ReleaseDate          string       `json:"release_date"`
	ReleaseDatePrecision string       `json:"release_date_precision"`
	Restrictions         Restrictions `json:"restrictions"`
	Type                 string       `json:"type"`
	URI                  string       `json:"uri"`
	Copyrights           []Copyright  `json:"copyrights"`
	ExternalIDs          ExternalIDs  `json:"external_ids"`
	Genres               []string     `json:"genres"`
	Label                string       `json:"label"`
	Popularity           int          `json:"popularity"`
	AlbumGroup           string       `json:"album_group"`
	Artists              []Artist     `json:"artists"`
}

// Image represents the image information.
type Image struct {
	URL    string `json:"url"`
	Height int    `json:"height"`
	Width  int    `json:"width"`
}

// Artist represents the artist information.
type Artist struct {
	ExternalURL ExternalURL `json:"external_urls"`
	Followers   Followers   `json:"followers"`
	Genres      []string    `json:"genres"`
	Href        string      `json:"href"`
	ID          string      `json:"id"`
	Images      []Image     `json:"images"`
	Name        string      `json:"name"`
	Popularity  int         `json:"popularity"`
	Type        string      `json:"type"`
	URI         string      `json:"uri"`
}

// Followers represents the followers information.
type Followers struct {
	Href  string `json:"href"`
	Total int    `json:"total"`
}

// ExternalIDs represents the external IDs information.
type ExternalIDs struct {
	ISRC string `json:"isrc"`
	EAN  string `json:"ean"`
	UPC  string `json:"upc"`
}

// Copyright represents the copyright information.
type Copyright struct {
	Text string `json:"text"`
	Type string `json:"type"`
}

// LinkedFrom represents the linked from information.
type LinkedFrom struct {
	// Empty for now
}

// Restrictions represents the restrictions information.
type Restrictions struct {
	Reason string `json:"reason"`
}

// Actions represents the actions information.
type Actions struct {
	InterruptingPlayback  bool `json:"interrupting_playback"`
	Pausing               bool `json:"pausing"`
	Resuming              bool `json:"resuming"`
	Seeking               bool `json:"seeking"`
	SkippingNext          bool `json:"skipping_next"`
	SkippingPrev          bool `json:"skipping_prev"`
	TogglingRepeatContext bool `json:"toggling_repeat_context"`
	TogglingShuffle       bool `json:"toggling_shuffle"`
	TogglingRepeatTrack   bool `json:"toggling_repeat_track"`
	TransferringPlayback  bool `json:"transferring_playback"`
}
