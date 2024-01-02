type Props = {
  name: string;
  id: string;
  type: string;
  value: string;
  onChangeFn: (e: React.ChangeEvent<HTMLInputElement>) => void;
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
        className="h-12 rounded-none text-lg px-2 w-60 bg-transparent border-b-2 border-white text-white placeholder:text-gray-200"
      />
    </div>
  );
}

export default FormInput;
