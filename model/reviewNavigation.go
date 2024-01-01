package model

type ReviewNavigation struct {
	Review        *Review
	PrevID        int
	NextID        int
	IsFirstReview bool
	IsLastReview  bool
}
