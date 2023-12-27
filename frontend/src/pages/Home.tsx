import { useEffect } from 'react';
import instance from '../api/axiosIstance';

const getAccounts = async () => {
  const response = await instance.get('test/accounts');
  console.log(response);
};

function Home() {
  useEffect(() => {
    getAccounts();
  }, []);

  return (
    <div className="h-screen w-screen">
      <div>
        <h1>Home</h1>
      </div>
    </div>
  );
}

export default Home;
