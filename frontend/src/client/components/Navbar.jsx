import React, { Component } from "react"
import { Link } from "react-router-dom"

import "./Navbar.scss"

export default class Navbar extends Component {
    render() {
        return (
            <div className="nav-wrapper">
                <div className="links">
                    <Link to="/">HOME</Link>
                </div>
                <div className="logo-container">
                    <span className="logo">K</span>
                </div>
            </div>
        )
    }
}