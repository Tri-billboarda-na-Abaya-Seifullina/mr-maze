import React from "react";
import { Link } from "react-router-dom";
import "./Header.scss";

const Header = () => {
    return (
        <header>
            <div className="header-left">
                <img src="/images/logo.png" alt="Mr. Maze" />
            </div>
            <div className="header-right">
                <nav>
                    <Link to="/about">About us</Link>
                    <Link to="/contact">Contact</Link>
                </nav>
            </div>
        </header>
    )
}

export default Header;