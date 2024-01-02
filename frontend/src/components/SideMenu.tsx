import { useQuery } from '@tanstack/react-query';
import { useAtom } from 'jotai';
import axios from 'axios';
import { useNavigate } from 'react-router-dom';
import { motion } from 'framer-motion';
import Division from './Division';
import { accountAtom, showMenuAtom, tokenAtom } from '../store/atom';
import { apiURL, queryParams } from '../consts';

function SideMenu() {
  const [id] = useAtom(accountAtom);
  const [token, setToken] = useAtom(tokenAtom);
  const [showMenu, setShowMenu] = useAtom(showMenuAtom);
  const navigate = useNavigate();
  const { data: admin } = useQuery({
    queryKey: ['user'],
    queryFn: () => axios.get(`${apiURL}/api/accounts/${id}`, queryParams(token)),
    select: ({ data: { admin } }) => admin,
  });

  const handleLogout = () => {
    setToken('');
    setShowMenu(false);
    navigate('/signin');
  };

  const goToDashboard = () => {
    navigate('/');
    setShowMenu(false);
  };

  const goToAdmin = () => {
    navigate('/admin');
    setShowMenu(false);
  };

  return (
    <motion.div
      className={ `fixed z-40 bg-black opacity-90 h-dvh w-full top-0 right-0 origin-top-right transition-all
      ${showMenu ? '' : 'translate-x-full'}` }
    >
      <ul
        className="flex flex-col items-center gap-12 mt-20
        text-white text-2xl font-bold"
      >
        <li>
          <button
            onClick={ goToDashboard }
          >
            Dashboard
          </button>
        </li>
        <Division />
        <li>
          <button>Profile</button>
        </li>
        <Division />
        {admin && (
          <>
            <li>
              <button
                onClick={ goToAdmin }
              >
                Admin
              </button>
            </li>
            <Division />
          </>
        )}
        <li>
          <button onClick={ handleLogout }>Logout</button>
        </li>
      </ul>
    </motion.div>
  );
}

export default SideMenu;
