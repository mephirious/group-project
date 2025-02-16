import React from 'react';
import { Link } from 'react-router-dom';
import { formatPrice } from '../../utils/helpers';
import "./LaptopCard.scss";

const DetailRow = ({ label, value }) => (
    <div className="detail-row">
      <span className="detail-label">{label}:</span>
    <span className="detail-value" title={value}>
      {value}
    </span>
  </div>
);

const LaptopCard = ({ laptop, allSpecKeys, onRemove }) => (
  <div className="laptop-card bg-white">
    {laptop.images && laptop.images.length > 0 && (
      <div className="laptop-image">
        <img src={laptop.images[0]} alt={laptop.model_name} />
      </div>
    )}
    <div className="laptop-details">
      <DetailRow label="Model" value={laptop.model_name} />
      <DetailRow label="Price" value={formatPrice(laptop.price)} />
      <DetailRow label="Brand" value={laptop.brand} />
      <DetailRow label="Category" value={laptop.category} />
      <DetailRow label="Type" value={laptop.type} />
      
      {allSpecKeys.map(key => {
        const value = (laptop.specifications && laptop.specifications[key]) || 'N/A';
        const label = key.charAt(0).toUpperCase() + key.slice(1);
        return (
          <DetailRow
            key={key}
            label={label}
            value={value}
            title={value}
          />
        );
      })}
    </div>
    <div className="laptop-actions">
      <Link to={`/product/${laptop.id}`} className="view-details-btn">
        View Details
      </Link>
      <button
        type="button"
        className="delete-btn text-dark"
        onClick={() => onRemove(laptop.id)}
      >
        Remove
      </button>
    </div>
  </div>
);

export default LaptopCard;