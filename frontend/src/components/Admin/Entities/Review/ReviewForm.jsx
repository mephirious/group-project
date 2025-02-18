import React, { useState, useEffect } from 'react';
import { useDispatch } from 'react-redux';
import OperationTabs from '../../Common/OperationTabs';
import ConfirmModal from '../../Common/ConfirmModal';
import { updateReview, deleteReview } from '../../../../store/adminSlice';
import './ReviewForm.scss';
import FormField from '../../Common/FormFields/FormField';

const ReviewForm = ({ selectedData, onActionSuccess }) => {
  const dispatch = useDispatch();
  const [operation, setOperation] = useState('update');
  const [formData, setFormData] = useState({
    customer_id: '',
    product_id: '',
    content: '',
    rating: '',
    verified: false
  });
  const [showConfirm, setShowConfirm] = useState(false);
  const [pendingAction, setPendingAction] = useState(null);

  useEffect(() => {
    if (selectedData) {
      setOperation('update');
      setFormData({
        customer_id: selectedData.customer_id,
        product_id: selectedData.product_id,
        content: selectedData.content,
        rating: selectedData.rating,
        verified: selectedData.verified
      });
    } else {
      setFormData({
        customer_id: '',
        product_id: '',
        content: '',
        rating: '',
        verified: false
      });
      setOperation('update');
    }
  }, [selectedData]);

  const handleChange = (e) => {
    const { name, value, type, checked } = e.target;
    setFormData(prev => ({
      ...prev,
      [name]: type === 'checkbox' ? checked : value
    }));
  };

  const handleSubmit = (e) => {
    e.preventDefault();
    setShowConfirm(true);
    setPendingAction(() => () => {
      if (operation === 'update') {
        dispatch(updateReview({ id: selectedData.id, data: formData }));
      } else if (operation === 'delete') {
        dispatch(deleteReview(selectedData.id));
      }
    });
  };

  return (
    <div className="review-form">
      <div className="form-header">
        <h2>{operation === 'update' ? 'Update Review' : 'Delete Review'}</h2>
        <OperationTabs currentOperation={operation} onOperationChange={setOperation} />
      </div>
      <form onSubmit={handleSubmit}>
        {operation === 'update' && (
          <>
            <FormField label="Customer ID" name="customer_id" value={formData.customer_id} onChange={handleChange} required />
            <FormField label="Product ID" name="product_id" value={formData.product_id} onChange={handleChange} required />
            <div className="form-field">
              <label>Content</label>
              <textarea name="content" value={formData.content} onChange={handleChange} required />
            </div>
            <FormField label="Rating" name="rating" type="number" value={formData.rating} onChange={handleChange} required />
            <div className="form-field checkbox-field">
              <label>
                <input type="checkbox" name="verified" checked={formData.verified} onChange={handleChange} /> Verified
              </label>
            </div>
          </>
        )}
        {operation === 'delete' && (
          <div className="delete-warning">Are you sure you want to delete this review?</div>
        )}
        <button type="submit" className="submit-btn">
          {operation === 'update' ? 'Update Review' : 'Delete Review'}
        </button>
      </form>
      {showConfirm && (
        <ConfirmModal
          message={`Are you sure you want to ${operation} this review?`}
          onConfirm={() => {
            pendingAction();
            setShowConfirm(false);
            if (onActionSuccess) onActionSuccess();
          }}
          onCancel={() => setShowConfirm(false)}
        />
      )}
    </div>
  );
};

export default ReviewForm;
