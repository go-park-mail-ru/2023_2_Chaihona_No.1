package analytics

import "github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/model"

type AnalyticsRepository interface {
	GetLastAnalytics(userId int) (model.Analitycs, error)
	GetMounthAgoAnalytics(userId int) (model.Analitycs, error)
	CountPosts(userId int) (int, error)
	CountLikes(userId int) (int, error)
	CountComments(userId int) (int, error)
	CountDonations(userId int) (int, int, int, error)
	CountEarned(userId int) (int, int, error)
	CountSubscribers(userId int) (int, error)
	AddNewAnalytics(analytics model.Analitycs) (int, error)
}