import React, { useState, useEffect } from 'react';
import api from '../api/api';
import './PersonnelList.css';

const PersonnelList = () => {
    const [personnelList, setPersonnelList] = useState([]);

    useEffect(() => {
        const fetchPersonnel = async () => {
            try {
                const response = await api.get('/personnel');
                setPersonnelList(response.data);
            } catch (error) {
                console.error('Personel listesi alınamadı:', error);
            }
        };

        fetchPersonnel();
    }, []);

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
        <div className="personnel-list-container">
            <h2 className="personnel-list-title">Personel Listesi</h2>
            <div className="personnel-list">
                {personnelList.map((personnel) => (
                    <div key={personnel.id} className="personnel-item">
                        <h3>{personnel.name}</h3>
                        <p>Maaş: {personnel.salary} TL</p>
                        <p>Açıklama: {personnel.description}</p>
                        <p>Durum: {getStatusButton(personnel.status)}</p>
                    </div>
                ))}
            </div>
        </div>
    );
};

export default PersonnelList;
