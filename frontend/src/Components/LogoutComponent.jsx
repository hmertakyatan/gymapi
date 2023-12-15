import React from 'react';
import api from '../api/api';
import { useNavigate } from 'react-router-dom'; // React Router kullanılıyorsa

const LogoutComponent = () => {
  const navigate = useNavigate();

  const handleLogout = async () => {
    try {
      const response = await api.get('/api/auth/logout');
      console.log(response.data); 
      navigate('/login');
    } catch (error) {
      console.error('Logout error:', error);
    }
  };
  return (
    <div>
      <button onClick={handleLogout}>Logout</button>
    </div>
  );
};

export default LogoutComponent;
