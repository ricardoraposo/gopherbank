import axios from 'axios';
import { apiURL } from '../consts';

const instance = axios.create({
  baseURL: `${apiURL}`,
  headers: {
    Authorization: `Bearer ${localStorage.getItem('token')}`,
    'Content-Type': 'application/json',
  },
});

export default instance;
