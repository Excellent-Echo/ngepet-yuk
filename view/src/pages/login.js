import React from 'react';
import { Link, useHistory } from 'react-router-dom';
import { useSelector, useDispatch } from "react-redux";

import userLoginAction from "../redux/user/login/userLoginAction";

import '../assets/css/auth.css'
import Navbar from '../components/navbar'
import securityIllustration from '../assets/illustrations/security.svg'
import Footer from '../components/footer'

function Login() {
    const userLoginData = useSelector(state =>state.userLogin);
    const dispatch = useDispatch();
    const history = useHistory();

    const loginSubmitHandler = e => {
        e.preventDefault();
    
        dispatch(userLoginAction.login(userLoginData.email, userLoginData.password, history));
      };

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
                                        <form onSubmit={loginSubmitHandler}>
                                            <div className="mb-3">
                                                <label for="username" className="form-label title">Username</label>
                                                <input 
                                                type="username" 
                                                className="form-control" 
                                                id="username" 
                                                aria-describedby="emailHelp"
                                                value={userLoginData.email}
                                                onChange={e => {
                                                    dispatch(userLoginAction.setEmail(e.target.value));
                                                }}
                                                />
                                            </div>
                                            <div className="mb-3">
                                                <label for="exampleInputPassword1" className="form-label title">Password</label>
                                                <input 
                                                type="password" 
                                                className="form-control" 
                                                id="exampleInputPassword1"
                                                value={userLoginData.password}
                                                onChange={e => {
                                                    dispatch(userLoginAction.setPassword(e.target.value));
                                                }}
                                                />
                                            </div>
                                            <button type="submit" 
                                            className="btn btnAuth btnLime"
                                            value={userLoginData.isLoading ? "Loading..." : "Login"}
                                            disabled={userLoginData.isLoading ? true : false}
                                            >Login</button>
                                        </form>
                                        <br/>
                                        <h6>Baru mau mulai ngepet?</h6>
                                        <Link className="btn btnAuth btnLimeOutline" to="/register">Daftar</Link>
                                    </div>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
            <Footer/>
        </>
    )
}

export default Login
