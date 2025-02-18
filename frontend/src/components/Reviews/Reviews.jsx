import React from 'react';
import { useSelector } from 'react-redux';
import { getAllProductsByReview, getReviewProductsStatus } from '../../store/reviewSlice';
import Loader from '../Loader/Loader';
import './Reviews.scss';

const Reviews = () => {
  const reviewsData = useSelector(getAllProductsByReview);
  const reviewStatus = useSelector(getReviewProductsStatus);

  if (reviewStatus === 'loading') {
    return <Loader />;
  }

  // Assuming reviewsData has { average_rating, reviews }.
  const { average_rating, reviews } = reviewsData || {};

  return (
    <div className="reviews">
      <h3>Customer Reviews</h3>
      <div className="average-rating">
        <span className="rating-label">Average Rating:</span>
        <span className="rating-value">{average_rating ? average_rating.toFixed(1) : 'N/A'}</span>
      </div>
      <div className="reviews-list">
        {reviews && reviews.length > 0 ? (
          reviews.map((review, idx) => (
            <div className="review-item" key={idx}>
              <div className="review-header">
                <span className="reviewer">{review.reviewer || 'Anonymous'}</span>
                <span className="review-date">{new Date(review.created_at).toLocaleDateString()}</span>
              </div>
              <div className="review-rating">Rating: {review.rating}</div>
              <div className="review-comment">{review.content}</div>
            </div>
          ))
        ) : (
          <p className="no-reviews">No reviews available.</p>
        )}
      </div>
    </div>
  );
};

export default Reviews;
