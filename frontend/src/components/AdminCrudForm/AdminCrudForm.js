import React, { useState } from 'react';
import { useDispatch } from 'react-redux';
import {
  createBrand,
  updateBrand,
  deleteBrand,
  createType,
  updateType,
  deleteType
} from '../../store/adminSlice';
import './AdminCrudForm.scss';
import ConfirmModal from '../ConfirmModal/ConfirmModal';

const AdminCrudForm = ({ entity }) => {
  const dispatch = useDispatch();
  const [operation, setOperation] = useState('create');
  const [formData, setFormData] = useState({});
  const [showConfirm, setShowConfirm] = useState(false);
  const [pendingAction, setPendingAction] = useState(null);

  const handleChange = (e) => {
    setFormData({ ...formData, [e.target.name]: e.target.value });
  };

  const handleSubmit = (e) => {
    e.preventDefault();
    setShowConfirm(true);
    setPendingAction(() => () => {
      if (entity === 'brands') {
        if (operation === 'create') {
          dispatch(createBrand({ brand_name: formData.brand_name }));
        } else if (operation === 'update') {
          dispatch(updateBrand({ id: formData.id, brandData: { brand_name: formData.brand_name } }));
        } else if (operation === 'delete') {
          dispatch(deleteBrand(formData.id));
        }
      } else if (entity === 'types' || entity === 'categories') {
        if (operation === 'create') {
          dispatch(createType({ type_name: formData.type_name }));
        } else if (operation === 'update') {
          dispatch(updateType({ id: formData.id, typeData: { type_name: formData.type_name } }));
        } else if (operation === 'delete') {
          dispatch(deleteType(formData.id));
        }
      }
      setFormData({});
    });
  };

  const renderFormFields = () => {
    if (operation === 'create') {
      return (
        <div className="form-group">
          {entity === 'brands' ? (
            <>
              <label>Brand Name</label>
              <input type="text" name="brand_name" value={formData.brand_name || ''} onChange={handleChange} />
            </>
          ) : (
            <>
              <label>{entity === 'types' ? 'Type Name' : 'Category Name'}</label>
              <input type="text" name="type_name" value={formData.type_name || ''} onChange={handleChange} />
            </>
          )}
        </div>
      );
    } else if (operation === 'update') {
      return (
        <>
          <div className="form-group">
            <label>ID</label>
            <input type="text" name="id" value={formData.id || ''} onChange={handleChange} />
          </div>
          <div className="form-group">
            {entity === 'brands' ? (
              <>
                <label>New Brand Name</label>
                <input type="text" name="brand_name" value={formData.brand_name || ''} onChange={handleChange} />
              </>
            ) : (
              <>
                <label>{entity === 'types' ? 'New Type Name' : 'New Category Name'}</label>
                <input type="text" name="type_name" value={formData.type_name || ''} onChange={handleChange} />
              </>
            )}
          </div>
        </>
      );
    } else if (operation === 'delete') {
      return (
        <div className="form-group">
          <label>ID</label>
          <input type="text" name="id" value={formData.id || ''} onChange={handleChange} />
        </div>
      );
    }
  };

  return (
    <div className="admin-crud-form">
      <div className="operation-tabs">
        <button className={operation === 'create' ? 'active' : ''} onClick={() => setOperation('create')}>
          Create
        </button>
        <button className={operation === 'update' ? 'active' : ''} onClick={() => setOperation('update')}>
          Update
        </button>
        <button className={operation === 'delete' ? 'active' : ''} onClick={() => setOperation('delete')}>
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
          message={`Are you sure you want to ${operation} ${entity}?`}
          onConfirm={() => {
            pendingAction();
            setShowConfirm(false);
          }}
          onCancel={() => setShowConfirm(false)}
        />
      )}
    </div>
  );
};

export default AdminCrudForm;
