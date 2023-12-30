import TProfilePic from './TProfilePic';

type Props = {
  name: string;
  profileURL: string;
  amount: number;
  type: 'Transfer' | 'Deposit' | 'Withdraw';
};

function Transaction({ name, profileURL, amount, type }: Props) {
  const usFormat = new Intl.NumberFormat('en-US', { style: 'currency', currency: 'USD' })

  return (
    <div className="px-5 py-4 flex bg-white w-full justify-between">
      <div className="flex gap-2">
        <TProfilePic profileURL={profileURL} className="h-10 w-10" />
        <div className="flex flex-col">
          <h3 className="font-semibold">{name}</h3>
          <p className="text-gray-200">20 - 01 2023, 06:00</p>
        </div>
      </div>
      <div className="flex flex-col items-end">
        {amount > 0 ? (
          <p className="text-green">
            +
            {usFormat.format(amount)}
          </p>
        ) : (
          <p className="text-red">
            {(usFormat.format(amount))}
          </p>
        )}
        <p className="text-gray-200">{type}</p>
      </div>
    </div>
  );
}

export default Transaction;
