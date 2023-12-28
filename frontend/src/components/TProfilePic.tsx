type Props = {
  profileURL: string;
};

function TProfilePic({ profileURL }: Props) {
  return (
    <img
      src={ profileURL }
      alt="transaction profile pic"
      className="h-9 w-9 object-cover rounded-full"
    />
  );
}

export default TProfilePic;
