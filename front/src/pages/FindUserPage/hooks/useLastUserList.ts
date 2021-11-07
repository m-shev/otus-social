import {useCallback, useEffect, useState} from 'react';
import {userListGet} from '../../../api';
import {HttpStatus, UserFriend} from '../../../types';

const TAKE_LAST_USER = 20;

export const useLastUserList = (): UserFriend[] => {
    const [userList, setUserList] = useState([]);

    const loadLastUserList = useCallback(async () => {
        const resp = await userListGet({skip: 0, take: TAKE_LAST_USER});

        if (resp.status === HttpStatus.Ok) {
            const list = await resp.json();
            setUserList(list);
        }
    }, []);

    useEffect(() => {
        // noinspection JSIgnoredPromiseFromCall
        loadLastUserList();

        // eslint-disable-next-line
    }, []);

    return userList;
};
