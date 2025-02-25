package models

type SpotifyIsPlaying struct {
	Song            string
	Artist          string
	ProgressMs      int
	DurationMs      int
	SecondsLeft     int
	TimeLeft        string
	Image64PixelUrl string
}

func (s *SpotifyIsPlaying) PercentageLeft() int {
	pct := s.ProgressMs * 100 / s.DurationMs
	return pct
}
