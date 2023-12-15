import { BrowserRouter, Routes, Route, Navigate } from 'react-router-dom';
import Login from './Components/Login';
import CustomerList from './Components/CustomerList';
import CreateCustomer from './Components/CreateCustomer';
import NotFoundPage from './Components/NotFoundPage';
import React, { useState, useEffect } from 'react';
import './Components/Sidebar.css';
import './Components/Content.css'
import Logo from './Components/Logo';
import MenuList from './Components/MenuList';
import { MenuUnfoldOutlined, MenuFoldOutlined } from '@ant-design/icons';
import { Button, Layout, theme } from 'antd';
import ToggleThemeButton from './Components/ToggleThemeButton';
import CreateMembershipType from './Components/CreateMembershipType';
import MembershipTypeList from './Components/MembershipTypeList';
import CreateMembership from './Components/CreateMembership';
import ListMembership from './Components/ListMembership';
import { ToastContainer } from 'react-toastify';
import PaymentList from './Components/PaymentList';
import PaymentCreate from './Components/PaymentCreate';
import PersonnelCreate from './Components/PersonnelCreate';
import PersonnelList from './Components/PersonnelList';
import Dashboard from './Components/Dashboard';
import CreatePersonalTraining from './Components/CreatePersonalTraining';

const { Header, Sider, Content } = Layout;

function App() {
  const [darkTheme, setDarkTheme] = useState(true);
  const [collapsed, setCollapsed] = useState(false);
  const {
    token: { colorBgContainer },
  } = theme.useToken();
  const toggleTheme = () => {
    setDarkTheme(!darkTheme);
  };


  return (
    
    <BrowserRouter>
      <ToastContainer />
      <Routes>
        <Route path="/login" element={<Login />} />
      </Routes>
        <Layout>
          <Sider
            className="sidebar"
            collapsed={collapsed}
            collapsible
            trigger={null}
            theme={darkTheme ? 'dark' : 'light'}
            width={300}
          >
            <Logo />
            <MenuList darkTheme={darkTheme} />
            <ToggleThemeButton darkTheme={darkTheme} toggleTheme={toggleTheme} />
          </Sider>

          <Layout>
            <Header style={{ padding: 0, background: '#001529' }}>
              <Button
                type="text"
                className="toggle toggle-button"
                onClick={() => setCollapsed(!collapsed)}
                icon={collapsed ? <MenuUnfoldOutlined /> : <MenuFoldOutlined />}
              />
            </Header>

            <Content className='content'>
              <Routes>
                <Route path="/customer-list" element={<CustomerList />} />
                <Route path="/create-customer" element={<CreateCustomer />} />
                <Route path="/create-membershiptype" element={<CreateMembershipType/>} />
                <Route path="/membershiptype-list" element={<MembershipTypeList/>} />
                <Route path="/create-membership" element={<CreateMembership/>} />
                <Route path="/membership-list" element={<ListMembership/>} />
                <Route path="/payment-list" element={<PaymentList/>} />
                <Route path="/create-payment" element={<PaymentCreate/>} />
                <Route path="/create-personnel" element={<PersonnelCreate/>} />
                <Route path="/personnel-list" element={<PersonnelList/>} />
                <Route path="/dashboard" element={<Dashboard/>} />
                <Route path="/create-pt" element={<CreatePersonalTraining/>} />
                <Route path="/*" element={<NotFoundPage />} />
              </Routes>
            </Content>
          </Layout>
        </Layout>
    </BrowserRouter>
  );
}

export default App;
