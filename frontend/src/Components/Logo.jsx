import React   from "react";
import custom_icon from "../Assets/sidebarlogo.png"
import "./Sidebar.css"

const SidebarLogo = () => {
    return (
        <div className="sidebarlogo">
            <div className="sidebarlogoicon">
            <img src={custom_icon} alt="" style={{ width: '75px', height: '75px' }} />
            </div>
        </div>
    )
}

export default SidebarLogo;