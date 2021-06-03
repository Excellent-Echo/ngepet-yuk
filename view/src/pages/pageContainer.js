import React from 'react'
import { BrowserRouter as Router, Switch, Route } from "react-router-dom";

import Navbar2 from '../components/navbar2'
import Sidebar from '../components/sidebar'

import '../assets/css/pagecontainer.css'
import CourseList from './courselist';

function PageContainer() {
    return (
        <>
            <div className="container-fluid">
                <Navbar2/>
            </div>                 
            <Sidebar/>
            <Router>
                <div className="container-fluid mainpage-container">
                    <Switch>
                        <Route exact path="/user/courses">
                            <CourseList/>
                        </Route>
                    </Switch>
                </div>
            </Router>
        </>
    )
}

export default PageContainer
