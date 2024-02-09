import { useQuery } from '@tanstack/react-query';
import axios from 'axios';
import { useAtom } from 'jotai';
import { apiURL, queryParams } from '../consts';
import { accountAtom, tokenAtom } from '../store/atom';
import Loading from './Loading';
import ToggleMenu from './ToggleMenu';
import ToggleNotifications from './ToggleNotifications';

function Header() {
  const [id] = useAtom(accountAtom);
  const [token] = useAtom(tokenAtom);
  const { data, isLoading } = useQuery({
    queryKey: ['user'],
    queryFn: () => axios.get(`${apiURL}/api/accounts/${id}`, queryParams(token)),
    select: ({ data }) => data,
  });

  if (isLoading) return <Loading />;

  return (
    <header className="flex justify-between items-center">
      <div className="flex gap-3">
        <div>
          <img
            src={ data ? data?.edges.user.pictureUrl : '' }
            alt="profile"
            className="h-11 w-11 object-cover rounded-full"
          />
        </div>
        <div>
          <p className="text-gray-200 text-sm font-medium">Hi, welcome</p>
          <p className="text-white text-lg font-semibold">{data ? data?.edges.user.firstName : ''}</p>
        </div>
      </div>
      <div className="flex gap-6">
        <ToggleNotifications />
        <ToggleMenu />
      </div>
    </header>
  );
}

export default Header;
