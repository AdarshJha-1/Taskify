import axios from "axios";
import { TodoType } from "../types/todoTypes";
import { useState } from "react";
import { todoListState } from "../store/atom";
import { useSetRecoilState } from "recoil";

export default function TodoCard({
  id,
  title,
  description,
  is_completed,
}: TodoType) {
  console.log(is_completed);
  const setTodoList = useSetRecoilState(todoListState);

  const [localIsCompleted, setLocalIsCompletes] = useState(is_completed);
  const handleIsComplete = async () => {
    setLocalIsCompletes(!localIsCompleted);
    try {
      const response = await axios.put(
        import.meta.env.VITE_BASE_URL + "/todos/" + id,
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
        console.log(response.data);
      }
    } catch (error) {
      if (axios.isAxiosError(error)) {
        console.error("Axios error:", error.message);
      } else {
        console.error("Unexpected error:", error);
      }
    }
  };
  const handleDelete = async () => {
    setTodoList((prevTodos: TodoType[]) =>
      prevTodos.filter((todo) => todo.id != id)
    );
    try {
      const response = await axios.delete(
        import.meta.env.VITE_BASE_URL + "/todos/" + id,
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
        console.log(response.data);
      }
    } catch (error) {
      if (axios.isAxiosError(error)) {
        console.error("Axios error:", error.message);
      } else {
        console.error("Unexpected error:", error);
      }
    }
  };

  return (
    <div className="block w-[360px] shrink-0 max-w-sm p-6 bg-white border border-gray-200 rounded-lg shadow hover:bg-gray-100 dark:bg-gray-800 dark:border-gray-700 dark:hover:bg-gray-700">
      <h1 className="mb-2 text-2xl font-bold tracking-tight break-words text-gray-900 dark:text-white">
        {title}
      </h1>
      <div className="flex items-center justify-between gap-5">
        <input
          type="checkbox"
          checked={localIsCompleted}
          className="size-4"
          onChange={handleIsComplete}
        />
        <button onClick={handleDelete}>🗑️</button>
      </div>
      <p className="font-normal text-gray-700 dark:text-gray-400 break-words">
        {description}
      </p>
    </div>
  );
}
