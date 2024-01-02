import { useQuery } from '@tanstack/react-query';
import axios from 'axios';
import { useAtom } from 'jotai';
import { Outlet } from 'react-router-dom';
import { apiURL, queryParams } from '../consts';
import { accountAtom, showMenuAtom, tokenAtom } from '../store/atom';
import Loading from './Loading';
import ToggleMenu from './ToggleMenu';

function Header() {
  const [id] = useAtom(accountAtom);
  const [token] = useAtom(tokenAtom);
  const [, setShowMenu] = useAtom(showMenuAtom);
  const { data, isLoading } = useQuery({
    queryKey: ['user', token],
    queryFn: () => axios.get(`${apiURL}/api/accounts/${id}`, queryParams(token)),
    select: ({ data }) => data,
  });

  if (isLoading) return <Loading />;

  return (
    <>
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
        <button
          className="bg-gray-500 w-11 h-11 flex justify-center items-center rounded-full z-50"
          onClick={ () => setShowMenu((prev) => !prev) }
        >
          <ToggleMenu />
        </button>
      </header>
      <Outlet />
    </>
  );
}

export default Header;
