package multistorage

import (
	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/repositories/analytics"
	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/repositories/attaches"
	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/repositories/comments"
	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/repositories/likes"
	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/repositories/payments"
	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/repositories/posts"
	subscriptionlevels "github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/repositories/subscription_levels"
	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/repositories/subscriptions"
	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/repositories/users"
)

type MultiStorage struct {
	Attaches attaches.AttachRepository
	Comments *comments.CommentManager
	Likes likes.LikeRepository
	Payments payments.PaymentRepository
	Posts posts.PostRepository
	SubscriptionLevels subscriptionlevels.SubscribeLevelRepository
	Subscriptions subscriptions.SubscriptionRepository
	Users users.UserRepository
	Analytics analytics.AnalyticsRepository
}