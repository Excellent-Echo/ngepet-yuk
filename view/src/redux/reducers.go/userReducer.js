const Initialstate = {
    users : [],
    loading: true,
    user : {}
}

const userReducer = (state = Initialstate, action) => {
    switch(action.type) {
        case "GET_USERS":
            return{
                ...state,
                users: action.payload,
                loading:false
            }
        default:
            return state
    }
}

export default userReducer;