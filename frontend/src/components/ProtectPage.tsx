import { useRecoilValue } from "recoil";
import { loginState } from "../store/atom";
import { useNavigate } from "react-router-dom";
import { useEffect } from "react";

export default function ProtectPage({
  children,
}: {
  children: React.ReactNode;
}) {
  const { isLogin, token } = useRecoilValue(loginState);
  const navigate = useNavigate();

  useEffect(() => {
    if (!isLogin || !token) {
      navigate("/signin", { replace: true });
    }
  }, [isLogin, token, navigate]);
  return <div className="flex-grow flex">{children}</div>;
}
