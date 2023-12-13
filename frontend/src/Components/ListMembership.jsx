import React, { useState, useEffect } from 'react';
import api from '../api/api';
import './ListMembership.css';

const ListMembership = () => {
  const [memberships, setMemberships] = useState([]);
  const [searchTerm, setSearchTerm] = useState('');
  const [showDebtorsOnly, setShowDebtorsOnly] = useState(false);

  useEffect(() => {
    const fetchMemberships = async () => {
      try {
        const response = await api.get('/membership');
        setMemberships(response.data);
      } catch (error) {
        console.error('Bir hata oluştu:', error);
      }
    };

    fetchMemberships();
  }, []);

  const filteredMemberships = memberships.filter((membership) => {
    if (showDebtorsOnly) {
      return membership.amount_remaining > 0;
    }
    return (
      membership.edges.customer &&
      membership.edges.customer.full_name.toLowerCase().includes(searchTerm.toLowerCase())
    );
  });

  const sortMembershipsByDebt = () => {
    const sortedMemberships = [...filteredMemberships];
    sortedMemberships.sort((a, b) => b.amount_remaining - a.amount_remaining);
    setMemberships(sortedMemberships);
  };
  

  const handleSortChange = (e) => {
    if (e.target.value === 'debt') {
      sortMembershipsByDebt();
    }
  };

  return (
    <>
    <h2 className="membership-title">Üyelik Listesi</h2>
    <div className="membership-list-container">
      <div className="filter-and-sort">
        <div className="sort-memberships">
          <select onChange={handleSortChange}>
            <option value="">Sıralama</option>
            <option value="debt">Borç Miktarına Göre (En Fazla Borçludan En Az Borçluya)</option>
          </select>
        </div>
        <div className="search-bar-container">
          <input
            type="text"
            placeholder="Müşteri adıyla ara..."
            value={searchTerm}
            onChange={(e) => setSearchTerm(e.target.value)}
            className="search-bar"
          />
        </div>
        <div className="filter-debtor">
          <label>
            <input
              type="checkbox"
              checked={showDebtorsOnly}
              onChange={() => setShowDebtorsOnly(!showDebtorsOnly)}
            />
            Borçlu Müşterileri Göster
          </label>
        </div>
      </div>
      
        
      <div className="membership-list">
      
        {filteredMemberships.map((membership) => (
          <div key={membership.id} className="membership-item">
            {membership.edges.customer && (
              <div className="customer-info">
                <img src={membership.edges.customer.customer_picture_url} alt="" />
                <p className="customer-name">Müşteri Adı: {membership.edges.customer.full_name}</p>
                <p className='customer-birthdate'>Doğum Günü: {new Date(membership.edges.customer.birth_date).toLocaleDateString('tr-TR')}</p>
              </div>
            )}
            <div className="membership-info">
            <p className="start-date">Başlangıç Tarihi: {new Date(membership.start_date).toLocaleDateString('tr-TR')}</p>
            <p className="end-date">Bitiş Tarihi: {new Date(membership.end_date).toLocaleDateString('tr-TR')}</p>
            <p className="amount-paid" style={{ fontWeight: 'bold', color: membership.amount_paid > 0 ? 'green' : 'inherit' }}>
  Ödenen Miktar: {Number(membership.amount_paid).toFixed(2)}
</p>
<p className="amount-remaining" style={{ fontWeight: 'bold', color: membership.amount_remaining > 0 ? 'red' : 'inherit' }}>
  {!isNaN(membership.amount_remaining) && membership.amount_remaining !== 0
    ? `Kalan Miktar: ${Number(membership.amount_remaining).toFixed(2)}`
    : 'Tamamı Ödendi'}
</p>

            </div>
            
          </div>
        ))}
      </div>
    </div>
    </>
  );
};

export default ListMembership;
