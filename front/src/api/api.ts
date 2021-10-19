import {CreateUserForm} from '../types';

const baseUrl = (): string => {
    return 'http://localhost:3005';
};

export const createUserPost = async (createUserForm: CreateUserForm): Promise<Response> => {
    return fetch(`${baseUrl()}/user/registration`, {
        method: 'post',
        body: JSON.stringify(createUserForm),
    });
};
