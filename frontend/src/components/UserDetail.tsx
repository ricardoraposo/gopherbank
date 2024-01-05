type Props = {
  label: string;
  detail: string;
};

function UserDetail({ label, detail }: Props) {
  return (
    <div className="px-5 py-6 flex bg-white w-full justify-between">
      <p className="text-gray-500">{label}</p>
      <p className="text-gray-300 pr-6">{detail}</p>
    </div>
  );
}

export default UserDetail;
