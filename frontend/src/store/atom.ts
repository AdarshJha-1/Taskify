import { atom } from "recoil";

const todoListState = atom({
    key: 'todoListState',
    default: [],
  });

const loginState = atom({
    key: 'loginState',
    default: {
      isLogin: false,
      token: null
    },
  });

const userState = atom({
    key: "userState",
    default: null
})

export {
  todoListState,
  loginState,
  userState
}