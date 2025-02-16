import React, {useEffect} from 'react';
import "./TypeProductsPage.scss";
import ProductList from "../../components/ProductList/ProductList";
import { useDispatch, useSelector } from 'react-redux';
import { useParams } from 'react-router-dom';
import { getAllProductsByType, fetchAsyncProductsOfType, getTypeProductsStatus } from '../../store/typesSlice';
import Loader from '../../components/Loader/Loader';
import { STATUS } from '../../utils/status';

const TypeProductPage = () => {
  const dispatch = useDispatch();
  const { type } = useParams();
  const typeProducts = useSelector(getAllProductsByType);
  const typeProductsStatus = useSelector(getTypeProductsStatus);

  useEffect(() => {
    dispatch(fetchAsyncProductsOfType(type));
  }, [dispatch, type]);

  return (
    <div className='cat-products py-5 bg-whitesmoke'>
      <div className='container'>
        <div className='cat-products-content'>
          <div className='title-md'>
            <h3>See our <span className='text-capitalize'>{type.replace("-", " ")}</span></h3>
          </div>

          {
            typeProductsStatus === STATUS.LOADING ? <Loader /> : <ProductList products = {typeProducts} />
          }
        </div>
      </div>
    </div>
  )
}

export default TypeProductPage