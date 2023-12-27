package analytics

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/v2/dbscan"
	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/configs"
	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/model"
)

func SelectAnalyticsMonthAgoSQL(userId int, date string) squirrel.SelectBuilder {
	return squirrel.Select("*").
		From(configs.AnalitycsTable + " a").
		Where(squirrel.Eq{"user_id": userId}, "created_at >= " + date, "created_at < " + date + " interval '1 day'").
		GroupBy("a.user_id, a.id").
		OrderBy("a.created_at DESC").
		Limit(1).
		PlaceholderFormat(squirrel.Dollar)	
}

func SelectLastAnalyticsSQL(userId int) squirrel.SelectBuilder {
	return squirrel.Select("*").
		From(configs.AnalitycsTable + " a").
		Where(squirrel.Eq{"user_id": userId}).
		GroupBy("a.user_id, a.id").
		OrderBy("a.created_at DESC").
		Limit(1).
		PlaceholderFormat(squirrel.Dollar)	
}

func SelectFirstAnalyticsSQL(userId int) squirrel.SelectBuilder {
	return squirrel.Select("*").
		From(configs.AnalitycsTable + " a").
		Where(squirrel.Eq{"user_id": userId}).
		GroupBy("a.user_id, a.id").
		OrderBy("a.created_at ASC").
		Limit(1).
		PlaceholderFormat(squirrel.Dollar)	
}

func CountPostsSQL(userId int) squirrel.SelectBuilder {
	return squirrel.Select("COUNT(*)").
		From(configs.PostTable).
		Where(squirrel.Eq{"creator_id": userId}).
		PlaceholderFormat(squirrel.Dollar)
}

func CountLikesSQL(userId int) squirrel.SelectBuilder {
	return squirrel.Select("COUNT(*)").
		From(configs.LikeTable + " l").
		InnerJoin(fmt.Sprintf("%s p ON p.id = l.post_id", configs.PostTable)).
		Where(squirrel.Eq{"p.creator_id": userId}).
		PlaceholderFormat(squirrel.Dollar)
}

func CountCommentsSQL(userId int) squirrel.SelectBuilder {
	return squirrel.Select("COUNT(*)").
		From(configs.CommentTable + " c").
		InnerJoin(fmt.Sprintf("%s p ON p.id = c.post_id", configs.PostTable)).
		Where(squirrel.Eq{"p.creator_id": userId}).
		PlaceholderFormat(squirrel.Dollar)
}

func CountDonationsSQL(userId int) squirrel.SelectBuilder {
	return squirrel.Select("COUNT(*) as total_donations, coalesce(SUM(p.payment_integer), 0) as total_donations_earned_integer, coalesce(SUM(p.payment_fractional), 0) as total_donations_earned_fractional").
		From(configs.PaymentTable + " p").
		Where(squirrel.And{
			squirrel.Eq{"p.creator_id": userId}, 
			squirrel.Eq{"p.status": model.PaymentSucceededStatus},
			squirrel.Eq{"p.payment_type": model.PaymentTypeDonate},
		}).
		PlaceholderFormat(squirrel.Dollar)
}

func CountEarnedSQL(userId int) squirrel.SelectBuilder {
	return squirrel.Select("coalesce(SUM(p.payment_integer), 0) as total_donations_earned_integer, coalesce(SUM(p.payment_fractional), 0) as total_donations_earned_fractional").
		From(configs.PaymentTable + " p").
		Where(squirrel.And{
			squirrel.Eq{"p.creator_id": userId}, 
			squirrel.Eq{"p.status": model.PaymentSucceededStatus},
		}).
		PlaceholderFormat(squirrel.Dollar)
}

func CountSubscribersSQL(userId int) squirrel.SelectBuilder {
	return squirrel.Select("COUNT(*)").
		From(configs.SubscriptionTable + " s").
		Where(squirrel.Eq{"s.creator_id": userId}).
		PlaceholderFormat(squirrel.Dollar)
}

func InsertAnalyticsSQL(analytics model.Analitycs) squirrel.InsertBuilder {
	return squirrel.Insert(configs.AnalitycsTable).
		Columns("user_id", "total_posts", "total_likes", "total_comments",
			"total_donations",  "total_donations_earned_integer", "total_donations_earned_fractional",
			"total_earned_integer", "total_earned_fractional", "total_subscribers",
			"difference_posts", "difference_likes", "difference_comments", 
			"difference_donations", "difference_donations_earned_integer", 
			"difference_donations_earned_fractional", "difference_earned_integer",
			"difference_earned_fractional", "difference_subscribers",
		).
		Values(analytics.UserId, analytics.TotalPosts, analytics.TotalLikes,
			analytics.TotalComments, analytics.TotalDonations, analytics.TotalDonationsEarnedInteger,
			analytics.TotalDonationsEarnedFractional, analytics.TotalEarnedInteger,
			analytics.TotalEarnedFractional, analytics.TotalSubscribers,
			analytics.DifferencePosts, analytics.DifferenceLikes, analytics.DifferenceComments,
			analytics.DifferenceDonations, analytics.DifferenceDonationsEarnedInteger,
			analytics.DifferenceDonationsEarnedFractional, analytics.DifferenceEarnedInteger,
			analytics.DifferenceEarnedFractional, analytics.DifferenceSubscribers).
		Suffix("RETURNING \"id\"").
		PlaceholderFormat(squirrel.Dollar)
}

