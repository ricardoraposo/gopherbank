import TransferIcon from '../assets/transfer.svg';
import OpButton from './OpButton';
import DepositIcon from '../assets/deposit.svg';
import WithdrawIcon from '../assets/withdraw.svg';

function OperationSection() {
  return (
    <div className="w-full flex justify-around mt-7">
      <OpButton label="Transfer" icon={ TransferIcon } />
      <OpButton label="Deposit" icon={ DepositIcon } />
      <OpButton label="Withdraw" icon={ WithdrawIcon } className="w-8 h-8" />
    </div>
  );
}

export default OperationSection;
