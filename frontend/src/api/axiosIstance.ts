import axios from 'axios';

const token = localStorage.getItem('jwt-token');

const instance = axios.create({
  baseURL: 'http://192.168.0.141:5432/',
  headers: {
    'Content-Type': 'application/json',
    Authorization: `Bearer ${token}`,
  },
});

export default instance;
