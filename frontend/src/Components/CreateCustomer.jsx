import React, { useState } from 'react';
import api from '../api/api';
import './CreateCustomer.css'

const CreateCustomer = () => {
  const [fullName, setFullName] = useState('');
  const [birthDate, setBirthDate] = useState('');
  const [customerPicture, setCustomerPicture] = useState(null); 
  const [status, setStatus] = useState(true);

  const handleCreateCustomer = () => {
    const newCustomerRequest = {
      full_name: fullName,
      birth_date: birthDate,
      status: status,
      customer_picture_url: customerPicture // Burada base64 verisi olduğunu varsayıyoruz
    };

    api.post('/customer', newCustomerRequest)
      .then(response => {
        console.log('Yeni müşteri oluşturuldu:', response.data);
      })
      .catch(error => {
        console.error('Müşteri oluşturulurken bir hata oluştu:', error);
      });
  };

  const handleFileChange = (e) => {
    const file = e.target.files[0];
    const reader = new FileReader();
  
    reader.onloadend = () => {
      const base64String = reader.result;
      setCustomerPicture(base64String);
    };
  
    reader.readAsDataURL(file);
  };

  return (
    <div className="create-customer-container">
      <div>
      {customerPicture && (
        <img src={customerPicture} alt="" className="preview-image" />
      )}
      <input
        type="text"
        value={fullName}
        onChange={(e) => setFullName(e.target.value)}
        placeholder="Ad Soyad"
      />
      <input
        type="text"
        value={birthDate}
        onChange={(e) => setBirthDate(e.target.value)}
        placeholder="Doğum Tarihi (DD/MM/YYYY)"
      />
      <input
        type="file" 
        onChange={handleFileChange}
      />
      <label>
        <input
          type="checkbox"
          checked={status}
          onChange={(e) => setStatus(e.target.checked)}
        />
        Aktif
      </label>
      
      <button onClick={handleCreateCustomer}>Müşteri Oluştur</button>
    </div>
    </div>
    
  );
};

export default CreateCustomer;
