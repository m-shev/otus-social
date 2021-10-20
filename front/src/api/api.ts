import {CreateUserForm, LoginForm} from '../types';

const baseUrl = (): string => {
    return 'http://localhost:3005';
};

export const createUserPost = async (createUserForm: CreateUserForm): Promise<Response> => {
    return fetch(`${baseUrl()}/user/registration`, {
        method: 'post',
        body: JSON.stringify(createUserForm),
        credentials: 'include',
    });
};

export const loginPost = async (loginForm: LoginForm): Promise<Response> => {
    return fetch(`${baseUrl()}/user/auth`, {
        method: 'post',
        body: JSON.stringify(loginForm),
        credentials: 'include',
    });
};
