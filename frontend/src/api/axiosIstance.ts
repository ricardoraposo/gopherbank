import axios from 'axios';
import { apiURL } from '../consts';

const token = localStorage.getItem('token');

const instance = axios.create({
  baseURL: `${apiURL}`,
  headers: {
    Authorization: `Bearer ${token}`,
    'Content-Type': 'application/json',
  },
});

export default instance;
