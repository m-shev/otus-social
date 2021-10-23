import {useCallback, useEffect, useState} from 'react';
import {userProfileGet} from '../../api';
import {HttpStatus} from '../../types';
import {userAuthEvent} from '../../store/user';

export type IUseInit = {
    isPending: boolean;
};

export const useInit = () => {
    const [pending, setPending] = useState(true);
    const loadUserProfile = useCallback(async () => {
        try {
            const resp = await userProfileGet();

            if (resp.status === HttpStatus.Ok) {
                userAuthEvent(await resp.json());
            }
        } catch (e) {
            console.log(e);
        } finally {
            setPending(false);
        }
    }, []);

    useEffect(() => {
        // noinspection JSIgnoredPromiseFromCall
        loadUserProfile();
    }, []);

    return {
        pending,
    };
};
