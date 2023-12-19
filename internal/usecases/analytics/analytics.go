package analytics

import (
	"fmt"
	"log"

	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/model"
	"github.com/go-park-mail-ru/2023_2_Chaihona_No.1/internal/repositories/multistorage"
	"github.com/robfig/cron/v3"
)

func MakeCronAnalytics(multiStorage *multistorage.MultiStorage) {
	c := cron.New()
	_, err := c.AddFunc("* * * * *", func() {
		fmt.Println("job5")
		MakeAnalytics(multiStorage)
	})
	if err != nil {
		log.Println(err)
		return
	}
	c.Start()
}

func MakeAnalytics(multiStorage *multistorage.MultiStorage) {
	users, err := multiStorage.Users.GetUsers()
	if err != nil {
		log.Println(err)
	}
	for _, user := range users {
		if (user.Is_author) {
			lastAnalytics, err := multiStorage.Analytics.GetLastAnalytics(int(user.ID))
			if err != nil {
				fmt.Println("here")
				log.Println(err)
			}

			var newAnalytics model.Analitycs
			newAnalytics.Id = int(user.ID)
			newAnalytics.TotalPosts, err = multiStorage.Analytics.CountPosts(int(user.ID))
			if err != nil {
				log.Println(err)
			}
			newAnalytics.TotalLikes, err = multiStorage.Analytics.CountLikes(int(user.ID))
			if err != nil {
				log.Println(err)
			}
			newAnalytics.TotalComments, err = multiStorage.Analytics.CountComments(int(user.ID))
			if err != nil {
				log.Println(err)
			}
			newAnalytics.TotalDonations, 
			newAnalytics.TotalDonationsEarnedInteger, 
			newAnalytics.TotalDonationsEarnedFractional, 
			err = multiStorage.Analytics.CountDonations(int(user.ID))
			if err != nil {
				log.Println(err)
			}
			newAnalytics.TotalEarnedInteger, 
			newAnalytics.TotalEarnedFractional,
			err = multiStorage.Analytics.CountEarned(int(user.ID))
			if err != nil {
				log.Println(err)
			}
			newAnalytics.TotalSubscribers, err = multiStorage.Analytics.CountSubscribers(int(user.ID))
			if err != nil {
				log.Println(err)
			}


			newAnalytics.DifferencePosts = newAnalytics.TotalPosts - lastAnalytics.TotalPosts
			newAnalytics.DifferenceLikes = newAnalytics.TotalLikes - lastAnalytics.TotalLikes
			newAnalytics.DifferenceComments = newAnalytics.TotalComments - lastAnalytics.TotalComments
			newAnalytics.DifferenceDonations = newAnalytics.DifferenceDonations - lastAnalytics.DifferenceDonations
			newAnalytics.DifferenceDonationsEarnedInteger = newAnalytics.TotalDonationsEarnedInteger - lastAnalytics.TotalDonationsEarnedInteger
			newAnalytics.DifferenceDonationsEarnedFractional = newAnalytics.TotalDonationsEarnedFractional - lastAnalytics.TotalDonationsEarnedFractional
			if newAnalytics.DifferenceDonationsEarnedFractional < 0 {
				newAnalytics.DifferenceDonationsEarnedInteger -= 1
				newAnalytics.DifferenceDonationsEarnedFractional += 100
			}
			newAnalytics.DifferenceEarnedInteger = newAnalytics.TotalEarnedInteger - lastAnalytics.TotalEarnedInteger
			newAnalytics.DifferenceEarnedFractional = newAnalytics.TotalEarnedFractional - lastAnalytics.TotalEarnedFractional
			if newAnalytics.DifferenceEarnedFractional < 0 {
				newAnalytics.DifferenceEarnedInteger -= 1
				newAnalytics.DifferenceEarnedFractional += 100
			}
			newAnalytics.DifferenceSubscribers = newAnalytics.TotalSubscribers - lastAnalytics.TotalSubscribers

			_, err = multiStorage.Analytics.AddNewAnalytics(newAnalytics)
			if err != nil {
				log.Println(err)
				return
			}
		}
	}


}