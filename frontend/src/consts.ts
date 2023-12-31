export const apiURL = `http://${import.meta.env.VITE_API_HOST}:${import.meta.env.VITE_API_PORT}`;

export const defaultPic = 'https://i.imgur.com/fLjGgnc.png';

export const queryParams = (token: string) => ({
  headers: {
    Authorization: `Bearer ${token}`,
    'Content-Type': 'application/json',
  },
});
