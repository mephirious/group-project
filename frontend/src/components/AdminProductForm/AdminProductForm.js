import React, { useState, useEffect } from 'react';
import { useDispatch } from 'react-redux';
import { createProduct, updateProduct, deleteProduct } from '../../store/adminSlice';
import ConfirmModal from '../ConfirmModal/ConfirmModal';
import { BASE_URL } from '../../utils/apiURL';
import './AdminProductForm.scss';

const defaultSpecifications = [
  { key: 'cpu', value: '' },
  { key: 'cpu_cores', value: '' },
  { key: 'operating_system', value: '' },
  { key: 'screen_size', value: '' },
  { key: 'screen_refresh_rate', value: '' },
  { key: 'screen_brightness', value: '' },
  { key: 'screen_type', value: '' },
  { key: 'storage', value: '' },
  { key: 'ram', value: '' },
  { key: 'dimensions', value: '' },
  { key: 'weight', value: '' },
];

const AdminProductForm = ({ selectedProduct, onActionSuccess }) => {
  const dispatch = useDispatch();
  const [operation, setOperation] = useState('create');
  const [modelName, setModelName] = useState('');
  const [price, setPrice] = useState('');
  const [selectedCategory, setSelectedCategory] = useState({ id: '', name: '' });
  const [selectedBrand, setSelectedBrand] = useState({ id: '', name: '' });
  const [selectedType, setSelectedType] = useState({ id: '', name: '' });
  const [content, setContent] = useState('');
  const [images, setImages] = useState([]);
  const [newImageUrl, setNewImageUrl] = useState('');
  const [specifications, setSpecifications] = useState(defaultSpecifications);
  const [showConfirm, setShowConfirm] = useState(false);
  const [pendingAction, setPendingAction] = useState(null);

  // Dropdown options
  const [brands, setBrands] = useState([]);
  const [types, setTypes] = useState([]);
  const [categories, setCategories] = useState([]);

  useEffect(() => {
    fetch(`${BASE_URL}products/brands`, { credentials: 'include' })
      .then(res => res.json())
      .then(data => setBrands(data))
      .catch(err => console.error(err));
    fetch(`${BASE_URL}products/types`, { credentials: 'include' })
      .then(res => res.json())
      .then(data => setTypes(data))
      .catch(err => console.error(err));
    fetch(`${BASE_URL}products/categories`, { credentials: 'include' })
      .then(res => res.json())
      .then(data => setCategories(data))
      .catch(err => console.error(err));
  }, []);

  useEffect(() => {
    if (selectedProduct) {
      setOperation('update');
      setModelName(selectedProduct.model_name);
      setPrice(selectedProduct.price);
      const cat = categories.find(c => c.category_name === selectedProduct.category) || { id: '', category_name: selectedProduct.category };
      setSelectedCategory({ id: cat.id, name: cat.category_name });
      const br = brands.find(b => b.brand_name === selectedProduct.brand) || { id: '', brand_name: selectedProduct.brand };
      setSelectedBrand({ id: br.id, name: br.brand_name });
      const ty = types.find(t => t.type_name === selectedProduct.type) || { id: '', type_name: selectedProduct.type };
      setSelectedType({ id: ty.id, name: ty.type_name });
      setContent(selectedProduct.content);
      setImages(selectedProduct.images || []);
      const specsArray = Object.entries(selectedProduct.specifications || {}).map(([key, value]) => ({ key, value }));
      setSpecifications(specsArray);
    }
  }, [selectedProduct, brands, types, categories]);

  const handleAddImage = () => {
    if (newImageUrl.trim() !== '') {
      setImages([...images, newImageUrl.trim()]);
      setNewImageUrl('');
    }
  };

  const handleRemoveImage = (index) => {
    setImages(images.filter((_, i) => i !== index));
  };

  const handleSpecificationChange = (index, field, value) => {
    const newSpecs = [...specifications];
    newSpecs[index][field] = value;
    setSpecifications(newSpecs);
  };

  const handleAddSpecification = () => {
    setSpecifications([...specifications, { key: '', value: '' }]);
  };

  const handleRemoveSpecification = (index) => {
    setSpecifications(specifications.filter((_, i) => i !== index));
  };

  const handleSubmit = (e) => {
    e.preventDefault();
    if (images.length === 0) {
      alert('Please add at least one image.');
      return;
    }
    const specsObj = {};
    specifications.forEach(spec => {
      if (spec.key.trim() !== '') {
        specsObj[spec.key.trim()] = spec.value;
      }
    });
    const productData = {
      model_name: modelName,
      price: Number(price),
      category_id: selectedCategory.id,
      brand_id: selectedBrand.id,
      type_id: selectedType.id,
      specifications: specsObj,
      content,
      images,
    };
    setShowConfirm(true);
    setPendingAction(() => () => {
      if (operation === 'create') {
        dispatch(createProduct(productData));
      } else if (operation === 'update') {
        dispatch(updateProduct({ id: selectedProduct.id, data: productData }));
      } else if (operation === 'delete') {
        dispatch(deleteProduct(selectedProduct.id));
      }
      // Reset fields
      setModelName('');
      setPrice('');
      setSelectedCategory({ id: '', name: '' });
      setSelectedBrand({ id: '', name: '' });
      setSelectedType({ id: '', name: '' });
      setContent('');
      setImages([]);
      setSpecifications(defaultSpecifications);
    });
  };

  return (
    <div className="admin-product-form">
      <div className="form-header">
        <h2>{operation === 'create' ? 'Create Product' : 'Update Product'}</h2>
        <div className="operation-switch">
          <button className={operation === 'create' ? 'active' : ''} onClick={() => setOperation('create')}>
            Create
          </button>
          <button className={operation === 'update' ? 'active' : ''} onClick={() => setOperation('update')}>
            Update
          </button>
          {operation === 'update' && (
            <button className={operation === 'delete' ? 'active' : ''} onClick={() => setOperation('delete')}>
              Delete
            </button>
          )}
        </div>
      </div>
      <form onSubmit={handleSubmit}>
        <div className="form-group">
          <label>Model Name</label>
          <input type="text" value={modelName} onChange={(e) => setModelName(e.target.value)} required />
        </div>
        <div className="form-group">
          <label>Price</label>
          <input type="number" value={price} onChange={(e) => setPrice(e.target.value)} required />
        </div>
        <div className="form-group dropdown-group">
          <label>Brand</label>
          <select
            value={selectedBrand.id}
            onChange={(e) => {
              const selected = brands.find(b => b.id === e.target.value);
              setSelectedBrand({ id: selected.id, name: selected.brand_name });
            }}
            required
          >
            <option value="">Select Brand</option>
            {brands.map(b => (
              <option key={b.id} value={b.id}>{b.brand_name}</option>
            ))}
          </select>
        </div>
        <div className="form-group dropdown-group">
          <label>Type</label>
          <select
            value={selectedType.id}
            onChange={(e) => {
              const selected = types.find(t => t.id === e.target.value);
              setSelectedType({ id: selected.id, name: selected.type_name });
            }}
            required
          >
            <option value="">Select Type</option>
            {types.map(t => (
              <option key={t.id} value={t.id}>{t.type_name}</option>
            ))}
          </select>
        </div>
        <div className="form-group dropdown-group">
          <label>Category</label>
          <select
            value={selectedCategory.id}
            onChange={(e) => {
              const selected = categories.find(c => c.id === e.target.value);
              setSelectedCategory({ id: selected.id, name: selected.category_name });
            }}
            required
          >
            <option value="">Select Category</option>
            {categories.map(c => (
              <option key={c.id} value={c.id}>{c.category_name}</option>
            ))}
          </select>
        </div>
        <div className="form-group">
          <label>Content</label>
          <textarea value={content} onChange={(e) => setContent(e.target.value)} required />
        </div>
        <div className="images-section">
          <label>Images (at least one required)</label>
          <div className="image-input">
            <input
              type="text"
              placeholder="Enter image URL"
              value={newImageUrl}
              onChange={(e) => setNewImageUrl(e.target.value)}
            />
            <button type="button" onClick={handleAddImage}>Add Image</button>
          </div>
          <div className="image-preview">
            {images.map((img, index) => (
              <div key={index} className="image-item">
                <img src={img} alt={`preview-${index}`} />
                <button type="button" onClick={() => handleRemoveImage(index)}>Remove</button>
              </div>
            ))}
          </div>
        </div>
        <div className="specifications-section">
          <label>Specifications</label>
          {specifications.map((spec, index) => (
            <div key={index} className="specification-row">
              <input
                type="text"
                placeholder="Key"
                value={spec.key}
                onChange={(e) => handleSpecificationChange(index, 'key', e.target.value)}
                required
              />
              <input
                type="text"
                placeholder="Value"
                value={spec.value}
                onChange={(e) => handleSpecificationChange(index, 'value', e.target.value)}
                required
              />
              <button type="button" onClick={() => handleRemoveSpecification(index)}>Delete</button>
            </div>
          ))}
          <button type="button" onClick={handleAddSpecification}>Add Specification</button>
        </div>
        <button type="submit" className="submit-btn">
          {operation === 'create'
            ? 'Create Product'
            : operation === 'update'
            ? 'Update Product'
            : 'Delete Product'}
        </button>
      </form>
      {showConfirm && (
        <ConfirmModal
          message={`Are you sure you want to ${operation} this product?`}
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

export default AdminProductForm;
