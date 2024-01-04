import { atom } from 'jotai';
import { atomWithStorage } from 'jotai/utils';
import { defaultPic } from '../consts';

export const tokenAtom = atomWithStorage('token', '');

export const accountAtom = atom('');

export const amountAtom = atom('');

export const accountNumberAtom = atom('');

export const showMenuAtom = atom(false);

export const showNotificationAtom = atom(false);

export const initialFormValues = {
  firstName: '',
  lastName: '',
  email: '',
  password: '',
  confirm: '',
  pictureUrl: defaultPic,
};

export const signUpAtom = atom(initialFormValues);

export const selectedAtom = atom('');
