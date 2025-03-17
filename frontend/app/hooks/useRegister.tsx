// src/hooks/useRegister.ts
import { useState } from 'react';
import axios from 'axios';
import { useLoading } from '../Contexts/LoadingContext';

export const useRegister = () => {
  const { setLoading } = useLoading();
  const [formData, setFormData] = useState({
    email: '',
    password: '',
    confirmPassword: '',
  });
  const [error, setError] = useState<string | null>(null);

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setFormData({
      ...formData,
      [e.target.name]: e.target.value,
    });
  };

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setLoading(true);

    // Check if the passwords match
    if (formData.password !== formData.confirmPassword) {
      setError("Passwords don't match");
      setLoading(false);
      return;
    }

    try {
      await axios.post('http://localhost:8080/register', formData);
      window.location.href = '/login'; // Redirect to login page after registration
    } catch (err) {
      setError('Error during registration');
    } finally {
      setLoading(false); // Hide the spinner after the API call finishes
    }
  };

  return { formData, error, handleChange, handleSubmit };
};
