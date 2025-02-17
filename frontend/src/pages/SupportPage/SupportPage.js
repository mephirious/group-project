import React from 'react';
import "./SupportPage.scss";

const SupportPage = () => {
  return (
    <main className="support-page py-5 bg-whitesmoke">
      <div className="container">
        <div className="support-content bg-white p-4 rounded shadow">
          <h1 className="title">Academic Project Support</h1>
          <p className="intro">
            As this is a student project, please allow 24-48 hours for responses. 
            We appreciate your understanding and patience.
          </p>

          <h2>Technical Assistance</h2>
          <div className="contact-methods">
            <div className="contact-card">
              <h3>Project Supervisor</h3>
              <p>Team Lead: Nursultan Nurgaliyev - <a href='mailto:230810@astanait.edu.kz?subject=RESTInRehab%20Support&body=Hello%2C%20i%20need%20a%20help%20with%20%5BINPUT%5D.'>230810@astanait.edu.kz</a><br/></p>
            </div>
            <div className="contact-card">
              <h3>Development Team</h3>
              <p>Team Lead: Ilya Gussak - <a href='mailto:231089@astanait.edu.kz?subject=RESTInRehab%20Support&body=Hello%2C%20i%20need%20a%20help%20with%20%5BINPUT%5D.'>231089@astanait.edu.kz</a><br/>
              GitHub Issues: <a href='https://github.com/mephirious/group-project/issues'>'click'</a></p>
            </div>
          </div>

          <h2>Academic Resources</h2>
          <ul className="faq-list">
            <li>
              <strong>Project Documentation:</strong><br/>
              <a href="https://github.com/mephirious/group-project/blob/main/README.md" target="_blank" rel="noreferrer">
                Technical Specification & API Docs
              </a>
            </li>
            <li>
              <strong>Deployment Guide:</strong><br/>
              <a href="https://github.com/mephirious/group-project/blob/main/README.md" target="_blank" rel="noreferrer">
                Docker setup instructions and environment variables
              </a>
            </li>
            <li>
              <strong>Academic Integrity:</strong><br/>
              <a href="https://github.com/mephirious/group-project" target="_blank" rel="noreferrer">
                All code available for review via our GitHub repository
              </a>
            </li>
          </ul>

          <h2>Office Hours</h2>
          <p>
            Monday 15:00-17:00<br/>Wednesdays 14:00-16:00<br/>
            EXPO C4.5, Floor 3, Meeting room<br/>
            (By appointment only)
          </p>

          <p className="updated">Support SLA: 48hr response time | Monitored during academic terms</p>
        </div>
      </div>
    </main>
  );
};

export default SupportPage;