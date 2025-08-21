import axios from 'axios';

const API_BASE_URL = process.env.NEXT_PUBLIC_API_URL || 'http://localhost:8080'\;

export const api = axios.create({
  baseURL: API_BASE_URL,
  timeout: 5000,
});

export const healthCheck = async () => {
  try {
    const response = await api.get('/health');
    return response.data;
  } catch (error) {
    throw error;
  }
};

export const dbCheck = async () => {
  try {
    const response = await api.get('/api/db-check');
    return response.data;
  } catch (error) {
    throw error;
  }
};

export const apiInfo = async () => {
  try {
    const response = await api.get('/api');
    return response.data;
  } catch (error) {
    throw error;
  }
};
