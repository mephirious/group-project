import React from 'react';
import './ReviewListItem.scss';

const ReviewListItem = ({ item, onSelect }) => {
  return (
    <div className="review-list-item" onClick={() => onSelect(item)}>
      <div className="content">{item.content}</div>
      <div className="details">
        <span className="rating">Rating: {item.rating}</span>
        <span className="verified">{item.verified ? 'Verified' : 'Not Verified'}</span>
      </div>
      <div className="meta">
        <span className="customer">Customer: {item.customer_id}</span>
        <span className="product">Product: {item.product_id}</span>
        <span className="created-at">{new Date(item.created_at).toLocaleString()}</span>
      </div>
    </div>
  );
};

export default ReviewListItem;
