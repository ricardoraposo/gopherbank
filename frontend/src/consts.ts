export const apiURL = `http://${import.meta.env.VITE_API_HOST}:${import.meta.env.VITE_API_PORT}`;

export const profileUrl = 'https://pbs.twimg.com/media/Ew7TGJJXMAIFkDD?format=jpg&name=large';

export const profilepic1 = 'https://github.com/ashleymcnamara/gophers/blob/master/GoDZILLA.png?raw=true';
export const profilepic2 = 'https://github.com/ashleymcnamara/gophers/blob/master/GIRL_GOPHER.png?raw=true';
export const profilepic3 = 'https://github.com/ashleymcnamara/gophers/blob/master/Biker_Gopher.png?raw=true';
export const profilepic4 = 'https://github.com/ashleymcnamara/gophers/blob/master/GopherSpaceMentor.png?raw=true';

type Transaction = {
  id: number;
  name: string;
  profileURL: string;
  amount: number;
  type: 'Transfer' | 'Withdraw' | 'Deposit';
}

export const transactionMock: Transaction[] = [
  {
    id: 1,
    name: "Liz Bomtempo",
    profileURL: profilepic1,
    amount: 1200,
    type: "Transfer"
  },
  {
    id: 4,
    name: "Netflix",
    profileURL: "https://images.ctfassets.net/y2ske730sjqp/5QQ9SVIdc1tmkqrtFnG9U1/de758bba0f65dcc1c6bc1f31f161003d/BrandAssets_Logos_02-NSymbol.jpg?w=940",
    amount: -45.99,
    type: "Transfer"
  },
  {
    id: 2,
    name: "Isabela Raposo",
    profileURL: profilepic2,
    amount: -1200,
    type: "Transfer"
  },
  {
    id: 3,
    name: "Myself",
    profileURL: profileUrl,
    amount: -500,
    type: "Withdraw"
  },
]
