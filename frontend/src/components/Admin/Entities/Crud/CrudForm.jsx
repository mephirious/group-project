import React, { useState, useEffect } from 'react';
import { useDispatch } from 'react-redux';
import OperationTabs from '../../Common/OperationTabs';
import ConfirmModal from '../../Common/ConfirmModal';
import FormField from '../../Common/FormFields/FormField';
import {
  createBrand,
  updateBrand,
  deleteBrand,
  createType,
  updateType,
  deleteType,
  createCategory,
  updateCategory,
  deleteCategory,
} from '../../../../store/adminSlice';
import './CrudForm.scss';

const entityConfig = {
  brands: {
    field: 'brand_name',
    label: 'Brand',
    actions: {
      create: createBrand,
      update: updateBrand,
      delete: deleteBrand,
    },
  },
  types: {
    field: 'type_name',
    label: 'Type',
    actions: {
      create: createType,
      update: updateType,
      delete: deleteType,
    },
  },
  categories: {
    field: 'category_name',
    label: 'Category',
    actions: {
      create: createCategory,
      update: updateCategory,
      delete: deleteCategory,
    },
  },
};

const CrudForm = ({ entity, selectedData, onActionSuccess }) => {
  const dispatch = useDispatch();
  const [operation, setOperation] = useState('create');
  const [formData, setFormData] = useState({});
  const [showConfirm, setShowConfirm] = useState(false);
  const [pendingAction, setPendingAction] = useState(null);
  const config = entityConfig[entity];

  useEffect(() => {
    if (selectedData) {
      setOperation('update');
      setFormData({ id: selectedData.id, [config.field]: selectedData[config.field] });
    } else {
      setFormData({});
      setOperation('create');
    }
  }, [selectedData, config]);

  const handleChange = e => {
    setFormData({ ...formData, [e.target.name]: e.target.value });
  };

  const handleSubmit = e => {
    e.preventDefault();
    setShowConfirm(true);
    setPendingAction(() => () => {
      if (operation === 'create') {
        dispatch(config.actions.create({ [config.field]: formData[config.field] }));
      } else if (operation === 'update') {
        dispatch(config.actions.update({ id: formData.id, data: { [config.field]: formData[config.field] } }));
      } else if (operation === 'delete') {
        dispatch(config.actions.delete(formData.id));
      }
      setFormData({});
    });
  };

  return (
    <div className="crud-form">
      <div className="form-header">
        <h2>
          {operation === 'create'
            ? `Create ${config.label}`
            : operation === 'update'
            ? `Update ${config.label}`
            : `Delete ${config.label}`}
        </h2>
        <OperationTabs currentOperation={operation} onOperationChange={setOperation} />
      </div>
      <form onSubmit={handleSubmit}>
        {operation !== 'delete' && (
          <>
            {operation === 'update' && (
              <FormField label="ID" name="id" value={formData.id || ''} onChange={handleChange} readOnly={!!selectedData} />
            )}
            <FormField
              label={`${config.label} Name`}
              name={config.field}
              value={formData[config.field] || ''}
              onChange={handleChange}
              required
            />
          </>
        )}
        {operation === 'delete' && (
          <FormField label="ID" name="id" value={formData.id || ''} onChange={handleChange} required />
        )}
        <button type="submit" className="submit-btn">
          {operation === 'create'
            ? `Create ${config.label}`
            : operation === 'update'
            ? `Update ${config.label}`
            : `Delete ${config.label}`}
        </button>
      </form>
      {showConfirm && (
        <ConfirmModal
          message={`Are you sure you want to ${operation} ${config.label}?`}
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

export default CrudForm;
