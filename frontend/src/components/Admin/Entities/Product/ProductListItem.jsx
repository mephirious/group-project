import React from "react";
import "./ProductListItem.scss";

const ProductListItem = ({ item, onSelect }) => {
  const imageUrl = item?.images?.[0] || "";
  const price = item?.price ?? 0;

  return (
    <div className="product-list-item" onClick={() => onSelect(item)}>
      {imageUrl && (
        <img src={imageUrl} alt={item.model_name} className="product-image" />
      )}
      <div className="product-details">
        <div className="model-name">{item?.model_name || "Unknown Model"}</div>
        <div className="price">{price.toLocaleString()} KZT</div>
        <div className="brand-category">{item?.brand} • {item?.category}</div>
        <div className="created-at">{new Date(item.created_at).toLocaleString()}</div>
      </div>
    </div>
  );
};

export default ProductListItem;
