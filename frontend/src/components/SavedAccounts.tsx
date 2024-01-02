import { useAtom } from 'jotai';
import { useQuery } from '@tanstack/react-query';
import axios from 'axios';
import TProfilePic from './TProfilePic';
import { tokenAtom } from '../store/atom';
import { apiURL, queryParams } from '../consts';

function SavedAccounts() {
  const [token] = useAtom(tokenAtom);
  const { data } = useQuery({
    queryKey: ['favorites', token],
    queryFn: async () => {
      return axios.get(`${apiURL}/api/favorite/`, queryParams(token));
    },
    select: ({ data: { favorites } }) => favorites,
  });

  return (
    <div className="mt-4 flex flex-col items-center gap-2">
      <div className="w-4/5 flex flex-col items-start gap-3">
        <h2 className="text-white">Saved Accounts</h2>
        <div className="flex gap-3">
          {data ? data?.map((favorites: any) => (
            <TProfilePic
              key={ favorites.number }
              profileURL={ favorites.edges.user.pictureUrl }
              className="w-16 h-16"
            />
          )) : <div>No saved accounts</div>}
        </div>
      </div>
      <div className="w-screen border-t border-gray-500 mt-2" />
    </div>
  );
}

export default SavedAccounts;
