import BellIcon from '../assets/notification.svg';
import { useQuery } from '@tanstack/react-query';
import instance from '../api/axiosIstance';

type Props = {
  id: string;
}

function Header({ id }: Props) {
  const { data } = useQuery({
    queryKey: ["user"],
    queryFn: () => instance.get(`/api/accounts/${id}`),
    select: ({ data }) => data,
  })

  return (
    <header className="flex justify-between items-center">
      <div className="flex gap-3">
        <div>
          <img
            src={data.edges.user.pictureUrl}
            alt="profile"
            className="h-11 w-11 object-cover rounded-full"
          />
        </div>
        <div>
          <p className="text-gray-200 text-sm font-medium">Hi, welcome</p>
          <p className="text-white text-lg font-semibold">{data.edges.user.firstName}</p>
        </div>
      </div>
      <div className="bg-gray-500 w-10 h-10 flex justify-center items-center rounded-full">
        <img src={BellIcon} alt="notification icon" className="w-5 h-5 object-cover" />
      </div>
    </header>
  );
}

export default Header;
