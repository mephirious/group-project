// Reviews.jsx
import React, { useState, useEffect } from 'react';
import { useSelector, useDispatch } from 'react-redux';
import { getAllProductsByReview, getReviewProductsStatus, createAsyncReview, fetchAsyncReviewsOfProduct } from '../../store/reviewSlice';
import Loader from '../Loader/Loader';
import './Reviews.scss';

const Reviews = ({ productId }) => {
  const dispatch = useDispatch();
  const reviewsData = useSelector(getAllProductsByReview);
  const reviewStatus = useSelector(getReviewProductsStatus);
  const [showReviewForm, setShowReviewForm] = useState(false);
  const [reviewContent, setReviewContent] = useState("");
  const [reviewRating, setReviewRating] = useState(5);
  const [isSubmittingReview, setIsSubmittingReview] = useState(false);

  useEffect(() => {
    if (productId) {
      dispatch(fetchAsyncReviewsOfProduct(productId));
    }
  }, [productId, dispatch]);

  const submitReview = async (e) => {
    e.preventDefault();
    setIsSubmittingReview(true);
    const user = { user_id: "dummy-user-id" };
    if (!user || !productId) return;
    const reviewData = {
      customer_id: user.user_id,
      product_id: productId,
      content: reviewContent,
      rating: Number(reviewRating),
    };
    try {
      await dispatch(createAsyncReview(reviewData)).unwrap();
      dispatch(fetchAsyncReviewsOfProduct(productId));
      setReviewContent("");
      setReviewRating(5);
      setShowReviewForm(false);
    } catch (error) {
      console.error("Error creating review:", error);
    } finally {
      setIsSubmittingReview(false);
    }
  };

  return (
    <div className="reviews">
      <h3>Customer Reviews</h3>
      <div className="average-rating">
        <span className="rating-label">Average Rating:</span>
        <span className="rating-value">
          {reviewsData?.average_rating ? reviewsData.average_rating.toFixed(1) : 'N/A'}
        </span>
      </div>

      {!showReviewForm ? (
        <div className="btns">
          <button
            type="button"
            className="btn"
            onClick={() => setShowReviewForm(true)}
          >
            Write a Review
          </button>
        </div>
      ) : (
        <form className="review-form" onSubmit={submitReview}>
          <h4>Submit Your Review</h4>
          <div className="form-group">
            <label>Rating (1-10):</label>
            <input
              type="number"
              min="1"
              max="10"
              value={reviewRating}
              onChange={(e) => setReviewRating(e.target.value)}
              required
            />
          </div>
          <div className="form-group">
            <label>Review:</label>
            <textarea
              value={reviewContent}
              onChange={(e) => setReviewContent(e.target.value)}
              required
            />
          </div>
          <div className="btns">
            <button
              type="button"
              className="btn"
              onClick={() => setShowReviewForm(false)}
            >
              Cancel
            </button>
            <button
              type="submit"
              className="btn"
              disabled={isSubmittingReview}
            >
              {isSubmittingReview ? "Submitting..." : "Submit Review"}
            </button>
          </div>
        </form>
      )}

      {reviewStatus === 'loading' ? (
        <Loader />
      ) : (
        <div className="reviews-list">
          {reviewsData?.reviews && reviewsData.reviews.length > 0 ? (
            reviewsData.reviews.map((review, idx) => (
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
      )}
    </div>
  );
};

export default Reviews;