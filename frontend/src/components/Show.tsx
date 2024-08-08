import { useRecoilValue } from "recoil";
import { loginState } from "../store/atom";
import { useNavigate } from "react-router-dom";
import { useEffect, useState } from "react";

export default function Show() {
  const { isLogin, token } = useRecoilValue(loginState);
  const [isLoading, setIsLoading] = useState(true);
  const navigate = useNavigate();

  useEffect(() => {
    if (isLogin && token) {
      navigate("/dashboard", { replace: true });
    } else {
      setIsLoading(false);
    }
  }, [isLogin, token, navigate]);

  if (isLoading) {
    <div className="w-full flex-grow bg-black flex flex-col items-center justify-center text-white"></div>;
    return (
      <div className="w-full flex-grow bg-black flex flex-col items-center justify-center text-white">
        <h1>loading...</h1>
      </div>
    );
  }
  return (
    <div className="w-full flex-grow bg-black flex flex-col items-center justify-center text-white">
      <h1 className="text-5xl py-5 font-bold text-center bg-gradient-to-r from-blue-600 via-green-500 to-indigo-400 inline-block text-transparent bg-clip-text">
        Welcome to Taskify!
      </h1>
      <p className="w-2/3 p-y text-lg text-center bg-gradient-to-r from-yellow-500 via-pink-500 to-indigo-200 inline-block text-transparent bg-clip-text">
        Your productivity partner. Taskify helps you keep track of your tasks
        and stay organized effortlessly. Whether you're managing personal goals
        or professional projects, Taskify is here to assist you.
      </p>
      <p className="w-1/2 py-5 text-center bg-gradient-to-r from-indigo-500 via-white to-indigo-300 inline-block text-transparent bg-clip-text">
        Ready to start organizing your tasks? Create an account now and take the
        first step towards a more organized life!
      </p>
    </div>
  );
}
