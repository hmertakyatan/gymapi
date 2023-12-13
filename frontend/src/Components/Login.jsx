import React, { useState } from "react";
import api from '../api/api';
import {useNavigate } from "react-router-dom";
import user_icon from '../Assets/usernameicon.png'
import password_icon from '../Assets/passwordicon.png'
import "./Login.css"

const Login = () => {
    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');
    const [error, setError] = useState(null);
    const navigate = useNavigate();
  
    const handleSubmit = async (e) => {
      e.preventDefault();
  
      if (!username || !password) {
        setError('Kullanıcı adı ve şifre alanları boş bırakılamaz!');
        return;
      }
  
      try {
        const response = await api.post('/auth/login', {
          username: username,
          password: password,
        });
        setError(null);
        navigate('/dashboard');
      } catch (error) {
        setError('Kullanıcı adı veya şifre yanlış!');
        console.error('API isteği başarısız oldu:', error);
      }
    };

    return (
        <div className="login-bg">
            <div className="login-container">
                <div className="header">
                    <div className="text">Giriş</div>
                    <div className="underline"></div>
                </div>
                {error && <div className="error-message">{error}</div>}
                <form onSubmit={handleSubmit}>
                    <div className="inputs">
                        <div className="input">
                            <img src={user_icon} alt="" />
                            <input
                                type="text"
                                placeholder="Kullanıcı Adı"
                                value={username}
                                onChange={(e) => setUsername(e.target.value)}
                            />
                        </div>
                        <div className="input">
                            <img src={password_icon} alt="" />
                            <input
                                type="password"
                                placeholder="Şifre"
                                value={password}
                                onChange={(e) => setPassword(e.target.value)}
                            />
                        </div>
                    </div>
                    <div className="submit-container">
                        <button type="submit" className="submit">Giriş Yap</button>
                    </div>
                </form>
            </div>
        </div>
    );
}

export default Login;
