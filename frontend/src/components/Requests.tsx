import { useQuery } from '@tanstack/react-query';
import { useAtom } from 'jotai';

import axios from 'axios';
import { tokenAtom } from '../store/atom';
import { apiURL, queryParams } from '../consts';

import Line from './Line';
import Request from './Request';
import Warning from '../assets/warning.svg';
import { usFormat } from '../utils/helpers';

function Requests() {
  const [token] = useAtom(tokenAtom);
  const { data, refetch } = useQuery({
    queryKey: ['requests'],
    queryFn: () => axios.get(`${apiURL}/api/deposit-request`, queryParams(token)),
    select: ({ data: { requests } }) => requests,
  });

  return (
    <div
      className="relative flex flex-col gap-4 w-screen h-auto min-h-dvh mt-6 pt-12
      bg-gray-100 rounded-t-[40px]"
    >
      <Line />
      {
        data?.length === 0 ? (
          <div className="text-2xl text-center font-bold text-gray-200">
            <p className="my-8">No requests</p>
            <img src={ Warning } alt="warning symbol" className="w-12 h-12 mx-auto" />
          </div>
        ) : (
          data?.map((request: any) => (
            <Request
              key={ request.id }
              account={ request.edges.account.number }
              id={ request.id }
              name={ `${request.edges.account.edges.user.firstName} ${request.edges.account.edges.user.lastName}` }
              profileURL={ request.edges.account.edges.user.pictureUrl }
              date={ usFormat.format(request.amount) }
              status={ request.status }
              refetch={ refetch }
            />
          ))

        )
      }
    </div>
  );
}

export default Requests;
