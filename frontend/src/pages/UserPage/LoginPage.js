import React, { useState } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import { useNavigate } from 'react-router-dom';
import { loginAsync, verifyAuth, selectLoginStatus, selectLoginError } from '../../store/authSlice';
import './UserPage.scss';

const LoginPage = () => {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const dispatch = useDispatch();
  const navigate = useNavigate();
  const loginStatus = useSelector(selectLoginStatus);
  const loginError = useSelector(selectLoginError);

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      await dispatch(loginAsync({ email, password })).unwrap();
      await dispatch(verifyAuth()).unwrap();
      navigate('/');
    } catch (error) {
      console.error(error);
    }
  };

  return (
    <div className="user-page">
      <form onSubmit={handleSubmit}>
        <h2>Login</h2>
        {loginError && <p className="error">{loginError}</p>}
        <div>
          <label>Email:</label>
          <input type="email" value={email} onChange={(e) => setEmail(e.target.value)} required />
        </div>
        <div>
          <label>Password:</label>
          <input type="password" value={password} onChange={(e) => setPassword(e.target.value)} required />
        </div>
        <button type="submit" disabled={loginStatus === 'loading'}>Login</button>
      </form>
    </div>
  );
};

export default LoginPage;
