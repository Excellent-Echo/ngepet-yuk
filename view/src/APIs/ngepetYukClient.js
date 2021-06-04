import Axios from "axios";

const ngepetYukClient = Axios.create({
    baseURL: "http://localhost:8080",
});

export default ngepetYukClient;