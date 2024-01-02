import axios from 'axios';
import { motion } from 'framer-motion';
import { useAtom } from 'jotai';
import { useQuery } from '@tanstack/react-query';
import { useEffect } from 'react';
import { useNavigate } from 'react-router-dom';
import { apiURL, queryParams } from '../consts';
import { accountAtom, tokenAtom } from '../store/atom';

import Header from '../components/Header';
import Loading from '../components/Loading';
import InvestBanner from '../components/InvestBanner';
import Transactions from '../components/Transactions';
import BalanceDisplay from '../components/BalanceDisplay';
import OperationSection from '../components/OperationSection';
import RecentTransactions from '../components/RecentTransactions';
import SideMenu from '../components/SideMenu';

function Dashboard() {
  const navigate = useNavigate();
  const [, setID] = useAtom(accountAtom);
  const [token, setToken] = useAtom(tokenAtom);
  const { isLoading, error } = useQuery({
    queryKey: ['user', token],
    queryFn: async () => {
      const { data: { number } } = await axios.get(`${apiURL}/api/jwt/`, queryParams(token));
      setID(number);
      return axios.get(`${apiURL}/api/accounts/${number}`, queryParams(token));
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
      initial={ { x: 300, opacity: 0 } }
      animate={ { x: 0, opacity: 1 } }
      exit={ { x: -300, opacity: 0, transition: { duration: 0.1 } } }
    >
      <SideMenu />
      <div className="mx-5 py-2">
        <Header />
        <BalanceDisplay />
        <OperationSection />
        <div className="flex gap-4 mt-6">
          <InvestBanner />
          <RecentTransactions />
        </div>
      </div>
      <Transactions />
    </motion.div>
  );
}

export default Dashboard;
