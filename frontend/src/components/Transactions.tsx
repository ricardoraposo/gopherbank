import { useQuery } from '@tanstack/react-query';
import instance from '../api/axiosIstance';
import Line from './Line';
import Transaction from './Transaction';
import { chooseName, choosePicture, makeCapitalized } from '../utils/transactionHelpers';

type Props = {
  id: string;
}

function Transactions({ id }: Props) {
  const { data, isLoading } = useQuery({
    queryKey: ["transactions"],
    queryFn: () => instance.get(`/api/transaction/${id}`),
    select: ({ data: { transactions } }) => transactions,
  })

  const usFormat = new Intl.DateTimeFormat('en-US', { dateStyle: 'medium', timeStyle: 'short' })

  if (isLoading) return <div>Let him cook</div>

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
            data.map((transaction: any) => (
              <Transaction
                key={transaction.id}
                name={chooseName(transaction.edges)}
                profileURL={choosePicture(transaction.edges)}
                amount={transaction.edges.detail.amount}
                date={usFormat.format(new Date(transaction.edges.detail.transactedAt))}
                type={makeCapitalized(transaction.edges.detail.type) as "Transfer" | "Withdraw" | "Deposit"}
              />
            ))
          }
        </div>
      </div>
    </div>
  );
}

export default Transactions;