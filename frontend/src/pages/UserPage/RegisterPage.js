import React, { useState } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import { registerAsync, selectRegisterStatus, selectRegisterError } from '../../store/authSlice';
import './UserPage.scss';

const RegisterPage = () => {
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const [confirmPassword, setConfirmPassword] = useState('');
    const dispatch = useDispatch();
    const registerStatus = useSelector(selectRegisterStatus);
    const registerError = useSelector(selectRegisterError);

    const handleSubmit = (e) => {
        e.preventDefault();
        if (password !== confirmPassword) {
            alert('Passwords do not match.');
            return;
        }
        dispatch(registerAsync({ email, password, confirmPassword }));
    };

    return (
        <div className="user-page">
            <form onSubmit={handleSubmit}>
                <h2>Register</h2>
                {registerError && <p className="error">{registerError}</p>}
                <div>
                    <label>Email:</label>
                    <input type="email" value={email} onChange={(e) => setEmail(e.target.value)} required />
                </div>
                <div>
                    <label>Password:</label>
                    <input type="password" value={password} onChange={(e) => setPassword(e.target.value)} required />
                </div>
                <div>
                    <label>Confirm Password:</label>
                    <input type="password" value={confirmPassword} onChange={(e) => setConfirmPassword(e.target.value)} required />
                </div>
                <button type="submit" disabled={registerStatus === 'loading'}>Register</button>
            </form>
        </div>
    );
};

export default RegisterPage;
