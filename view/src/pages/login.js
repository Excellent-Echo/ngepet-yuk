import React from 'react'
import Navbar from '../components/navbar'

import '../assets/css/auth.css'

import securityIllustration from '../assets/illustrations/security.svg'

function Login() {
    return (
        <>
            <div className="container-fluid">
                <Navbar/>
            </div>
            <div className="container-fluid">
                <div className="container">
                    <div className="auth-page-container">
                        <div className="row align-items-center">
                            <div className="col-md-8">
                                <div className="auth-img-container">
                                    <img className="img-fluid auth-img" src={ securityIllustration } alt=""/>
                                </div>
                            </div>
                            <div className="col-md-4">
                                <div className="auth-form-container">
                                    <h1>
                                        Halo !
                                    </h1>
                                    <p>
                                        Pe-ngepet sejati adalah mereka yang bisa ngepet dengan halal, mari lanjut belajar bersama kami.
                                    </p>
                                    <div className="auth-form">
                                        <form>
                                            <div className="mb-3">
                                                <label for="username" className="form-label title">Username</label>
                                                <input type="username" className="form-control" id="username" aria-describedby="emailHelp"/>
                                            </div>
                                            <div className="mb-3">
                                                <label for="exampleInputPassword1" className="form-label title">Password</label>
                                                <input type="password" className="form-control" id="exampleInputPassword1"/>
                                            </div>
                                            <button type="submit" className="btn btnAuth btnLime">Login</button>
                                        </form>
                                        <br/>
                                        <h6>Baru mau mulai ngepet?</h6>
                                        <button type="button" className="btn btnAuth btnLimeOutline">Daftar</button>
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </>
    )
}

export default Login
