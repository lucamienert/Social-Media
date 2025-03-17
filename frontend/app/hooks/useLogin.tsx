// src/hooks/useLogin.ts
import { useState } from 'react';
import axios from 'axios';
import { useLoading } from '../Contexts/LoadingContext';

interface LoginState {
  email: string;
  password: string;
}

export const useLogin = () => {
  const { setLoading } = useLoading();
  const [formData, setFormData] = useState<LoginState>({ email: '', password: '' });
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
    setError(null);

    try {
      const response = await axios.post('http://localhost:8080/login', formData);

      // Store JWT token in localStorage
      localStorage.setItem('token', response.data.token);
      
      // Redirect to home page
      window.location.href = '/home';
    } catch (err) {
      setError('Invalid email or password');
    } finally {
      setLoading(false); // Set loading to false when API call completes
    }
  };

  return { formData, error, handleChange, handleSubmit };
};
