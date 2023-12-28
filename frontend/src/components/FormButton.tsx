type Props = {
  label: string;
  onSubmitFn: (e: React.MouseEvent<HTMLButtonElement>) => void;
};

function FormButton({ label, onSubmitFn: handleSubmit }: Props) {
  return (
    <button
      type="button"
      onClick={ handleSubmit }
      className="mt-2 h-12 w-40 bg-orange rounded-xl
            text-white text-2xl font-semibold"
    >
      {label}
    </button>
  );
}

export default FormButton;
