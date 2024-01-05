import axios from 'axios';
import { motion } from 'framer-motion';
import { useAtom } from 'jotai';
import { useEffect, useState } from 'react';
import { useNavigate, useParams } from 'react-router-dom';
import { apiURL } from '../consts';
import { accountNumberAtom, amountAtom } from '../store/atom';
import { usFormat } from '../utils/helpers';

function OperationAccount() {
  const navigate = useNavigate();
  const [errorMsg, setErrorMsg] = useState('');

  const [accountNumber, setAccountNumber] = useAtom(accountNumberAtom);
  const [amount] = useAtom(amountAtom);

  const { type } = useParams();

  const handleSend = async () => {
    try {
      const { data: { number } } = await axios.get(`${apiURL}/api/jwt/`);
      await axios.post(`${apiURL}/api/${type}`, {
        fromAccountNumber: number,
        toAccountNumber: accountNumber,
        amount: parseFloat(amount),
      });
      navigate('/operation/transfer/success');
    } catch (e: any) {
      setErrorMsg(e.response.data);
    }
  };

  useEffect(() => {
    if (amount === '') {
      navigate(`/operation/${type}`);
    }
  }, []);

  return (
    <motion.div
      className="flex flex-col h-dvh justify-between items-center"
      initial={ { x: 300, opacity: 0 } }
      animate={ { x: 0, opacity: 1 } }
      exit={ { x: -300, opacity: 0, transition: { duration: 0.1 } } }
    >
      <div className="flex flex-col items-center gap-4 mt-8">
        <h1 className="w-4/5 text-2xl text-white font-light">To which account would you like to transfer ?</h1>
        <div className="w-screen border-t border-gray-500" />
      </div>
      <div className="flex flex-col items-center gap-4">
        <input
          inputMode="decimal"
          type="number"
          value={ accountNumber }
          name="amount"
          maxLength={ 8 }
          onChange={ (e) => setAccountNumber(e.target.value) }
          placeholder="Type here"
          className={ `text-4xl text-white text-center font-semibold w-3/5 py-6 bg-transparent
        placeholder:text-gray placeholder:font-light
        ${errorMsg ? 'border-4 border-red rounded-full' : ''}` }
        />
        {errorMsg && <p className="text-xl text-red">{errorMsg}</p>}
      </div>
      <button
        className="bg-orange text-white text-lg font-semibold w-full h-24
        flex justify-center items-center cursor-pointer"
        onClick={ () => handleSend() }
      >
        Send
        {' '}
        {amount ? usFormat.format(parseFloat(amount)) : '$0.00'}
      </button>
    </motion.div>
  );
}

export default OperationAccount;
