import React, {useEffect} from 'react';
import "./BrandProductsPage.scss";
import ProductList from "../../components/ProductList/ProductList";
import { useDispatch, useSelector } from 'react-redux';
import { useParams } from 'react-router-dom';
import { getAllProductsByBrand, fetchAsyncProductsOfBrand, getBrandProductsStatus } from '../../store/brandsSlice';
import Loader from '../../components/Loader/Loader';
import { STATUS } from '../../utils/status';

const BrandProductPage = () => {
  const dispatch = useDispatch();
  const { brand } = useParams();
  const brandProducts = useSelector(getAllProductsByBrand);
  const brandProductsStatus = useSelector(getBrandProductsStatus);

  useEffect(() => {
    dispatch(fetchAsyncProductsOfBrand(brand));
  }, [dispatch, brand]);

  return (
    <div className='cat-products py-5 bg-whitesmoke'>
      <div className='container'>
        <div className='cat-products-content'>
          <div className='title-md'>
            <h3>See our <span className='text-capitalize'>{brand.replace("-", " ")}</span></h3>
          </div>

          {
            brandProductsStatus === STATUS.LOADING ? <Loader /> : <ProductList products = {brandProducts} />
          }
        </div>
      </div>
    </div>
  )
}

export default BrandProductPage