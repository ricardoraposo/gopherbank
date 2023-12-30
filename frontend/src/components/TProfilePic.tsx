type Props = {
  profileURL: string;
  className?: string;
};

function TProfilePic({ profileURL, className = 'h-9 w-9' }: Props) {
  return (
    <img
      src={ profileURL }
      alt="transaction profile pic"
      className={ `object-cover rounded-full ${className}` }
    />
  );
}

export default TProfilePic;
