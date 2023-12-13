import React, { useState } from 'react';
import api from '../api/api';
import './PersonnelCreate.css'; // Stil dosyasını burada içe aktarıyoruz

const PersonnelCreate = () => {
    const [name, setName] = useState('');
    const [salary, setSalary] = useState('');
    const [description, setDescription] = useState('');
    const [birthDate, setBirthDate] = useState('');
    const [startDate, setStartDate] = useState('');
    const [status, setStatus] = useState('');

    const handleSubmit = async (e) => {
        e.preventDefault();

        const newPersonnel = {
            name: name,
            salary: parseFloat(salary),
            description: description,
            birth_date: birthDate,
            start_date: startDate,
            status: status === 'true' ? true : false,
        };

        try {
            const response = await api.post('/personnel', newPersonnel);
            console.log('Yeni personel oluşturuldu:', response.data);
        } catch (error) {
            console.error('Personel oluşturulamadı:', error);
        }
    };

    return (
        <div className="personnel-create">
            <h2>Yeni Personel Oluştur</h2>
            <form onSubmit={handleSubmit}>
                <div className="form-group">
                    <label>Ad Soyad:</label>
                    <input
                        className="input-field"
                        type="text"
                        value={name}
                        placeholder="Ad Soyad"
                        onChange={(e) => setName(e.target.value)}
                    />
                </div>
                <div className="form-group">
                    <label>Maaş:</label>
                    <input
                        className="input-field"
                        type="number"
                        value={salary}
                        placeholder="Maaş"
                        onChange={(e) => setSalary(e.target.value)}
                    />
                </div>
                <div className="form-group">
    <label>Açıklama:</label>
    <textarea
        className="input-field input-field-textarea"
        value={description}
        placeholder="Açıklama girin"
        onChange={(e) => setDescription(e.target.value)}
    />
</div>
                <div className="form-group">
                    <label>Doğum Tarihi:</label>
                    <input
                        className="input-field"
                        type="text"
                        value={birthDate}
                        placeholder="Doğum günü (gg/aa/yy)"
                        onChange={(e) => setBirthDate(e.target.value)}
                    />
                </div>
                <div className="form-group">
                    <label>Başlangıç Tarihi:</label>
                    <input
                        className="input-field"
                        type="text"
                        value={startDate}
                        placeholder="Başlangıç Tarihi (gg/aa/yy)"
                        onChange={(e) => setStartDate(e.target.value)}
                    />
                </div>
                <div className="form-group">
                    <label>Durum:</label>
                    <select
                        className="input-field"
                        value={status}
                        onChange={(e) => setStatus(e.target.value)}
                    >
                        <option value="true">Aktif</option>
                        <option value="false">Pasif</option>
                    </select>
                </div>
                <button className="submit-button" type="submit">Personel Ekle</button>
            </form>
        </div>
    );
};

export default PersonnelCreate;
