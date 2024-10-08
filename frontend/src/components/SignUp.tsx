import axios from "axios";
import { useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";
import { SignUpForm } from "../types/userTypes";
import { useRecoilValue } from "recoil";
import { loginState } from "../store/atom";

export default function SignUp() {
  const [formData, setFormData] = useState<SignUpForm>({
    username: "",
    email: "",
    password: "",
  });

  const navigate = useNavigate();

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const { name, value } = e.target;
    setFormData((prevData) => ({ ...prevData, [name]: value }));
  };

  const handleSubmit = async (e: React.ChangeEvent<HTMLFormElement>) => {
    e.preventDefault();
    try {
      const response = await axios.post(
        import.meta.env.VITE_BASE_URL + "/sign-up",
        formData,
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


  const { isLogin, token } = useRecoilValue(loginState);
  const [isLoading, setIsLoading] = useState(true);

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
    <div className="flex-grow flex flex-col justify-center items-center gap-10">
      <h1 className="text-white text-4xl font-medium">Sign Up</h1>
      <form
        onSubmit={handleSubmit}
        className=" flex flex-col items-center gap-5"
      >
        <input
          type="text"
          placeholder="username"
          required
          name="username"
          value={formData?.username}
          onChange={(e) => handleChange(e)}
        />
        <input
          type="email"
          placeholder="email"
          required
          name="email"
          value={formData?.email}
          onChange={(e) => handleChange(e)}
        />
        <input
          type="text"
          placeholder="password"
          required
          name="password"
          value={formData?.password}
          onChange={(e) => handleChange(e)}
        />
        <button className="relative inline-flex items-center justify-center p-0.5 mb-2 me-2 overflow-hidden text-sm font-medium text-gray-900 rounded-lg group bg-gradient-to-br from-orange-500 to-pink-500 group-hover:from-purple-500 group-hover:to-pink-500 hover:text-white dark:text-white focus:ring-4 focus:outline-none focus:ring-purple-200 dark:focus:ring-purple-800">
          <span className="flex items-center relative px-5 py-2.5 transition-all ease-in duration-75 bg-white dark:bg-gray-900 rounded-md group-hover:bg-opacity-0">
            Sign Up
          </span>
        </button>
      </form>
    </div>
  );
}
