import { useRecoilState, useRecoilValue } from "recoil";
import { addTodoState, todoListState } from "../store/atom";
import AddTodo from "./AddTodo";
import { useEffect } from "react";
import axios from "axios";

export default function Dashboard() {
  const [todoList, setTodoList] = useRecoilState(todoListState);
  
  const AddTodoState = useRecoilValue(addTodoState);
  useEffect(() => {
    const fetchTodos = async () => {
      try {
        const res = await axios.get(import.meta.env.VITE_BASE_URL + "/todos", {withCredentials: true})
      
        setTodoList(res.data.data.todos)
      } catch (error)  {
        console.error('Error fetching todo list:', error);
      }
    }

    fetchTodos()
  }, [setTodoList]) 

  return (
    <div className="flex-grow flex text-white relative">
      {AddTodoState && <AddTodo />}
      <div
        className={`flex-grow flex ${AddTodoState ? "blur-sm" : "blur-none"}`}
      >
        {todoList.length == 0 ? (
          <div className="flex-grow flex flex-col gap-4 justify-center items-center text-8xl font-bold text-gray-300 opacity-80">
            <h1>No Todo</h1>
            <span className="text-xl text-white">create one now</span>
          </div>
        ) : (
          <div className="text-white">{
            todoList.map((todo: any) => <div key={todo.id}>{todo.title}</div>)
          }</div>
        )}
      </div>
    </div>
  );
}
