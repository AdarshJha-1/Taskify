import { Route, Routes } from "react-router-dom"
import { Show } from "./pages/Show"
import { SignUp } from "./pages/SignUp"
import { SignIn } from "./pages/SignIn"
import { Dashboard } from "./pages/Dashboard"

export const App = () => {
  return (
    <Routes>
      <Route path="/" element={<Show/>}/>
      <Route path="/signup" element={<SignUp/>}/>
      <Route path="/signin" element={<SignIn/>}/>
      <Route path="/dashboard" element={<Dashboard/>}/>
    </Routes>
  )
}
