import { z } from 'zod';
import axios from 'axios';
import { motion } from 'framer-motion';
import { useAtom } from 'jotai';
import { useQuery } from '@tanstack/react-query';
import { useEffect, useState } from 'react';
import { useNavigate } from 'react-router-dom';

import { EditError, type User } from '../types';
import { apiURL, queryParams } from '../consts';
import { accountAtom, showNotificationAtom, tokenAtom } from '../store/atom';

import Blur from '../components/Blur';
import Line from '../components/Line';
import Check from '../assets/checkmark.svg';
import Arrow from '../assets/arrow_left.svg';
import Loading from '../components/Loading';
import SideMenu from '../components/SideMenu';
import ToggleMenu from '../components/ToggleMenu';
import ToggleNotifications from '../components/ToggleNotifications';
import { usDateFormat } from '../utils/helpers';
import UserDetail from '../components/UserDetail';
import EditInput from '../components/EditInput';

const editSchema = z.object({
  firstName: z.union([z.string().length(0), z.string().min(2, 'First name must be at least 2 characters')])
    .optional()
    .transform((e) => (e === '' ? undefined : e)),
  lastName: z.union([z.string().length(0), z.string().min(2, 'Last name must be at least 2 characters')])
    .optional()
    .transform((e) => (e === '' ? undefined : e)),
  email: z.union([z.string().length(0), z.string().email('Invalid email address')])
    .optional()
    .transform((e) => (e === '' ? undefined : e)),
  pictureUrl: z.union([z.string().length(0), z.string().url('Invalid URL')])
    .optional()
    .transform((e) => (e === '' ? undefined : e)),
});

function EditProfile() {
  const navigate = useNavigate();
  const [show] = useAtom(showNotificationAtom);
  const [, setID] = useAtom(accountAtom);
  const [token, setToken] = useAtom(tokenAtom);

  const { data: user, isLoading, error } = useQuery({
    queryKey: ['user'],
    queryFn: async () => {
      const { data } = await axios.get(`${apiURL}/api/jwt/`, queryParams(token));
      setID(data.number);
      return axios.get(`${apiURL}/api/accounts/${data.number}`, queryParams(token));
    },
    select: ({ data }) => {
      return { number: data.number, join: data.createdAt, ...data.edges.user } as User;
    },
    retry: 2,
  });

  const [form, setForm] = useState({
    firstName: '',
    lastName: '',
    email: '',
    pictureUrl: '',
  });

  const [errors, setErrors] = useState<EditError>({
    firstName: null,
    lastName: null,
    email: null,
    pictureUrl: null,
  });

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setForm({ ...form, [e.target.name]: e.target.value });
  };

  const handleSubmit = async (e: React.MouseEvent<HTMLButtonElement>) => {
    e.preventDefault();
    try {
      editSchema.parse(form);
      await axios.put(`${apiURL}/api/user`, form, queryParams(token));
      navigate('/profile');
    } catch (e: any) {
      if (e instanceof z.ZodError) {
        e.issues.forEach((issue) => {
          setErrors({ ...errors, [issue.path[0]]: 'Invalid value' });
        });
      }
    }
  };

  useEffect(() => {
    if (axios.isAxiosError(error) && error.response?.status === 401) {
      setToken('');
      navigate('/signin');
    }
  }, [error]);

  if (isLoading) return <Loading />;

  return (
    <motion.div
      className="pt-2"
      initial={{ x: 300, opacity: 0 }}
      animate={{ x: 0, opacity: 1 }}
      exit={{ x: -300, opacity: 0, transition: { duration: 0.1 } }}
    >
      {show && <Blur />}
      <SideMenu />
      <header className="flex justify-between items-center mx-5">
        <button
          onClick={() => navigate(-1)}
          className="bg-gray-500 w-11 h-11 flex justify-center items-center rounded-full z-30"
        >
          <img src={Arrow} alt="go back" />
        </button>
        <div className="flex gap-6">
          <ToggleNotifications />
          <ToggleMenu />
        </div>
      </header>
      <div className="flex flex-col items-center gap-5 mt-24 mb-8">
        <img src={user?.pictureUrl} alt="user profile" className="w-52 h-52 rounded-full" />
        <div>
          <EditInput
            label="Picture URL"
            onChange={handleChange}
            name="pictureUrl"
            value={form.pictureUrl}
            error={errors.pictureUrl}
          />
          {errors.pictureUrl && <p className="text-red">{errors.pictureUrl}</p>}
        </div>
        <div className="flex justify-center gap-4">
          <div>
            <EditInput
              label={`${user?.firstName}`}
              onChange={handleChange}
              type="name"
              name="firstName"
              value={form.firstName}
              error={errors.firstName}
            />
            {errors.firstName && <p className="text-red">{errors.firstName}</p>}
          </div>
          <div>
            <EditInput
              label={`${user?.lastName}`}
              onChange={handleChange}
              type="name"
              name="lastName"
              value={form.lastName}
              error={errors.lastName}
            />
            {errors.lastName && <p className="text-red text-sm">{errors.lastName}</p>}
          </div>
        </div>
      </div>
      <div
        className="relative w-screen h-auto min-h-[60dvh] mt-6 pt-16 bg-slate-50 rounded-t-[40px] overflow-hidden"
      >
        <Line />
        <button
          className="absolute top-2 right-4 bg-gray-100 w-11 h-11
          flex justify-center items-center rounded-full"
          onClick={handleSubmit }
        >
          <img src={ Check } alt="edit" />
        </button>
        {
          user && (
            <div className="flex flex-col gap-4">
              <UserDetail label="Account Number" detail={ user?.number } />
              <div className="px-5 py-6 flex bg-white w-full justify-between">
                <p className="text-gray-500">E-mail</p>
                <div>
                  <input
                    type="text"
                    placeholder={ user?.email }
                    className={ `text-gray-400 border-b-2 border-gray-300
                  focus:outline-none placeholder:text-gray-300 ${errors.email && 'border-red'}` }
                    name="email"
                    value={ form.email }
                    onChange={ handleChange }
                  />
                  {errors.email && <p className="text-red text-sm">{errors.email}</p>}
                </div>
              </div>
              <UserDetail label="Joined at" detail={ usDateFormat.format(new Date(user?.join)) } />
            </div>
          )
        }
      </div>
    </motion.div>
  );
}

export default EditProfile;
