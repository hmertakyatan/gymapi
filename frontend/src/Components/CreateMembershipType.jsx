import React, { useState } from 'react';
import api from '../api/api';
import './CreateMembershipType.css';

const CreateMembershipType = () => {
  const [name, setName] = useState('');
  const [description, setDescription] = useState('');
  const [membershipMonth, setMembershipMonth] = useState('');
  const [price, setPrice] = useState('');

  const handleCreateMembershipType = () => {
    const newMembershipTypeRequest = {
      name: name,
      description: description,
      membership_month: parseInt(membershipMonth),
      price: parseFloat(price),
    };

    api.post('/membership-type', newMembershipTypeRequest)
      .then(response => {
        console.log('Yeni üyelik türü oluşturuldu:', response.data);
      })
      .catch(error => {
        console.error('Üyelik türü oluşturulurken bir hata oluştu:', error);
      });
  };

  return (
    <div className="create-membership-type-container">
      <input
        type="text"
        value={name}
        onChange={(e) => setName(e.target.value)}
        placeholder="Üyelik Türü Adı"
      />
      <textarea
  value={description}
  onChange={(e) => setDescription(e.target.value)}
  placeholder="Açıklama"
></textarea>

      <input
        type="number"
        value={membershipMonth}
        onChange={(e) => setMembershipMonth(e.target.value)}
        placeholder="Üyelik Süresi (Ay)"
      />
      <input
        type="number"
        value={price}
        onChange={(e) => setPrice(e.target.value)}
        placeholder="Fiyat"
      />
      <button onClick={handleCreateMembershipType}>Üyelik Türü Oluştur</button>
    </div>
  );
};

export default CreateMembershipType;
