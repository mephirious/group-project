import React from 'react';
import './ConfirmModal.scss';

const ConfirmModal = ({ message, onConfirm, onCancel }) => {
  return (
    <div className="confirm-modal-overlay" onClick={onCancel}>
      <div className="confirm-modal-content" onClick={(e) => e.stopPropagation()}>
        <p>{message}</p>
        <div className="confirm-modal-buttons">
          <button onClick={onConfirm} className='confirm-ok'>Confirm</button>
          <button onClick={onCancel} className='confirm-cancel'>Cancel</button>
        </div>
      </div>
    </div>
  );
};

export default ConfirmModal;
