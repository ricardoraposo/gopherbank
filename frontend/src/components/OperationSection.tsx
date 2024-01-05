import { useNavigate } from 'react-router-dom';
import TransferIcon from '../assets/transfer.svg';
import OpButton from './OpButton';
import DepositIcon from '../assets/deposit.svg';
import WithdrawIcon from '../assets/withdraw.svg';

function OperationSection() {
  const navigate = useNavigate();

  return (
    <div className="w-full flex justify-around mt-7">
      <OpButton onClick={ () => navigate('/operation/deposit') } label="Deposit" icon={ DepositIcon } />
      <OpButton onClick={ () => navigate('/operation/transfer') } label="Transfer" icon={ TransferIcon } />
      <OpButton onClick={ () => navigate('/operation/withdraw') } label="Withdraw" icon={ WithdrawIcon } className="w-8 h-8" />
    </div>
  );
}

export default OperationSection;
