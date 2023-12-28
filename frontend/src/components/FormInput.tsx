type Props = {
  label: string;
  name: string;
  id: string;
  type: string;
  value: string;
  onChangeFn: (e: React.ChangeEvent<HTMLInputElement>) => void;
  inputMode?: 'numeric' | 'text';
};

function FormInput({
  label,
  type,
  name,
  id,
  value,
  onChangeFn,
  inputMode = 'text',
}: Props) {
  return (
    <div className="flex flex-col items-center w-4/5">
      <label
        htmlFor="number"
        className="text-white text-base self-start pl-2 pb-1"
      >
        {label}
      </label>
      <input
        type={ type }
        id={ id }
        inputMode={ inputMode }
        name={ name }
        value={ value }
        onChange={ onChangeFn }
        className="h-12 rounded-xl text-2xl px-2 w-60"
      />
    </div>
  );
}

export default FormInput;
