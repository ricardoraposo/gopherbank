import z from 'zod';
import { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { motion } from 'framer-motion';
import { useAtom } from 'jotai';
import FormButton from '../components/FormButton';
import FormInput from '../components/FormInput';
import { signUpAtom } from '../store/atom';
import { SignUpError } from '../types';

const signUpSchema = z.object({
  firstName: z.string().min(2, "First name can't be shorter than 2 characters"),
  lastName: z.string().min(2, "Last name can't be shorter than 2 characters"),
  email: z.string().email('Please enter a valid e-mail address'),
  password: z.string().min(8, "Password can't be shorter than 8 characters"),
  confirm: z.string(),
}).refine((data) => data.password === data.confirm, {
  path: ['confirm'],
  message: 'Passwords do not match',
});

function SignUp() {
  const navigate = useNavigate();
  const [isLoading, setIsLoading] = useState(false);
  const [formValues, setFormValues] = useAtom(signUpAtom);
  const [errors, setErrors] = useState<SignUpError>({
    firstName: null,
    lastName: null,
    email: null,
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
      setIsLoading(true);
      signUpSchema.parse(formValues);
      navigate('/signup/picture');
    } catch (error: any) {
      if (error instanceof z.ZodError) {
        error.issues.forEach((err: any) => {
          setErrors((prev) => ({ ...prev, [err.path[0]]: err.message }));
        });
      }
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
        flex flex-col justify-center items-center gap-6
        w-4/5 h-[90%]"
      >
        <h1 className="text-white text-3xl font-bold">REGISTER</h1>
        <div className="w-4/5">
          <FormInput
            label="First name"
            name="firstName"
            id="firstName"
            type="text"
            value={ formValues.firstName }
            onChangeFn={ handleFormChange }
            error={ errors.firstName }
          />
          {errors.firstName && <p className="text-red text-sm px-2">{errors.firstName}</p>}
        </div>
        <div className="w-4/5">
          <FormInput
            label="Last name"
            name="lastName"
            id="lastName"
            type="text"
            value={ formValues.lastName }
            onChangeFn={ handleFormChange }
            error={ errors.lastName }
          />
          {errors.lastName && <p className="text-red text-sm px-2">{errors.lastName}</p>}
        </div>
        <div className="w-4/5">
          <FormInput
            label="E-mail"
            name="email"
            id="email"
            type="email"
            value={ formValues.email }
            onChangeFn={ handleFormChange }
            error={ errors.email }
          />
          {errors.email && <p className="text-red text-sm px-2">{errors.email}</p>}
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
            label="Confirm password"
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
          label="Sign up"
          onSubmitFn={ handleSubmit }
          isLoading={ isLoading }
        />
      </form>
    </motion.div>
  );
}

export default SignUp;
