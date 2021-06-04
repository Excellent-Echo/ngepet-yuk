import React from 'react'
import { Link } from "react-router-dom"

function Sidebar() {
    return (
        <>
            <div class="offcanvas offcanvas-start mainpage-sidebar" data-bs-scroll="true" tabindex="-1" id="sidebarMenu" aria-labelledby="sidebarMenuTitle">
                <div class="offcanvas-header sidebar-title">
                    <h3 class="offcanvas-title" id="sidebarMenuTitle">NGEPET YUK!</h3>
                    <button type="button" class="btn-close text-reset" data-bs-dismiss="offcanvas" aria-label="Close"></button>
                </div>
                <div class="offcanvas-body sidebar-body">
                    <ul className="list-unstyled sidebar-menubar">
                        <li>
                            <div className="container-fluid sidebar-link">
                                <Link className="row align-items-center"  to="user/courses">
                                    <div className="col-3"><i class="fas fa-book"></i></div>
                                    <div className="col">
                                        <span className="title">
                                            Pelatihan
                                        </span>
                                    </div>
                                </Link>
                            </div>
                        </li>
                        <li>
                            <div className="container-fluid sidebar-link">
                                <Link className="row align-items-center" to="/mentorchat">
                                    <div className="col-3"><i class="fas fa-comment-alt"></i></div>
                                    <div className="col">
                                        <span className="title">Chat Mentor</span>
                                    </div>
                                </Link>
                            </div>
                        </li>
                    </ul>
                </div>
            </div>
        </>
    )
}

export default Sidebar
