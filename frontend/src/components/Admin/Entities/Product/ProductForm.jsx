import React, { useState, useEffect } from 'react';
import { useDispatch } from 'react-redux';
import OperationTabs from '../../Common/OperationTabs';
import ConfirmModal from '../../Common/ConfirmModal';
import FormField from '../../Common/FormFields/FormField';
import { createProduct, updateProduct, deleteProduct } from '../../../../store/adminSlice';
import { BASE_URL } from '../../../../utils/apiURL';
import './ProductForm.scss';
import ImageField from '../../Common/FormFields/ImageField';

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

const ProductForm = ({ selectedData, onActionSuccess }) => {
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
    if (selectedData) {
      setOperation('update');
      setModelName(selectedData.model_name);
      setPrice(selectedData.price);
      const cat = categories.find(c => c.category_name === selectedData.category) || { id: '', category_name: selectedData.category };
      setSelectedCategory({ id: cat.id, name: cat.category_name });
      const br = brands.find(b => b.brand_name === selectedData.brand) || { id: '', brand_name: selectedData.brand };
      setSelectedBrand({ id: br.id, name: br.brand_name });
      const ty = types.find(t => t.type_name === selectedData.type) || { id: '', type_name: selectedData.type };
      setSelectedType({ id: ty.id, name: ty.type_name });
      setContent(selectedData.content);
      setImages(selectedData.images || []);
      const specsArray = Object.entries(selectedData.specifications || {}).map(([key, value]) => ({ key, value }));
      setSpecifications(specsArray);
    } else {
      setOperation('create');
      resetFields();
    }
  }, [selectedData, brands, types, categories]);

  const resetFields = () => {
    setModelName('');
    setPrice('');
    setSelectedCategory({ id: '', name: '' });
    setSelectedBrand({ id: '', name: '' });
    setSelectedType({ id: '', name: '' });
    setContent('');
    setImages([]);
    setSpecifications(defaultSpecifications);
  };

  const handleSpecificationChange = (index, field, value) => {
    const newSpecs = [...specifications];
    newSpecs[index][field] = value;
    setSpecifications(newSpecs);
  };

  const handleAddSpecification = () => {
    setSpecifications([...specifications, { key: '', value: '' }]);
  };

  const handleRemoveSpecification = index => {
    setSpecifications(specifications.filter((_, i) => i !== index));
  };

  const handleSubmit = e => {
    e.preventDefault();
    if (operation !== 'delete' && images.length === 0) {
      alert('Please add at least one image.');
      return;
    }
    const specsObj = {};
    specifications.forEach(spec => {
      if (spec.key.trim()) specsObj[spec.key.trim()] = spec.value;
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
        dispatch(updateProduct({ id: selectedData.id, data: productData }));
      } else if (operation === 'delete') {
        dispatch(deleteProduct(selectedData.id));
      }
      resetFields();
    });
  };

  return (
    <div className="product-form">
      <div className="form-header">
        <h2>
          {operation === 'create'
            ? 'Create Product'
            : operation === 'update'
            ? 'Update Product'
            : 'Delete Product'}
        </h2>
        <OperationTabs currentOperation={operation} onOperationChange={setOperation} />
      </div>
      <form onSubmit={handleSubmit}>
        {operation !== 'delete' && (
          <>
            <FormField label="Model Name" name="modelName" value={modelName} onChange={e => setModelName(e.target.value)} required />
            <FormField label="Price" name="price" type="number" value={price} onChange={e => setPrice(e.target.value)} required />
            <div className="dropdowns">
              <FormField
                label="Brand"
                type="select"
                name="brand"
                value={selectedBrand.id}
                onChange={e => {
                  const selected = brands.find(b => b.id === e.target.value);
                  setSelectedBrand({ id: selected.id, name: selected.brand_name });
                }}
                options={brands.map(b => ({ id: b.id, name: b.brand_name }))}
                required
              />
              <FormField
                label="Type"
                type="select"
                name="type"
                value={selectedType.id}
                onChange={e => {
                  const selected = types.find(t => t.id === e.target.value);
                  setSelectedType({ id: selected.id, name: selected.type_name });
                }}
                options={types.map(t => ({ id: t.id, name: t.type_name }))}
                required
              />
              <FormField
                label="Category"
                type="select"
                name="category"
                value={selectedCategory.id}
                onChange={e => {
                  const selected = categories.find(c => c.id === e.target.value);
                  setSelectedCategory({ id: selected.id, name: selected.category_name });
                }}
                options={categories.map(c => ({ id: c.id, name: c.category_name }))}
                required
              />
            </div>
            <div className="form-field">
              <label>Content</label>
              <textarea value={content} onChange={e => setContent(e.target.value)} required />
            </div>
            <ImageField label='Images (at least one required)' images={images} setImages={setImages}></ImageField>
            <div className="specifications-section">
              <label>Specifications</label>
              {specifications.map((spec, index) => (
                <div key={index} className="specification-row">
                  <input
                    type="text"
                    placeholder="Key"
                    value={spec.key}
                    onChange={e => handleSpecificationChange(index, 'key', e.target.value)}
                    required
                  />
                  <input
                    type="text"
                    placeholder="Value"
                    value={spec.value}
                    onChange={e => handleSpecificationChange(index, 'value', e.target.value)}
                    required
                  />
                  <button type="button" onClick={() => handleRemoveSpecification(index)}>Delete</button>
                </div>
              ))}
              <button type="button" onClick={handleAddSpecification}>Add Specification</button>
            </div>
          </>
        )}
        {operation === 'delete' && (
          <div className="delete-warning">
            <p>Are you sure you want to delete this product?</p>
          </div>
        )}
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

export default ProductForm;
