import { useQuery } from '@tanstack/react-query';
import { useAtom } from 'jotai';
import axios from 'axios';
import Division from './Division';
import { accountAtom, showMenuAtom, tokenAtom } from '../store/atom';
import { apiURL, queryParams } from '../consts';

function SideMenu() {
  const [id] = useAtom(accountAtom);
  const [token] = useAtom(tokenAtom);
  const [showMenu] = useAtom(showMenuAtom);
  const { data: admin } = useQuery({
    queryKey: ['user'],
    queryFn: () => axios.get(`${apiURL}/api/accounts/${id}`, queryParams(token)),
    select: ({ data: { admin } }) => admin,
  });

  return (
    <div
      className={ `fixed z-40 bg-black h-dvh w-full top-0 right-0 opacity-90 transition-all
      ${showMenu ? '' : 'translate-x-full'}` }
    >
      <ul
        className="flex flex-col items-center gap-12 mt-20
        text-white text-2xl font-bold"
      >
        <li>
          <button>Profile</button>
        </li>
        <Division />
        {admin && (
          <>
            <li>Admin</li>
            <Division />
          </>
        )}
        <li>Logout</li>
      </ul>
    </div>
  );
}

export default SideMenu;