type AnalyticsStorage struct {
	db *sql.DB
}

func CreateAnalyticsStorage(db *sql.DB) AnalyticsRepository {
	return &AnalyticsStorage{
		db: db,
	}
}

func (storage *AnalyticsStorage) GetMounthAgoAnalytics(userId int) (model.Analitycs, error) {
	now := time.Now()
	monthAgo := now.AddDate(0, -1, 0)
	var analyticsMonth []model.Analitycs	
	rowsMonth, err := SelectAnalyticsMonthAgoSQL(userId, monthAgo.Format("2006-01-02")).RunWith(storage.db).Query()
	
	if err != nil {
		log.Println(err)
		return model.Analitycs{}, err
	}
	err = dbscan.ScanAll(&analyticsMonth, rowsMonth)
	fmt.Println(analyticsMonth)
	if err != nil {
		log.Println(err)
		return model.Analitycs{}, err
	}
	if len(analyticsMonth) == 0 {
		lastRows, err := SelectFirstAnalyticsSQL(userId).RunWith(storage.db).Query()
		if err != nil {
			log.Println(err)
			return model.Analitycs{}, err
		}
		var analyticsLast []model.Analitycs	
		err = dbscan.ScanAll(&analyticsLast, lastRows)
		if err != nil || len(analyticsLast) == 0 {
			log.Println(err)
			return model.Analitycs{}, err
		}
		fmt.Println(analyticsLast)
		return analyticsLast[0], nil
	}
	return analyticsMonth[0], nil
}

func (storage *AnalyticsStorage) GetLastAnalytics(userId int) (model.Analitycs, error) {
	lastRows, err := SelectLastAnalyticsSQL(userId).RunWith(storage.db).Query()
	if err != nil {
		log.Println(err)
		return model.Analitycs{}, err
	}
	var analyticsLast []model.Analitycs	
	err = dbscan.ScanAll(&analyticsLast, lastRows)
	if err != nil || len(analyticsLast) == 0 {
		log.Println(err)
		return model.Analitycs{}, err
	}
	return analyticsLast[0], nil
}

func (storage *AnalyticsStorage) CountPosts(userId int) (int, error) {
	var countPosts int
	err := CountPostsSQL(userId).RunWith(storage.db).QueryRow().Scan(&countPosts)
	if err != nil {
		return 0, err
	}
	return countPosts, nil
}

func (storage *AnalyticsStorage) CountLikes(userId int) (int, error) {
	var countLikes int
	err := CountLikesSQL(userId).RunWith(storage.db).QueryRow().Scan(&countLikes)
	if err != nil {
		return 0, err
	}
	return countLikes, nil
}

func (storage *AnalyticsStorage) CountComments(userId int) (int, error) {
	var countComments int
	err := CountCommentsSQL(userId).RunWith(storage.db).QueryRow().Scan(&countComments)
	if err != nil {
		return 0, err
	}
	return countComments, nil
}

func (storage *AnalyticsStorage) CountDonations(userId int) (int, int, int, error) {
	var totalDonations, DonationsInteger, DonationsFractional int
	err := CountDonationsSQL(userId).RunWith(storage.db).QueryRow().Scan(&totalDonations, &DonationsInteger, &DonationsFractional)
	if err != nil {
		return 0, 0, 0, err
	}
	return  totalDonations, DonationsInteger, DonationsFractional, nil
}

func (storage *AnalyticsStorage) CountEarned(userId int) (int, int, error) {
	var DonationsInteger, DonationsFractional int
	err := CountEarnedSQL(userId).RunWith(storage.db).QueryRow().Scan(&DonationsInteger, &DonationsFractional)
	if err != nil {
		return 0, 0, err
	}
	return DonationsInteger, DonationsFractional, nil
}

func (storage *AnalyticsStorage) CountSubscribers(userId int) (int, error) {
	var countSubscribers int
	err := CountSubscribersSQL(userId).RunWith(storage.db).QueryRow().Scan(&countSubscribers)
	if err != nil {
		return 0, err
	}
	return countSubscribers, nil
}

func (storage *AnalyticsStorage) AddNewAnalytics(analytics model.Analitycs) (int, error) {
	var analyticsId int
	err := InsertAnalyticsSQL(analytics).RunWith(storage.db).QueryRow().Scan(&analyticsId)
	if err != nil {
		return 0, err
	}
	return analyticsId, nil
}