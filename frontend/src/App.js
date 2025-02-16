import './App.scss';
import { BrowserRouter, Routes, Route } from 'react-router-dom';
import { Home, CategoryProduct, ProductSingle, Cart, Search, Login, Register, User, LaptopComparison } from "./pages/index";
import Header from "./components/Header/Header";
import Sidebar from "./components/Sidebar/Sidebar";
import Footer from "./components/Footer/Footer";
import ProtectedRoute from './components/ProtectedRoute/ProtectedRoute';
import AdminEditPage from './pages/AdminEditPage/AdminEditPage';
import React, { useEffect } from 'react';
import { useDispatch } from 'react-redux';
import { verifyAuth } from './store/authSlice';
import BrandProductPage from './pages/BrandProductsPage/BrandProductsPage';
import TypeProductPage from './pages/TypeProductsPage/TypeProductsPage';

function App() {
  const dispatch = useDispatch();
  useEffect(() => {
    dispatch(verifyAuth());
  }, [dispatch]);

  return (
    <div className="App">
      <BrowserRouter>
        <Header />
        <Sidebar />
        <Routes>
          <Route path="/" element={<Home />} />
          <Route path="/product/:id" element={<ProductSingle />} />
          <Route path="/category/:category" element={<CategoryProduct />} />
          <Route path="/brands/:brand" element={<BrandProductPage />} />
          <Route path="/types/:type" element={<TypeProductPage />} />
          <Route path="/cart" element={<Cart />} />
          <Route path="/search/:searchTerm" element={<Search />} />
          <Route path="/login" element={<Login />} />
          <Route path="/register" element={<Register />} />
          <Route path="/user" element={<User />} />
          <Route path="/compare" element={<LaptopComparison />} />
          <Route
            path="/admin/edit"
            element={
              <ProtectedRoute adminOnly={true}>
                <AdminEditPage />
              </ProtectedRoute>
            }
          />
        </Routes>
        <Footer />
      </BrowserRouter>
    </div>
  );
}

export default App;
