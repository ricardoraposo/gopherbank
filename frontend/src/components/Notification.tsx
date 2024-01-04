import { useMutation, useQueryClient } from '@tanstack/react-query';
import axios from 'axios';
import { useAtom } from 'jotai';
import Cancel from '../assets/cancel.svg';
import { apiURL, queryParams } from '../consts';
import { tokenAtom } from '../store/atom';

type Props = {
  id: number;
  title: string;
  content: string;
};

function Notification({ id, title, content }: Props) {
  const [token] = useAtom(tokenAtom);
  const queryClient = useQueryClient();
  const notification = useMutation({
    mutationFn: (id: number) => axios.delete(`${apiURL}/api/notification/${id}`, queryParams(token)),
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ['notifications'] });
    },
  });

  return (
    <div className="px-6 py-2 flex justify-center items-center bg-white">
      <div className="flex flex-col items-start">
        <p className="text-gray-400 text-base font-semibold">{title}</p>
        <p className="text-gray-800 text-xs text-start">{content}</p>
      </div>
      <button onClick={ () => notification.mutate(id) }>
        <img src={ Cancel } alt="clean notification" />
      </button>
    </div>
  );
}

export default Notification;
