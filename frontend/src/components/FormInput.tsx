type Props = {
  name: string;
  id: string;
  type: string;
  value: string;
  onChangeFn: (e: React.ChangeEvent<HTMLInputElement>) => void;
  error?: string | null;
  label?: string;
  inputMode?: 'numeric' | 'text';
};

function FormInput({
  label = '',
  type,
  name,
  id,
  value,
  onChangeFn,
  inputMode = 'text',
  error = null,
}: Props) {
  return (
    <div className="flex flex-col items-center w-4/5">
      <input
        type={ type }
        id={ id }
        inputMode={ inputMode }
        name={ name }
        value={ value }
        onChange={ onChangeFn }
        placeholder={ label }
        className={ `h-12 rounded-none text-lg px-2 mx-auto w-60 bg-transparent border-b-2 text-white placeholder:text-gray-200 
        ${error ? 'border-red' : 'border-white'}` }
      />
    </div>
  );
}

export default FormInput;
