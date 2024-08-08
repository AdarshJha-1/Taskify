import { Link, useNavigate } from "react-router-dom";
import { useRecoilValue, useSetRecoilState } from "recoil";
import { addTodoState, loginState, todoListState } from "../store/atom";
import axios from "axios";

export default function Header() {
  const { isLogin, token } = useRecoilValue(loginState);
  const setLoginState = useSetRecoilState(loginState);
  const setAddTodoState = useSetRecoilState(addTodoState);
  const navigate = useNavigate();
  const handleSignOut = async () => {
    try {
      const response = await axios.post(
        import.meta.env.VITE_BASE_URL + "/sign-out",
        {},
        {
          headers: {
            "Content-Type": "application/json",
          },
          withCredentials: true,
        }
      );
      if (!response.data.success) {
        console.error("ERROR::", response.data.message);
      } else {
        setLoginState({ isLogin: false, token: null });
        navigate("/signin");
      }
    } catch (error) {
      if (axios.isAxiosError(error)) {
        console.error("Axios error:", error.message);
      } else {
        console.error("Unexpected error:", error);
      }
    }
  };

  const handleAddTodo = () => {
    setAddTodoState((prevState: boolean) => !prevState);
  };
  const setTodoList = useSetRecoilState(todoListState);
  // const [refresh, setRefresh] = useRecoilState(refreshState);
  const handleRefresh = async () => {
    try {
      const res = await axios.get(import.meta.env.VITE_BASE_URL + "/todos", {
        withCredentials: true,
      });

      setTodoList(res.data.data.todos);
    } catch (error) {
      console.error("Error fetching todo list:", error);
    }
  };

  return (
    <header className="w-full h-[70px] flex justify-center item-center">
      <nav className="w-1/2 h-full flex items-center justify-between">
        <Link
          to={"/"}
          className="text-5xl h-[70px] flex items-center font-bold tracking-tight bg-gradient-to-r from-pink-600 via--500 to-indigo-400 text-transparent bg-clip-text"
        >
          Taskify
        </Link>

        {!isLogin && !token ? (
          <ul className="h-full flex items-end justify-between gap-10 text-xl  font-semibold  tracking-tight text-white">
            <li>
              <button className="relative inline-flex items-center justify-center p-0.5 mb-2 me-2 overflow-hidden text-sm font-medium text-gray-900 rounded-lg group bg-gradient-to-br from-purple-500 to-pink-500 group-hover:from-purple-500 group-hover:to-pink-500 hover:text-white dark:text-white focus:ring-4 focus:outline-none focus:ring-purple-200 dark:focus:ring-purple-800">
                <Link
                  to={"/signup"}
                  className="flex items-center relative px-5 py-2.5 transition-all ease-in duration-75 bg-white dark:bg-gray-900 rounded-md group-hover:bg-opacity-0"
                >
                  Signup
                </Link>
              </button>
            </li>
            <li>
              <button className="relative inline-flex items-center justify-center p-0.5 mb-2 me-2 overflow-hidden text-sm font-medium text-gray-900 rounded-lg group bg-gradient-to-br from-purple-500 to-pink-500 group-hover:from-purple-500 group-hover:to-pink-500 hover:text-white dark:text-white focus:ring-4 focus:outline-none focus:ring-purple-200 dark:focus:ring-purple-800">
                <Link
                  to={"/signin"}
                  className="flex items-center relative px-5 py-2.5 transition-all ease-in duration-75 bg-white dark:bg-gray-900 rounded-md group-hover:bg-opacity-0"
                >
                  Signin
                </Link>
              </button>
            </li>
          </ul>
        ) : (
          <div className="flex gap-5 justify-center items-center">
            <button
              onClick={handleAddTodo}
              className="relative inline-flex items-center justify-center p-0.5 mb-2 me-2 overflow-hidden text-sm font-medium text-gray-900 rounded-lg group bg-gradient-to-br from-purple-500 to-pink-500 group-hover:from-purple-500 group-hover:to-pink-500 hover:text-white dark:text-white focus:ring-4 focus:outline-none focus:ring-purple-200 dark:focus:ring-purple-800"
            >
              <span className="flex items-center relative px-5 py-2.5 transition-all ease-in duration-75 bg-white dark:bg-gray-900 rounded-md group-hover:bg-opacity-0">
                Add Todo
              </span>
            </button>
            <button
              onClick={handleSignOut}
              className="relative inline-flex items-center justify-center p-0.5 mb-2 me-2 overflow-hidden text-sm font-medium text-gray-900 rounded-lg group bg-gradient-to-br from-purple-500 to-pink-500 group-hover:from-purple-500 group-hover:to-pink-500 hover:text-white dark:text-white focus:ring-4 focus:outline-none focus:ring-purple-200 dark:focus:ring-purple-800"
            >
              <span className="flex items-center relative px-5 py-2.5 transition-all ease-in duration-75 bg-white dark:bg-gray-900 rounded-md group-hover:bg-opacity-0">
                Sign Out
              </span>
            </button>
            <button
              onClick={handleRefresh}
              className="relative inline-flex items-center justify-center p-0.5 mb-2 me-2 overflow-hidden text-sm font-medium text-gray-900 rounded-lg group bg-gradient-to-br from-purple-500 to-pink-500 group-hover:from-purple-500 group-hover:to-pink-500 hover:text-white dark:text-white focus:ring-4 focus:outline-none focus:ring-purple-200 dark:focus:ring-purple-800"
            >
              <span className="flex items-center relative px-5 py-2.5 transition-all ease-in duration-75 bg-white dark:bg-gray-900 rounded-md group-hover:bg-opacity-0">
                refresh
              </span>
            </button>
          </div>
        )}
      </nav>
    </header>
  );
}
