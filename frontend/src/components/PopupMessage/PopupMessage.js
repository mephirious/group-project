import React from 'react';
import "./PopupMessage.scss";
import { correct } from "../../utils/images";

const PopupMessage = ({ message }) => {
  return (
    <div className='popup-message text-center'>
      <div className='popup-message-icon'>
        <img src = {correct} alt = "" />
      </div>
      <h6 className='text-white fs-14 fw-5'>{message}</h6>
    </div>
  )
}

export default PopupMessage