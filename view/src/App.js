import './App.css';
import React from "react";
import { BrowserRouter as Router, Switch, Route } from "react-router-dom";

// COMPONENTS
import Navbar from "./components/navbar";

// PAGES
import Register from './pages/register';
import Login from './pages/login';
import LandingPage from './pages/landingPage';

function App() {
  return (
    <>
      <Router>
        <div className="container-fluid">
          <Navbar/>
        </div>
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
        </Switch>
      </Router>
    </>
  );
}

export default App;
