import React from 'react'

import Navbar2 from '../components/navbar2'
import Sidebar from '../components/sidebar'

import '../assets/css/userdetail.css'
import maleAvatarIllustration from '../assets/illustrations/maleavatar.svg'
import femaleAvatarIllustration from '../assets/illustrations/femaleavatar.svg'
import Footer from '../components/footer'

function UserDetail() {
    return (
        <>
            <div className="container-fluid">
                <Navbar2/>
            </div>            
            <Sidebar/>
            <div className="container-fluid userdetail-container">
                <div className="container-fluid userdetail-title-container">
                    <div className="container">
                        <div className="row align-items-center">
                            <div className="col-sm-8">
                                <h1>
                                    PROFIL 
                                </h1>
                            </div>
                        </div>
                    </div>
                </div>
                <div className="container-fluid userdetail-section1-container">
                    <div className="container">
                        <div className="row align-items-center d-flex user-info">
                            <div className="col-sm-3 user-avatar-container">
                                <img src={ maleAvatarIllustration } className="img-fluid user-avatar" alt=""/>
                                <div className="user-avatar-container">

                                </div>
                            </div>
                            <div className="col-sm-3">
                                <h2>[ YOUR NAME ]</h2>
                                <h4 className="accent-title">EASY</h4>
                            </div>
                        </div>
                        <div className="row justify-content-center d-flex text-center">
                            <h3>
                                EDIT INFO
                            </h3>
                            <h1 className="display-4 accent-title" style={{ fontWeight: 700, lineHeight: "3rem", marginBottom: "25px" }}>—</h1>
                        </div>
                        <div className="row justify-content-center d-flex">
                            <div className="col-sm-10">
                                <form action="">
                                    <div className="row justify-content-center form-row">
                                        <div className="col">
                                            <label for="input" className="title mb-2">NAMA DEPAN</label>
                                            <input type="text" name="first_name" className="search-bar search-bar-dark"   placeholder="[CURRENT FIRST NAME]"/>
                                        </div>
                                        <div className="col">
                                        <label for="input" className="title mb-2">NAMA BELAKANG</label>
                                            <input type="text" name="last_name" className="search-bar search-bar-dark"   placeholder="[CURRENT LAST NAME]"/>
                                        </div>
                                    </div>
                                    <div className="row justify-content-center form-row">
                                        <div className="col">
                                            <label for="input" className="title mb-2">JENIS KELAMIN</label>
                                            <select class="search-bar search-bar-dark" name="gender">
                                                <option selected>Pilih Jenis Kelamin</option>
                                                <option value="1">Laki-Laki</option>
                                                <option value="2">Perempuan</option>
                                            </select>
                                        </div>
                                        <div className="col">
                                        <label for="input" className="title mb-2">NO HANDPHONE</label>
                                            <input type="text" name="no_handphone" className="search-bar search-bar-dark"   placeholder="[CURRENT FULL NAME]"/>
                                        </div>
                                    </div>
                                    <div className="row justify-content-center form-row">
                                        <div className="col">
                                        <label for="input" className="title mb-2">ALAMAT</label>
                                            <textarea type="password" name="password" className="search-bar search-bar-dark"  />
                                        </div>
                                    </div>
                                    <div className="row justify-content-center form-row">
                                        <div className="col">
                                            <label for="input" className="title mb-2">EMAIL</label>
                                            <input type="email" name="email" className="search-bar search-bar-dark"   placeholder="[CURRENT EMAIL]"/>
                                        </div>
                                        <div className="col">
                                        <label for="input" className="title mb-2">PASSWORD</label>
                                            <input type="password" name="password" className="search-bar search-bar-dark"  />
                                        </div>
                                    </div>
                                    <div className="row d-flex justify-content-center form-row btn-container">
                                        <div className="col">
                                            <button className="btn btn-lg btnLime" type="submit">Edit Profil</button>
                                        </div>
                                    </div>
                                </form>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
            <Footer/>
        </>
    )
}

export default UserDetail
