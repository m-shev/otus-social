import {
    CreateUserForm,
    FindUserForm,
    FriendForm,
    LoginForm,
    CreatePostForm,
    PostListQuery,
} from '../types';
import * as queryString from 'query-string';

const baseUrl = (): string => {
    // eslint-disable-next-line @typescript-eslint/ban-ts-comment
    // @ts-ignore
    return process.env.REACT_APP_BASE_URL;
};

export const createUser = async (createUserForm: CreateUserForm): Promise<Response> => {
    return fetch(`${baseUrl()}/user/registration`, {
        method: 'post',
        body: JSON.stringify(createUserForm),
        credentials: 'include',
    });
};

export const login = async (loginForm: LoginForm): Promise<Response> => {
    return fetch(`${baseUrl()}/user/auth`, {
        method: 'post',
        body: JSON.stringify(loginForm),
        credentials: 'include',
    });
};

export const getProfile = async (id: string): Promise<Response> => {
    return fetch(`${baseUrl()}/user/${id}/profile`, {
        method: 'get',
    });
};

export const getUserProfile = async (): Promise<Response> => {
    return fetch(`${baseUrl()}/user/profile`, {
        method: 'get',
        credentials: 'include',
    });
};

export const logout = async (): Promise<Response> => {
    return fetch(`${baseUrl()}/user/logout`, {
        method: 'get',
        credentials: 'include',
    });
};

export const addFriend = async (addFriendForm: FriendForm): Promise<Response> => {
    return fetch(`${baseUrl()}/user/friend`, {
        method: 'post',
        body: JSON.stringify(addFriendForm),
        credentials: 'include',
    });
};

export const removeFriend = async (removeFriendForm: FriendForm): Promise<Response> => {
    return fetch(`${baseUrl()}/user/friend`, {
        method: 'delete',
        body: JSON.stringify(removeFriendForm),
        credentials: 'include',
    });
};

export const getFriendList = async (userId: string): Promise<Response> => {
    return fetch(`${baseUrl()}/user/${userId}/friends?skip=0&take=1000`, {
        method: 'get',
    });
};

export const getUserList = async (findUserForm: FindUserForm): Promise<Response> => {
    const query = queryString.stringify(findUserForm);
    return fetch(`${baseUrl()}/user/list?${query}`, {
        method: 'get',
    });
};

export const createPost = async (createPostForm: CreatePostForm): Promise<Response> => {
    return fetch(`${baseUrl()}/post`, {
        method: 'post',
        body: JSON.stringify(createPostForm),
        credentials: 'include',
    });
};

export const getPostList = async (listQuery: PostListQuery): Promise<Response> => {
    const query = queryString.stringify(listQuery);
    return fetch(`${baseUrl()}/post?${query}`, {
        method: 'get',
        credentials: 'include',
    });
};
