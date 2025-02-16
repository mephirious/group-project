import React, { useEffect, useState } from 'react';
import "./ProductSinglePage.scss";
import { useParams } from "react-router-dom";
import { useSelector, useDispatch } from "react-redux";
import { fetchAsyncProductSingle, getProductSingle, getSingleProductStatus } from '../../store/productSlice';
import { STATUS } from '../../utils/status';
import Loader from "../../components/Loader/Loader";
import { formatPrice } from "../../utils/helpers";
import { addToCart, getCartMessageStatus, setCartMessageOff, setCartMessageOn } from '../../store/cartSlice';
import { addToComparison, getComparisonMessageStatus, setComparisonMessageOff, setComparisonMessageOn } from '../../store/comparisonSlice';
import PopupMessage from "../../components/PopupMessage/PopupMessage";

const DetailRow = ({ label, value }) => (
  <div className="detail-row">
    <span className="detail-label">{label}:</span>
    <span className="detail-value" title={value}>{value}</span>
  </div>
);

const ProductSinglePage = () => {
  const { id } = useParams();
  const dispatch = useDispatch();
  const product = useSelector(getProductSingle);
  const productSingleStatus = useSelector(getSingleProductStatus);
  const [quantity, setQuantity] = useState(1);
  const [selectedImage, setSelectedImage] = useState("");
  const [isPopupOpen, setIsPopupOpen] = useState(false);
  const cartMessageStatus = useSelector(getCartMessageStatus);
  const comparisonMessageStatus = useSelector(getComparisonMessageStatus);

  useEffect(() => {
    dispatch(fetchAsyncProductSingle(id));
  }, [id]);
  
  useEffect(() => {
    if (cartMessageStatus) {
      const timer = setTimeout(() => {
        dispatch(setCartMessageOff());
      }, 2000);
      return () => clearTimeout(timer);
    }
  }, [cartMessageStatus, dispatch]);
  
  useEffect(() => {
    if (comparisonMessageStatus) {
      const timer = setTimeout(() => {
        dispatch(setComparisonMessageOff());
      }, 2000);
      return () => clearTimeout(timer);
    }
  }, [comparisonMessageStatus, dispatch]);
  

  useEffect(() => {
    if (product?.images?.length) {
      setSelectedImage(product.images[0]);
    }
  }, [product]);

  let discountedPrice = product?.price - (product?.price * (product?.discountPercentage || 5) / 100);

  if (productSingleStatus === STATUS.LOADING) {
    return <Loader />;
  }

  const increaseQty = () => {
    setQuantity(prevQty => {
      let tempQty = prevQty + 1;
      if (tempQty > product?.stock) tempQty = product?.stock;
      return tempQty;
    });
  };

  const decreaseQty = () => {
    setQuantity(prevQty => {
      let tempQty = prevQty - 1;
      if (tempQty < 1) tempQty = 1;
      return tempQty;
    });
  };

  const addToCartHandler = (product) => {
    let discountedPrice = product?.price - (product?.price * (product?.discountPercentage || 5) / 100);
    let totalPrice = quantity * discountedPrice;
    dispatch(addToCart({ ...product, quantity, totalPrice, discountedPrice }));
    dispatch(setCartMessageOn(true));
  };

  const addToComparisonHandler = (product) => {
    dispatch(addToComparison(product));
    dispatch(setComparisonMessageOn(true));
  };

  return (
    <main className='py-5 bg-whitesmoke'>
      <div className='product-single'>
        <div className='container'>
          <div className='product-single-content bg-white grid'>
            <div className='product-single-l'>
              <div className='product-img'>
                <div className='product-img-zoom'>
                  <img src={selectedImage || ""} alt="" className='img-cover' onClick={() => setIsPopupOpen(true)} />
                </div>
                <div className='product-img-thumbs flex align-center my-2'>
                  {product?.images?.map((img, idx) => (
                    <div
                      className={`thumb-item ${selectedImage === img ? 'active' : ''}`}
                      key={idx}
                      onClick={() => setSelectedImage(img)}
                    >
                      <img src={img} alt="" className='img-cover' />
                    </div>
                  ))}
                </div>
              </div>
            </div>
            <div className='product-single-r'>
              <div className='product-details font-manrope'>
                <div className='title fs-20 fw-5'>{product?.title}</div>
                <p className='para fw-3 fs-15'>{product?.description}</p>
                <div className='info flex align-center flex-wrap fs-14'>
                  <div className='rating'>
                    <span className='text-orange fw-5'>Rating:</span>
                    <span className='mx-1'>{product?.rating}</span>
                  </div>
                  <div className='vert-line'></div>
                  <div className='brand'>
                    <span className='text-orange fw-5'>Brand:</span>
                    <span className='mx-1'>{product?.brand}</span>
                  </div>
                  <div className='vert-line'></div>
                  <div className='brand'>
                    <span className='text-orange fw-5'>Category:</span>
                    <span className='mx-1 text-capitalize'>{product?.category?.replace("-", " ")}</span>
                  </div>
                </div>
                <div className='price'>
                  <div className='flex align-center'>
                    <div className='old-price text-gray'>{formatPrice(product?.price)}</div>
                    <span className='fs-14 mx-2 text-dark'>Inclusive of all taxes</span>
                  </div>
                  <div className='flex align-center my-1'>
                    <div className='new-price fw-5 font-poppins fs-24 text-orange'>
                      {formatPrice(discountedPrice)}
                    </div>
                    <div className='discount bg-orange fs-13 text-white fw-6 font-poppins'>
                      {product?.discountPercentage}% OFF
                    </div>
                  </div>
                  <div className='brand'>
                    <span className='mx-1'>{product?.content?.replace("-", " ")}</span>
                  </div>
                  <div className='specifications'>
                    {product?.specifications && Object.entries(product.specifications).map(([key, value]) => {
                      const label = key.charAt(0).toUpperCase() + key.slice(1);
                      return (
                        <DetailRow key={key} specKey={key} label={label} value={value || '-'} />
                      );
                    })}
                  </div>
                </div>
                <div className='qty flex align-center my-4'>
                  <div className='qty-text'>Quantity:</div>
                  <div className='qty-change flex align-center mx-3'>
                    <button type='button' className='qty-decrease flex align-center justify-center' onClick={decreaseQty}>
                      <i className='fas fa-minus'></i>
                    </button>
                    <div className='qty-value flex align-center justify-center'>{quantity}</div>
                    <button type='button' className='qty-increase flex align-center justify-center' onClick={increaseQty}>
                      <i className='fas fa-plus'></i>
                    </button>
                  </div>
                  {product?.stock === 0 && <div className='qty-error text-uppercase bg-danger text-white fs-12 ls-1 mx-2 fw-5'>out of stock</div>}
                </div>
                <div className='btns'>
                  <button type='button' className='add-to-cart-btn btn' onClick={() => addToCartHandler(product)}>
                    <i className='fas fa-shopping-cart'></i>
                    <span className='btn-text mx-2'>add to cart</span>
                  </button>
                  <button type='button' className='add-to-comparison-btn btn mx-3' onClick={() => addToComparisonHandler(product)}>
                    <i className='fas fa-balance-scale'></i>
                    <span className='btn-text mx-2'>add to comparison</span>
                  </button>
                  <button type='button' className='buy-now btn'>
                    <span className='btn-text'>buy now</span>
                  </button>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
      {(cartMessageStatus || comparisonMessageStatus) && <PopupMessage message={"Item added"} />}
      {isPopupOpen && (
        <div className="image-popup-overlay" onClick={() => setIsPopupOpen(false)}>
          <div className="image-popup-content" onClick={(e) => e.stopPropagation()}>
            <img src={selectedImage || ""} alt="Full view" />
          </div>
        </div>
      )}
    </main>
  );
};

export default ProductSinglePage;
