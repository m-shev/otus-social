import {useCallback, useEffect, useState} from 'react';
import {getUserProfile} from '../../api';
import {HttpStatus} from '../../types';
import {userAuthEvent} from '../../store/user';

export type IUseInit = {
    pending: boolean;
};

export const useInit = (): IUseInit => {
    const [pending, setPending] = useState(true);
    const loadUserProfile = useCallback(async () => {
        try {
            const resp = await getUserProfile();

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
        // eslint-disable-next-line react-hooks/exhaustive-deps
    }, []);

    return {
        pending,
    };
};
