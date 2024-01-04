import axios from 'axios';
import { useQuery } from '@tanstack/react-query';
import { useAtom } from 'jotai';

import { apiURL, queryParams } from '../consts';
import { showNotificationAtom, tokenAtom } from '../store/atom';

import Loading from './Loading';
import Notification from './Notification';
import WarningIcon from '../assets/warning.svg';

function NotificationMenu() {
  const [token] = useAtom(tokenAtom);
  const [show] = useAtom(showNotificationAtom);
  const { data, isLoading } = useQuery({
    queryKey: ['notifications'],
    queryFn: () => axios.get(`${apiURL}/api/notification`, queryParams(token)),
    select: ({ data: { notifications } }) => notifications,
  });

  if (isLoading) return <Loading />;

  return (
    <div
      className={ `absolute w-72 h-72 bg-gray-100 -bottom-[18.5rem] -left-52 rounded-3xl shadow-xl
      flex flex-col py-4 gap-4 overflow-scroll ${show ? 'scale-100' : 'scale-0'} transition-all origin-[80%_0%]` }
    >
      {
        data ? (
          data?.map((notification: any) => (
            <Notification
              key={ notification.id }
              id={ notification.id }
              title={ notification.title }
              content={ notification.content }
            />
          ))
        ) : (
          <div className="text-2xl text-center font-bold text-gray-200">
            <p className="my-8">No notification</p>
            <img src={ WarningIcon } alt="warning symbol" className="w-12 h-12 mx-auto" />
          </div>
        )
      }
    </div>
  );
}

export default NotificationMenu;
