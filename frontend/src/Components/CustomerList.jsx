import React, { useState, useEffect } from 'react';
import api from '../api/api';
import './CustomerList.css';
import { confirmAlert } from 'react-confirm-alert';
import 'react-confirm-alert/src/react-confirm-alert.css';
import CustomerEditForm from './CustomerEditForm';
import { formatDate } from '../func/timeconverter';


const CustomerList = () => {
  const [customers, setCustomers] = useState([]);
  const [searchTerm, setSearchTerm] = useState('');
  const [editCustomerId, setEditCustomerId] = useState(null);
  const [isEditMode, setIsEditMode] = useState(false);
  const [editedCustomerData, setEditedCustomerData] = useState({
    full_name: '',
    birth_date: '',
    customer_picture_url: '',
    status: false,
  });
  const [imagePreview, setImagePreview] = useState('');

  useEffect(() => {
    const fetchCustomers = async () => {
      try {
        const response = await api.get('/customer');
        setCustomers(response.data);
      } catch (error) {
        console.error('Bir hata oluştu:', error);
      }
    };

    fetchCustomers();
  }, []);

  const handleDelete = async (customerId, customerName) => {
    confirmAlert({
      title: 'Müşteri Silme',
      message: `${customerName} isimli müşteriyi silmek istediğinizden emin misiniz?`,
      buttons: [
        {
          label: 'Evet',
          onClick: async () => {
            try {
              await api.delete(`/customer/${customerId}`);
              setCustomers(customers.filter((customer) => customer.id !== customerId));
            } catch (error) {
              console.error('Silme işlemi başarısız oldu:', error);
            }
          }
        },
        {
          label: 'Hayır',
          onClick: () => {}
        }
      ]
    });
  };

  const handleEdit = (customerId) => {
    setEditCustomerId(customerId);
    const selectedCustomer = customers.find(customer => customer.id === customerId);
    const formattedBirthDate = formatDate(selectedCustomer.birth_date);
    setEditedCustomerData({
      ...selectedCustomer,
      customer_picture_url: selectedCustomer.customer_picture_url,
      birth_date: formattedBirthDate,

    });
  };

  const handleSaveEdit = async () => {
    try {
      await api.put(`/customer/${editCustomerId}`, {
        ...editedCustomerData,
        birth_date: editedCustomerData.birth_date
      });
      setEditCustomerId(null);

      setCustomers(customers.map(customer =>
        customer.id === editCustomerId ? { ...customer, ...editedCustomerData } : customer
      ));
    } catch (error) {
      console.error('Düzenleme işlemi başarısız oldu:', error);
    }
  };

  const handleImageChange = (e) => {
    const file = e.target.files[0];
    if (file) {
      const reader = new FileReader();
      reader.onloadend = () => {
        setImagePreview(reader.result);
        setEditedCustomerData({ ...editedCustomerData, customer_picture_url: reader.result });
      };
      reader.readAsDataURL(file);
    }
  };

  const filteredCustomers = customers.filter(customer =>
    customer.full_name.toLowerCase().includes(searchTerm.toLowerCase())
  );

  const getStatusButton = (status) => {
    if (status) {
        return (
            <span className="status-button green-button"></span>
        );
    } else {
        return (
            <span className="status-button red-button"></span>
        );
    }
};

  return (
    <div className={`customer-list-container ${isEditMode ? 'blur-background' : ''}`}>
      <div className="customer-list-title">
        <h1>Müşteri Listesi</h1>
      </div>
      
      <input
        type="text"
        placeholder="Müşteri ara..."
        value={searchTerm}
        onChange={(e) => setSearchTerm(e.target.value)}
      />
      
      
      <div className="customer-cards">
        {filteredCustomers.map((customer) => (
          <div className="customer-card" key={customer.id}>
            {editCustomerId === customer.id ? (
              <CustomerEditForm
                handleImageChange={handleImageChange}
                imagePreview={imagePreview}
                editedCustomerData={editedCustomerData}
                setEditedCustomerData={setEditedCustomerData}
                handleSaveEdit={handleSaveEdit}
              />
            ) : (
              <div className="customer-details">
                <div className="single-customer">
                  <img src={customer.customer_picture_url} alt="" />
                  <h3>{customer.full_name}</h3>
                  <p style={{ color: 'black' }}>Doğum Tarihi: {new Date(customer.birth_date).toLocaleDateString('tr-TR')}</p>
<p style={{ color: 'black' }}>Kayıt Tarihi: {new Date(customer.registration_date).toLocaleDateString('tr-TR')}</p>
<p style={{ color: 'black' }}>Aktiflik Durumu: {getStatusButton(customer.status)}</p>
                </div>

                <div className="button-container">
                  <button className="delete-button" onClick={() => handleDelete(customer.id, customer.full_name)}></button>
                  <button className="edit-button" onClick={() => handleEdit(customer.id)}></button>
                </div>
              </div>
            )}
          </div>
        ))}
      </div>
    </div>
  );
};

export default CustomerList;
