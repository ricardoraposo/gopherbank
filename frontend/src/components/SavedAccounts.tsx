import { useAtom } from 'jotai';
import { useQuery } from '@tanstack/react-query';
import axios from 'axios';
import TProfilePic from './TProfilePic';
import { accountNumberAtom, selectedAtom, tokenAtom } from '../store/atom';
import { apiURL, queryParams } from '../consts';

function SavedAccounts() {
  const [token] = useAtom(tokenAtom);
  const [selected, setSelected] = useAtom(selectedAtom);
  const [, setAccountNumber] = useAtom(accountNumberAtom);
  const { data } = useQuery({
    queryKey: ['favorites', token],
    queryFn: async () => {
      return axios.get(`${apiURL}/api/favorite/`, queryParams(token));
    },
    select: ({ data: { favorites } }) => favorites,
  });

  const handleSelect = (account: string) => {
    setSelected(account);
    setAccountNumber(account);
  };

  return (
    <div className="mt-4 flex flex-col items-center gap-2">
      <div className="w-4/5 flex flex-col items-start gap-3">
        <h2 className="text-white">Saved Accounts</h2>
        <div className="flex gap-3">
          {data?.length > 0 ? data.map((favorites: any) => (
            <button
              key={ favorites.number }
              onClick={ () => handleSelect(favorites.number) }
            >
              <div
                className={ `border-4 rounded-full transition-colors ${selected === favorites.number ? 'border-white' : 'border-gray-600'}` }
              >
                <TProfilePic
                  profileURL={ favorites.edges.user.pictureUrl }
                  className="w-16 h-16"
                />
              </div>
            </button>
          )) : <p className="text-gray-400">No saved accounts</p>}
        </div>
      </div>
      <div className="w-screen border-t border-gray-500 mt-2" />
    </div>
  );
}

export default SavedAccounts;
