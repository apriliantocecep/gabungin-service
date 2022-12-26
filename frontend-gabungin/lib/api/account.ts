export const me = async (token: string) => {
    const request = await fetch(`${process.env.NEXT_PUBLIC_API_URL}/account/me`, {
        method: 'GET',
        headers: new Headers({
            'Authorization': 'Bearer '+token, 
        }),
    })

    return request
}

export const browseAccount = async (token: string, page: number = 1) => {
    const request = await fetch(`${process.env.NEXT_PUBLIC_API_URL}/account/browse?page=${page}`, {
        method: 'GET',
        headers: new Headers({
            'Authorization': 'Bearer '+token, 
        }),
    })

    return request
}