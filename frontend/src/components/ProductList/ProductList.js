import React from 'react';
import "./ProductList.scss";
import Product from "../Product/Product";
import BrandCard from "../BrandCard/BrandCard";

const ProductList = ({brand, products}) => {
  return (
    <div className='product-lists grid bg-whitesmoke my-3'>
      {brand && <BrandCard key = {brand.id} brand = {{...brand}} />}
      {
        products.map(product => {
          
    let discountedPrice = (product?.price) - (product?.price * (product?.discountPercentage ? product?.discountPercentage : 5) / 100);

          return (
            <Product key = {product.id} product = {{...product, discountedPrice}} />
          )
        })
      }
    </div>
  )
}

export default ProductList