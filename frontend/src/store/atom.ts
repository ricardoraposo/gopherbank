import { atom } from 'jotai';
import { atomWithStorage } from 'jotai/utils';

export const tokenAtom = atomWithStorage('token', '');

export const accountAtom = atom('');

export const amountAtom = atom('');

export const accountNumberAtom = atom('');

export const showMenuAtom = atom(false);
