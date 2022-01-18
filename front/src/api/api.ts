import {CreateUserForm, FindUserForm, FriendForm, LoginForm, CreatePostForm} from '../types';
import * as queryString from 'query-string';

const baseUrl = (): string => {
    // eslint-disable-next-line @typescript-eslint/ban-ts-comment
    // @ts-ignore
    return process.env.REACT_APP_BASE_URL;
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

export const addFriendPost = async (addFriendForm: FriendForm): Promise<Response> => {
    return fetch(`${baseUrl()}/user/friend`, {
        method: 'post',
        body: JSON.stringify(addFriendForm),
        credentials: 'include',
    });
};

export const removeFriendDelete = async (removeFriendForm: FriendForm): Promise<Response> => {
    return fetch(`${baseUrl()}/user/friend`, {
        method: 'delete',
        body: JSON.stringify(removeFriendForm),
        credentials: 'include',
    });
};

export const friendListGet = async (userId: string): Promise<Response> => {
    return fetch(`${baseUrl()}/user/${userId}/friends?skip=0&take=1000`, {
        method: 'get',
    });
};

export const userListGet = async (findUserForm: FindUserForm): Promise<Response> => {
    const query = queryString.stringify(findUserForm);
    return fetch(`${baseUrl()}/user/list?${query}`, {
        method: 'get',
    });
};

export const createPostPost = async (createPostForm: CreatePostForm): Promise<Response> => {
    return fetch(`${baseUrl()}/post`, {
        method: 'post',
        body: JSON.stringify(createPostForm),
        credentials: 'include',
    });
};
