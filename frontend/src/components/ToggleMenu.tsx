import { useAtom } from 'jotai';
import { showMenuAtom } from '../store/atom';

function ToggleMenu() {
  const [showMenu, setShowMenu] = useAtom(showMenuAtom);

  return (
    <button
      className="bg-gray-500 w-11 h-11 flex justify-center items-center rounded-full z-50"
      onClick={ () => setShowMenu((prev) => !prev) }
    >
      <div
        className="flex flex-col gap-1"
      >
        <div className={ `w-5 h-[2px] bg-white transition-transform duration-700 ${showMenu && 'rotate-[405deg] translate-y-[140%]'}` } />
        <div className={ `w-5 h-[2px] bg-white ${showMenu && 'hidden'}` } />
        <div className={ `w-5 h-[2px] bg-white transition-transform duration-700 ${showMenu && '-rotate-[405deg] -translate-y-[140%]'}` } />
      </div>
    </button>
  );
}

export default ToggleMenu;
