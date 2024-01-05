import axios from 'axios';
import { useEffect } from 'react';
import { useAtom } from 'jotai';
import { motion } from 'framer-motion';
import { useQuery } from '@tanstack/react-query';
import { useLocation, useNavigate } from 'react-router-dom';

import { apiURL } from '../consts';
import { accountAtom, tokenAtom } from '../store/atom';

import SavedAccounts from '../components/SavedAccounts';
import AmountDisplay from '../components/AmountDisplay';
import SendBtn from '../components/SendBtn';
import Loading from '../components/Loading';

function Operation() {
  const navigate = useNavigate();
  const { pathname } = useLocation();
  const [, setToken] = useAtom(tokenAtom);
  const [, setID] = useAtom(accountAtom);
  const { isLoading, error } = useQuery({
    queryKey: ['user'],
    queryFn: async () => {
      const { data: { number } } = await axios.get(`${apiURL}/api/jwt/`);
      setID(number);
      return axios.get(`${apiURL}/api/accounts/${number}`);
    },
    retry: 2,
  });

  useEffect(() => {
    if (axios.isAxiosError(error) && error.response?.status === 401) {
      setToken('');
      navigate('/signin');
    }
  }, []);

  if (isLoading) return <Loading />;

  return (
    <motion.div
      className="w-full h-dvh flex flex-col justify-between items-center"
      initial={ { x: 300, opacity: 0 } }
      animate={ { x: 0, opacity: 1 } }
      exit={ { x: -300, opacity: 0, transition: { duration: 0.1 } } }
    >
      {pathname.includes('transfer') ? <SavedAccounts /> : <div />}
      <AmountDisplay />
      <SendBtn />
    </motion.div>
  );
}

export default Operation;
