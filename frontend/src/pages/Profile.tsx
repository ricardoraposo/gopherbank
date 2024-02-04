import axios from 'axios';
import { motion } from 'framer-motion';
import { useAtom } from 'jotai';
import { useQuery } from '@tanstack/react-query';
import { useEffect } from 'react';
import { useNavigate } from 'react-router-dom';

import type { User } from '../types';
import { apiURL, queryParams } from '../consts';
import { accountAtom, showNotificationAtom, tokenAtom } from '../store/atom';

import Blur from '../components/Blur';
import Line from '../components/Line';
import Edit from '../assets/edit.svg';
import Arrow from '../assets/arrow_left.svg';
import Loading from '../components/Loading';
import SideMenu from '../components/SideMenu';
import ToggleMenu from '../components/ToggleMenu';
import ToggleNotifications from '../components/ToggleNotifications';
import { usDateFormat } from '../utils/helpers';
import UserDetail from '../components/UserDetail';

function Profile() {
  const [token, setToken] = useAtom(tokenAtom);
  const [, setID] = useAtom(accountAtom);
  const [show] = useAtom(showNotificationAtom);
  const navigate = useNavigate();

  const { data: user, isLoading, error } = useQuery({
    queryKey: ['user', token],
    queryFn: async () => {
      const { data } = await axios.get(`${apiURL}/api/jwt/`, queryParams(token));
      setID(data.number);
      return axios.get(`${apiURL}/api/accounts/${data.number}`, queryParams(token));
    },
    select: ({ data }) => {
      return { number: data.number, join: data.createdAt, ...data.edges.user } as User;
    },
    retry: 2,
  });

  useEffect(() => {
    if (axios.isAxiosError(error) && error.response?.status === 401) {
      setToken('');
      navigate('/signin');
    }
  }, [error]);

  if (isLoading) return <Loading />;

  return (
    <motion.div
      className="pt-2"
      initial={ { x: 300, opacity: 0 } }
      animate={ { x: 0, opacity: 1 } }
      exit={ { x: -300, opacity: 0, transition: { duration: 0.1 } } }
    >
      {show && <Blur />}
      <SideMenu />
      <header className="flex justify-between items-center mx-5">
        <button
          onClick={ () => navigate('/') }
          className="bg-gray-500 w-11 h-11 flex justify-center items-center rounded-full z-30"
        >
          <img src={ Arrow } alt="go back" />
        </button>
        <div className="flex gap-6">
          <ToggleNotifications />
          <ToggleMenu />
        </div>
      </header>
      <div className="flex flex-col items-center gap-5 mt-24 mb-8">
        <img src={ user?.pictureUrl } alt="user profile" className="w-52 h-52 rounded-full" />
        <p className="text-white text-lg font-semibold">{`${user?.firstName} ${user?.lastName}`}</p>
      </div>
      <div
        className="relative w-screen h-auto min-h-[60dvh] mt-6 pt-16 bg-slate-50 rounded-t-[40px] overflow-hidden"
      >
        <Line />
        <button
          className="absolute top-2 right-4 bg-gray-100 w-11 h-11
          flex justify-center items-center rounded-full"
          onClick={ () => navigate('/profile/edit') }
        >
          <img src={ Edit } alt="edit" />
        </button>
        {
          user && (
            <div className="flex flex-col gap-4">
              <UserDetail label="Account Number" detail={ user?.number } />
              <UserDetail label="Email" detail={ user?.email } />
              <UserDetail label="Joined at" detail={ usDateFormat.format(new Date(user?.join)) } />
            </div>
          )
        }
      </div>
    </motion.div>
  );
}

export default Profile;
