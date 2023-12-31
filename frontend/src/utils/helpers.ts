export const sleep = (time: number) => {
  return new Promise<void>((resolve) => {
    setTimeout(() => {
      resolve();
    }, time);
  });
};

export const usFormat = new Intl.NumberFormat('en-US', { style: 'currency', currency: 'USD' });

export const usDateFormat = new Intl.DateTimeFormat('en-US', { dateStyle: 'medium', timeStyle: 'short' });

const date = new Date();

export const getHourAndMinutes = () => {
  return `${date.getHours()}h${date.getMinutes()}`;
};

export const getVerbFromType = (type: string) => {
  switch (type) {
    case 'transfer':
      return 'Send';
    case 'withdraw':
      return 'Withdraw';
    default:
      return 'Deposit';
  }
};
