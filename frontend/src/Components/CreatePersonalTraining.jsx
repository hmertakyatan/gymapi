import React, { useState, useEffect } from 'react';
import api from '../api/api';
import { formatDate } from '../func/timeconverter';
import './CreatePersonalTraining.css';

const CreatePersonalTraining = () => {
  const [customers, setCustomers] = useState([]);
  const [filteredCustomers, setFilteredCustomers] = useState([]);
  const [personnels, setPersonnels] = useState([]);
  const [filteredPersonnels, setFilteredPersonnels] = useState([]);
  const [selectedCustomer, setSelectedCustomer] = useState('');
  const [selectedPersonnel, setSelectedPersonnel] = useState('');
  const [description, setDescription] = useState('');
  const [price, setPrice] = useState('');
  const maxObjectToShow = 5;

  useEffect(() => {
    api.get('/customer')
      .then((response) => {
        setCustomers(response.data);
    
      })
      .catch((error) => {
        console.error('Müşteriler alınamadı:', error);
      });

    api.get('/personnel')
      .then((response) => {
        setPersonnels(response.data);
      })
      .catch((error) => {
        console.error('Personel alınamadı:', error);
      });
  }, []);

  const handleCustomerSelect = (customerId) => {
    setSelectedCustomer(customerId);
    setFilteredCustomers(customers.filter(customer => customer.id === customerId));
  };

  const handlePersonnelSelect = (personnelId) => {
    setSelectedPersonnel(personnelId);
    setFilteredPersonnels(personnels.filter(personnel => personnel.id === personnelId));
  };

  const handleRequestSubmit = () => {
    const requestData = {
      customer_id: selectedCustomer,
      personnel_id: selectedPersonnel,
      description: description,
      price: parseFloat(price),
    };

    api.post('/pt', requestData)
      .then((response) => {
        console.log('İstek gönderildi:', response);
        // İstek başarılıysa yapılacak işlemler buraya eklenebilir
      })
      .catch((error) => {
        console.error('İstek gönderilemedi:', error);
        // Hata durumunda yapılacak işlemler buraya eklenebilir
      });
  };

  const handleCustomerSearch = (e) => {
    const searchValue = e.target.value.toLowerCase();
    const filteredC = searchValue
      ? customers.filter((customer) =>
          customer.full_name.toLowerCase().includes(searchValue)
        )
      : [];
    setFilteredCustomers(filteredC.slice(0, maxObjectToShow));
  };
  
  const handlePersonnelSearch = (e) => {
    const searchValue = e.target.value.toLowerCase();
    const filteredP = searchValue
      ? personnels.filter((personnel) =>
          personnel.name.toLowerCase().includes(searchValue)
        )
      : [];
    setFilteredPersonnels(filteredP.slice(0, maxObjectToShow));
  };
  

  return (
    <div className="pt-container">
      <div className="customer-list">
        <h2>Müşteri Seçimi</h2>
        <input
          type="text"
          placeholder="Müşteri Ara..."
          onChange={handleCustomerSearch}
        />
        <div className="customers-container">
          {filteredCustomers.slice(0, maxObjectToShow).map((customer) => (
            <div
              key={customer.id}
              onClick={() => handleCustomerSelect(customer.id)}
              className={`customer-container ${selectedCustomer === customer.id ? 'selected-customer' : ''}`}
            >
              <div className="customer-details">
                <img src={customer.customer_picture_url} alt="" />
                <h3>{customer.full_name}</h3>
                <p>{formatDate(customer.birth_date)}</p>
              </div>
            </div>
          ))}
        </div>
      </div>

      <div className="personnel-list">
        <h2>Personel Seçimi</h2>
        <input
          type="text"
          placeholder="Personel Ara..."
          onChange={handlePersonnelSearch}
        />
        <div className="personnels-container">
          {filteredPersonnels.slice(0, maxObjectToShow).map((personnel) => (
            <div
              key={personnel.id}
              onClick={() => handlePersonnelSelect(personnel.id)}
              className={`personnel-container ${selectedPersonnel === personnel.id ? 'selected-personnel' : ''}`}
            >
              <h3>{personnel.name}</h3>
              <p>{personnel.description}</p>
            </div>
          ))}
        </div>
      </div>

      <div className="inputs-container">
        <textarea
          placeholder="Açıklama"
          value={description}
          onChange={(e) => setDescription(e.target.value)}
        />
        <input
          type="number"
          placeholder="Fiyat"
          value={price}
          onChange={(e) => setPrice(e.target.value)}
        />
        <button className="create-button" onClick={handleRequestSubmit}>
          Kaydet
        </button>
      </div>
    </div>
  );
};

export default CreatePersonalTraining;
