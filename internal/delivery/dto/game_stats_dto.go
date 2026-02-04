package dto

type GameStatsDTO struct {
	TotalGames   int     `json:"total_games"`
	TotalPrice   float64 `json:"total_price"`
	AvgGamePrice float64 `json:"avg_game_price"`
	MinGamePrice float64 `json:"min_game_price"`
	MaxGamePrice float64 `json:"max_game_price"`
}
