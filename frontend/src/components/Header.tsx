import { profileUrl } from '../consts';
import BellIcon from '../assets/notification.svg';

function Header() {
  return (
    <header className="flex justify-between items-center">
      <div className="flex gap-3">
        <div>
          <img
            src={ profileUrl }
            alt="profile"
            className="h-11 w-11 object-cover rounded-full"
          />
        </div>
        <div>
          <p className="text-grayish text-sm font-medium">Hi, welcome</p>
          <p className="text-white text-lg font-semibold">Ricardo</p>
        </div>
      </div>
      <div className="bg-gray w-10 h-10 flex justify-center items-center rounded-full">
        <img src={ BellIcon } alt="notification icon" className="w-5 h-5 object-cover" />
      </div>
    </header>
  );
}

export default Header;
