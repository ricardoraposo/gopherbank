import axios from 'axios';
import { apiURL } from '../consts';

const axiosInstance = axios.create({
  baseURL: `${apiURL}`,
});

export default axiosInstance;
