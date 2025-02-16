import React, { useState, useEffect } from 'react';
import { useDispatch } from 'react-redux';
import {
  createCategory,
  updateCategory,
  deleteCategory,
  createBrand,
  updateBrand,
  deleteBrand,
  createType,
  updateType,
  deleteType,
} from '../../store/adminSlice';
import ConfirmModal from '../ConfirmModal/ConfirmModal';
import './AdminCrudForm.scss';

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

const AdminCrudForm = ({ entity, selectedData, onActionSuccess }) => {
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
    }
  }, [selectedData, config]);

  const handleChange = (e) => {
    setFormData({ ...formData, [e.target.name]: e.target.value });
  };

  const handleSubmit = (e) => {
    e.preventDefault();
    setShowConfirm(true);
    setPendingAction(() => () => {
      if (operation === 'create') {
        dispatch(config.actions.create({ [config.field]: formData[config.field] }));
      } else if (operation === 'update') {
        dispatch(
          config.actions.update({
            id: formData.id,
            data: { [config.field]: formData[config.field] },
          })
        );
      } else if (operation === 'delete') {
        dispatch(config.actions.delete(formData.id));
      }
      setFormData({});
    });
  };

  const renderFormFields = () => {
    if (operation === 'create') {
      return (
        <div className="form-group">
          <label>{config.label} Name</label>
          <input
            type="text"
            name={config.field}
            value={formData[config.field] || ''}
            onChange={handleChange}
          />
        </div>
      );
    } else if (operation === 'update') {
      return (
        <>
          <div className="form-group">
            <label>ID</label>
            <input
              type="text"
              name="id"
              value={formData.id || ''}
              onChange={handleChange}
              readOnly={!!selectedData}
            />
          </div>
          <div className="form-group">
            <label>New {config.label} Name</label>
            <input
              type="text"
              name={config.field}
              value={formData[config.field] || ''}
              onChange={handleChange}
            />
          </div>
        </>
      );
    } else if (operation === 'delete') {
      return (
        <div className="form-group">
          <label>ID</label>
          <input
            type="text"
            name="id"
            value={formData.id || ''}
            onChange={handleChange}
          />
        </div>
      );
    }
  };

  return (
    <div className="admin-crud-form">
      <div className="operation-tabs">
        <button
          className={operation === 'create' ? 'active' : ''}
          onClick={() => setOperation('create')}
        >
          Create
        </button>
        <button
          className={operation === 'update' ? 'active' : ''}
          onClick={() => setOperation('update')}
        >
          Update
        </button>
        <button
          className={operation === 'delete' ? 'active' : ''}
          onClick={() => setOperation('delete')}
        >
          Delete
        </button>
      </div>
      <form onSubmit={handleSubmit}>
        {renderFormFields()}
        <button type="submit" className="submit-btn">
          Submit
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

export default AdminCrudForm;
