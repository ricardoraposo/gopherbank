import { z } from 'zod';
import { motion } from 'framer-motion';
import { useState } from 'react';
import axios, { AxiosError } from 'axios';
import { useNavigate } from 'react-router-dom';
import FormInput from '../components/FormInput';
import FormButton from '../components/FormButton';
import { apiURL } from '../consts';

const recoverySchema = z.object({
  accountNumber: z.string(),
  password: z.string().min(8, 'Password must be at least 8 characters'),
  confirm: z.string().min(8, 'Password must be at least 8 characters'),
}).refine((data) => data.password === data.confirm, {
  message: 'Passwords do not match',
  path: ['confirm'],
});

function Recover() {
  const navigate = useNavigate();
  const [formValues, setFormValues] = useState({
    accountNumber: '',
    password: '',
    confirm: '',
  });

  const [errors, setErrors] = useState({
    accountNumber: null,
    password: null,
    confirm: null,
  });

  const handleFormChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value } = e.target;
    setFormValues((prev) => ({ ...prev, [name]: value }));
  };

  const handleSubmit = async (e: React.MouseEvent<HTMLButtonElement>) => {
    e.preventDefault();
    try {
      recoverySchema.parse(formValues);
      await axios.patch(`${apiURL}/auth/recover`, formValues);
      navigate('/signin');
    } catch (error) {
      if (error instanceof AxiosError && error.response?.status === 404) {
        navigate('/signin');
      }
      if (error instanceof z.ZodError) {
        error.issues.forEach((err) => {
          setErrors((prev) => ({ ...prev, [err.path[0]]: err.message }));
        });
      }
    }
  };

  return (
    <motion.div
      className="h-dvh flex justify-center items-center
      bg-login bg-cover"
      initial={ { x: 300, opacity: 0 } }
      animate={ { x: 0, opacity: 1 } }
      exit={ { x: -300, opacity: 0, transition: { duration: 0.1 } } }
    >
      <form
        className="bg-gray-500 rounded-[40px] opacity-85
        flex flex-col justify-center items-center gap-6
        w-4/5 h-[90%]"
      >
        <h1 className="text-white text-3xl font-bold">REGISTER</h1>
        <div className="w-4/5">
          <FormInput
            label="Account Number"
            name="accountNumber"
            id="accountNumber"
            type="number"
            value={ formValues.accountNumber }
            onChangeFn={ handleFormChange }
            error={ errors.accountNumber }
          />
          {errors.accountNumber && <p className="text-red text-sm px-2">{errors.accountNumber}</p>}
        </div>
        <div className="w-4/5">
          <FormInput
            label="Password"
            name="password"
            id="password"
            type="password"
            value={ formValues.password }
            onChangeFn={ handleFormChange }
            error={ errors.password }
          />
          {errors.password && <p className="text-red text-sm px-2">{errors.password}</p>}
        </div>
        <div className="w-4/5">
          <FormInput
            label="Confirm Password"
            name="confirm"
            id="confirm"
            type="password"
            value={ formValues.confirm }
            onChangeFn={ handleFormChange }
            error={ errors.confirm }
          />
          {errors.confirm && <p className="text-red text-sm px-2">{errors.confirm}</p>}
        </div>
        <FormButton
          label="Recover"
          onSubmitFn={ handleSubmit }
          isLoading={ false }
        />
      </form>
    </motion.div>
  );
}

export default Recover;
