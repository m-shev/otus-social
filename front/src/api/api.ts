import {AddFriendForm, CreateUserForm, LoginForm} from '../types';

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

export const profileGet = async (id: string): Promise<Response> => {
    return fetch(`${baseUrl()}/user/${id}/profile`, {
        method: 'get',
    });
};

export const userProfileGet = async (): Promise<Response> => {
    return fetch(`${baseUrl()}/user/profile`, {
        method: 'get',
        credentials: 'include',
    });
};

export const logoutGet = async (): Promise<Response> => {
    return fetch(`${baseUrl()}/user/logout`, {
        method: 'get',
        credentials: 'include',
    });
};

export const addFriendPost = async (addFriendForm: AddFriendForm): Promise<Response> => {
    return fetch(`${baseUrl()}/user/friend`, {
        method: 'post',
        body: JSON.stringify(addFriendForm),
        credentials: 'include',
    });
};

export const friendListGet = async (userId: string): Promise<Response> => {
    return fetch(`${baseUrl()}/user/${userId}/friends?skip=0&take=1000`, {
        method: 'get',
    });
};
