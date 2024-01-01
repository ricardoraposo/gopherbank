import Loading from "./Loading";

type Props = {
  label: string;
  onSubmitFn: (e: React.MouseEvent<HTMLButtonElement>) => void;
  isLoading: boolean;
};

function FormButton({ label, onSubmitFn: handleSubmit, isLoading }: Props) {
  return (
    <button
      type="button"
      onClick={handleSubmit}
      className="mt-2 h-12 w-40 bg-orange rounded-xl
            text-white text-2xl font-semibold"
    >
      {isLoading ? <Loading /> : label}
    </button>
  );
}

export default FormButton;
