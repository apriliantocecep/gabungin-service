import { ApiRequest } from "../../types/api"

export interface LoginRequest {
    email: string;
    password: string;
}

export const login = async ({body}:ApiRequest<LoginRequest>) => {
    const request = await fetch(`${process.env.NEXT_PUBLIC_API_URL}/auth/login`, {
        method: 'POST',
        body: JSON.stringify({
            email: body.email,
            password: body.password,
        })
    })

    return request
}

export interface RegisterRequest {
    email: string;
    password: string;
    first_name: string;
    last_name: string;
    gender: string;
}

export const register = async ({body}:ApiRequest<RegisterRequest>) => {
    const request = await fetch(`${process.env.NEXT_PUBLIC_API_URL}/auth/register`, {
        method: 'POST',
        body: JSON.stringify({
            email: body.email,
            password: body.password,
            first_name: body.first_name,
            last_name: body.last_name,
            gender: body.gender,
        })
    })

    return request
}