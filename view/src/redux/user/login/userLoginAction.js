import ngepetYukClient from "../../../APIs/ngepetYukClient";

const setEmail = email => {
    return {
      type: "USER_LOGIN_SET_EMAIL",
      payload: {
        email: email,
      },
    };
  };
  
  const setPassword = password => ({
    type: "USER_LOGIN_SET_PASSWORD",
    payload: {
      password: password,
    },
  });
  
  const setErrorMessage = errorMessage => ({
    type: "USER_LOGIN_SET_ERROR_MESSAGE",
    payload: {
      errorMessage: errorMessage,
    },
  });
  
  const startLoading = () => ({
    type: "USER_LOGIN_START_LOADING",
  });
  
  const stopLoading = () => ({
    type: "USER_LOGIN_STOP_LOADING",
  });
  
  const login = (email, password, history) => async dispatch => {
    try {
      dispatch(startLoading());
      dispatch(setErrorMessage(""));
  
      const loginData = {
        email: email,
        password: password,
      };
  
      const user = await ngepetYukClient({
        method: "POST",
        url: "/users/login",
        data: loginData,
      });
  
      const accessToken = user.data.data.token;
    
  
      localStorage.setItem("accessToken", accessToken);

      // const submitDetailData = {
      //   first_name:"",
      //   last_name:"",
      //   sub_type: "free",
      //   no_hanphone: "",
      //   gender: "",
      //   address: "",
      //   period: 0,
      // }

      // const serDetail = await ngepetYukClient({
      //   method: "POST",
      //   url:"userdetails",
      //   headers: {
      //     Authorization: accessToken,
      //   },
      //   data: submitDetailData
      // })
  
      // const userProfile = await ngepetYukClient({
      //   method: "GET",
      //   url: "/userdetails",
      //   headers: {
      //     Authorization: accessToken,
      //   },
      // });
      //console.log(userProfile.data.data)
      //dispatch(userProfileAction.setProfileData(userProfile.data.data));
  
      history.push("/user/courses");
  
      dispatch(stopLoading());
    } catch (error) {
      console.log(error);
      // dispatch(setErrorMessage(error.response.data.data.errors || ["internal server error"]));
      dispatch(stopLoading());
    }
  };
  
  const userLoginAction = {
    setEmail,
    setPassword,
    login,
  };
  export default userLoginAction;