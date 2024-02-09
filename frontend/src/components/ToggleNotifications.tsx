import axios from 'axios';
import { useAtom } from 'jotai';
import { useQuery } from '@tanstack/react-query';

import { showNotificationAtom, tokenAtom } from '../store/atom';
import { apiURL, queryParams } from '../consts';

import Bell from '../assets/notification.svg';
import NotificationMenu from './NotificationMenu';

function ToggleNotifications() {
  const [, setShow] = useAtom(showNotificationAtom);
  const [token] = useAtom(tokenAtom);
  const { data } = useQuery({
    queryKey: ['notifications', token],
    queryFn: () => axios.get(`${apiURL}/api/notification`, queryParams(token)),
    select: ({ data: { notifications } }) => notifications,
  });

  return (
    <div
      className="relative bg-gray-500 w-11 h-11 flex justify-center items-center rounded-full z-30"
    >
      <button onClick={ () => setShow((prev) => !prev) }>
        <img src={ Bell } alt="notifications" />
      </button>
      <div
        className={ `absolute text-xs w-5 h-5 bg-red top-0 -right-2 rounded-full flex
        justify-center items-center active:animate-pulse ${data?.length > 0 ? 'block' : 'hidden'}` }
      >
        {data?.length}
      </div>
      <NotificationMenu />
    </div>
  );
}

export default ToggleNotifications;
