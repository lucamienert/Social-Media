// src/pages/Home.tsx
import React from 'react';
import './Home.css'; // You can put extra custom styles here if needed.
import { Link } from 'react-router';

export default function HomeComponent() {
  return (
    <div className="home-container">
      {/* Top Navigation */}
      <nav className="navbar navbar-expand-lg navbar-light bg-white shadow-sm fixed-top">
        <div className="container-fluid">
          <a className="navbar-brand" href="#">SocialApp</a>
          <button className="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarNav" aria-controls="navbarNav" aria-expanded="false" aria-label="Toggle navigation">
            <span className="navbar-toggler-icon"></span>
          </button>
          <div className="collapse navbar-collapse" id="navbarNav">
            <ul className="navbar-nav ms-auto">
              <li className="nav-item">
                <Link to="/login" className="nav-link">Login</Link>
              </li>
              <li className="nav-item">
                <Link to="/register" className="nav-link">Register</Link>
              </li>
              <li className="nav-item">
                <a className="nav-link" href="#about">About</a>
              </li>
              <li className="nav-item">
                <a className="nav-link" href="#contact">Contact</a>
              </li>
            </ul>
          </div>
        </div>
      </nav>

      {/* Main Content */}
      <header className="hero-section bg-primary text-white text-center py-5">
        <div className="container">
          <h1 className="display-4">Welcome to SocialApp</h1>
          <p className="lead">Connect with people, share your moments, and grow your social circle.</p>
          <a href="#register" className="btn btn-light btn-lg">Get Started</a>
        </div>
      </header>

      {/* About Section */}
      <section id="about" className="py-5 bg-light">
        <div className="container">
          <h2 className="text-center mb-4">About Us</h2>
          <p className="text-center">
            SocialApp is the ultimate platform to stay connected with friends, family, and like-minded individuals.
            Join us to explore a world of possibilities and share your unique experiences with the world.
          </p>
        </div>
      </section>

      {/* Contact Section */}
      <section id="contact" className="py-5">
        <div className="container">
          <h2 className="text-center mb-4">Contact Us</h2>
          <p className="text-center">
            If you have any questions or feedback, feel free to reach out to us at <a href="mailto:support@socialapp.com" className="text-decoration-none">support@socialapp.com</a>.
          </p>
        </div>
      </section>

      {/* Footer */}
      <footer className="bg-dark text-white text-center py-4 mt-auto">
        <p>&copy; 2025 SocialApp. All Rights Reserved.</p>
      </footer>
    </div>
  );
}
