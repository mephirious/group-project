import React, { useEffect } from 'react';
import './UserPage.scss';
import { useSelector, useDispatch } from 'react-redux';
import { fetchUserData } from '../../store/authSlice';
import Loader from '../../components/Loader/Loader';
import { formatDate } from '../../utils/helpers';

const UserPage = () => {
  const dispatch = useDispatch();
  const { user, authStatus } = useSelector((state) => state.auth);

  useEffect(() => {
    dispatch(fetchUserData());
  }, [dispatch]);

  if (authStatus === 'loading') return <Loader />;
  if (!user)
    return (
      <div className="user-page container">
        <p className="error">No user data available.</p>
      </div>
    );

  return (
    <div className="user-page container">
      <h2>User Profile</h2>
      <div className="user-details">
        <p><strong>ID:</strong> {user._id}</p>
        <p><strong>Email:</strong> {user.email}</p>
        <p><strong>Username:</strong> {user.username || 'N/A'}</p>
        <p><strong>Full Name:</strong> {user.firstName} {user.lastName}</p>
        <p><strong>Role:</strong> {user.role}</p>
        <p><strong>Verified:</strong> {user.verified ? 'Yes' : 'No'}</p>
        <p><strong>Phone:</strong> {user.phone || 'N/A'}</p>
        <p><strong>Address:</strong> {user.address || 'N/A'}</p>
        <p><strong>Company:</strong> {user.company || 'N/A'}</p>
        <p><strong>Created At:</strong> {user?.created_at?formatDate(user?.created_at):"N/A"}</p>
        <p><strong>Updated At:</strong> {user?.updated_at?formatDate(user?.updated_at):"N/A"}</p>
      </div>
    </div>
  );
};

export default UserPage;
