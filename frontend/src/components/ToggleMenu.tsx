import { useAtom } from 'jotai';
import { showMenuAtom } from '../store/atom';

function ToggleMenu() {
  const [showMenu] = useAtom(showMenuAtom);

  return (
    <div
      className="flex flex-col gap-1"
    >
      <div className={ `w-5 h-[2px] bg-white transition-transform duration-500 ${showMenu && 'rotate-[405deg] translate-y-[140%]'}` } />
      <div className={ `w-5 h-[2px] bg-white ${showMenu && 'hidden'}` } />
      <div className={ `w-5 h-[2px] bg-white transition-transform duration-500 ${showMenu && '-rotate-[405deg] -translate-y-[140%]'}` } />
    </div>
  );
}

export default ToggleMenu;
