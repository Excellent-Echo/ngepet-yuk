import './App.css';
import React from "react";
import { BrowserRouter as Router, Switch, Route } from "react-router-dom";

// COMPONENTS
import Navbar from "./components/navbar";

// PAGES
import Register from './pages/register';
import Login from './pages/login';
import LandingPage from './pages/landingPage';
import MainPage from './pages/mainPage';

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
          <Route exact path="/mainmenu">
            <MainPage/>
          </Route>
        </Switch>
      </Router>
    </>
  );
}

export default App;
