import { atom } from "recoil";
import { recoilPersist } from "recoil-persist";

const { persistAtom } = recoilPersist();
const todoListState = atom({
  key: "todoListState",
  default: [],
  effects_UNSTABLE: [persistAtom],
});

const loginState = atom({
  key: "loginState",
  default: {
    isLogin: false,
    token: null,
  },
  effects_UNSTABLE: [persistAtom],
});

const userState = atom({
  key: "userState",
  default: null,
});

const addTodoState = atom({
  key: "addTodoState",
  default: false,
  effects_UNSTABLE: [persistAtom],
});

export { todoListState, loginState, userState, addTodoState};
