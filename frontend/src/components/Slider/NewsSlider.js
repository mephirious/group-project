import React from 'react';
import Slider from 'react-slick';
import "slick-carousel/slick/slick.css";
import "slick-carousel/slick/slick-theme.css";
import Post from '../Post/Post';
import "./HeaderSlider.scss";

const NewsSlider = ({ posts }) => {
  const settings = {
    autoplay: false,
    arrows: true, 
    dots: false,
    infinite: false,
    speed: 500,
    slidesToShow: 4,
    slidesToScroll: 1,
    responsive: [
      {
        breakpoint: 1200,
        settings: { slidesToShow: 3 }
      },
      {
        breakpoint: 992,
        settings: { slidesToShow: 2 }
      },
      {
        breakpoint: 768,
        settings: { slidesToShow: 1 }
      }
    ]
  };

  return (
    <div className="news-slider">
      <div className='title-md'>
        <h3>News</h3>
      </div>
      <Slider {...settings}>
        {posts.map((post) => (
          <div key={post.id} className="news-slider-item">
            <Post post={post} />
          </div>
        ))}
      </Slider>
    </div>
  );
};

export default NewsSlider;
