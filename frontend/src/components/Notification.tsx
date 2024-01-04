import { useMutation, useQueryClient } from '@tanstack/react-query';
import axios from 'axios';
import { useAtom } from 'jotai';
import { useState } from 'react';
import Cancel from '../assets/cancel.svg';
import { apiURL, queryParams } from '../consts';
import { tokenAtom } from '../store/atom';
import { sleep } from '../utils/helpers';

type Props = {
  id: number;
  title: string;
  content: string;
};

function Notification({ id, title, content }: Props) {
  const [token] = useAtom(tokenAtom);
  const [shrink, setShrink] = useState(false);
  const queryClient = useQueryClient();
  const notification = useMutation({
    mutationFn: (id: number) => axios.delete(`${apiURL}/api/notification/${id}`, queryParams(token)),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['notifications'] });
    },
  });

  const handleClick = async (id: number) => {
    setShrink(true);
    await sleep(500);
    notification.mutate(id);
  };

  return (
    <div
      className={ `px-6 py-2 flex justify-center items-center bg-white origin-top
                  transition-all duration-700 ${shrink ? 'scale-0' : ''}` }
    >
      <div className="flex flex-col items-start">
        <p className="text-gray-400 text-base font-semibold">{title}</p>
        <p className="text-gray-800 text-xs text-start">{content}</p>
      </div>
      <button onClick={ () => handleClick(id) }>
        <img src={ Cancel } alt="clean notification" />
      </button>
    </div>
  );
}

export default Notification;
