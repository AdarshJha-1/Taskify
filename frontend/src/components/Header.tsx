import { Link } from "react-router-dom";

export default function Header() {
  return (
    <header className="w-full h-[70px] flex justify-center item-center">
      <nav className="w-1/2 h-full flex items-center justify-between">
        <Link
          to={"/"}
          className="text-5xl h-[70px] flex items-center font-bold tracking-tight bg-gradient-to-r from-pink-600 via--500 to-indigo-400 text-transparent bg-clip-text"
        >
          Taskify
        </Link>

        <ul className="h-full flex items-end justify-between gap-10 text-xl  font-semibold  tracking-tight text-white">
          <li>
            <button className="relative inline-flex items-center justify-center p-0.5 mb-2 me-2 overflow-hidden text-sm font-medium text-gray-900 rounded-lg group bg-gradient-to-br from-purple-500 to-pink-500 group-hover:from-purple-500 group-hover:to-pink-500 hover:text-white dark:text-white focus:ring-4 focus:outline-none focus:ring-purple-200 dark:focus:ring-purple-800">
              <Link to={"/signup"} className="flex items-center relative px-5 py-2.5 transition-all ease-in duration-75 bg-white dark:bg-gray-900 rounded-md group-hover:bg-opacity-0">
              Signup
              </Link>
            </button>
          </li>
          <li>
            <button className="relative inline-flex items-center justify-center p-0.5 mb-2 me-2 overflow-hidden text-sm font-medium text-gray-900 rounded-lg group bg-gradient-to-br from-purple-500 to-pink-500 group-hover:from-purple-500 group-hover:to-pink-500 hover:text-white dark:text-white focus:ring-4 focus:outline-none focus:ring-purple-200 dark:focus:ring-purple-800">
              <Link to={"/signin"} className="flex items-center relative px-5 py-2.5 transition-all ease-in duration-75 bg-white dark:bg-gray-900 rounded-md group-hover:bg-opacity-0">
              Signin
              </Link>
            </button>
          </li>
        </ul>
      </nav>
    </header>
  );
}
