// src/hooks/useRegister.ts
import { useState } from 'react';
import axios from 'axios';
import { useLoading } from '../Contexts/LoadingContext';

export const useRegister = () => {
  const { setLoading } = useLoading();
  const [formData, setFormData] = useState({
    username: '',
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
      await axios.post("http://localhost:8080/register", formData, {
        headers: {
          "Content-Type": "application/json"
        },
        withCredentials: true  // Falls der Server Cookies oder Auth-Header braucht
      })
      .then(response => console.log("Success:", response.data))
      .catch(error => console.error("Error:", error));
      window.location.href = '/login'; // Redirect to login page after registration
    } catch (err) {
      setError('Error during registration');
    } finally {
      setLoading(false); // Hide the spinner after the API call finishes
    }
  };

  return { formData, error, handleChange, handleSubmit };
};
