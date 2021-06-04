import './App.css';
import React from "react";
import { BrowserRouter as Router, Switch, Route } from "react-router-dom";

// COMPONENTS
// import Navbar from "./components/navbar";

// PAGES
import Register from './pages/register';
import Login from './pages/login';
import LandingPage from './pages/landingPage';
import PageContainer from './pages/pageContainer';
import CourseList from './pages/courselist';
import UserDetail from './pages/userdetail';

function App() {
  return (
    <>
      <Router>
        <Switch>
          <Route exact path="/">
            <LandingPage/>
          </Route>
          <Route exact path="/login">
            <Login/>
          </Route>
          <Route exact path="/register">
            <Register/>
          </Route>
          <Route exact path="/user">
            <PageContainer/>
          </Route>
          <Route exact path="/user/courses">
            <CourseList/>
          </Route>
          <Route exact path="/user/detail">
            <UserDetail/>
          </Route>
        </Switch>
      </Router>
    </>
  );
}

export default App;
