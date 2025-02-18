import React from 'react';
import './OperationTabs.scss';

const OperationTabs = ({ currentOperation, onOperationChange }) => {
  const operations = ['create', 'update', 'delete'];
  return (
    <div className="operation-tabs">
      {operations.map(op => (
        <button
          key={op}
          className={currentOperation === op ? 'active' : ''}
          onClick={() => onOperationChange(op)}
        >
          {op.charAt(0).toUpperCase() + op.slice(1)}
        </button>
      ))}
    </div>
  );
};

export default OperationTabs;
