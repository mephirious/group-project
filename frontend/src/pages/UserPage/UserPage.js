import React, { useEffect, useState } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import { validateTokenAsync, refreshTokenAsync, selectAuthStatus, selectAuthError } from '../../store/authSlice';
import './UserPage.scss';

const UserPage = () => {
    const dispatch = useDispatch();
    const authStatus = useSelector(selectAuthStatus);
    const authError = useSelector(selectAuthError);
    const [user, setUser] = useState(null);

    useEffect(() => {
        const validateToken = async () => {
            const result = await dispatch(validateTokenAsync());
            if (result.error) {
                await dispatch(refreshTokenAsync());
            } else {
                setUser(result.payload);
            }
        };
        validateToken();
    }, [dispatch]);

    if (authStatus === 'loading') {
        return <div>Loading...</div>;
    }

    if (authError) {
        return <div>Error: {authError}</div>;
    }

    return (
        <div className="user-page">
            <h2>User Page</h2>
            {user && (
                <div>
                    <p>Welcome, {user.email}</p>
                </div>
            )}
        </div>
    );
};

export default UserPage;
