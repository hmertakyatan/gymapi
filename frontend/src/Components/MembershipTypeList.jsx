import React, { useState, useEffect } from 'react';
import { confirmAlert } from 'react-confirm-alert';
import api from '../api/api';
import MembershipTypeEditForm from './MemberShipTypeEditForm';
import './MembershipTypeList.css'
const MembershipTypeList = () => {
    const [membershipTypes, setMembershipTypes] = useState([]);
    const [searchTerm, setSearchTerm] = useState('');
    const [editMembershipTypeId, setEditMembershipTypeId] = useState(null);
    const [isEditMode, setIsEditMode] = useState(false);
    const [editedMembershipTypeData, setEditedMembershipTypeData] = useState({
      name: '',
      description: '',
      membership_month: 0,
      price: 0.0,
    });
    useEffect(() => {
      const fetchMembershipTypes = async () => {
        try {
          const response = await api.get('/membership-type');
          setMembershipTypes(response.data);
        } catch (error) {
          console.error('Bir hata oluştu:', error);
        }
      };
  
      fetchMembershipTypes();
    }, []);
  
    const handleDelete = async (membershipTypeId, membershipTypeName) => {
      confirmAlert({
        title: 'Üyelik Türü Silme',
        message: `${membershipTypeName} isimli üyelik türünü silmek istediğinizden emin misiniz?`,
        buttons: [
          {
            label: 'Evet',
            onClick: async () => {
              try {
                await api.delete(`/membership-type/${membershipTypeId}`);
                setMembershipTypes(membershipTypes.filter((membershipType) => membershipType.id !== membershipTypeId));
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
  
    const handleEdit = (membershipTypeId) => {
      setEditMembershipTypeId(membershipTypeId);
      const selectedMembershipType = membershipTypes.find(membershipType => membershipType.id === membershipTypeId);
      setEditedMembershipTypeData({
        ...selectedMembershipType,
      });
    };
  
    const handleSaveEdit = async () => {
      try {
        await api.put(`/membership-type/${editMembershipTypeId}`, editedMembershipTypeData);
        setEditMembershipTypeId(null);
  
        setMembershipTypes(membershipTypes.map(membershipType =>
          membershipType.id === editMembershipTypeId ? editedMembershipTypeData : membershipType
        ));
      } catch (error) {
        console.error('Düzenleme işlemi başarısız oldu:', error);
      }
    };
  
  
    return (
      <div className={`membership-type-list-container ${isEditMode ? 'blur-background' : ''}`}>
        <div className="membership-type-list-title">
          <h1>Üyelik Türü Listesi</h1>
        </div>
        <div className="membership-type-list-search">
          <input
            type="text"
            placeholder="Üyelik türü ara..."
            value={searchTerm}
            onChange={(e) => setSearchTerm(e.target.value)}
          />
        </div>
        <div className="membership-type-cards">
          {membershipTypes.map((membershipType) => (
            <div className="membership-type-card" key={membershipType.id}>
              {editMembershipTypeId === membershipType.id ? (
                <MembershipTypeEditForm
                  editedMembershipTypeData={editedMembershipTypeData}
                  setEditedMembershipTypeData={setEditedMembershipTypeData}
                  handleSaveEdit={handleSaveEdit}
                />
              ) : (
                <div className="membership-type-details">
                  <div className="single-membership-type">
                  <div className="button-container">
                    <button className="delete-button" onClick={() => handleDelete(membershipType.id, membershipType.name)}></button>
                    <button className="edit-button" onClick={() => handleEdit(membershipType.id)}></button>
                  </div>
                    <h3>{membershipType.name}</h3>
                    <p>Açıklama: {membershipType.description}</p>
                    <p>Üyelik Süresi: {membershipType.membership_month} ay</p>
                    <p>Fiyat: {membershipType.price} TL</p>
                  </div>
                 
                </div>
              )}
            </div>
          ))}
        </div>
      </div>
    );
  };
  
  export default MembershipTypeList;
  