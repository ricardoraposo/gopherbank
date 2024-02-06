import { useAtom } from 'jotai';
import { amountAtom } from '../store/atom';

function AmountInput() {
  const [amount, setAmount] = useAtom(amountAtom);
  return (
    <input
      inputMode="decimal"
      type="number"
      value={ amount }
      name="amount"
      onChange={ (e) => setAmount(e.target.value) }
      placeholder="0.00"
      className="text-4xl text-white text-center font-semibold w-4/5 bg-transparent placeholder:text-gray"
    />
  );
}

export default AmountInput;
