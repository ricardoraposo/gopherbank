import { AnimatePresence } from 'framer-motion';
import { Route, Routes, useLocation } from 'react-router-dom';

import Admin from './pages/Admin';
import SignIn from './pages/SignIn';
import SignUp from './pages/SignUp';
import Profile from './pages/Profile';
import Success from './pages/Success';
import Recover from './pages/Recover';
import ErrorPage from './pages/Error';
import Dashboard from './pages/Dashboard';
import Operation from './pages/Operation';
import AddPicture from './pages/AddPicture';
import EditProfile from './pages/EditProfile';
import SignUpSuccess from './pages/SignUpSuccess';
import OperationAccount from './pages/OperationAccount';

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
        <Route path="*" Component={ ErrorPage } />
      </Routes>
    </AnimatePresence>
  );
}

export default App;
