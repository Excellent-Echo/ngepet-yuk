import React from 'react'
import Navbar from '../components/navbar'

import '../assets/css/landing.css'
import savingIllustration from "../assets/illustrations/savings.svg"
import ethereumIllustration from "../assets/illustrations/ethereum.svg"
import learningIllustration from "../assets/illustrations/learning.svg"

function LandingPage() {
    return (
        <>
        <div className="container-fluid">
          <Navbar/>
        </div>        
        <div className="container-fluid landingpage-container">
          <div className="container-fluid landing-title-container">
            <div className="container">
              <div className="row align-items-center">
                <div className="col-sm-6">
                  <h1 className="display-3" style={{ fontWeight: 700, lineHeight: "3rem" }}>—</h1>
                  <h1>
                    Perjalanan ngepet <i className="title">halal</i> kamu akan dimulai dari sini, bersama kami.
                  </h1>
                  <div className="btn-lime-container">
                    <button type="button" className="btn btnLime">Mulai Sekarang</button>
                  </div>
                </div>
                <div className="col-sm-6 landing-title-image-container">
                  <img src={ ethereumIllustration } className="img-fluid landing-title-image" alt="Ethereum"/>
                </div>
              </div>
            </div>
            <div class="custom-shape-divider-bottom-1622601846">
              <svg data-name="Layer 1" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 1200 120" preserveAspectRatio="none">
                  <path d="M321.39,56.44c58-10.79,114.16-30.13,172-41.86,82.39-16.72,168.19-17.73,250.45-.39C823.78,31,906.67,72,985.66,92.83c70.05,18.48,146.53,26.09,214.34,3V0H0V27.35A600.21,600.21,0,0,0,321.39,56.44Z" class="shape-fill"></path>
              </svg>
            </div>
          </div>
          <div className="container-fluid sections-container section1-container">
            <div className="container">
              <div className="row align-items-center">
                <div className="col-sm-6 section1-img">
                  <img src={ savingIllustration } className="img-fluid section-image" alt="Saving"/>
                </div>
                <div className="col-sm-6 section-text section1-text">
                  <h3>Ngepet <i className="title">halal</i> itu bagaimana?</h3>
                    <p className="subtitle">
                      Pasti kamu pernah menemui kasus atau seseorang yang terlihat tak bekerja
                      tapi terlihat kaya bukan? Kemungkinan besar orang tersebut menjalankan <span><i>"ngepet halal"</i></span> dengan cara mengelola uangnya
                      melalui <span><b>trading,</b></span> atau melakukan perdagangan <span><i><b>cryptocurrency.</b></i></span>
                      <span><i> Atau mungkin orang tersebut ngepet beneran (haram).</i></span>
                      <br /><br />
                      Ingin belajar caranya <span><i><b>ngepet halal?</b></i></span>
                    </p>
                    <button type="button" className="btn btnLime">Belajar Sekarang</button>
                </div>
              </div>
            </div>
          </div>
          <div className="container-fluid sections-container section2-container">
            <div className="container">
              <div className="row section2">
                <div className="col-sm-6 section2-img">
                  <img src={ learningIllustration } className="img-fluid section-image" alt="Learning"/>
                </div>
                <div className="col-sm-6 section-text section2-text">
                  <h3 className="light-title">Apa itu <span className="title">"NGEPET YUK!"</span> ?</h3>
                    <p className="subtitle">
                      <span><b>NGEPET YUK!</b></span> adalah sebuah platform untuk kamu yang ingin menjadi seorang
                      <span><i> pe-ngepet handal. </i></span> Di sini kamu dapat mempelajari seluk-beluk trading atau cryptocurrency, serta mempelajari 
                      strategi yang ampuh dalam mengambil keuntungan, ditemani & dibimbing oleh para mentor <span><i>ngepet</i></span> yang kompeten.
                      <br /><br />
                      Tertarik?
                    </p>
                    <button type="button" className="btn btnLime">Belajar Sekarang</button>
                </div>
              </div>
            </div>
          </div>
          <div className="container-fluid sections-container section3-container">
            <div className="container">
              <div className="row">
                <div className="col-sm-6 section-text section3-text">
                  <h3 className="light-title">WHY <span className="title">"NGEPET YUK!"</span> ?</h3>
                  <p className="subtitle">
                    Jika kamu masih ragu dengan <span><b>NGEPET YUK!,</b></span> berikut beberapa alasan
                    yang bisa kami berikan agar kamu lebih memilih platform kami, <span><i>alasan nomor 2 bikin tercengang!</i></span>
                  </p>
                </div>
                <div className="col-sm-6 section3-img">
                  <div className="container-fluid">
                    <div className="row section3-icon-container">
                      <div className="col">
                        <div className="row">
                          <div className="col-6">
                            <div className="row">
                              <div className="col-3">
                                <i class="fas fa-dollar-sign landing-icon"></i>
                              </div>
                              <div className="col">
                                <div className="row">
                                  <h5 className="title">Murah</h5>
                                </div>
                              </div>
                            </div>
                          </div>
                          <div className="col-6">
                            <div className="row">
                              <div className="col-3">
                                <i class="fas fa-dollar-sign landing-icon"></i>
                              </div>
                              <div className="col">
                                <div className="row">
                                  <h5 className="title">Tak Ada Biaya Lain</h5>
                                </div>
                              </div>
                            </div>
                          </div>
                        </div>
                        <div className="row">
                          <div className="col-6">
                            <i class="fas fa-dollar-sign landing-icon"></i>
                            <span><b>Halal</b></span>
                          </div>
                          <div className="col-6">
                            <i class="fas fa-dollar-sign landing-icon"></i>
                            <span><b>Lengkap</b></span>
                          </div>
                        </div>
                      </div>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>
          <div className="container-fluid sections-container section4-container">
            <div className="container">
              <div className="row text-center">
                <div className="col-sm">
                  <h6 className="title accent-title">
                    PILIH PAKET MU
                  </h6>
                  <h3>
                    PILIH PAKET BELAJARMU, DAN NGEPET YUK!
                  </h3>
                  <h1 className="display-4" style={{ fontWeight: 700, lineHeight: "3rem" }}>—</h1>
                </div>
              </div>
              <div className="row row-cols-sm-3">
                <div className="col-sm landing-price-card-container">
                  <div class="card h-100 landing-price-card">
                    <div class="card-body">
                      <h1 class="card-title">
                        <sup className="light-title fs-4">IDR&nbsp;</sup>
                          49K
                        <sub className="light-title fs-5">/bln</sub>
                      </h1>
                      <div className="card-text">
                        <h4>
                          BEGINNER
                        </h4>
                        <p className="subtitle">
                          Buat kamu yang cuma pengen tau dasar-dasar <i>per-ngepetan</i> aja.
                        </p>
                        <br />
                        <ul className="price-card-feature-list">
                          <li>15 Materi Kelas</li>
                          <li>Rekomendasi Situs Trading/Cryptocurrency</li>
                          <li>Mentor Ahli</li>
                          <li>Mentor Chat Support Terjadwal</li>
                        </ul>
                      </div>
                    </div>
                    <div className="card-footer">
                      <button type="button" className="btn btnLime">Pilih Paket Ini</button>
                    </div>
                  </div>
                </div>
                <div className="col-sm landing-price-card-container">
                  <div class="card h-100 landing-price-card middle-card">
                    <div class="card-body">
                      <h1 class="card-title">
                        <sup className="light-title fs-4">IDR&nbsp;</sup>
                          149K
                        <sub className="light-title fs-5">/bln</sub>
                      </h1>
                      <div className="card-text">
                        <h4>
                          NOVICE
                        </h4>
                        <p className="subtitle">
                          Buat kamu yang pengen tau lebih soal ngepet, dan bisa sombong ke temen-temen.
                        </p>
                        <br />
                        <ul className="price-card-feature-list">
                          <li>25 Materi Kelas</li>
                          <li>Mentor Ahli</li>
                          <li>Sesi one-on-one</li>
                          <li>Rekomendasi Situs Trading/Cryptocurrency</li>
                          <li>Mentor Chat Support 24/5</li>
                        </ul>
                      </div>
                    </div>
                    <div className="card-footer">
                      <button type="button" className="btn btnLime">Pilih Paket Ini</button>
                    </div>
                  </div>
                </div>
                <div className="col-sm landing-price-card-container">
                  <div class="card h-100 landing-price-card">
                    <div class="card-body">
                      <h1 class="card-title">
                        <sup className="light-title fs-4">IDR&nbsp;</sup>
                          249K
                        <sub className="light-title fs-5">/bln</sub>
                      </h1>
                      <div className="card-text">
                        <h4>
                          ADVANCED
                        </h4>
                        <p className="subtitle">
                          Buat kamu yang pengen jadi <i>pe-ngepet handal</i>, sekaligus jadi bahan gosip tetangga.
                        </p>
                        <br />
                        <ul className="price-card-feature-list">
                          <li>Semua Materi Kelas</li>
                          <li>Mentor Ahli</li>
                          <li>Sesi One-on-One</li>
                          <li>Rekomendasi Situs Trading/Cryptocurrency</li>
                          <li>Mentor Chat 24/7</li>
                        </ul>
                      </div>
                    </div>
                    <div className="card-footer">
                      <button type="button" className="btn btnLime">Pilih Paket Ini</button>
                    </div>                  
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
        <div className="container-fluid footer-container">
          <div className="container">
            <div className="row align-items-center">
              <div className="col">
                <b>Copyright © 2021 - Ngepet Yuk! Team</b>
              </div>
            </div>
          </div>
        </div>
        </>
    )
}

export default LandingPage
