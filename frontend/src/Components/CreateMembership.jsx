import React, { useState, useEffect } from 'react';
import api, { fetchCustomers, fetchMembershipTypes, createMembership } from '../api/api';
import './CreateMembership.css';
import { formatDate } from '../func/timeconverter';
import ErrorNotification from './ErrorNotification';

const CreateMembership = () => {
  const [customers, setCustomers] = useState([]);
  const [filteredCustomers, setFilteredCustomers] = useState([]);
  const [membershipTypes, setMembershipTypes] = useState([]);
  const [filteredMembershipTypes, setFilteredMembershipTypes] = useState([]);
  const [selectedCustomer, setSelectedCustomer] = useState('');
  const [selectedMembershipType, setSelectedMembershipType] = useState('');
  const [amountPaid, setAmountPaid] = useState('');
  const [startDateStr, setStartDateStr] = useState('');
  const [errorMessage, setErrorMessage] = useState('');
  const maxObjectToShow = 5;

  useEffect(() => {
    fetchCustomers()
      .then((data) => {
        setCustomers(data);
      })
      .catch((error) => console.error('Müşteriler alınamadı: ', error));

    fetchMembershipTypes()
      .then((data) => {
        setMembershipTypes(data);
      })
      .catch((error) => console.error('Üyelik tipleri alınamadı: ', error));
  }, []);

  const handleCustomerClick = (customerId) => {
    setSelectedCustomer(customerId);
    setFilteredCustomers(customers.filter(customer => customer.id === customerId));
  };
  const handleMembershipTypeClick = (typeId) => {
    setSelectedMembershipType(typeId);
    setFilteredMembershipTypes(membershipTypes.filter(type => type.id === typeId));
  };

  const handleCreateMembership = () => {
    const membershipData = {
      start_date: startDateStr,
      amount_paid: parseFloat(amountPaid),
      customer_id: selectedCustomer,
      membership_type_id: selectedMembershipType,
    };
    setErrorMessage('');
    createMembership(membershipData)
      .then((response) => {
        console.log('Üyelik oluşturuldu: ', response);
      })
      .catch((error) => {
        console.error('Hata:', error);
        setErrorMessage(error.message);
      });
  };

  const handleCustomerSearch = (e) => {
    const filteredC = customers.filter((customer) =>
      customer.full_name.toLowerCase().includes(e.target.value.toLowerCase())
    );
    setFilteredCustomers(filteredC);
  };

  const handleMembershipTypeSearch = (e) => {
    const filteredTypes = membershipTypes.filter((type) =>
      type.name.toLowerCase().includes(e.target.value.toLowerCase())
    );
    setFilteredMembershipTypes(filteredTypes);
  };

  return (
    <div className="membership-container">
      {errorMessage && <ErrorNotification message={errorMessage} onClose={() => setErrorMessage('')} />}
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
    onClick={() => handleCustomerClick(customer.id)}
    className="customer-container"
    style={{
      backgroundColor: selectedCustomer === customer.id ? '#3498db' : 'white',
      color: selectedCustomer === customer.id ? 'white' : '#555',
      border: selectedCustomer === customer.id ? '2px solid #2980b9' : '1px solid #ccc',
      borderRadius: '5px',
      padding: '10px',
      marginBottom: '20px',
    }}
  >
    <div className="customer-details">
      <img src={customer.customer_picture_url} alt="" style={{maxWidth: '100px',
                        maxHeight: '100px',
                        borderRadius: '50%',
                        marginRight: '10px'}}/>
      <h3>{customer.full_name}</h3>
      <p>{formatDate(customer.birth_date)}</p>
    </div>
    
  </div>
))}



        </div>
      </div>

      <div className="membership-type-list">
        <h2>Üyelik Tipleri</h2>
        <input
          type="text"
          placeholder="Üyelik Tipi Ara..."
          onChange={handleMembershipTypeSearch}
        />
        <div className="membership-types-container">
        {filteredMembershipTypes.slice(0, maxObjectToShow).map((type) => (
  <div
    key={type.id}
    onClick={() => handleMembershipTypeClick(type.id)}
    className="membership-type-container"
    style={{
      backgroundColor: selectedMembershipType === type.id ? '#3498db' : 'white',
      color: selectedMembershipType === type.id ? 'white' : '#555',
      border: selectedMembershipType === type.id ? '2px solid #2980b9' : '1px solid #ccc',
      borderRadius: '5px',
      padding: '10px',
      marginBottom: '20px',
    }}
  >
    <h3>{type.name}</h3>
    <p>{type.description}</p>
    <p>{type.price}</p>
  </div>
))}
        </div>
      </div>

      <div className="inputs-container">
        <input
          type="number"
          value={amountPaid}
          onChange={(e) => setAmountPaid(e.target.value)}
          placeholder="Ödenen Miktar"
        />

        <input
          type="text"
          value={startDateStr}
          onChange={(e) => setStartDateStr(e.target.value)}
          placeholder="Başlangıç Tarihi (gg/aa/yyyy)"
        />

        <button className="create-button" onClick={handleCreateMembership}>
          Üyelik Oluştur
        </button>
      </div>
    </div>
  );
};

export default CreateMembership;
