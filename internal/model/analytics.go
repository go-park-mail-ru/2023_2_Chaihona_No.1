package model

type Analitycs struct {
	Id       int    `json:"id" db:"id"`
	UserId int `json:"user_id" db:"user_id"`
	TotalPosts int `json:"total_posts" db:"total_posts"`
	TotalLikes int `json:"total_likes" db:"total_likes"`
	TotalComments int `json:"total_comments" db:"total_comments"`
	TotalDonations int `json:"total_donations" db:"total_donations"`
	TotalDonationsEarnedInteger int `json:"total_donations_earned_integer" db:"total_donations_earned_integer"`
	TotalDonationsEarnedFractional int `json:"total_donations_earned_fractional" db:"total_donations_earned_fractional"`
	TotalEarnedInteger int `json:"total_earned_integer" db:"total_earned_integer"`
	TotalEarnedFractional int `json:"total_earned_fractional" db:"total_earned_fractional"`
	TotalSubscribers int `json:"total_subscribers" db:"total_subscribers"`
	//TotalSubscriberGroups []SubscribeGroup 
	DifferencePosts int `json:"difference_posts" db:"difference_posts"`
	DifferenceLikes int `json:"difference_likes" db:"difference_likes"`
	DifferenceComments int `json:"difference_comments" db:"difference_comments"`
	DifferenceDonations int `json:"difference_donations" db:"difference_donations"`
	DifferenceDonationsEarnedInteger int `json:"difference_donations_earned_integer" db:"difference_donations_earned_integer"`
	DifferenceDonationsEarnedFractional int `json:"difference_donations_earned_fractional" db:"difference_donations_earned_fractional"`
	DifferenceEarnedInteger int `json:"difference_earned_integer" db:"difference_earned_integer"`
	DifferenceEarnedFractional int `json:"difference_earned_fractional" db:"difference_earned_fractional"`
	DifferenceSubscribers int `json:"difference_subscribers" db:"difference_subscribers"`
	CreatedAt string `json:"created_at" db:"created_at"`
}