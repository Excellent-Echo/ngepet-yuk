import React from 'react'
import Navbar2 from '../components/navbar2'

import '../assets/css/mainpage.css'

function MainPage() {
    return (
        <>
            <div className="container-fluid">
                <Navbar2/>
            </div>                 
            <div class="offcanvas offcanvas-start mainpage-sidebar" data-bs-scroll="true" tabindex="-1" id="sidebarMenu" aria-labelledby="sidebarMenuTitle">
                <div class="offcanvas-header sidebar-title">
                    <h3 class="offcanvas-title" id="sidebarMenuTitle">NGEPET YUK!</h3>
                    <button type="button" class="btn-close text-reset" data-bs-dismiss="offcanvas" aria-label="Close"></button>
                </div>
                <div class="offcanvas-body sidebar-body">
                    <ul className="list-unstyled sidebar-menubar">
                        <li className="sidebar-link">
                            <div className="row">
                                <div className="col-3"><i class="fas fa-home"></i></div>
                                <div className="col">
                                    <p className="title">Dashboard</p>
                                </div>
                            </div>
                        </li>
                        <li className="sidebar-link">
                            <div className="row">
                                <div className="col-3"><i class="fas fa-book"></i></div>
                                <div className="col">
                                    <p className="title">
                                        Pelatihan
                                    </p>
                                </div>
                            </div>
                        </li>
                        <li className="sidebar-link">
                            <div className="row">
                                <div className="col-3"><i class="fas fa-comment-alt"></i></div>
                                <div className="col">
                                    <p className="title">Chat Mentor</p>
                                </div>
                            </div>
                        </li>
                    </ul>
                </div>
                <div className="offcanvas-body sidebar-footer">
                    <ul className="list-unstyled footer-bar">
                        <li className="sidebar-link">
                            <div className="row">
                                <div className="col-4"><i class="fas fa-sign-out-alt"></i></div>
                                <div className="col">
                                    <p className="title">Logout </p>
                                </div>
                            </div>
                        </li>
                    </ul>
                </div>
            </div>
            <div class="offcanvas offcanvas-end user-sidebar" data-bs-scroll="true" tabindex="-1" id="sidebarUser" aria-labelledby="sidebarUser">
                <div class="offcanvas-header sidebar-title">
                    <h3 class="offcanvas-title" id="sidebarMenuTitle">NGEPET YUK!</h3>
                    <button type="button" class="btn-close text-reset" data-bs-dismiss="offcanvas" aria-label="Close"></button>
                </div>
                <div class="offcanvas-body sidebar-body">
                    <ul className="list-unstyled sidebar-menubar">
                        <li className="sidebar-link">
                            <div className="row">
                                <div className="col-3"><i class="fas fa-home"></i></div>
                                <div className="col">
                                    <p className="title">Dashboard</p>
                                </div>
                            </div>
                        </li>
                        <li className="sidebar-link">
                            <div className="row">
                                <div className="col-3"><i class="fas fa-book"></i></div>
                                <div className="col">
                                    <p className="title">
                                        Pelatihan
                                    </p>
                                </div>
                            </div>
                        </li>
                        <li className="sidebar-link">
                            <div className="row">
                                <div className="col-3"><i class="fas fa-comment-alt"></i></div>
                                <div className="col">
                                    <p className="title">Chat Mentor</p>
                                </div>
                            </div>
                        </li>
                    </ul>
                </div>
                <div className="offcanvas-body sidebar-footer">
                    <ul className="list-unstyled footer-bar">
                        <li className="sidebar-link">
                            <div className="row">
                                <div className="col-4"><i class="fas fa-sign-out-alt"></i></div>
                                <div className="col">
                                    <p className="title">Logout </p>
                                </div>
                            </div>
                        </li>
                    </ul>
                </div>
            </div>
            <div className="container-fluid mainpage-container">
                <div className="container">
                    <h1>DASHBOARD</h1>
                </div>
            </div>
        </>
    )
}

export default MainPage
