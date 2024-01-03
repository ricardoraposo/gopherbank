import { useRef, useState } from 'react';
import { useAtom } from 'jotai';
import axios from 'axios';
import { useMutation, useQuery, useQueryClient } from '@tanstack/react-query';
import { chooseAccount, choosePicture } from '../utils/transactionHelpers';
import TProfilePic from './TProfilePic';
import { accountAtom, tokenAtom } from '../store/atom';
import { apiURL, queryParams } from '../consts';

type Props = {
  edges: any;
};

function UserPopover({ edges }: Props) {
  const [id] = useAtom(accountAtom);
  const [token] = useAtom(tokenAtom);
  const [show, setShow] = useState(false);

  const account = chooseAccount(edges);

  const queryClient = useQueryClient();
  const { data: favorites } = useQuery({
    queryKey: ['favorites'],
    queryFn: async () => {
      return axios.get(`${apiURL}/api/favorite/`, queryParams(token));
    },
    select: ({ data: { favorites } }) => favorites,
  });

  const favoriteMutation = useMutation({
    mutationFn: async (accountNumber) => {
      return axios.post(`${apiURL}/api/favorite/`, { favoritedID: accountNumber }, queryParams(token));
    },
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['favorites'] });
    },
  });

  const isFriend = favorites?.find((acc: any) => acc.number === account.number);

  return (
    <div className="relative">
      <button
        onClick={ () => setShow((prev) => !prev) }
      >
        <TProfilePic profileURL={ choosePicture(edges) } className="w-10 h-10" />
      </button>
      <div
        onMouseLeave={ () => setShow(false) }
        className={ `absolute w-72 h-72 bg-gray-600 -top-72 left-10 rounded-t-3xl rounded-br-3xl drop-shadow-md
        shadow-lg flex flex-col items-center justify-center px-8 py-4 origin-bottom-left z-50 overflow-hidden
        ${show ? 'scale-100' : 'scale-0'} transition-all duration-300 ease-in-out will-change-transform` }
      >
        <div className="absolute -top-10 bg-orange h-28 w-full z-20 rounded-2xl" />
        <img
          src={ account.edges.user.pictureUrl }
          alt="user profile pic"
          className="w-20 h-20 rounded-full z-30 border-4 border-white"
        />
        <div className="mt-5">
          <p className="text-gray-100 text-xl text-center">{`${account.edges.user.firstName} ${account.edges.user.lastName}`}</p>
          <p className="text-gray-200 text-center">{`${account.edges.user.email}`}</p>
          <p className="text-white text-center text-xl font-bold mt-2">{`${account.number}`}</p>
        </div>
        {
          id !== account.number && (
            <button
              className="text-white bg-orange rounded-3xl py-2 px-6 mt-4 active:scale-90 active:bg-pink-500 transition-all duration-300 ease-in-out"
              onClick={ () => favoriteMutation.mutate(account.number) }
            >
              {isFriend ? 'Remove account' : 'Save account'}
            </button>
          )
        }
      </div>
    </div>
  );
}

export default UserPopover;
