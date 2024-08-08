import { useRecoilState, useRecoilValue } from "recoil";
import { addTodoState, todoListState } from "../store/atom";
import AddTodo from "./AddTodo";
import { useEffect } from "react";
import axios from "axios";
import { TodoType } from "../types/todoTypes";
import TodoCard from "./TodoCard";

export default function Dashboard() {
  const [todoList, setTodoList] = useRecoilState(todoListState);

  const AddTodoState = useRecoilValue(addTodoState);
  useEffect(() => {
    const fetchTodos = async () => {
      try {
        const res = await axios.get(import.meta.env.VITE_BASE_URL + "/todos", {
          withCredentials: true,
        });

        setTodoList(res.data.data.todos);
      } catch (error) {
        console.error("Error fetching todo list:", error);
      }
    };

    fetchTodos();
  }, [setTodoList]);

  return (
    <div className="flex-grow flex bg-zinc-950 text-white relative">
      {AddTodoState && <AddTodo />}
      <div
        className={`flex-grow flex ${AddTodoState ? "blur-sm" : "blur-none"}`}
      >
        {todoList == null ? (
          <div className="flex-grow flex flex-col gap-4 justify-center items-center text-8xl font-bold text-gray-300 opacity-80">
            <h1>No Todo</h1>
            <span className="text-xl text-white">create one now</span>
          </div>
        ) : (
          <div className="text-white w-11/12 mx-auto py-5 grid grid-cols-3 items-start gap-3">
            {todoList?.map(({ id, title,description, user_id,is_completed}: TodoType) => (
              <TodoCard
                key={id}
                id={id}
                title={title}
                description={description}
                is_completed={is_completed}
                user_id={user_id}
              />
            ))}
          </div>
        )}
      </div>
    </div>
  );
}
