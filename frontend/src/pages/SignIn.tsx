import axios from 'axios';
import { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { apiURL } from '../consts';
import FormInput from '../components/FormInput';
import FormButton from '../components/FormButton';

function SignIn() {
  const navigate = useNavigate();
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
      console.log(formValues)
      const { data } = await axios.post(`${apiURL}/auth`, formValues);
      const { token } = data;
      localStorage.setItem('token', token);
      navigate('/dashboard');
    } catch (error: any) {
      console.error(error.response.data);
    }
  };

  return (
    <div
      className="h-dvh flex flex-col justify-center items-center
      bg-bg bg-login bg-cover font-primary"
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
        flex flex-col justify-center items-center gap-2
        w-4/5 h-96 mt-16"
      >
        <h1 className="text-white text-3xl font-bold">LOGIN</h1>
        <FormInput
          label="Account Number"
          name="number"
          id="number"
          type="number"
          value={formValues.number}
          onChangeFn={handleFormChange}
          inputMode="numeric"
        />
        <FormInput
          label="Password"
          name="password"
          id="password"
          type="password"
          value={formValues.password}
          onChangeFn={handleFormChange}
        />
        <FormButton label="Sign in" onSubmitFn={handleSubmit} />
        <div>
          <div className="flex justify-between gap-4 text-xs text-white my-1">
            Don't have an account ?
            {' '}
            <button
              className="text-purple font-semibold"
              onClick={() => navigate('/signup')}
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
            >
              Recover it
            </button>
          </div>
        </div>
      </form>
    </div>
  );
}

export default SignIn;
