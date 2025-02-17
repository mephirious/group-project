import React from 'react';
import "./TermsOfServicePage.scss";

const TermsOfServicePage = () => {
  return (
    <main className="terms-page py-5 bg-whitesmoke">
      <div className="container">
        <div className="terms-content bg-white p-4 rounded shadow">
          <h1 className="title">Terms of Service</h1>
          <p className="intro">
            Welcome to our website! We are a team of three students from Astana IT University working on a group project. This website represents our first attempt at commercial-like development, where we are exploring microservice architecture and Docker from scratch. By accessing or using our services, you agree to abide by these Terms of Service.
          </p>

          <h2>Use of Our Services</h2>
          <p>
            Our services are provided for educational and demonstrative purposes. We ask that you use our website only for lawful purposes and in accordance with these terms. Please do not misuse our services or interfere with their proper functioning.
          </p>

          <h2>Intellectual Property</h2>
          <p>
            All content on this website, including text, images, logos, and other materials, is owned by our team or used with permission. As part of our academic project, we value creativity and respect for intellectual property. You may not reproduce or distribute any content without our explicit consent.
          </p>

          <h2>Disclaimer</h2>
          <p>
            As this project is part of our academic journey, our services are provided on an "as-is" basis. While we strive to deliver a quality experience, we cannot guarantee that the website will be free from errors or interruptions.
          </p>

          <h2>Changes to These Terms</h2>
          <p>
            We reserve the right to modify these Terms of Service at any time. Any changes will be posted on this page, and your continued use of our services will constitute your acceptance of the revised terms.
          </p>

          <p className="updated">Last updated: February 17, 2025</p>
        </div>
      </div>
    </main>
  );
};

export default TermsOfServicePage;
