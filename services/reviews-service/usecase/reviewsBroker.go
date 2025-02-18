package usecase

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

var ratingsData map[primitive.ObjectID]float64 = make(map[primitive.ObjectID]float64)

func StartReviewRatingsBroker(sleep time.Duration, u *reviewUseCase) {
	verified := true
	for {
		time.Sleep(sleep)

		// Update Data
		for productID := range ratingsData {
			// Re-calculate all ratings for each product
			stats, err := u.reviewRepository.GetReviewStatsByProductID(context.Background(), productID, &verified)
			if err != nil {
				continue
			}
			ratingsData[productID] = stats
		}
	}
}
