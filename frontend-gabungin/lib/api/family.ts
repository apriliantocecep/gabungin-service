import { ApiRequest } from "../../types/api"
import Cookies from 'js-cookie';
import { Family } from "../../types/family";

export const getAllFamily = async (token: string) => {
    const request = await fetch(`${process.env.NEXT_PUBLIC_API_URL}/family/all`, {
        method: 'GET',
        headers: new Headers({
            'Authorization': 'Bearer '+token, 
        }),
    })

    return request
}