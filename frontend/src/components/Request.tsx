import axios from 'axios';
import { useAtom } from 'jotai';
import { useMutation, useQueryClient } from '@tanstack/react-query';
import TProfilePic from './TProfilePic';
import Check from '../assets/checkmark.svg';
import Cancel from '../assets/cancel.svg';
import { makeCapitalized } from '../utils/transactionHelpers';
import { apiURL, queryParams } from '../consts';
import { tokenAtom } from '../store/atom';

type Props = {
  name: string;
  id: number;
  account: string;
  profileURL: string;
  status: 'approved' | 'pending' | 'rejected';
  date: string;
};

type Params = {
  id: number;
  account: string;
};

function Request({ id, account, name, profileURL, date, status }: Props) {
  const [token] = useAtom(tokenAtom);
  const queryClient = useQueryClient();
  const approvalMutation = useMutation({
    mutationFn: ({ id, account }: Params) => {
      return axios.post(`${apiURL}/api/deposit-request/approve/${id}`, { account }, queryParams(token));
    },
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['requests'] });
    },
  });

  const rejectMutation = useMutation({
    mutationFn: (id: number) => {
      return axios.patch(`${apiURL}/api/deposit-request/reject/${id}`, {}, queryParams(token));
    },
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['requests'] });
    },
  });

  return (
    <div className="px-5 py-4 flex bg-white w-full justify-between">
      <div className="flex gap-2">
        <TProfilePic profileURL={ profileURL } className="h-10 w-10" />
        <div className="flex flex-col">
          <h3 className="font-semibold">{name}</h3>
          <p className="text-gray-200">{date}</p>
        </div>
      </div>
      <div className="flex flex-col items-end justify-center">
        {
          status === 'pending' ? (
            <div className="flex gap-6">
              <button
                className="flex justify-center items-center
            bg-red w-9 h-9 text-center text-lg rounded-full"
                onClick={ () => rejectMutation.mutate(id) }
              >
                <img src={ Cancel } alt="cancel" />
              </button>
              <button
                className="flex justify-center items-center
            bg-green w-9 h-9 text-center text-lg rounded-full"
                onClick={ () => approvalMutation.mutate({ id, account }) }
              >
                <img src={ Check } alt="checkmark" />
              </button>
            </div>
          ) : (
            <p className={ status === 'approved' ? 'text-green' : 'text-red' }>
              {makeCapitalized(status)}
            </p>
          )
        }
      </div>
    </div>
  );
}

export default Request;
