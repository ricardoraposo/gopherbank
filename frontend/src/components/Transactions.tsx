import { useQuery } from '@tanstack/react-query';
import instance from '../api/axiosIstance';
import { transactionMock } from '../consts';
import Line from './Line';
import Transaction from './Transaction';

function Transactions() {
  const { data, isLoading } = useQuery({
    queryKey: ["transactions"],
    queryFn: () => instance.get('/api/transaction/06182488'),
    select: ({ data }) => data,
  })

  if (isLoading) return <div>I'm loading</div>

  return (
    <div
      className="relative w-screen h-auto min-h-[60dvh] mt-6 bg-gray-100 rounded-t-[40px]"
    >
      {console.log(data)}
      <Line />
      <div>
        <br />
        <h2 className="mx-5 text-lg font-bold">Transactions</h2>
        <br />
        <div className="flex flex-col gap-2">
          {
            transactionMock.map((transaction) => (
              <Transaction
                key={transaction.id}
                name={transaction.name}
                profileURL={transaction.profileURL}
                amount={transaction.amount}
                type={transaction.type}
              />
            ))
          }
        </div>
      </div>
    </div>
  );
}

export default Transactions;
