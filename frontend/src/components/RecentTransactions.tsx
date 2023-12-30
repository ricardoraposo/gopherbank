import { useQuery } from '@tanstack/react-query';
import vector from '../assets/vector.svg';
import TProfilePic from './TProfilePic';
import instance from '../api/axiosIstance';
import { choosePicture } from '../utils/transactionHelpers';

type Props = {
  id: string;
}

function RecentTransactions({ id }: Props) {
  const { data, isLoading } = useQuery({
    queryKey: ["transactions"],
    queryFn: () => instance.get(`/api/transaction/${id}`),
    select: ({ data: { transactions } }) => transactions.slice(0, 5),
  })

  if (isLoading) return <div>Loading...</div>

  return (
    <div
      className="h-28 text-white bg-orange rounded-3xl px-5
      relative flex flex-col grow justify-center items-start gap-4"
    >
      <img src={vector} alt="little thingy" className="absolute right-2 top-0 h-14" />
      <div className="z-10">
        <h2 className="text-base font-normal">Recent Transactions</h2>
        <div className="flex gap-2">
          {
            data.map((transaction: any) => (
              <TProfilePic key={transaction.id} profileURL={choosePicture(transaction.edges)} />
            ))
          }
        </div>
      </div>
    </div>
  );
}

export default RecentTransactions;
