import axios from 'axios';
import { motion } from 'framer-motion';
import { useAtom } from 'jotai';
import { useNavigate, useParams } from 'react-router-dom';
import { useQuery } from '@tanstack/react-query';

import SuccessImg from '../assets/success.png';
import { apiURL, queryParams } from '../consts';
import { accountNumberAtom, amountAtom, tokenAtom } from '../store/atom';
import { getHourAndMinutes, usFormat } from '../utils/helpers';

function Success() {
  const { type } = useParams();
  const navigate = useNavigate();
  const [accountNumber] = useAtom(accountNumberAtom);
  const [amount] = useAtom(amountAtom);
  const [token] = useAtom(tokenAtom);

  const { data } = useQuery({
    queryKey: ['toUser', token],
    queryFn: () => axios.get(`${apiURL}/api/user/${accountNumber}`, queryParams(token)),
    select: ({ data: { user } }) => user,
  });

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
      <p className="w-[338px] text-white text-2xl font-bold mt-5">Done, operation concluded successfully!!!</p>
      <div
        className="flex flex-col justify-center items-center
        border border-gray-500 w-[338px] h-32 rounded-xl mt-5"
      >
        <p className="text-3xl text-white font-bold">
          {usFormat.format(parseFloat(amount)) && '$100.00'}
        </p>
        <p className="text-white text-sm">
          To
          {' '}
          {type === 'transfer' ? `${data?.firstName} ${data?.lastName}` : 'my wallet'}
        </p>
        <p className="text-white text-sm">
          Right now •
          {' '}
          {getHourAndMinutes()}
        </p>
      </div>
      <button
        className="w-52 h-14 font-bold text-white rounded-full bg-orange mt-8"
        onClick={ () => navigate('/') }
      >
        Back to Dashboard
      </button>
    </motion.div>
  );
}

export default Success;
