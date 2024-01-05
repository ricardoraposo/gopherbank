export type SignUpError = {
  firstName: string | null;
  lastName: string | null;
  email: string | null;
  password: string | null;
  confirm: string | null;
};

export type EditError = {
  firstName: string | null;
  lastName: string | null;
  email: string | null;
  pictureUrl: string | null;
};

export type User = {
  email: string;
  firstName: string;
  lastName: string;
  id: number;
  number: string;
  join: string;
  pictureUrl: string;
};
