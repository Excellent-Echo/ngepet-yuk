import React from 'react'
import Navbar from '../components/navbar'

import savingIllustration from "../assets/illustrations/savings.svg"
import ethereumIllustration from "../assets/illustrations/ethereum.svg"

function LandingPage() {
    return (
        <>
        <div className="container-fluid landing-title-container">
          <div className="container">
            <div className="row align-items-center">
              <div className="col-sm-6">
                <h1>
                  Perjalanan ngepet <i className="title">halal</i> kamu akan dimulai dari sini, bersama kami.
                </h1>
                <div className="btn-lime-container">
                  <button type="button" className="btn btnLime">Mulai Sekarang</button>
                </div>
              </div>
              <div className="col-sm-6 landing-title-image-container">
                <img src={ ethereumIllustration } className="img-fluid landing-title-image" alt="Saving"/>
              </div>
            </div>
          </div>
        </div>
        </>
    )
}

export default LandingPage
