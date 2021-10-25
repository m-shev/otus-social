import {createEvent, createStore} from 'effector';
import {User} from '../../types';

export type UserStore = {
    user: User | null;
};

const defaultState: UserStore = {
    user: null,
};

export const $userStore = createStore(defaultState);
export const userAuthEvent = createEvent<User>();
export const userLogoutEvent = createEvent();

$userStore.on(userAuthEvent, (state, user) => {
    return {user};
});

$userStore.on(userLogoutEvent, () => {
    return {user: null};
});
