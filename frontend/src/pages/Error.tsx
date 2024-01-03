import { motion } from 'framer-motion';
import { useLocation, useNavigate } from 'react-router-dom';
import Unhauthorized from '../assets/unhauthorized.png';
import NotFound from '../assets/not_found.png';
import ToDo from '../assets/to_be_done.png';

function Error() {
  const navigate = useNavigate();
  const { pathname } = useLocation();

  const pickSrc = (pathname: string) => {
    switch (pathname) {
      case '/unauth':
        return Unhauthorized;
      case '/todo':
        return ToDo;
      default:
        return NotFound;
    }
  };

  const pickText = (pathname: string) => {
    switch (pathname) {
      case '/unauth':
        return (
          <div className="w-80 text-center my-4">
            <p className="font-bold text-orange text-2xl">Error 401</p>
            <p className="font-bold text-white text-2xl">
              You don’t have permission to be here
            </p>
          </div>
        );
      case '/todo':
        return (
          <div className="w-80 text-center my-4">
            <p className="font-bold text-orange text-2xl">Error "I'm lazy"</p>
            <p className="font-bold text-white text-2xl">
              This is yet to be done
            </p>
          </div>
        );
      default:
        return (
          <div className="w-80 text-center my-4">
            <p className="font-bold text-orange text-2xl">Error 404</p>
            <p className="font-bold text-white text-2xl">
              Page not found
            </p>
          </div>
        );
    }
  };

  return (
    <motion.div
      className="h-dvh w-screen flex flex-col justify-center items-center"
      initial={ { x: 300, opacity: 0 } }
      animate={ { x: 0, opacity: 1 } }
      exit={ { x: -300, opacity: 0, transition: { duration: 0.1 } } }
    >
      <img
        src={ pickSrc(pathname) }
        alt="error"
        className="w-60"
      />
      {
        pickText(pathname)
      }
      <button
        className="font-bold text-white bg-orange rounded-full w-52 h-14 mt-12"
        onClick={ () => navigate(pathname === '/todo' ? '/' : '/signin') }
      >
        Go back
      </button>
    </motion.div>
  );
}

export default Error;
