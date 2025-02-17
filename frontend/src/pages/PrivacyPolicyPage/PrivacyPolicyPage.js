import React from 'react';
import "./PrivacyPolicyPage.scss";

const PrivacyPolicyPage = () => {
  return (
    <main className="privacy-policy-page py-5 bg-whitesmoke">
      <div className="container">
        <div className="privacy-content bg-white p-4 rounded shadow">
          <h1 className="title">Academic Project Privacy Statement</h1>
          <p className="intro">
            This privacy policy applies to our academic project prototype and 
            reflects our commitment to ethical data practices per Astana IT 
            University guidelines.
          </p>

          <h2>Data Collection Scope</h2>
          <p>
            We only collect essential operational data required for:
            <ul>
              <li>Core application functionality</li>
              <li>Academic performance metrics</li>
              <li>Technical error monitoring</li>
            </ul>
          </p>

          <h2>Academic Usage</h2>
          <p>
            Data may be used in:
            <ul>
              <li>Course project demonstrations</li>
              <li>Technical documentation</li>
              <li>Academic performance reviews</li>
            </ul>
            All usage will comply with university ethics guidelines.
          </p>

          <h2>Security Measures</h2>
          <p>
            Implemented protections include:
            <ul>
              <li>Docker container isolation</li>
              <li>JWT token authentication</li>
              <li>Environment variable configuration</li>
              <li>Regular dependency audits</li>
            </ul>
          </p>

          <h2>Data Retention</h2>
          <p>
            All non-essential data will be purged:
            <ul>
              <li>Immediately after course completion</li>
              <li>Upon user account deletion request</li>
              <li>At prototype decommissioning</li>
            </ul>
          </p>

          <p className="updated">
            Valid until: June 2025 | Reviewed by: Ilya G.<br/>
            Audit Trail: GitHub commit history
          </p>
        </div>
      </div>
    </main>
  );
};

export default PrivacyPolicyPage;