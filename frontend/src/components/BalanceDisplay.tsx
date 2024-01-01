import { useQuery } from '@tanstack/react-query';
import { useAtom } from 'jotai';
import StatisticIcon from '../assets/statistics.svg';
import DownArrow from '../assets/down_arrow.svg';
import instance from '../api/axiosIstance';
import { apiURL, queryParams } from '../consts';
import { accountAtom, tokenAtom } from '../store/atom';
import { usFormat } from '../utils/helpers';

function BalanceDisplay() {
  const [id] = useAtom(accountAtom);
  const [token] = useAtom(tokenAtom);

  const { data } = useQuery({
    queryKey: ['user', token],
    queryFn: () => instance.get(`${apiURL}/api/accounts/${id}`, queryParams(token)),
    select: ({ data }) => data,
  });

  return (
    <div className="flex justify-between mt-9">
      <div className="flex flex-col justify-between items-start h-20">
        <p className="text-gray-200 text-base font-medium">My Balance</p>
        <p className="text-white text-4xl font-bold">
          <span className="text-orange">$</span>
          {data ? usFormat.format(data?.balance).replace('$', '') : ''}
        </p>
      </div>
      <div className="flex flex-col justify-between items-end">
        <button
          type="button"
          className="w-28 h-9 bg-gray-500 text-sm text-white
          flex justify-center items-center gap-1 rounded-full"
        >
          <img src={ StatisticIcon } alt="statistics icon" />
          <p>Statistics</p>
        </button>
        <div className="flex mb-1">
          <p className="text-white text-base">USD</p>
          <img src={ DownArrow } alt="down arrow" />
        </div>
      </div>
    </div>
  );
}

export default BalanceDisplay;
