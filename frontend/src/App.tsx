import { Route, Routes } from 'react-router-dom';
import Home from './pages/Home';
import SignIn from './pages/SignIn';
import SignUp from './pages/SignUp';
import Dashboard from './pages/Dashboard';

function App() {
  return (
    <Routes>
      <Route path="/" element={ <Home /> } />
      <Route path="/signin" Component={ SignIn } />
      <Route path="/signup" Component={ SignUp } />
      <Route path="/dashboard" Component={ Dashboard } />
    </Routes>
  );
}

export default App;
