import BalanceDisplay from '../components/BalanceDisplay';
import Header from '../components/Header';
import InvestBanner from '../components/InvestBanner';
import OperationSection from '../components/OperationSection';
import RecentTransactions from '../components/RecentTransactions';
import Transactions from '../components/Transactions';
import instance from '../api/axiosIstance';
import { useNavigate } from 'react-router-dom';
import { useQuery } from '@tanstack/react-query';


function Dashboard() {
  const navigate = useNavigate();
  const id = "08533436"
  const { isError, isLoading } = useQuery({
    queryKey: ["user"],
    queryFn: () => instance.get(`/api/accounts/${id}`),
  })

  if (isError) {
    navigate('/signin');
  }

  if (isLoading) return <div>I'm loading!!!</div>

  return (
    < div >
      <div className="mx-5 my-2">
        <Header id={id} />
        <BalanceDisplay id={id} />
        <OperationSection />
        <div className="flex gap-4 mt-6">
          <InvestBanner />
          <RecentTransactions id={id} />
        </div>
      </div>
      <Transactions id={id} />
    </div >
  );
}

export default Dashboard;
