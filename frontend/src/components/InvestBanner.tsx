import { useNavigate } from 'react-router-dom';
import CreditCards from '../assets/credit_cards.svg';

function InvestBanner() {
  const navigate = useNavigate();

  return (
    <button
      className="h-28 w-24 py-4 bg-gray-500 flex flex-col justify-around items-center
      rounded-3xl"
      onClick={ () => navigate('/todo') }
    >
      <div
        className="bg-gray-400 w-12 h-12 rounded-full flex justify-center items-center"
      >
        <img src={ CreditCards } alt="credit cards" className="w-8 h-8" />
      </div>
      <h2 className="text-white">Cards</h2>
    </button>
  );
}

export default InvestBanner;
