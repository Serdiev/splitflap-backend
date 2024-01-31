package models

type SpotifyIsPlaying struct {
	Song        string
	Artist      string
	ProgressMs  int
	DurationMs  int
	SecondsLeft int
	TimeLeft    string
}

func (s *SpotifyIsPlaying) PercentageLeft() int {
	pct := s.ProgressMs * 100 / s.DurationMs
	return pct
}
