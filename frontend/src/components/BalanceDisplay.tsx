import StatisticIcon from '../assets/statistics.svg';
import DownArrow from '../assets/down_arrow.svg';

function BalanceDisplay() {
  return (
    <div className="flex justify-between mt-9">
      <div className="flex flex-col justify-between items-start h-20">
        <p className="text-grayish text-base font-medium">My Balance</p>
        <p className="text-white text-4xl font-bold">
          <span className="text-orange">$</span>
          5,332.18
        </p>
      </div>
      <div className="flex flex-col justify-between items-end">
        <button
          type="button"
          className="w-28 h-9 bg-gray text-sm text-white
          flex justify-center items-center gap-1 rounded-full"
        >
          <img src={ StatisticIcon } alt="statistics icon" />
          <p>Statistics</p>
        </button>
        <div className="flex mb-1">
          <p className="text-white text-base">USD</p>
          <img src={ DownArrow } alt="down arrow" />
        </div>
      </div>
    </div>
  );
}

export default BalanceDisplay;
