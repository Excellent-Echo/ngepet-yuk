import ngepetYukClient from "../../../APIs/ngepetYukClient";

const resetFrom = () => {
    return{
        type: "USER_REGISTER_RESET_FORM"
    };
};

const setUserName = userName => {
    return {
        type: "USER_REGISTER_SET_USERNAME",
        payload: {
          userName: userName,
        },
    };
};

const setEmail = email => {
    return {
      type: "USER_REGISTER_SET_EMAIL",
      payload: {
        email: email,
      },
    };
  };
  
const setPassword = password => {
    return {
      type: "USER_REGISTER_SET_PASSWORD",
      payload: {
        password: password,
      },
    };
};

const setRole = role => {
    return {
        type: "USER_REGISTER_SET_ROLE",
        payload: {
            role: role,
        },
    };
};

const setErrorMessage = errorMessage => {
    return {
        type: "USER_REGISTER_SET_ERROR_MESSAGE",
        payload: {
            errorMessage: errorMessage,
        },
    };
};
  
const setSuccessMessage = successMessage => {
    return {
        type: "USER_REGISTER_SET_SUCCESS_MESSAGE",
        payload: {
            successMessage: successMessage,
        },
    };
};
  
const startLoading = () => {
    return {
        type: "USER_REGISTER_START_LOADING",
    };
  };
  
const stopLoading = () => {
    return {
        type: "USER_REGISTER_STOP_LOADING",
    };
};

const register = (userName, email, password, role) => async dispatch => {
    try {
        console.log("registering");
        dispatch(startLoading());
        dispatch(setSuccessMessage(""));
        dispatch(setErrorMessage(""));

        const submitData = {
            username: userName,
            email: email,
            password: password,
            role:role
        };

        const user = await ngepetYukClient({
            method:"POST",
            url: "/users/register",
            data: submitData,
        });

        dispatch(setSuccessMessage("Registrasi Berhasil, Silakan Login"));
        dispatch(stopLoading())
    } catch (error){
        console.log(error.response);
        //console.log{error.response};
        //dispatch(setErrorMessage(error.response.data.data.errors || ["internal server error"]));
        dispatch(stopLoading());
    }
};

const userRegisterAction = {
    resetFrom,
    setUserName,
    setEmail,
    setPassword,
    setRole,
    register,
};

export default userRegisterAction