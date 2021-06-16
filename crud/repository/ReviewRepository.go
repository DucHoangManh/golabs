package repository

import (
	"errors"

	"github.com/DucHoangManh/golabs/crud/model"
)
type Observer interface {
	Update(interface{})
}
type ReviewRepo struct {
	Observers []Observer
	reviews  map[int64]*model.Review
	autoID int64
}
func (reviewRepo *ReviewRepo) Subscribe(ob Observer){
	reviewRepo.Observers = append(reviewRepo.Observers,ob)
}
func (reviewRepo *ReviewRepo) Notify(review *model.Review){
	for _, ob := range reviewRepo.Observers{
		ob.Update(review)
	}
}
var Reviews ReviewRepo

func init() {
	Reviews = ReviewRepo{autoID: 0}
	Reviews.reviews = make(map[int64]*model.Review)
}

func (r *ReviewRepo) getAutoID() int64 {
	r.autoID += 1
	return r.autoID
}
func (reviewRepo *ReviewRepo) CreateNewReview(review *model.Review) int64 {
	nextID := reviewRepo.getAutoID()
	review.Id = nextID
	reviewRepo.reviews[nextID] = review
	reviewRepo.Notify(review)
	return nextID
}
func (reviewRepo *ReviewRepo) GetAllReviews() map[int64]*model.Review{
	return reviewRepo.reviews
}
func (reviewRepo *ReviewRepo) DeleteReviewById(Id int64) error {
	if review, ok := reviewRepo.reviews[Id]; ok {
		delete(reviewRepo.reviews, Id)
		reviewRepo.Notify(review)
		return nil
	} else {
		return errors.New("book not found")
	}
}