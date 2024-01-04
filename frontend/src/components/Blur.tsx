import { useAtom } from 'jotai';
import { showNotificationAtom } from '../store/atom';

function Blur() {
  const [, setShow] = useAtom(showNotificationAtom);
  return (
    <div
      onClickCapture={ () => setShow(false) }
      className="fixed w-full h-full backdrop-blur-sm z-20"
    />
  );
}

export default Blur;
