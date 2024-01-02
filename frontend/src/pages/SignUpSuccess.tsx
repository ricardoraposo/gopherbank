import { motion } from 'framer-motion';
import { useNavigate } from 'react-router-dom';
import { useAtom } from 'jotai';
import SuccessImg from '../assets/success.png';
import { accountAtom } from '../store/atom';

function SignUpSuccess() {
  const navigate = useNavigate();
  const [account] = useAtom(accountAtom);

  const handleClick = () => {
    navigate('/signin');
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
      <p className="w-[338px] text-white text-center text-2xl font-bold mt-5">
        Account created successfully!!!
      </p>
      <div
        className="flex flex-col justify-center items-center
        border border-gray-500 w-[338px] h-32 rounded-xl mt-5"
      >
        <p className="text-3xl text-orange font-bold">
          {account}
        </p>
        <p className="text-white text-sm">
          ↑ Account number  ↑
        </p>
      </div>
      <button
        className="w-52 h-14 font-bold text-white rounded-full bg-orange mt-8"
        onClick={ handleClick }
      >
        Sign In
      </button>
    </motion.div>
  );
}

export default SignUpSuccess;
