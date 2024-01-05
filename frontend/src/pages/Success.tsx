import axios from 'axios';
import { motion } from 'framer-motion';
import { useAtom } from 'jotai';
import { useNavigate, useParams } from 'react-router-dom';
import { useQuery } from '@tanstack/react-query';

import SuccessImg from '../assets/success.png';
import { apiURL, queryParams } from '../consts';
import { accountNumberAtom, amountAtom, selectedAtom, tokenAtom } from '../store/atom';
import { getHourAndMinutes, usFormat } from '../utils/helpers';

function Success() {
  const { type } = useParams();
  const navigate = useNavigate();
  const [accountNumber, setAccountNumber] = useAtom(accountNumberAtom);
  const [amount, setAmount] = useAtom(amountAtom);
  const [, setSelected] = useAtom(selectedAtom);
  const [token] = useAtom(tokenAtom);

  const { data } = useQuery({
    queryKey: ['toUser', token],
    queryFn: () => axios.get(`${apiURL}/api/user/${accountNumber}`, queryParams(token)),
    select: ({ data: { user } }) => user,
    enabled: type === 'transfer',
  });

  const handleClick = () => {
    setAccountNumber('');
    setSelected('');
    setAmount('');
    navigate('/');
  };

  return (
    <motion.div
      initial={ { x: 300, opacity: 0 } }
      animate={ { x: 0, opacity: 1 } }
      exit={ { x: -300, opacity: 0, transition: { duration: 0.1 } } }
      className="h-dvh flex flex-col items-center justify-center"
    >
      <img
        src={ SuccessImg }
        alt="gopher in a party, celebrating the good transaction"
        className="w-44 h-44"
      />
      {
        type === 'deposit' ? (
          <div className="w-[338px]  mt-5 text-white">
            <p className="text-xl font-bold">
              Done, your deposit request has been sent!!!
            </p>
            <p className="text-lg text-white mt-3">
              In a few hours, you should get confirmation of the deposit
            </p>
          </div>
        ) : (
          <p className="w-[338px] text-white text-2xl font-bold mt-5">
            Done, operation concluded successfully!!!
          </p>
        )
      }
      <div
        className="flex flex-col justify-center items-center
        border border-gray-500 w-[338px] h-32 rounded-xl mt-5"
      >
        <p className="text-3xl text-white font-bold">
          {usFormat.format(parseFloat(amount))}
        </p>
        <p className="text-white text-sm">
          {type === 'transfer' ? `To ${data?.firstName} ${data?.lastName}` : `${type === 'deposit' ? 'To' : 'From'} my wallet`}
        </p>
        <p className="text-white text-sm">
          Right now •
          {' '}
          {getHourAndMinutes()}
        </p>
      </div>
      <button
        className="w-52 h-14 font-bold text-white rounded-full bg-orange mt-8"
        onClick={ handleClick }
      >
        Back to Dashboard
      </button>
    </motion.div>
  );
}

export default Success;
