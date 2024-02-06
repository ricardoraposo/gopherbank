type Props = {
  label: string;
  onSubmitFn: (e: React.MouseEvent<HTMLButtonElement>) => void;
  isLoading: boolean;
};

function FormButton({ label, onSubmitFn: handleSubmit, isLoading }: Props) {
  return (
    <button
      type="button"
      onClick={ handleSubmit }
      className="mt-2 h-12 w-40 rounded-xl bg-orange
      text-white text-2xl font-semibold"
    >
      {isLoading ? (
        <div
          className="h-6 w-6 border-4 rounded-full mx-auto
          border-t-white border-l-white border-r-orange border-b-orange
          animate-spin"
        />
      ) : label}
    </button>
  );
}

export default FormButton;
