const initialState = {
    email: "",
    password: "",
    errorMessage: "",
    isLoading: false,
};

const userLoginReducer = (state = initialState, action) => {
    switch (action.type) {
      case "USER_LOGIN_SET_EMAIL":
        return {
          ...state,
          email: action.payload.email,
        };
      case "USER_LOGIN_SET_PASSWORD":
        return {
          ...state,
          password: action.payload.password,
        };
      case "USER_LOGIN_SET_ERROR_MESSAGE":
        return {
          ...state,
          errorMessage: action.payload.errorMessage,
        };
  
      case "USER_LOGIN_START_LOADING":
        return {
          ...state,
          isLoading: true,
        };
  
      case "USER_LOGIN_STOP_LOADING":
        return {
          ...state,
          isLoading: false,
        };
  
      default:
        return state;
    }
};

export default userLoginReducer;