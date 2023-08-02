import {defineStore} from "pinia";
import {LoggedUser} from "../models/user.ts";
import {ref} from "vue";
import {GET_TOKEN} from "../../utils/token.ts";

let userStorage = defineStore("user", () => {
    let logged = ref(false)
    let loggedUser = ref<LoggedUser>({
        id: 0,
        name: 'test',
        avatar: '',
    })
    let token: String | null = GET_TOKEN()

    return {
        logged,
        loggedUser,
        token
    }
});

export default userStorage;