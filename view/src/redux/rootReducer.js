import {combineReducers, conbineReducer} from "redux";

import userRegisterReducer from "./user/register/userRegisterReducer";

const rootReducer = combineReducers({
    userRegister: userRegisterReducer
})

export default rootReducer;