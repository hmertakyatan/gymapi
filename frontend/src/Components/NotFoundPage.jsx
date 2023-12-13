import React from 'react';
import { Link } from 'react-router-dom';

const NotFoundPage = () => {
  return (
    <div style={{ textAlign: 'center', marginTop: '50px' }}>
      <h1>404 - Sayfa Bulunamadı</h1>
      <p>Aradığınız sayfa maalesef bulunamadı.</p>
      <p>
        Ana sayfaya dönmek için{' '}
        <Link to="/dashboard" style={{ color: 'blue' }}>
          buraya tıklayın
        </Link>
        .
      </p>
    </div>
  );
};

export default NotFoundPage;