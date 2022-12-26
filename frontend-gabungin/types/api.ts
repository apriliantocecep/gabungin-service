export interface ApiRequest<T> {
    body: T;
}

export interface ApiResponse<T> {
    status: number;
    error: string;
    data: T|null;
    page: number;
    pageSize: number;
}