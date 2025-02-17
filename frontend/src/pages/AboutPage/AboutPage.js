import React from 'react';
import "./AboutPage.scss";
import teamPhoto from '../../assets/images/team-photo.jpg'; // Add your team photo

const AboutPage = () => {
  return (
    <main className="about-page py-5 bg-whitesmoke">
      <div className="container">
        <div className="about-content bg-white p-4 rounded shadow">
          <h1 className="title">About Our Team</h1>
          <div className="team-intro">
            <img src={teamPhoto} alt="Our Team" className="team-photo" />
            <p className="intro">
              We are three ambitious Software Engineer students from Astana IT University 
              embarking on our first commercial-grade project. This initiative is part of 
              our Advanced Programming 1 course, where we're cutting our teeth on 
              microservice architecture and Docker deployment.
            </p>
          </div>

          <h2>Project Challenge</h2>
          <p>
            For the first time, we're implementing:
            <ul className="tech-list">
              <li>Microservice architecture with API gateways</li>
              <li>Docker containerization from scratch</li>
              <li>CI/CD pipeline implementation</li>
              <li>Production-grade database management</li>
            </ul>
          </p>

          <h2>Meet the Team</h2>
          <div className="team-members">
            <div className="member-card">
              <h3>Lead Fullstack Developer</h3>
              <p> <a target="_blank" href="mailto:231089@astanait.edu.kz?subject=Contact%20Me&body=Greetings%2C%20I've%20seen%20your%20golang%20project%20RESTInRehab%2C%20Good%20Work!">Ilya Gussak</a></p>
              <p>Specializing in Golang/React and Docker configuration</p>
            </div>
            <div className="member-card">
              <h3>Services Architect</h3>
              <p><a target="_blank" href="mailto:230810@astanait.edu.kz?subject=Contact%20Me&body=Greetings%2C%20I've%20seen%20your%20golang%20project%20RESTInRehab%2C%20Good%20Work!">Nursultan Nurgaliyev</a></p>
              <p>Project architecture management expert</p>
            </div>
            <div className="member-card">
              <h3>Content Creator</h3>
              <p><a target="_blank" href="mailto:230810@astanait.edu.kz?subject=Contact%20Me&body=Greetings%2C%20I've%20seen%20your%20golang%20project%20RESTInRehab%2C%20Good%20Work!">Miras Tulebayev</a></p>
              <p>Work distribution, database maintainer</p>
            </div>
          </div>
          <p className="intro">
            We are a team of three enthusiastic students from Astana IT University, embarking on our first commercial-like project.
          </p>

          <h2>Our Vision</h2>
          <p>
            Our goal is to build a robust microservice architecture using modern tools like Docker. We aim to learn, innovate, and deliver quality work while exploring industry best practices.
          </p>

          <h2>Our Journey</h2>
          <p>
            This project marks our first step into professional software development. With determination and collaboration, we’re learning the ins and outs of microservices, containerization, and scalable design.
          </p>

          <h2>Academic Context</h2>
          <p>
            This project serves as our capstone project for the <strong>Advanced Programming 1</strong> course (SE-2325), supervised by <b><i>Alshynov Shynggys</i></b>. 
            Our goal is to implement enterprise-level development practices while 
            maintaining academic rigor.
          </p>

          <p className="updated">Project Start: February 2025 | Version: 1.0.0-alpha</p>
        </div>
      </div>
    </main>
  );
};

export default AboutPage;