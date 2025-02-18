import React, { useState } from 'react';
import './ImageField.scss';

const ImageField = ({ label, limitImages, images, setImages }) => {
  const [newImageUrl, setNewImageUrl] = useState('');

  const handleAddImage = () => {
    if (limitImages && images.length >= limitImages) {
      alert(`You can only add up to ${limitImages} images.`);
      return;
    }

    if (newImageUrl.trim() && !images.includes(newImageUrl)) {
      setImages([...images, newImageUrl]);
      setNewImageUrl('');
    }
  };

  const handleRemoveImage = (index) => {
    setImages(images.filter((_, i) => i !== index));
  };

  return (
    <div className="image-field">
      <label>{label || 'Images (at least one required)'}</label>
      <div className="image-input">
        <input
          type="text"
          placeholder="Enter image URL"
          value={newImageUrl}
          onChange={(e) => setNewImageUrl(e.target.value)}
        />
        <button type="button" onClick={handleAddImage}>
          Add Image
        </button>
      </div>
      <div className="image-preview">
        {images.map((img, index) => (
          <div key={index} className="image-item">
            <img src={img} alt={`preview-${index}`} />
            <button type="button" onClick={() => handleRemoveImage(index)}>
              Remove
            </button>
          </div>
        ))}
      </div>
      {limitImages && (
        <div className="image-limit">
          {images.length}/{limitImages} images added
        </div>
      )}
    </div>
  );
};

export default ImageField;