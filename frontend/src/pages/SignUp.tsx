import { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { motion } from 'framer-motion';
import { useAtom } from 'jotai';
import FormButton from '../components/FormButton';
import FormInput from '../components/FormInput';
import { signUpAtom } from '../store/atom';

function SignUp() {
  const navigate = useNavigate();
  const [isLoading, setIsLoading] = useState(false);
  const [formValues, setFormValues] = useAtom(signUpAtom);

  const [confirmPassword, setConfirmPassword] = useState('');

  const handleFormChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value } = e.target;
    setFormValues((prev) => ({ ...prev, [name]: value }));
  };

  const handleSubmit = async (e: React.MouseEvent<HTMLButtonElement>) => {
    e.preventDefault();
    try {
      setIsLoading(true);
      navigate('/signup/picture');
    } catch (error: any) {
      console.error(error.response.data);
    } finally {
      setIsLoading(false);
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
        flex flex-col justify-center items-center gap-2
        w-4/5 h-[90%]"
      >
        <h1 className="text-white text-3xl font-bold">REGISTER</h1>
        <FormInput
          label="First name"
          name="firstName"
          id="firstName"
          type="text"
          value={ formValues.firstName }
          onChangeFn={ handleFormChange }
        />
        <FormInput
          label="Last name"
          name="lastName"
          id="lastName"
          type="text"
          value={ formValues.lastName }
          onChangeFn={ handleFormChange }
        />
        <FormInput
          label="E-mail"
          name="email"
          id="email"
          type="email"
          value={ formValues.email }
          onChangeFn={ handleFormChange }
        />
        <FormInput
          label="Password"
          name="password"
          id="password"
          type="password"
          value={ formValues.password }
          onChangeFn={ handleFormChange }
        />
        <FormInput
          label="Confirm password"
          name="confirm"
          id="confirm"
          type="password"
          value={ confirmPassword }
          onChangeFn={ (e) => setConfirmPassword(e.target.value) }
        />
        <FormButton
          label="Sign up"
          onSubmitFn={ handleSubmit }
          isLoading={ isLoading }
        />
      </form>
    </motion.div>
  );
}

export default SignUp;
