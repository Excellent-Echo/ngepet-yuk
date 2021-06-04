const initialState = {
    userName: "",
    email: "",
    password: "",
    role: "user",
    errorMessage: "",
    succesMessage: "",
    isLoading: false
}

const userRegisterReducer = (state = initialState, action) => {
    switch (action.type) {
        case "USER_REGISTER_RESET_FORM":
            return {
                ...initialState
            };

        case "USER_REGISTER_SET_USERNAME":
            return {
                ...state,
                userName: action.payload.userName,
            };

        case "USER_REGISTER_SET_EMAIL":
            return {
                ...state,
                email: action.payload.email,
            };

        case "USER_REGISTER_SET_PASSWORD":
            return {
              ...state,
              password: action.payload.password,
            };

        case "USER_REGISTER_SET_ROLE":
            return {
                ...state,
                role: action.payload.role,
            };

        case "USER_REGISTER_SET_ERROR_MESSAGE":
            return {
              ...state,
              errorMessage: action.payload.errorMessage,
            };
      
        case "USER_REGISTER_SET_SUCCESS_MESSAGE":
            return {
            ...state,
            successMessage: action.payload.successMessage,
            };
      
        case "USER_REGISTER_START_LOADING":
            return {
              ...state,
              isLoading: true,
            };
      
        case "USER_REGISTER_STOP_LOADING":
            return {
              ...state,
              isLoading: false,
            };
      
        default:
            return state;
    }
    
};

export default userRegisterReducer;