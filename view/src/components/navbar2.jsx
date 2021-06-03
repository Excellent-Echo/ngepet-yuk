import React from 'react'
import { Link } from "react-router-dom"

function Navbar2() {
    return (
        <>      
            <nav className="navbar navbar-expand-lg navbar-dark fixed-top">
                <div className="container">
                    <button className="btn btnLimeOutline" style={{ marginRight: "15px" }} data-bs-toggle="offcanvas" data-bs-target="#sidebarMenu" aria-controls="sidebarMenu">
                        <span class="fas fa-bars"></span>
                    </button>
                        <Link className="navbar-brand title" to="/">NGEPET YUK!</Link>
                        <button className="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarNav" aria-controls="navbarNav" aria-expanded="false" aria-label="Toggle navigation">
                        <span className="navbar-toggler-icon"></span>
                    </button>
                    <div className="collapse navbar-collapse" id="navbarNav">
                        <ul className="navbar-nav ms-auto">
                            <li className="nav-item dropdown">
                                <a class="nav-link dropdown-toggle" href="#" id="navbarDropdownMenuLink" role="button" data-bs-toggle="dropdown" aria-expanded="false">
                                    <span className="title">[YOUR NAME]</span>
                                </a>
                                <ul class="dropdown-menu dropdown-menu-dark" aria-labelledby="navbarDropdownMenuLink">
                                    <Link className="dropdown-item" to="/userdetail">
                                        Profil
                                    </Link>
                                    <Link className="dropdown-item" to="/logout">
                                        Log Out
                                    </Link>
                                </ul>
                            </li>
                        </ul>
                    </div>
                </div>
            </nav>
        </>
    )
}

export default Navbar2
