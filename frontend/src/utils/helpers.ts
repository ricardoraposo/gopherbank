export const sleep = (time: number) => {
  return new Promise<void>((resolve) => {
    setTimeout(() => {
      resolve();
    }, time);
  });
};

export const usFormat = new Intl.NumberFormat('en-US', { style: 'currency', currency: 'USD' });

const date = new Date();

export const getHourAndMinutes = () => {
  return `${date.getHours()}h${date.getMinutes()}`;
};
