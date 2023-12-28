import BalanceDisplay from '../components/BalanceDisplay';
import Header from '../components/Header';
import InvestBanner from '../components/InvestBanner';
import OperationSection from '../components/OperationSection';
import RecentTransactions from '../components/RecentTransactions';

function Dashboard() {
  return (
    <div className="mx-5 my-2">
      <Header />
      <BalanceDisplay />
      <OperationSection />
      <div className="flex gap-4 mt-6">
        <InvestBanner />
        <RecentTransactions />
      </div>
    </div>
  );
}

export default Dashboard;
