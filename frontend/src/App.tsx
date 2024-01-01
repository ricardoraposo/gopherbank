import { Route, Routes, useLocation } from 'react-router-dom';
import { AnimatePresence } from 'framer-motion';

import SignIn from './pages/SignIn';
import SignUp from './pages/SignUp';
import Success from './pages/Success';
import Dashboard from './pages/Dashboard';
import Operation from './pages/Operation';
import OperationAccount from './pages/OperationAccount';

function App() {
  const location = useLocation();

  return (
    <AnimatePresence>
      <Routes location={ location } key={ location.pathname }>
        <Route path="/" Component={ Dashboard } />
        <Route path="/signin" Component={ SignIn } />
        <Route path="/signup" Component={ SignUp } />
        <Route path="/operation/:type" Component={ Operation } />
        <Route path="/operation/:type/account" Component={ OperationAccount } />
        <Route path="/operation/:type/success" Component={ Success } />
      </Routes>
    </AnimatePresence>
  );
}

export default App;
