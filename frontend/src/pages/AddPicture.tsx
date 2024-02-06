import { motion } from 'framer-motion';
import { useState } from 'react';
import axios from 'axios';
import { useAtom } from 'jotai';
import { useNavigate } from 'react-router-dom';
import { apiURL, defaultPic } from '../consts';
import FormButton from '../components/FormButton';
import { accountAtom, initialFormValues, signUpAtom } from '../store/atom';

function AddPicture() {
  const navigate = useNavigate();
  const [formValues, setFormValues] = useAtom(signUpAtom);
  const [, setAccount] = useAtom(accountAtom);
  const [isLoading, setIsLoading] = useState(false);

  const [file, setFile] = useState<File | null>(null);

  const handleSubmit = async (e: React.MouseEvent<HTMLButtonElement>) => {
    e.preventDefault();
    try {
      const formData = new FormData();
      formData.append('file', file as File);
      formData.append(
        'formValues',
        new Blob([JSON.stringify(formValues)], { type: 'application/json' }),
      );

      setIsLoading(true);
      const { data: { number } } = await axios.post(`${apiURL}/auth/new`, formData, {
        headers: { 'Content-Type': 'multipart/form-data' },
        transformRequest: [(data) => data],
      });
      setAccount(number);
      setFormValues(initialFormValues);
      navigate('/signup/success');
    } catch (error: any) {
      console.error(error.response.data);
    } finally {
      setIsLoading(false);
    }
  };

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    if (e.target.files) {
      setFile(e.target.files[0]);
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
      <div
        className="bg-gray-500 rounded-[40px] opacity-85
        flex flex-col justify-center items-center gap-2
        w-4/5 h-[90%]"
      >
        <div className="flex flex-col justify-center items-center gap-4">
          <img
            src={ file ? URL.createObjectURL(file) : defaultPic }
            alt="default user profile"
            className="w-60 h-60 bg-white border-8 border-black rounded-full opacity-100"
          />
          <form className="flex flex-col items-center gap-4">
            <label
              htmlFor="profilePicture"
              className="text-white text-lg bg-purple p-2 rounded-lg cursor-pointer"
            >
              Add a profile picture
            </label>
            <input
              type="file"
              onChange={ handleChange }
              id="profilePicture"
              className="hidden"
            />
            <FormButton
              label="Add picture"
              isLoading={ isLoading }
              onSubmitFn={ handleSubmit }
            />
          </form>
        </div>
        <button
          className="text-lg text-purple mt-12"
        >
          Or do it later...
        </button>
      </div>
    </motion.div>
  );
}

export default AddPicture;
