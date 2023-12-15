import { Menu } from "antd";
import React from "react";
import { Link } from "react-router-dom";
// ICON 
import homeicon from "../Assets/icons/gymhomeicon.png";
import customersicon from "../Assets/icons/customersicon.png"
import tlicon from "../Assets/icons/tlicon.png"
import membershipicon from "../Assets/icons/membershipicon.png"
import membershiptypeicon from "../Assets/icons/membertshiptypeicon.png"
import personnelicon from "../Assets/icons/personnelicon.png"
import coach from "../Assets/icons/coach.png"
const iconStyle = { fontSize: '30px', width: '30px', height: '30px' };
const MenuList = ({darkTheme}) => {
    return (
        <Menu theme={darkTheme ? 'dark' : 'light'} mode="inline" className="menu-bar">
            <Menu.Item key="dashboard" icon={<img src={homeicon} alt="" style={{ ...iconStyle }} />}>
            <Link to="/dashboard">Anasayfa</Link>
            </Menu.Item>

            <Menu.SubMenu key="customertransactions" icon={<img src={customersicon} alt="" style={{ ...iconStyle }} />} title="Müşteri İşlemleri">

                <Menu.Item key="read-customers">
                    <Link to="/customer-list">Müşteri Listesi</Link>
                </Menu.Item>

                <Menu.Item key="create-customer">
                    <Link to="/create-customer">Müşteri Kaydı Oluştur</Link>
                </Menu.Item>

            </Menu.SubMenu>

            <Menu.SubMenu key="membershiptransactions" icon={<img src={membershipicon} alt="" style={{ ...iconStyle }} />} title="Üyelik İşlemleri">

                <Menu.Item key="read-memberships">
                    <Link to="/membership-list">Üyelik Listesi</Link>
                </Menu.Item>

                <Menu.Item key="create-membership">
                    <Link to="/create-membership">Üyelik Oluştur</Link>
                </Menu.Item>
            </Menu.SubMenu>

            <Menu.SubMenu key="membershiptypetransactions" icon={<img src={membershiptypeicon} alt="" style={{ ...iconStyle }} />} title="Üyelik Tipleri">
                <Menu.Item key="read-membershiptypes">
                    <Link to="/membershiptype-list">Üyelik Tipleri</Link>
                </Menu.Item>
                <Menu.Item key="create-membershiptype">
                    < Link to="/create-membershiptype">Yeni Üyelik Tipi Oluştur</Link>
                </Menu.Item>
            </Menu.SubMenu>



            <Menu.SubMenu key="paymenttransactions" icon={<img src={tlicon} alt="" style={{ ...iconStyle }} />} title="Ödeme İşlemleri">
                <Menu.Item key="read-payments">
                    <Link to="/payment-list">Ödemelerim</Link>
                </Menu.Item>
                <Menu.Item key="create-payment">
                    <Link to="/create-payment">Ödeme Oluştur</Link>
                </Menu.Item>
            </Menu.SubMenu>

            <Menu.SubMenu key="personneltransactions" icon={<img src={personnelicon} alt="" style={{ ...iconStyle }} />} title="Çalışanlar">
                <Menu.Item key="read-personnels">
                    <Link to="/personnel-list">Çalışanlar Listesi</Link>
                </Menu.Item>
                <Menu.Item key="create-personnels">
                    <Link to="/create-personnel">Yeni Çalışan Alımı</Link>
                </Menu.Item>
            </Menu.SubMenu>

            <Menu.SubMenu key="pttransactions" icon={<img src={coach} alt="" style={{ ...iconStyle }} />} title="Personal Training">
                <Menu.Item key="read-pt">
                    <Link to="/pt-list">Personal Training Listesi</Link>
                </Menu.Item>
                <Menu.Item key="create-personnels">
                    <Link to="/create-pt">Personal Training Kaydı Oluştur</Link>
                </Menu.Item>
            </Menu.SubMenu>

            

            



        </Menu>
    )
}

export default MenuList;