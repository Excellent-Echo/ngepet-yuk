import React from 'react'

import Navbar2 from '../components/navbar2'
import Sidebar from '../components/sidebar'

import '../assets/css/courselist.css'
import Footer from '../components/footer'

function CourseList() {
    return (
        <>
            <div className="container-fluid">
                <Navbar2/>
            </div>            
            <Sidebar/>
            <div className="container-fluid courselist-container">
                <div className="container-fluid courselist-title-container">
                    <div className="container">
                        <div className="row align-items-center">
                            <div className="col-sm-8 courselist-title">
                                <h1>
                                    PELATIHAN NGEPET
                                </h1>
                            </div>
                            <div className="col-sm-4">
                                <form className="d-flex">
                                    <input className="search-bar search-bar-dark" type="search" placeholder="Cari Sesuatu" aria-label="Cari"/>
                                    <button className="btn btnLimeOutline" type="submit">
                                        <span className="title">
                                            <i className="fas fa-search"></i>
                                        </span>
                                    </button>
                                </form>
                            </div>
                        </div>
                    </div>
                </div>
                <div className="container-fluid courselist-section1-container">
                    <div className="container">
                        <div className="row">
                            <h3>BERDASARKAN PAKET MU : <span className="accent-title title">EASY</span></h3>
                        </div>
                        <div className="row courselist-cards-container overflow-auto">
                            <div className="col-sm-3">
                                <div className="card courselist-card">
                                <img src="..." className="card-img-top" alt="..."/>
                                    <div className="card-body">
                                        <h5 className="card-title">[ COURSE TITLE ]</h5>
                                        <p className="card-text">
                                            Lorem ipsum dolor sit amet, consectetur adipisicing elit. Inventore qui unde consequatur dolore delectus fugiat! 
                                        </p>
                                    </div>
                                </div>
                            </div>
                            <div className="col-sm-3">
                                <div className="card courselist-card">
                                <img src="..." className="card-img-top" alt="..."/>
                                    <div className="card-body">
                                        <h5 className="card-title">[ COURSE TITLE ]</h5>
                                        <p className="card-text">
                                            Lorem ipsum dolor sit amet, consectetur adipisicing elit. Inventore qui unde consequatur dolore delectus fugiat! 
                                        </p>
                                    </div>
                                </div>
                            </div>
                            <div className="col-sm-3">
                                <div className="card courselist-card">
                                <img src="..." className="card-img-top" alt="..."/>
                                    <div className="card-body">
                                        <h5 className="card-title">[ COURSE TITLE ]</h5>
                                        <p className="card-text">
                                            Lorem ipsum dolor sit amet, consectetur adipisicing elit. Inventore qui unde consequatur dolore delectus fugiat! 
                                        </p>
                                    </div>
                                </div>
                            </div>
                            <div className="col-sm-3">
                                <div className="card courselist-card">
                                <img src="..." className="card-img-top" alt="..."/>
                                    <div className="card-body">
                                        <h5 className="card-title">[ COURSE TITLE ]</h5>
                                        <p className="card-text">
                                            Lorem ipsum dolor sit amet, consectetur adipisicing elit. Inventore qui unde consequatur dolore delectus fugiat! 
                                        </p>
                                    </div>
                                </div>
                            </div>
                            <div className="col-sm-3">
                                <div className="card courselist-card">
                                <img src="..." className="card-img-top" alt="..."/>
                                    <div className="card-body">
                                        <h5 className="card-title">[ COURSE TITLE ]</h5>
                                        <p className="card-text">
                                            Lorem ipsum dolor sit amet, consectetur adipisicing elit. Inventore qui unde consequatur dolore delectus fugiat! 
                                        </p>
                                    </div>
                                </div>
                            </div>
                            <div className="col-sm-3">
                                <div className="card courselist-card">
                                <img src="..." className="card-img-top" alt="..."/>
                                    <div className="card-body">
                                        <h5 className="card-title">[ COURSE TITLE ]</h5>
                                        <p className="card-text">
                                            Lorem ipsum dolor sit amet, consectetur adipisicing elit. Inventore qui unde consequatur dolore delectus fugiat! 
                                        </p>
                                    </div>
                                </div>
                            </div>
                            <div className="col-sm-3">
                                <div className="card courselist-card">
                                <img src="..." className="card-img-top" alt="..."/>
                                    <div className="card-body">
                                        <h5 className="card-title">[ COURSE TITLE ]</h5>
                                        <p className="card-text">
                                            Lorem ipsum dolor sit amet, consectetur adipisicing elit. Inventore qui unde consequatur dolore delectus fugiat! 
                                        </p>
                                    </div>
                                </div>
                            </div>

                        </div>
                    </div>
                </div>
                <div className="container-fluid courselist-sections-container">
                    <div className="container">
                        <div className="row">
                            <h3>KATEGORI : <span className="accent-title title">TRADING</span></h3>
                        </div>
                        <div className="row courselist-cards-container overflow-auto">
                            <div className="col-sm-3">
                                <div className="card courselist-card">
                                <img src="..." className="card-img-top" alt="..."/>
                                    <div className="card-body">
                                        <h5 className="card-title">[ COURSE TITLE ]</h5>
                                        <p className="card-text">
                                            Lorem ipsum dolor sit amet, consectetur adipisicing elit. Inventore qui unde consequatur dolore delectus fugiat! 
                                        </p>
                                    </div>
                                </div>
                            </div>
                            <div className="col-sm-3">
                                <div className="card courselist-card">
                                <img src="..." className="card-img-top" alt="..."/>
                                    <div className="card-body">
                                        <h5 className="card-title">[ COURSE TITLE ]</h5>
                                        <p className="card-text">
                                            Lorem ipsum dolor sit amet, consectetur adipisicing elit. Inventore qui unde consequatur dolore delectus fugiat! 
                                        </p>
                                    </div>
                                </div>
                            </div>
                            <div className="col-sm-3">
                                <div className="card courselist-card">
                                <img src="..." className="card-img-top" alt="..."/>
                                    <div className="card-body">
                                        <h5 className="card-title">[ COURSE TITLE ]</h5>
                                        <p className="card-text">
                                            Lorem ipsum dolor sit amet, consectetur adipisicing elit. Inventore qui unde consequatur dolore delectus fugiat! 
                                        </p>
                                    </div>
                                </div>
                            </div>
                            <div className="col-sm-3">
                                <div className="card courselist-card">
                                <img src="..." className="card-img-top" alt="..."/>
                                    <div className="card-body">
                                        <h5 className="card-title">[ COURSE TITLE ]</h5>
                                        <p className="card-text">
                                            Lorem ipsum dolor sit amet, consectetur adipisicing elit. Inventore qui unde consequatur dolore delectus fugiat! 
                                        </p>
                                    </div>
                                </div>
                            </div>
                            <div className="col-sm-3">
                                <div className="card courselist-card">
                                <img src="..." className="card-img-top" alt="..."/>
                                    <div className="card-body">
                                        <h5 className="card-title">[ COURSE TITLE ]</h5>
                                        <p className="card-text">
                                            Lorem ipsum dolor sit amet, consectetur adipisicing elit. Inventore qui unde consequatur dolore delectus fugiat! 
                                        </p>
                                    </div>
                                </div>
                            </div>
                            <div className="col-sm-3">
                                <div className="card courselist-card">
                                <img src="..." className="card-img-top" alt="..."/>
                                    <div className="card-body">
                                        <h5 className="card-title">[ COURSE TITLE ]</h5>
                                        <p className="card-text">
                                            Lorem ipsum dolor sit amet, consectetur adipisicing elit. Inventore qui unde consequatur dolore delectus fugiat! 
                                        </p>
                                    </div>
                                </div>
                            </div>
                            <div className="col-sm-3">
                                <div className="card courselist-card">
                                <img src="..." className="card-img-top" alt="..."/>
                                    <div className="card-body">
                                        <h5 className="card-title">[ COURSE TITLE ]</h5>
                                        <p className="card-text">
                                            Lorem ipsum dolor sit amet, consectetur adipisicing elit. Inventore qui unde consequatur dolore delectus fugiat! 
                                        </p>
                                    </div>
                                </div>
                            </div>

                        </div>
                    </div>
                </div>
                <div className="container-fluid courselist-sections-container courselist-sections-container">
                    <div className="container">
                        <div className="row">
                            <h3>KATEGORI : <span className="accent-title title">CRYPTOCURRENCY</span></h3>
                        </div>
                        <div className="row courselist-cards-container overflow-auto">
                            <div className="col-sm-3">
                                <div className="card courselist-card">
                                <img src="..." className="card-img-top" alt="..."/>
                                    <div className="card-body">
                                        <h5 className="card-title">[ COURSE TITLE ]</h5>
                                        <p className="card-text">
                                            Lorem ipsum dolor sit amet, consectetur adipisicing elit. Inventore qui unde consequatur dolore delectus fugiat! 
                                        </p>
                                    </div>
                                </div>
                            </div>
                            <div className="col-sm-3">
                                <div className="card courselist-card">
                                <img src="..." className="card-img-top" alt="..."/>
                                    <div className="card-body">
                                        <h5 className="card-title">[ COURSE TITLE ]</h5>
                                        <p className="card-text">
                                            Lorem ipsum dolor sit amet, consectetur adipisicing elit. Inventore qui unde consequatur dolore delectus fugiat! 
                                        </p>
                                    </div>
                                </div>
                            </div>
                            <div className="col-sm-3">
                                <div className="card courselist-card">
                                <img src="..." className="card-img-top" alt="..."/>
                                    <div className="card-body">
                                        <h5 className="card-title">[ COURSE TITLE ]</h5>
                                        <p className="card-text">
                                            Lorem ipsum dolor sit amet, consectetur adipisicing elit. Inventore qui unde consequatur dolore delectus fugiat! 
                                        </p>
                                    </div>
                                </div>
                            </div>
                            <div className="col-sm-3">
                                <div className="card courselist-card">
                                <img src="..." className="card-img-top" alt="..."/>
                                    <div className="card-body">
                                        <h5 className="card-title">[ COURSE TITLE ]</h5>
                                        <p className="card-text">
                                            Lorem ipsum dolor sit amet, consectetur adipisicing elit. Inventore qui unde consequatur dolore delectus fugiat! 
                                        </p>
                                    </div>
                                </div>
                            </div>
                            <div className="col-sm-3">
                                <div className="card courselist-card">
                                <img src="..." className="card-img-top" alt="..."/>
                                    <div className="card-body">
                                        <h5 className="card-title">[ COURSE TITLE ]</h5>
                                        <p className="card-text">
                                            Lorem ipsum dolor sit amet, consectetur adipisicing elit. Inventore qui unde consequatur dolore delectus fugiat! 
                                        </p>
                                    </div>
                                </div>
                            </div>
                            <div className="col-sm-3">
                                <div className="card courselist-card">
                                <img src="..." className="card-img-top" alt="..."/>
                                    <div className="card-body">
                                        <h5 className="card-title">[ COURSE TITLE ]</h5>
                                        <p className="card-text">
                                            Lorem ipsum dolor sit amet, consectetur adipisicing elit. Inventore qui unde consequatur dolore delectus fugiat! 
                                        </p>
                                    </div>
                                </div>
                            </div>
                            <div className="col-sm-3">
                                <div className="card courselist-card">
                                <img src="..." className="card-img-top" alt="..."/>
                                    <div className="card-body">
                                        <h5 className="card-title">[ COURSE TITLE ]</h5>
                                        <p className="card-text">
                                            Lorem ipsum dolor sit amet, consectetur adipisicing elit. Inventore qui unde consequatur dolore delectus fugiat! 
                                        </p>
                                    </div>
                                </div>
                            </div>

                        </div>
                    </div>
                </div>
                <div className="container-fluid before-footer">

                </div>
            </div>
            <Footer/>
        </>
    )
}

export default CourseList
