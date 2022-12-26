import { Account } from "../../types/account";

export const getUserName = (account: Account) => {
    return `${account.first_name||account.email} ${account.last_name||''}`.trim()
}