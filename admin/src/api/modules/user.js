import http from '../http'
import secret from "../../utils/secret";

const api = {
    url: "admin/admins",
    login_url: "admin/login"
}

export function login(data) {
    var row = JSON.parse(JSON.stringify(data))
    if (row["password"]) {
        row["password"] = secret.Base64Encode(row["password"])
    }
    return http.post(api.login_url, row)
}

export default {
    login
}
