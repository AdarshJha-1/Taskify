import axios from "axios";
import { useState } from "react";
import { useSetRecoilState } from "recoil";
import { addTodoState } from "../store/atom";

interface AddTodoType {
  title: string;
  description: string;
}
export default function AddTodo() {
    const setAddTodoState = useSetRecoilState(addTodoState);
  const [todoData, setTodoData] = useState<AddTodoType>({
    title: "",
    description: "",
  });
  const handleSubmit = async (e: React.ChangeEvent<HTMLFormElement>) => {
    e.preventDefault();
    try {
      const response = await axios.post(
        import.meta.env.VITE_BASE_URL + "/todos",
        todoData,
        {
          headers: {
            "Content-Type": "application/json",
          },
          withCredentials: true,
        }
      );
      if (!response.data.success){
        console.error("ERROR::",response.data.message);
      }else{
        console.log(response.data);
        
      }
    } catch (error) {
      if (axios.isAxiosError(error)) {
        console.error('Axios error:', error.message);
      } else {
        console.error('Unexpected error:', error);
      }
    } finally {
        setAddTodoState(false)
    }
  };

  return (
    <div className="flex justify-center items-center absolute w-full z-30">
      <form onSubmit={handleSubmit} className="flex flex-col items-center gap-5  bg-white bg-opacity-10 p-5">
        <input
          type="text"
          placeholder="Title"
          value={todoData.title}
          onChange={(e) => setTodoData({ ...todoData, title: e.target.value })}
        />
        <input
          type="text"
          placeholder="Description"
          value={todoData.description}
          onChange={(e) =>
            setTodoData({ ...todoData, description: e.target.value })
          }
        />
        <button
          className="relative inline-flex items-center justify-center p-0.5 mb-2 me-2 overflow-hidden text-sm font-medium text-gray-900 rounded-lg group bg-gradient-to-br from-purple-500 to-pink-500 group-hover:from-purple-500 group-hover:to-pink-500 hover:text-white dark:text-white focus:ring-4 focus:outline-none focus:ring-purple-200 dark:focus:ring-purple-800"
        >
          <span className="flex items-center relative px-5 py-2.5 transition-all ease-in duration-75 bg-white dark:bg-gray-900 rounded-md group-hover:bg-opacity-0">
            Add
          </span>
        </button>
      </form>
    </div>
  );
}
