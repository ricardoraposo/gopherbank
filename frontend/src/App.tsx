import { Route, Routes, useLocation } from 'react-router-dom';
import { AnimatePresence } from 'framer-motion';

import SignIn from './pages/SignIn';
import SignUp from './pages/SignUp';
import Success from './pages/Success';
import Dashboard from './pages/Dashboard';
import Operation from './pages/Operation';
import AddPicture from './pages/AddPicture';
import OperationAccount from './pages/OperationAccount';
import SignUpSuccess from './pages/SignUpSuccess';
import Admin from './pages/Admin';
import Error from './pages/Error';
import Recover from './pages/Recover';
import Profile from './pages/Profile';
import EditProfile from './pages/EditProfile';

function App() {
  const location = useLocation();

  return (
    <AnimatePresence>
      <Routes location={ location } key={ location.pathname }>
        <Route path="/" Component={ Dashboard } />
        <Route path="/admin" Component={ Admin } />
        <Route path="/signin" Component={ SignIn } />
        <Route path="/signup" Component={ SignUp } />
        <Route path="/signup/picture" Component={ AddPicture } />
        <Route path="/signup/success" Component={ SignUpSuccess } />
        <Route path="/recover" Component={ Recover } />
        <Route path="/profile" Component={ Profile } />
        <Route path="/profile/edit" Component={ EditProfile } />
        <Route path="/operation/:type" Component={ Operation } />
        <Route path="/operation/:type/account" Component={ OperationAccount } />
        <Route path="/operation/:type/success" Component={ Success } />
        <Route path="*" Component={ Error } />
      </Routes>
    </AnimatePresence>
  );
}

export default App;
