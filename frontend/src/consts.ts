export const apiURL = `${import.meta.env.VITE_API_PROTOCOL}://${import.meta.env.VITE_API_HOST}`;

export const defaultPic = 'https://i.imgur.com/fLjGgnc.png';

export const queryParams = (token: string) => ({
  headers: {
    Authorization: `Bearer ${token}`,
    'Content-Type': 'application/json',
  },
});
