require('dotenv').config();
const express = require('express');
const mongoose = require('mongoose');
const bcrypt = require('bcrypt');
const session = require('express-session');
const nodemailer = require('nodemailer');
const crypto = require('crypto');
const app = express();

const PORT = process.env.PORT || 3000;

// Middleware
app.use(express.static('public'));
app.use(express.urlencoded({ extended: true }));
app.use(session({
    secret: 'secureSecret',
    resave: false,
    saveUninitialized: false
}));
app.set('view engine', 'ejs');

// Database Connection
mongoose.connect(process.env.MONGODB_URI)
    .then(() => console.log('MongoDB Connected'))
    .catch(err => console.error('MongoDB Connection Error:', err));

// User Schema
const userSchema = new mongoose.Schema({
    name: String,
    email: { type: String, unique: true },
    password: String,
    resetToken: String,
    resetTokenExpiry: Date
});

const User = mongoose.model('User', userSchema);

// Nodemailer Email Configuration
const transporter = nodemailer.createTransport({
    service: 'gmail',
    auth: {
        user: process.env.EMAIL_USER, 
        pass: process.env.EMAIL_PASS  
    }
});

// Home Page
app.get('/', (req, res) => {
    res.render('index');
});

// Register Page
app.get('/register', (req, res) => {
    res.render('register');
});

// Handle User Registration
app.post('/register', async (req, res) => {
    const { name, email, password } = req.body;

    if (!name || !email || !password) {
        return res.status(400).render('error', { message: 'All fields are required!' });
    }

    try {
        const hashedPassword = await bcrypt.hash(password, 10);
        await User.create({ name, email, password: hashedPassword });
        res.redirect('/login');
    } catch (err) {
        console.error(err);
        res.status(500).render('error', { message: 'Internal Server Error. Please try again later.' });
    }
});

// Login Page
app.get('/login', (req, res) => {
    res.render('login');
});

// Handle User Login
app.post('/login', async (req, res) => {
    const { email, password } = req.body;

    if (!email || !password) {
        return res.status(400).render('error', { message: 'Email and Password are required!' });
    }

    const user = await User.findOne({ email });
    if (!user) {
        return res.status(404).render('error', { message: 'User not found. Please check your email or register a new account.' });
    }

    const isMatch = await bcrypt.compare(password, user.password);
    if (!isMatch) {
        return res.status(401).render('error', { message: 'Invalid credentials. Please try again.' });
    }

    req.session.userId = user._id;
    res.redirect('/dashboard');
});

// Dashboard
app.get('/dashboard', async (req, res) => {
    if (!req.session.userId) {
        return res.status(403).render('error', { message: 'Access Denied! Please log in first.' });
    }

    const user = await User.findById(req.session.userId);
    res.render('dashboard', { user });
});

// Logout
app.get('/logout', (req, res) => {
    req.session.destroy(() => {
        res.redirect('/');
    });
});

// Forgot Password Page
app.get('/forgot', (req, res) => {
    res.render('forgot');
});

// Handle Password Reset Request
app.post('/forgot', async (req, res) => {
    const { email } = req.body;
    const user = await User.findOne({ email });

    if (!user) {
        return res.status(404).render('error', { message: 'User not found. Please enter a registered email.' });
    }

    // Generate Secure Reset Token
    const resetToken = crypto.randomBytes(32).toString('hex');
    user.resetToken = resetToken;
    user.resetTokenExpiry = Date.now() + 3600000; // 1 hour expiry
    await user.save();

    // Send Reset Email
    const resetURL = `http://localhost:${PORT}/reset/${resetToken}`;
    const mailOptions = {
        to: user.email,
        subject: 'Password Reset Request',
        html: `<h3>Password Reset</h3>
               <p>Click <a href="${resetURL}">here</a> to reset your password.</p>
               <p>This link is valid for 1 hour.</p>`
    };

    transporter.sendMail(mailOptions, (err) => {
        if (err) {
            console.error(err);
            return res.status(500).render('error', { message: 'Error sending email. Try again later.' });
        }
        res.render('success', { message: 'Reset link sent! Check your email.' });
    });
});

// Reset Password Page
app.get('/reset/:token', async (req, res) => {
    const user = await User.findOne({
        resetToken: req.params.token,
        resetTokenExpiry: { $gt: Date.now() }
    });

    if (!user) {
        return res.status(400).render('error', { message: 'Invalid or expired reset link.' });
    }

    res.render('reset', { token: req.params.token });
});

// Handle Password Reset Submission
app.post('/reset/:token', async (req, res) => {
    const { password } = req.body;
    const user = await User.findOne({
        resetToken: req.params.token,
        resetTokenExpiry: { $gt: Date.now() }
    });

    if (!user) {
        return res.status(400).render('error', { message: 'Invalid or expired reset link.' });
    }

    user.password = await bcrypt.hash(password, 10);
    user.resetToken = undefined;
    user.resetTokenExpiry = undefined;
    await user.save();

    res.render('success', { message: 'Password reset successful! You can now log in.' });
});

// 404 Page Not Found
app.use((req, res) => {
    res.status(404).render('error', { message: 'Page Not Found! The page you are looking for does not exist.' });
});

// Start Server
app.listen(PORT, () => {
    console.log(`Server running on http://localhost:${PORT}`);
});
