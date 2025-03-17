// src/components/Spinner.tsx
import React from 'react';
import { useLoading } from '../Contexts/LoadingContext';
import SyncLoader from 'react-spinners/SyncLoader';

const Spinner = () => {
  const { isLoading } = useLoading();

  if (!isLoading) {
    return null;
  }

  return (
    <div style={spinnerOverlayStyles}>
      <SyncLoader color="#36d7b7" />
    </div>
  );
};

const spinnerOverlayStyles: React.CSSProperties = {
  position: 'fixed',
  top: 0,
  left: 0,
  right: 0,
  bottom: 0,
  display: 'flex',
  justifyContent: 'center',
  alignItems: 'center',
  backgroundColor: 'rgba(0, 0, 0, 0.5)',
  zIndex: 1000,
};

export default Spinner;
