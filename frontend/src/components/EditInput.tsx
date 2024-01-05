type Props = {
  label: string;
  onChange: (e: React.ChangeEvent<HTMLInputElement>) => void;
  name: string;
  value: string;
  type?: 'picture' | 'name';
  error?: string | null;
};

function EditInput({ label, onChange, name, value, type = 'picture', error = null }: Props) {
  const className = type === 'picture' ? 'text-sm' : 'text-lg w-32';
  const errorClass = error === null ? '' : 'border-red';

  return (
    <input
      placeholder={ label }
      onChange={ onChange }
      name={ name }
      value={ value }
      className={ `text-gray-100 text-center border-b-2 border-gray-300 bg-transparent
        focus:outline-none placeholder:text-gray-300 ${className} ${errorClass}` }
      type="text"
    />
  );
}

export default EditInput;
