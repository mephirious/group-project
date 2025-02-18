import React from 'react';
import './FormField.scss';

const FormField = ({ label, type = 'text', name, value, onChange, required, readOnly, placeholder, options }) => {
  return (
    <div className="form-field">
      {label && <label>{label}</label>}
      {type === 'select' ? (
        <select name={name} value={value} onChange={onChange} required={required}>
          <option value="">{`Select ${label}`}</option>
          {options && options.map(option => (
            <option key={option.id} value={option.id}>{option.name}</option>
          ))}
        </select>
      ) : (
        <input
          type={type}
          name={name}
          value={value}
          onChange={onChange}
          required={required}
          readOnly={readOnly}
          placeholder={placeholder}
        />
      )}
    </div>
  );
};

export default FormField;
