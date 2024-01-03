import { usFormat } from '../utils/helpers';
import { chooseName, makeCapitalized } from '../utils/transactionHelpers';
import UserPopover from './UserPopover';

type Props = {
  edges: any;
};

function Transaction({ edges }: Props) {
  const usDateFormat = new Intl.DateTimeFormat('en-US', { dateStyle: 'medium', timeStyle: 'short' });

  return (
    <div className="px-5 py-4 flex bg-white w-full justify-between">
      <div className="flex gap-2">
        <UserPopover edges={ edges } />
        <div className="flex flex-col">
          <h3 className="font-semibold">{chooseName(edges)}</h3>
          <p className="text-gray-200">{usDateFormat.format(new Date(edges.detail.transactedAt))}</p>
        </div>
      </div>
      <div className="flex flex-col items-end">
        {edges.detail.amount > 0 ? (
          <p className="text-green">
            +
            {usFormat.format(edges.detail.amount)}
          </p>
        ) : (
          <p className="text-red">
            {(usFormat.format(edges.detail.amount))}
          </p>
        )}
        <p className="text-gray-200">{makeCapitalized(edges.detail.type)}</p>
      </div>
    </div>
  );
}

export default Transaction;
