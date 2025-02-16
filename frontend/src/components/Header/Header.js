// src/components/Header/Header.js
import React from 'react';
import { useDispatch, useSelector } from 'react-redux';
import { Link } from 'react-router-dom';
import { logoutAsync, selectAuthStatus } from '../../store/authSlice';
import "./Header.scss";
import Navbar from "../Navbar/Navbar";

const Header = () => {
  const dispatch = useDispatch();
  const user = useSelector((state) => state.auth.user);

  const handleLogout = () => {
    dispatch(logoutAsync());
  };

  return (
    <header className='header text-white'>
      <div className='container'>
        <div className='header-cnt'>
          <div className='header-cnt-top fs-13 py-2 flex align-center justify-between'>
            <div className='header-cnt-top-l'>
              {/* ...other header content... */}
            </div>
            <div className='header-cnt-top-r'>
              <ul className='top-links flex align-center'>
                <li>
                  <Link to="/" className='top-link-itm'>
                    <span className='top-link-itm-ico mx-2'>
                      <i className='fa-solid fa-circle-question'></i>
                    </span>
                    <span className='top-link-itm-txt'>Support</span>
                  </Link>
                </li>
                <li className='vert-line'></li>
                {!user ? (
                  <>
                    <li>
                      <Link to="/register">
                        <span className='top-link-itm-txt'>Register</span>
                      </Link>
                    </li>
                    <li className='vert-line'></li>
                    <li>
                      <Link to="/login">
                        <span className='top-link-itm-txt'>Log in</span>
                      </Link>
                    </li>
                  </>
                ) : (
                  <li>
                    <Link to="/">
                      <span onClick={handleLogout} className='top-link-itm-txt'>Logout</span>
                    </Link>
                  </li>
                )}
              </ul>
            </div>
          </div>
          <div className='header-cnt-bottom'>
            <Navbar />
          </div>
        </div>
      </div>
    </header>
  );
};

export default Header;
