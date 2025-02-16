import React, {useEffect} from 'react';
import "./Sidebar.scss";
import {Link} from 'react-router-dom';
import {useSelector, useDispatch} from 'react-redux';
import { getSidebarStatus, setSidebarOff } from '../../store/sidebarSlice';
import { fetchAsyncCategories, getAllCategories } from '../../store/categorySlice';
import { getAllBrands } from '../../store/brandsSlice';
import { getAllTypes } from '../../store/typesSlice';

const Sidebar = () => {

  const dispatch = useDispatch();
  const isSidebarOn = useSelector(getSidebarStatus);
  const categories = useSelector(getAllCategories);
  const brands = useSelector(getAllBrands);
  const types = useSelector(getAllTypes);

  useEffect(() => {
    dispatch(fetchAsyncCategories())
  }, [dispatch])

  return (
    <aside className={`sidebar ${isSidebarOn ? 'hide-sidebar' : ""}`}>
      <button type = "button" className='sidebar-hide-btn' onClick={() => dispatch(setSidebarOff())}>
        <i className='fas fa-times'></i>
      </button>
      <div className='sidebar-cnt'>
        <div className='cat-title fs-17 text-uppercase fw-6 ls-1h'>All Categories</div>
        <ul className='cat-list'>
          {
            categories.map((item, idx) => {
              return (
                <li key = {idx} onClick = {() => dispatch(setSidebarOff())}>
                  <Link to = {`category/${item}`} className='cat-list-link text-capitalize'>{item.replace("-", " ")}</Link>
                </li>
              )
            })
          }
        </ul>
        
        <div className='cat-title fs-17 text-uppercase fw-6 ls-1h'>All Brands</div>
        <ul className='cat-list'>
          {
            brands.map((item, idx) => {
              return (
                <li key = {idx} onClick = {() => dispatch(setSidebarOff())}>
                  <Link to = {`brands/${item}`} className='cat-list-link text-capitalize'>{item.replace("-", " ")}</Link>
                </li>
              )
            })
          }
        </ul>

        <div className='cat-title fs-17 text-uppercase fw-6 ls-1h'>All Types</div>
        <ul className='cat-list'>
          {
            types.map((item, idx) => {
              return (
                <li key = {idx} onClick = {() => dispatch(setSidebarOff())}>
                  <Link to = {`types/${item}`} className='cat-list-link text-capitalize'>{item.replace("-", " ")}</Link>
                </li>
              )
            })
          }
        </ul>
      </div>
    </aside>
  )
}

export default Sidebar