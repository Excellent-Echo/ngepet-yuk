import React, {useEffect} from 'react'
import { Link } from 'react-router-dom'
import {useSelector, useDispatch} from "react-redux";

import userRegisterAction from "../redux/user/register/userRegisterAction";

import '../assets/css/auth.css'
import Navbar from '../components/navbar'
import profileIllustration from '../assets/illustrations/profile.svg'

function Register() {
    const userRegisterData = useSelector(state => state.userRegister);
    const dispatch = useDispatch();

    useEffect(()=> {
        console.log(userRegisterData.role)
        dispatch(userRegisterAction.resetFrom())
    }, []);

    const handlerRegisterSubmit = e => {
        e.preventDefault();
        //console.log(userRegisterData.userName)
        dispatch(
            userRegisterAction.register(
                userRegisterData.userName,
                userRegisterData.email,
                userRegisterData.password,
                userRegisterData.role
            )
        );
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
                                    <img className="img-fluid auth-img" src={ profileIllustration } alt=""/>
                                </div>
                            </div>
                            <div className="col-md-4">
                                <div className="auth-form-container">
                                    <h1>
                                        Halo !
                                    </h1>
                                    <p>
                                        Pe-ngepet sejati adalah mereka yang bisa ngepet dengan halal, mari belajar bersama kami.
                                    </p>
                                    <div className="auth-form">
                                        <form onSubmit={handlerRegisterSubmit}>
                                            <div className="mb-3">
                                                <label for="username" className="form-label title">Username</label>
                                                <input 
                                                type="text" 
                                                className="form-control" 
                                                id="username" 
                                                aria-describedby="emailHelp"
                                                value={userRegisterData.username}
                                                onChange={e => dispatch(userRegisterAction.setUserName(e.target.value))}
                                                />
                                            </div>
                                            <div className="mb-3">
                                                <label for="email" className="form-label title">Alamat Email</label>
                                                <input 
                                                type="email" 
                                                className="form-control" 
                                                id="email" 
                                                aria-describedby="emailHelp"
                                                value={userRegisterData.email}
                                                onChange={e => dispatch(userRegisterAction.setEmail(e.target.value))}
                                                />
                                            </div>
                                            <div className="mb-3">
                                                <label for="exampleInputPassword1" className="form-label title">Password</label>
                                                <input 
                                                type="password" 
                                                className="form-control" 
                                                id="exampleInputPassword1"
                                                value={userRegisterData.password}
                                                onChange={e => dispatch(userRegisterAction.setPassword(e.target.value))}
                                                />
                                            </div>
                                            <button 
                                            type="submit" 
                                            className="btn btnAuth btnLime"
                                            value={userRegisterData.isLoading ? "Loading.." : "Daftar"}
                                            disabled={userRegisterData.isLoading ? true : false}
                                            >Daftar</button>
                                        </form>
                                        <br/>
                                        <h6>Sudah pernah belajar ngepet disini?</h6>
                                        <Link className="btn btnAuth btnLimeOutline" to="/login">Login</Link>
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

export default Register
