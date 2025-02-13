import React, { useEffect } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import { Link } from 'react-router-dom';
import { validateTokenAsync, refreshTokenAsync, logoutAsync, selectAuthStatus } from '../../store/authSlice';
import "./Header.scss";
import Navbar from "../Navbar/Navbar";

const Header = () => {
    const dispatch = useDispatch();
    const authStatus = useSelector(selectAuthStatus);

    useEffect(() => {
        const validateToken = async () => {
            const result = await dispatch(validateTokenAsync());
            if (result.error) {
                await dispatch(refreshTokenAsync());
            }
        };
        validateToken();
    }, [dispatch]);

    const handleLogout = () => {
        dispatch(logoutAsync());
    };

    return (
        <header className='header text-white'>
            <div className='container'>
                <div className='header-cnt'>
                    <div className='header-cnt-top fs-13 py-2 flex align-center justify-between'>
                        <div className='header-cnt-top-l'>
                            <ul className='flex top-links align-center'>
                                <li className='vert-line'></li>
                                <li className='flex align-center'>
                                    <span className='fs-13'>Follow us on</span>
                                    <ul className='social-links flex align-center'>
                                        <li className='mx-2'>
                                            <a href="www.facebook.com" className='fs-15'>
                                                <i className='fab fa-facebook'></i>
                                            </a>
                                        </li> 
                                        <li className='mx-2'>
                                            <a href="www.instagram.com" className='fs-15'>
                                                <i className='fab fa-instagram'></i>
                                            </a>
                                        </li>
                                    </ul>
                                </li>
                            </ul>
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
                                {authStatus !== 'succeeded' ? (
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
                                    <>
                                        {/* <li>
                                            <Link to="/user">
                                                <span className='top-link-itm-txt'>User Page</span>
                                            </Link>
                                        </li>
                                        <li className='vert-line'></li> */}
                                        <li>
                                          <Link to="/">
                                            <span onClick={handleLogout} className='top-link-itm-txt'>Logout</span>
                                          </Link>
                                        </li>
                                    </>
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