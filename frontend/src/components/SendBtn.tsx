import { useAtom } from 'jotai';
import { useLocation, useNavigate, useParams } from 'react-router-dom';
import axios from 'axios';
import { amountAtom, tokenAtom } from '../store/atom';
import { apiURL, queryParams } from '../consts';
import { getVerbFromType, usFormat } from '../utils/helpers';

function SendBtn() {
  const { type } = useParams();
  const navigate = useNavigate();
  const { pathname } = useLocation();

  const [token] = useAtom(tokenAtom);
  const [amount] = useAtom(amountAtom);

  const handleWithdraw = async () => {
    try {
      const { data: { number } } = await axios.get(`${apiURL}/api/jwt/`, queryParams(token));
      await axios.post(`${apiURL}/api/withdraw`, {
        fromAccountNumber: number,
        amount: parseFloat(amount),
      }, queryParams(token));
      navigate('/operation/withdraw/success');
    } catch (e: any) {
      console.log(e.response.data);
    }
  };

  const handleDeposit = async () => {
    try {
      await axios.post(
        `${apiURL}/api/deposit-request`,
        { amount: parseFloat(amount) },
        queryParams(token),
      );
      navigate('/operation/deposit/success');
    } catch (e: any) {
      console.log(e.response.data);
    }
  };

  const handleSend = async () => {
    switch (type) {
      case 'withdraw':
        handleWithdraw();
        break;
      case 'deposit':
        handleDeposit();
        break;
      default:
        navigate(`${pathname}/account`);
    }
  };

  return (
    <button
      className="bg-orange text-white text-lg font-semibold w-full h-24
      flex justify-center items-center cursor-pointer"
      onClick={ handleSend }
    >
      {type && getVerbFromType(type)}
      {' '}
      {amount ? usFormat.format(parseFloat(amount)) : '$0.00'}
    </button>
  );
}

export default SendBtn;
