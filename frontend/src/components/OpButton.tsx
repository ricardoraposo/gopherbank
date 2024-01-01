type Props = {
  icon: string;
  label: string;
  onClick: () => void;
  className?: string
};

function OpButton({ icon, label, onClick, className = '' }: Props) {
  return (
    <div className="flex flex-col justify-center items-center">
      <button
        type="button"
        className="w-16 h-16 rounded-full bg-gray-500 text-white text-lg font-semibold
        flex justify-center items-center gap-2"
        onClick={onClick}
      >
        <img src={icon} alt="transfer icon" className={`w-6 h-6 ${className}`} />
      </button>
      <p className="text-white">{label}</p>
    </div>
  );
}

export default OpButton;
