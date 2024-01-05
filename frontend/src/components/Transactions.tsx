import { useQuery } from '@tanstack/react-query';
import axios from 'axios';
import { useAtom } from 'jotai';
import Line from './Line';
import Transaction from './Transaction';
import WarningIcon from '../assets/warning.svg';
import { apiURL } from '../consts';
import { accountAtom } from '../store/atom';

function Transactions() {
  const [id] = useAtom(accountAtom);
  const { data, isLoading } = useQuery({
    queryKey: ['transactions'],
    queryFn: () => axios.get(`${apiURL}/api/transaction/${id}`),
    select: ({ data: { transactions } }) => transactions,
  });

  if (isLoading) return <div>Let him cook</div>;

  return (
    <div
      className="relative w-screen h-auto min-h-[60dvh] mt-6 bg-gray-100 rounded-t-[40px]"
    >
      <Line />
      <div>
        <br />
        <h2 className="mx-5 text-lg font-bold">Transactions</h2>
        <br />
        <div className="flex flex-col gap-2">
          {
            data.length === 0 ? (
              <div className="text-2xl text-center font-bold text-gray-200">
                <p className="my-8">No transactions</p>
                <img src={ WarningIcon } alt="warning symbol" className="w-12 h-12 mx-auto" />
              </div>
            ) : (
              data?.map((transaction: any) => (
                <Transaction
                  key={ transaction.id }
                  edges={ transaction.edges }
                />
              ))
            )
          }
        </div>
      </div>
    </div>
  );
}

export default Transactions;
