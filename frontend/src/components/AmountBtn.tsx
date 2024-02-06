import { useAtom } from 'jotai';
import { amountAtom } from '../store/atom';

type Props = {
  value: string;
};

function AmountBtn({ value }: Props) {
  const [, setAmount] = useAtom(amountAtom);

  return (
    <button
      className="bg-gray-300 font-bold w-[3.25rem] h-8 text-white rounded-full"
      type="button"
      onClick={ () => setAmount(`${value}.00`) }
    >
      $
      {value}
    </button>
  );
}

export default AmountBtn;
