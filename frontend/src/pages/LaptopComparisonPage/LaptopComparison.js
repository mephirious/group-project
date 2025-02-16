import React from 'react';
import Slider from 'react-slick';
import "slick-carousel/slick/slick.css";
import "slick-carousel/slick/slick-theme.css";
import "./LaptopComparison.scss";
import { useSelector, useDispatch } from 'react-redux';
import { Link } from 'react-router-dom';
import { shopping_cart } from '../../utils/images';
import { getComparisonList, removeFromComparison, clearComparison } from '../../store/comparisonSlice';
import LaptopCard from '../../components/LaptopCard/LaptopCard';

const LaptopComparison = () => {
  const dispatch = useDispatch();
  const laptops = useSelector(getComparisonList);

  if (laptops.length === 0) {
    return (
      <div className="container my-5">
        <div className="empty-comparison flex justify-center align-center flex-column font-manrope">
          <img src={shopping_cart} alt="" />
          <span className="fw-6 fs-15 text-gray">
            Your laptop comparison list is empty.
          </span>
          <Link to="/" className="compare-btn bg-orange text-white fw-5">
            Browse Laptops
          </Link>
        </div>
      </div>
    );
  }

  const allSpecKeysSet = new Set();
  laptops.forEach(laptop => {
    if (laptop.specifications) {
      Object.keys(laptop.specifications).forEach(key => allSpecKeysSet.add(key));
    }
  });
  const allSpecKeys = Array.from(allSpecKeysSet);

  const settings = {
    autoplay: false,
    arrows: true,
    dots: false,
    infinite: false,
    speed: 500,
    slidesToShow: 4,
    slidesToScroll: 1,
    responsive: [
      { breakpoint: 1200, settings: { slidesToShow: 3 } },
      { breakpoint: 992, settings: { slidesToShow: 2 } },
      { breakpoint: 768, settings: { slidesToShow: 1 } },
    ],
  };

  return (
    <div className="laptop-comparison bg-whitesmoke">
      <div className="container">
        <h3 className="comparison-title">Laptop Comparison</h3>
        <Slider {...settings}>
          {laptops.map(laptop => (
            <div key={laptop.id} className="comparison-slide">
              <LaptopCard 
                laptop={laptop} 
                allSpecKeys={allSpecKeys} 
                onRemove={(id) => dispatch(removeFromComparison(id))}
              />
            </div>
          ))}
        </Slider>
        <div className="comparison-footer">
          <button
            type="button"
            className="clear-comparison-btn text-danger fs-15 text-uppercase fw-4"
            onClick={() => dispatch(clearComparison())}
          >
            <i className="fas fa-trash"></i>
            <span className="mx-1">Clear Comparison</span>
          </button>
        </div>
      </div>
    </div>
  );
};

export default LaptopComparison;
