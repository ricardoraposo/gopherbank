export const chooseName = (edges: any) => {
  const { detail, from_account: fromAccount, to_account: toAccount } = edges;
  switch (detail.type) {
    case 'transfer':
      if (detail.amount > 0) {
        return `${fromAccount.edges.user.firstName} ${fromAccount.edges.user.lastName}`;
      }
      return `${toAccount.edges.user.firstName} ${toAccount.edges.user.lastName}`;

    case 'deposit':
      return `${toAccount.edges.user.firstName} ${toAccount.edges.user.lastName}`;
    case 'withdraw':
      return `${fromAccount.edges.user.firstName} ${fromAccount.edges.user.lastName}`;
    default:
      return 'Unknown';
  }
};

export const chooseAccount = (edges: any) => {
  const { detail, from_account: fromAccount, to_account: toAccount } = edges;
  switch (detail.type) {
    case 'transfer':
      if (detail.amount > 0) {
        return fromAccount;
      }
      return toAccount;

    case 'deposit':
      return toAccount;
    case 'withdraw':
      return fromAccount;
    default:
      return 'Unknown';
  }
};

export const choosePicture = (edges: any) => {
  const { detail, from_account: fromAccount, to_account: toAccount } = edges;
  switch (detail.type) {
    case 'transfer':
      if (detail.amount > 0) {
        return fromAccount.edges.user.pictureUrl;
      }
      return toAccount.edges.user.pictureUrl;

    case 'deposit':
      return toAccount.edges.user.pictureUrl;
    case 'withdraw':
      return fromAccount.edges.user.pictureUrl;
    default:
      return 'Unknown';
  }
};

export const turnUnique = (transactions: any) => {
  const pictures = transactions.map((transaction: any) => choosePicture(transaction.edges));
  const unique = [...new Set(pictures)];
  return unique;
};

export const makeCapitalized = (str: string) => {
  return str.charAt(0).toUpperCase() + str.slice(1);
};
