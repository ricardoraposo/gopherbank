import axios from 'axios';
import { useEffect, useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { useAtom } from 'jotai';
import { motion } from 'framer-motion';
import { apiURL, queryParams } from '../consts';
import FormInput from '../components/FormInput';
import FormButton from '../components/FormButton';
import { tokenAtom } from '../store/atom';

function SignIn() {
  const navigate = useNavigate();
  const [token, setToken] = useAtom(tokenAtom);
  const [error, setError] = useState<string | null>(null);
  const [isLoading, setIsLoading] = useState(false);
  const [formValues, setFormValues] = useState({
    number: '',
    password: '',
  });

  const handleFormChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value } = e.target;
    setFormValues({
      ...formValues,
      [name]: value,
    });
  };

  const handleSubmit = async (e: React.MouseEvent<HTMLButtonElement>) => {
    e.preventDefault();
    try {
      setIsLoading(true);
      const { data: { token } } = await axios.post(`${apiURL}/auth`, formValues);
      setToken(token);
      navigate('/');
    } catch (error: any) {
      setError("Couldn't sign in. Please check your credentials.");
    } finally {
      setIsLoading(false);
    }
  };

  useEffect(() => {
    const checkToken = async () => {
      if (token) {
        try {
          await axios.get(`${apiURL}/api/jwt/`, queryParams(token));
          navigate('/');
        } catch (e: any) {
          console.log('Keep going');
        }
      }
    };
    checkToken();
  }, [token]);

  return (
    <motion.div
      className="h-dvh flex flex-col justify-center items-center
      bg-bg bg-login bg-cover font-primary"
      initial={ { x: 300, opacity: 0 } }
      animate={ { x: 0, opacity: 1 } }
      exit={ { x: -300, opacity: 0, transition: { duration: 0.1 } } }
    >
      <div className="w-4/5">
        <div className="text-white text-[2.8rem] leading-snug font-extrabold">
          <h1>
            STEP INTO THE BANKING FUTURE
            {' '}
            <span className="text-orange">NOW</span>
          </h1>
        </div>
        <div className="text-gray-200">
          <p>We will do everything but count coins.</p>
          <p>Faster, easier and more convenient</p>
        </div>
      </div>
      <form
        className="bg-gray-500 rounded-[40px] opacity-85
        flex flex-col justify-center items-center gap-4
        w-4/5 h-96 mt-16"
      >
        <h1 className="text-white text-3xl font-bold">LOGIN</h1>
        <FormInput
          label="Account Number"
          name="number"
          id="number"
          type="number"
          value={ formValues.number }
          onChangeFn={ handleFormChange }
          inputMode="numeric"
          error={ error }
        />
        <FormInput
          label="Password"
          name="password"
          id="password"
          type="password"
          value={ formValues.password }
          onChangeFn={ handleFormChange }
          error={ error }
        />
        {error && <p className="text-red text-sm w-4/5 px-2">{error}</p>}
        <FormButton label="Sign in" onSubmitFn={ handleSubmit } isLoading={ isLoading } />
        <div>
          <div className="flex justify-between gap-4 text-xs text-white my-1">
            Don't have an account ?
            {' '}
            <button
              className="text-purple font-semibold"
              onClick={ () => navigate('/signup') }
              type="button"
            >
              Create one
            </button>
          </div>
          <div className="flex justify-between gap-4 text-xs text-white">
            Forgot your password ?
            {' '}
            <button
              className="text-purple font-semibold"
              type="button"
              onClick={ () => navigate('/recover') }
            >
              Recover it
            </button>
          </div>
        </div>
      </form>
    </motion.div>
  );
}

export default SignIn;
