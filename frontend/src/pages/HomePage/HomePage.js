import React, {useEffect} from 'react';
import "./HomePage.scss";
import HeaderSlider from "../../components/Slider/HeaderSlider";
import { useSelector, useDispatch } from 'react-redux';
import { getAllCategories } from '../../store/categorySlice';
import ProductList from "../../components/ProductList/ProductList";
import { fetchAsyncProducts, getAllProducts, getAllProductsStatus } from '../../store/productSlice';
import Loader from "../../components/Loader/Loader";
import { STATUS } from '../../utils/status';
import { fetchAsyncBrands, getAllBrands } from '../../store/brandsSlice';

const HomePage = () => {
  const dispatch = useDispatch();
  const categories = useSelector(getAllCategories);
  const brands = useSelector(getAllBrands);
  useEffect(() => {
    dispatch(fetchAsyncProducts(50));
    dispatch(fetchAsyncBrands());
  }, []);

  const products = useSelector(getAllProducts);
  const productStatus = useSelector(getAllProductsStatus);

  // randomizing the products in the list
  const tempProducts = [];
  if(products.length > 0){
    for(let i in products){
      let randomIndex = Math.floor(Math.random() * products.length);

      while(tempProducts.includes(products[randomIndex])){
        randomIndex = Math.floor(Math.random() * products.length);
      }
      tempProducts[i] = products[randomIndex];
    }
  }

  let catProductsOne = products.filter(product => product.brand === brands[0]);
  let catProductsTwo = products.filter(product => product.brand === brands[1]);
  let catProductsThree = products.filter(product => product.brand === brands[2]);
  let catProductsFour = products.filter(product => product.brand === brands[3]);

  return (
    <main>
      <div className='slider-wrapper'>
        <HeaderSlider />
      </div>
      <div className='main-content bg-whitesmoke'>
        <div className='container'>
          <div className='categories py-5'>
            <div className='categories-item'>
              <div className='title-md'>
                <h3>See our products</h3>
              </div>
              { productStatus === STATUS.LOADING ? <Loader /> : <ProductList products = {tempProducts.filter((v, i) => i < 4)} />}
            </div>

            {brands.map((brand, index) => {
              const brandProducts = products.filter(product => product.brand === brand).slice(0, 4);
              return brandProducts.length > 0 && (
                <div className='categories-item' key={index}>
                  <div className='title-md'>
                    <h3>{brand}</h3>
                  </div>
                  {productStatus === STATUS.LOADING ? <Loader /> : <ProductList products={brandProducts} />}
                </div>
              );
            })}
          </div>
        </div>
      </div>
    </main>
  )
}

export default HomePage