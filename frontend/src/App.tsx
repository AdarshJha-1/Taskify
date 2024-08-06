import { Route, Routes } from "react-router-dom";
import Show from "./components/Show";
import SignUp from "./components/SignUp";
import SignIn from "./components/SignIn";
import Dashboard from "./components/Dashboard";
import Header from "./components/Header";
import Footer from "./components/Footer";
import ProtectPage from "./components/ProtectPage";

export default function App() {
  return (
    <main className="w-full min-h-screen flex flex-col bg-black px-4">
      <Header />
      <Routes>
        <Route path="/" element={<Show />} />
        <Route path="/signup" element={<SignUp />} />
        <Route path="/signin" element={<SignIn />} />
        <Route path="/dashboard" element={<ProtectPage>
          <Dashboard />
        </ProtectPage>} />
      </Routes>
      <Footer/>
    </main>
  );
}
