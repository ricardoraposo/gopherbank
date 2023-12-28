import vector from '../assets/vector.svg';
import { profilepic1, profilepic2, profilepic3, profilepic4 } from '../consts';
import TProfilePic from './TProfilePic';

function RecentTransactions() {
  return (
    <div
      className="h-28 text-white bg-orange rounded-3xl px-5
      relative flex flex-col grow justify-center items-center gap-4"
    >
      <img src={ vector } alt="little thingy" className="absolute right-2 top-0 h-14" />
      <div>
        <h2 className="text-base font-normal">Recent Transactions</h2>
        <div className="flex gap-2">
          <TProfilePic profileURL={ profilepic1 } />
          <TProfilePic profileURL={ profilepic2 } />
          <TProfilePic profileURL={ profilepic3 } />
          <TProfilePic profileURL={ profilepic4 } />
          <TProfilePic profileURL={ profilepic1 } />
        </div>
      </div>
    </div>
  );
}

export default RecentTransactions;
